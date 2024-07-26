package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMainFunc(t *testing.T) {
	cases := []struct {
		name    string
		asmFile string
		cmpFile string
	}{
		// without any labels
		{"add", "../test-files/Add.asm", "../test-files/Add.cmp"},
		{"maxL", "../test-files/MaxL.asm", "../test-files/MaxL.cmp"},
		{"rectL", "../test-files/RectL.asm", "../test-files/RectL.cmp"},
		{"pongL", "../test-files/PongL.asm", "../test-files/PongL.cmp"},

		// only goto label
		{"max", "../test-files/Max.asm", "../test-files/Max.cmp"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			orgArgs := os.Args
			newArgs := append(orgArgs, "-f", c.asmFile)
			os.Args = newArgs

			main()
			binFile := strings.Replace(c.asmFile, ".asm", ".hack", 1)
			err := CompareFiles(binFile, c.cmpFile)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestAbc(t *testing.T) {
	// Open the file
	file, err := os.Open("../test-files/Add.asm")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file the first time
	scanner := bufio.NewScanner(file)
	fmt.Println("First read:")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Seek to the beginning of the file
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}

	// Read the file the second time
	scanner = bufio.NewScanner(file)
	fmt.Println("Second read:")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}
