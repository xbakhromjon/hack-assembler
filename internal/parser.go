package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Parser struct {
	file    *os.File
	scanner *bufio.Scanner
	current string
}

func NewParser(sc *bufio.Scanner) (*Parser, error) {
	return &Parser{scanner: sc}, nil
}

func (p *Parser) Scan() bool {
	for {
		if !p.scanner.Scan() {
			return false
		}
		text := p.scanner.Text()
		if text == "" || strings.HasPrefix(text, "//") {
			continue
		}
		p.current = p.scanner.Text()
		return true
	}
}

func (p *Parser) ScanErr() error {
	return p.scanner.Err()
}

func (p *Parser) A() bool {
	return strings.HasPrefix(p.current, "@")
}

func (p *Parser) Aval() (uint32, error) {
	val, err := strconv.ParseUint(p.current[1:], 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(val), nil
}

func (p *Parser) C() bool {
	return strings.Contains(p.current, "=")
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
	parts = strings.Split(parts[1], ";")
	return parts[0], nil
}

func (p *Parser) Jump() (string, error) {
	parts := strings.Split(p.current, "=")
	parts = strings.Split(parts[1], ";")
	if len(parts) == 1 {
		return "", nil
	}
	return parts[1], nil
}
