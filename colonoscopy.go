package main

import (
	"os"
	"fmt"
	"log"
	"strings"
	"io/ioutil"
	"github.com/pkg/term"
)

type stack []int

var memory [1024]int
var tokens []func()
var ip = 0
var ptr = 0
var while = make(stack, 0)

var symbols = map[string]func(){
	";};": func() {
		tokens = append(tokens, func() {
			ptr++
		})
	},
	";{;": func() {
		tokens = append(tokens, func() {
			ptr--
		})
	},
	";;};": func() {
		tokens = append(tokens, func() {
			memory[ptr]++
		})
	},
	";;{;": func() {
		tokens = append(tokens, func() {
			memory[ptr]--
		})
	},
	";;;};": func() {
		tokens = append(tokens, func() {
			fmt.Printf("%c", memory[ptr])
		})
	},
	";;;{;": func() {
		tokens = append(tokens, func() {
			memory[ptr], _ = getChar()
		})
	},
	"{{;": func() {
		while = while.push(len(tokens))
		tokens = append(tokens, func() {
			panic("Error: {{ without }}")
		})
	},
	"}};": func() {
		var jmp int
		while, jmp = while.pop()

		tokens[jmp] = func() {
			if memory[ptr] == 0 {
				ip = len(tokens)
			}
		}

		tokens = append(tokens, func() {
			if memory[ptr] != 0 {
				ip = jmp
			}
		})
	},
}

func getChar() (ascii int, err error) {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bytes := make([]byte, 3)

	var numRead int
	numRead, err = t.Read(bytes)

	if err != nil {
		return
	}

	if numRead == 1 {
		ascii = int(bytes[0])
	}

	t.Restore()
	t.Close()
	return
}

func (s stack) push(v int) stack {
	return append(s, v)
}

func (s stack) pop() (stack, int) {
	l := len(s)

	if l == 0 {
		panic("Error: }} without {{")
	}

	return s[:l - 1], s[l - 1]
}

func main() {
	filename := os.Args[1]
	buffer, err := ioutil.ReadFile(filename)
	source := string(buffer)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Tokenising...")

	source = strings.TrimSpace(source)

	for len(source) > 0 {
		found := false
		for symbol, fn := range symbols {
			if strings.HasPrefix(source, symbol) {
				source = strings.Replace(source, symbol, "", 1)
				found = true
				fn()
				break
			}
		}

		if (!found) {
			panic("Syntax Error");
		}
	}

	fmt.Println("done.")

	for ip < len(tokens) {
		fn := tokens[ip]
		ip++
		fn()
	}

	fmt.Println("Execution complete!")
}
