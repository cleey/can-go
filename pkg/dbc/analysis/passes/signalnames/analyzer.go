package signalnames

import (
	"github.com/cleey/can-go/pkg/dbc"
	"github.com/cleey/can-go/pkg/dbc/analysis"
	"github.com/cleey/can-go/pkg/identifiers"
)

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "signalnames",
		Doc:  "check that signal names are valid CamelCase identifiers",
		Run:  run,
	}
}

func run(pass *analysis.Pass) error {
	for _, d := range pass.File.Defs {
		messageDef, ok := d.(*dbc.MessageDef)
		if !ok {
			continue
		}
		for _, signalDef := range messageDef.Signals {
			if !identifiers.IsCamelCase(string(signalDef.Name)) {
				pass.Reportf(signalDef.Pos, "signal names must be CamelCase")
			}
		}
	}
	return nil
}
