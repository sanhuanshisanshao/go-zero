package validate

import (
	"errors"
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/sanhuanshisanshao/go-zero/tools/goctl/api/parser"
	"github.com/spf13/cobra"
)

// VarStringAPI describes an API.
var VarStringAPI string

// GoValidateApi verifies whether the api has a syntax error
func GoValidateApi(_ *cobra.Command, _ []string) error {
	apiFile := VarStringAPI

	if len(apiFile) == 0 {
		return errors.New("missing -api")
	}

	spec, err := parser.Parse(apiFile)
	if err != nil {
		return err
	}

	err = spec.Validate()
	if err == nil {
		fmt.Println(aurora.Green("api format ok"))
	}
	return err
}
