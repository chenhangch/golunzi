package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const configFlagName = "config"

var configFile string

func init()  {
	pflag.StringVarP(&configFile,configFlagName,"c",configFile,"set you config file || file type is yaml")
}

func addConfigFlag(basename string, fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(configFlagName))

	// viper 首先读取环境变量 前缀为basename改大写，且环境变量采用"_"而非"-"或者"."
	viper.AutomaticEnv()
	viper.SetEnvPrefix(strings.Replace(strings.ToUpper(basename),"-","_",-1))
	viper.SetEnvKeyReplacer(strings.NewReplacer(",","_","-","_"))

	// 设置在调用每个命令的 Execute 方法时要运行的传递函数
	cobra.OnInitialize(func() {
		if configFile != ""{
			// 需要显示输入配置文件的路径、文件名称以及扩展名、
			viper.SetConfigFile(configFile)
		} else {
			// 当前文件的位置
			viper.AddConfigPath(".")
			viper.AddConfigPath("./bestEx/config")
			// TODO: 多配置文件路径选择

			viper.SetConfigName(basename)
			viper.SetConfigType("yaml")
		}

		if err := viper.ReadInConfig(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: failed to read configuration file(%s): %v\n", configFile, err)
			os.Exit(1)
		}
	})
}