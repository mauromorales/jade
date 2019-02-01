package evaluator

import (
	"fmt"

	"github.com/mauromorales/jade/object"
)

var builtins = map[string]*object.Builtin{
	"largo": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("número de argumentos inválido. Recibí %d en lugar de 1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("la función `largo` no soporta un argumento de tipo %s", args[0].Type())
			}
		},
	},
	"adjuntar": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return newError("número de argumenots inválido. Recibí %d en lugar de 2", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Array{Elements: append(arg.Elements, args[1:]...)}
			}

			return nil
		},
	},
	"imprimir": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
}
