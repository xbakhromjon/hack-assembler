package main

import (
	"errors"
	"log"
	"os"
	"strings"
)

func main() {
	filepath := GetArg("-f")

	asmFile, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("open asm file: %s", err.Error())
	}

	defer func() {
		err := asmFile.Close()
		if err != nil {
			log.Fatalf("close asm file: %s", err.Error())
		}
	}()

	parser, err := NewParser(asmFile)
	if err != nil {
		log.Fatal(err)
	}

	code := NewCode()

	outfilepath := strings.Replace(filepath, ".asm", ".hack", 1)

	file, err := os.OpenFile(outfilepath, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Fatalf("open file: %s", err.Error())
		}
		log.Println("outfilepath: " + outfilepath)
		file, err = os.Create(outfilepath)
		if err != nil {
			log.Fatalf("create output file: %s", err.Error())
		}
	} else {
		err := file.Truncate(0)
		if err != nil {
			log.Fatalf("truncate output file: %s", err.Error())
		}
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalf("close hack file: %s", err.Error())
		}
	}()

	symTable := NewSymbolTable()

	// first-pass
	for parser.ScanLabel() {
		if !parser.IsLabel() {
			continue
		}
		label := parser.Label()
		symTable.Add(label, parser.line)
	}

	err = parser.Reset()
	if err != nil {
		log.Fatal(err)
	}

	// second-pass
	for parser.Scan() {
		binary := ""
		if parser.A() {
			var addr uint32 = 0
			if parser.isANumeric() {
				log.Printf("A numeric: %s", parser.current)
				addrVal, err := parser.Address()
				if err != nil {
					log.Fatalf("extract a val: %s", err.Error())
				}
				addr = addrVal
			} else {
				symbol, err := parser.Symbol()
				if err != nil {
					log.Fatalf("extract symbol from A instruction: %s", err.Error())
				}
				log.Printf("A symbol: %s", parser.current)
				addr = symTable.Get(symbol)
			}
			binAddr, err := ConvertTo15BitBinary(addr)
			if err != nil {
				log.Fatalf("convert a val to 15 bit: %s", err.Error())
			}
			log.Printf("A binary: %s", binAddr)
			binary = "0" + binAddr
		}

		if parser.C() {
			dest, _ := parser.Dest()
			comp, _ := parser.Comp()
			jump, _ := parser.Jump()

			destBin, _ := code.Dest(dest)
			compBin, _ := code.Comp(comp)
			jumpBin, _ := code.Jump(jump)

			binary = "111" + compBin + destBin + jumpBin
		}
		_, err = file.Write([]byte(binary + "\r\n"))
		if err != nil {
			log.Fatalf("write file: %s", err.Error())
		}
	}
	err = parser.ScanErr()
	if err != nil {
		log.Fatalf("scan error: %s", err.Error())
	}

}
