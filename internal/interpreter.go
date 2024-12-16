package interpreter

import (
	"fmt"
	"os"
)

type Interpreter struct {
	Tape  [30000]uint8
	Ptr   int
	Input []byte
}

func InterpreterError(msg string, idx int) string {
	return fmt.Sprintf("an error occured at %d character: %s", idx, msg)
}

func NewInterpreter(input []byte) *Interpreter {
	tape := [30000]uint8{}

	return &Interpreter{
		Tape:  tape,
		Ptr:   0,
		Input: input,
	}
}

func (p *Interpreter) Run() {
	bracketCounter := 0

	for i := 0; i < len(p.Input); i++ {
		switch p.Input[i] {
		case '>':
			p.Ptr++
		case '<':
			p.Ptr--
		case '+':
			p.Tape[p.Ptr]++
		case '-':
			p.Tape[p.Ptr]--
		case '.':
			fmt.Print(string(p.Tape[p.Ptr]))
		case ',':
			var input string
			if _, err := fmt.Scan(&input); err != nil {
				fmt.Println(InterpreterError("failed to read input", i))
				os.Exit(1)
			}
			runes := []rune(input)
			if len(runes) > 1 {
				fmt.Println(InterpreterError(fmt.Sprintf("recieved %s. expected single character input", input), i))
				os.Exit(1)
			}
			p.Tape[p.Ptr] = byte(runes[0])
		case '[':
			if p.Tape[p.Ptr] == 0 {
				bracketCounter++

				for p.Input[i] != ']' || bracketCounter != 0 {
					i++

					if p.Input[i] == '[' {
						bracketCounter++
					} else if p.Input[i] == ']' {
						bracketCounter--
					}
				}
			}
		case ']':
			if p.Tape[p.Ptr] != 0 {
				bracketCounter++

				for p.Input[i] != '[' || bracketCounter != 0 {
					i--

					if p.Input[i] == ']' {
						bracketCounter++
					} else if p.Input[i] == '[' {
						bracketCounter--
					}
				}
			}
		}
	}
}
