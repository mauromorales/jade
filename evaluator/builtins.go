package evaluator

import "github.com/mauromorales/jade/object"

var builtins = map[string]*object.Builtin{
	"largo": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("número de argumenots inválido. Recibí %d en lguar de 1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("la función `largo` no soporta un argumento de tipo %s", args[0].Type())
			}
		},
	},
}
