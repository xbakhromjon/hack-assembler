package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
)

var digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

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

	l := 0
	for cmpScanner.Scan() {
		if !binScanner.Scan() {
			return errors.New("bin file is smaller")
		}

		if cmpScanner.Text() != binScanner.Text() {
			return errors.New(fmt.Sprintf("line=%d expected: %s; actual: %s", l, cmpScanner.Text(), binScanner.Text()))
		}
		l++
	}
	if binScanner.Scan() {
		return errors.New("bin file is larger")
	}
	return nil
}

func IsNumeric(val string) bool {
	runes := []rune(val)
	for _, c := range runes {
		contains := slices.Contains(digits, string(c))
		if !contains {
			return false
		}
	}
	return true
}

func ParseStrToUint(str string) (uint32, error) {
	val, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(val), nil
}
