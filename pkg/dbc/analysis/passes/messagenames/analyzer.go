package messagenames

import (
	"github.com/cleey/can-go/pkg/dbc"
	"github.com/cleey/can-go/pkg/dbc/analysis"
	"github.com/cleey/can-go/pkg/identifiers"
)

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "messagenames",
		Doc:  "check that message names are valid CamelCase identifiers",
		Run:  run,
	}
}

func run(pass *analysis.Pass) error {
	for _, def := range pass.File.Defs {
		messageDef, ok := def.(*dbc.MessageDef)
		if !ok {
			continue // not a message
		}
		if !identifiers.IsCamelCase(string(messageDef.Name)) {
			pass.Reportf(messageDef.Pos, "message names must be CamelCase")
		}
	}
	return nil
}
