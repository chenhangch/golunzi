package commandlinetool

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	configFlagName        = "config"
	defaultConfigFileType = "yaml"
)

var configFile string

var configIn string

func init() {
	pflag.StringVarP(&configFile, configFlagName, "c", configFile, "set you config file && file type is yaml")
}

func addConfigFlag(basename string, fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(configFlagName))

	// viper 首先读取环境变量 前缀为basename改大写，且环境变量采用"_"而非"-"或者"."
	viper.AutomaticEnv()
	viper.SetEnvPrefix(strings.Replace(strings.ToUpper(basename), "-", "_", -1))
	viper.SetEnvKeyReplacer(strings.NewReplacer(",", "_", "-", "_"))

	// 设置在调用每个命令的 Execute 方法时要运行的传递函数
	cobra.OnInitialize(func() {
		fmt.Println(color.GreenString("===========> viper read config"))
		if configFile != "" {
			// 需要显示输入配置文件的路径、文件名称以及扩展名、
			viper.SetConfigFile(configFile)
		} else {
			if configIn != "" {
				//viper.SetConfigFile(configIn)
				viper.AddConfigPath(configIn)
			} else {
				// 默认读取文件的位置
				viper.AddConfigPath("./config")
			}
			viper.SetConfigName(basename)
			viper.SetConfigType(defaultConfigFileType)
		}

		if err := viper.ReadInConfig(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: failed to read configuration file(%s): %v\n", configFile, err)
			os.Exit(1)
		}
	})
}

func SetConfigIn(configFile string) {
	configIn = configFile
}
