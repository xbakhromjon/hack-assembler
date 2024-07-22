package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetArg(name string) string {
	args := os.Args
	next := false
	for _, a := range args {
		if next {
			return a
		}
		if a == name {
			next = true
		}
	}
	return ""
}

func ConvertTo15BitBinary(number uint32) (string, error) {
	if number > (1 << 15) {
		return "", errors.New("argument out of range 15 bit")
	}
	binary := strconv.FormatInt(int64(number), 2)
	for len(binary) < 15 {
		binary = "0" + binary
	}
	return binary, nil
}

func CompareFiles(binFilepath string, cmpFilepath string) error {
	cmpFile, err := os.Open(cmpFilepath)
	if err != nil {
		return err
	}
	cmpScanner := bufio.NewScanner(cmpFile)

	binFile, err := os.Open(binFilepath)
	if err != nil {
		return err
	}
	binScanner := bufio.NewScanner(binFile)

	for cmpScanner.Scan() {
		if !binScanner.Scan() {
			return errors.New("bin file is smaller")
		}

		if cmpScanner.Text() != binScanner.Text() {
			return errors.New(fmt.Sprintf("expected: %s; actual: %s", cmpScanner.Text(), binScanner.Text()))
		}
	}
	if binScanner.Scan() {
		return errors.New("bin file is larger")
	}
	return nil
}
