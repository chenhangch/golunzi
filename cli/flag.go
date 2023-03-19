package cli

import (
	"strings"

	"github.com/spf13/pflag"
)

// wordSepNornamlizeFunc 标准化参数名称 
func wordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	form := []string{"-","_"}
	to := "."
	for _, sep := range form {
		name = strings.Replace(name,sep,to,-1)
	}
	return pflag.NormalizedName(name)
}