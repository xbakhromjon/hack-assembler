package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Parser struct {
	file    *os.File
	scanner *bufio.Scanner
	current string
	line    uint32
}

func NewParser(file *os.File) (*Parser, error) {
	scanner := bufio.NewScanner(file)
	return &Parser{file: file, scanner: scanner}, nil
}

func (p *Parser) Scan() bool {
	for {
		if !p.scanner.Scan() {
			return false
		}
		text := p.scanner.Text()
		text = strings.TrimSpace(text)
		if text == "" || strings.HasPrefix(text, "//") || strings.HasPrefix(text, "(") {
			continue
		}
		p.line++
		p.current = text
		return true
	}
}

func (p *Parser) ScanLabel() bool {
	for {
		if !p.scanner.Scan() {
			return false
		}
		text := p.scanner.Text()
		text = strings.TrimSpace(text)
		if text == "" || strings.HasPrefix(text, "//") {
			continue
		}
		p.current = text
		if !p.IsLabel() {
			p.line++
		}
		return true
	}
}

func (p *Parser) ScanErr() error {
	return p.scanner.Err()
}

func (p *Parser) A() bool {
	return strings.HasPrefix(p.current, "@")
}

func (p *Parser) isANumeric() bool {
	return IsNumeric(p.current[1:])
}

func (p *Parser) Address() (uint32, error) {
	val, err := strconv.ParseUint(p.current[1:], 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(val), nil
}

func (p *Parser) Symbol() (string, error) {
	return p.current[1:], nil
}

func (p *Parser) C() bool {
	return strings.Contains(p.current, "=") || strings.Contains(p.current, ";")
}

func (p *Parser) Dest() (string, error) {
	parts := strings.Split(p.current, "=")
	if len(parts) == 1 {
		return "", nil
	}
	return parts[0], nil
}

func (p *Parser) Comp() (string, error) {
	parts := strings.Split(p.current, "=")
	if len(parts) == 1 {
		parts = strings.Split(parts[0], ";")
	} else {
		parts = strings.Split(parts[1], ";")
	}
	return parts[0], nil
}

func (p *Parser) Jump() (string, error) {
	parts := strings.Split(p.current, "=")
	if len(parts) == 1 {
		parts = strings.Split(parts[0], ";")
	} else {
		parts = strings.Split(parts[1], ";")
	}
	if len(parts) == 1 {
		return "", nil
	}
	return parts[1], nil
}

func (p *Parser) IsLabel() bool {
	return strings.HasPrefix(p.current, "(") && strings.HasSuffix(p.current, ")")
}

func (p *Parser) Label() string {

	return p.current[1 : len(p.current)-1]
}

func (p *Parser) Reset() error {
	_, err := p.file.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("error seeking file: %s", err.Error())
	}
	p.scanner = bufio.NewScanner(p.file)
	p.current = ""
	p.line = 0
	return nil
}
