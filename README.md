## About The Project

This project developed as part of the https://www.nand2tetris.org/ course.

## Usage

```
go build -o hack-asm ./internal
./hack-asm -f foo.asm
```

## Translating algorithm

# Init

- init symbol table
- add pre-defined symbols

# First-pass

- foreach over asm lines
    - read line
    - if line start with `(` and end with `)`
        - parse (XXX) into XXX
        - add to symbol table (XXX, line+1)

# Second-pass

- n = 16 (starter memory position for variables)
- foreach over asm lines
    - read line
    - if start with `@`
        - if `@symbol`
            - if not exist in symbol table
                - add (symbol, n)
                - n++
            - replace symbol with its address
        - translate A-instruction
    - if C-instruction
        - parse command into parts(dest, comp, jump)
        - translate dest, comp, jump
        - put all together
    - write to output file

## Contact

Email - [xbakhromjon@gmail.com](xbakhromjon@gmail.com) \
Linkedin - [linkedin.com/in/xbakhromjon](https://www.linkedin.com/in/xbakhromjon/)