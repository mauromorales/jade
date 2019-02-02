package runner

import (
	"bufio"
	"io"
	"os"

	"github.com/mauromorales/jade/evaluator"
	"github.com/mauromorales/jade/lexer"
	"github.com/mauromorales/jade/object"
	"github.com/mauromorales/jade/parser"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluator.Eval(program, env)
		if len(program.Errors) != 0 {
			printEvaluationErrors(out, program.Errors)
		}

		if len(p.Errors())+len(program.Errors) != 0 {
			os.Exit(1)
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Error(es) de parseo:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func printEvaluationErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Error(es) de corrida:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
