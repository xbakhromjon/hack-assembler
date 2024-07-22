package main

import (
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
		{"add", "../files/Add.asm", "../files/Add.cmp"},
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
