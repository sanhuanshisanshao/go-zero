package env

import (
	"fmt"

	"github.com/sanhuanshisanshao/go-zero/tools/goctl/pkg/env"
	"github.com/spf13/cobra"
)

func write(_ *cobra.Command, _ []string) error {
	if len(sliceVarWriteValue) > 0 {
		return env.WriteEnv(sliceVarWriteValue)
	}
	fmt.Println(env.Print())
	return nil
}
