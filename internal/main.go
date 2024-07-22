package main

import (
	"bufio"
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

	scanner := bufio.NewScanner(asmFile)
	parser, err := NewParser(scanner)
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

	for parser.Scan() {
		binary := ""
		if parser.A() {
			addr, err := parser.Aval()
			if err != nil {
				log.Fatalf("extract a val: %s", err.Error())
			}
			binAddr, err := ConvertTo15BitBinary(addr)
			if err != nil {
				log.Fatalf("convert a val to 15 bit: %s", err.Error())
			}
			binary = "0" + binAddr
		}

		if parser.C() {
			dest, _ := parser.Dest()
			comp, _ := parser.Comp()
			jump, _ := parser.Jump()

			log.Printf("dest: %s comp: %s jump: %s \n", dest, comp, jump)

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
