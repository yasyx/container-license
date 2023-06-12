package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yasyx/container-license/pkg/constants"
	"github.com/yasyx/container-license/pkg/utils/encrypt"
	myfile "github.com/yasyx/container-license/pkg/utils/file"
	"github.com/yasyx/container-license/pkg/utils/logger"
	"os"
	"strconv"
	"time"
)

var (
	debug    bool
	month    int
	duration string
)

var exampleRun = `
generate --duration=10m ，10 minutes license
such as "300ms", "-1.5h" or "2h45m".
generate --month=3 ，3 months license
"
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "generate",
	Short:   "generate is a command-line to generate the app license ",
	Example: exampleRun,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 默认3月有效期
		licenseDate := time.Now().AddDate(0, 3, 0)
		// 按月配置有效期
		if month > 0 {
			licenseDate = time.Now().AddDate(0, month, 0)
		}
		// 按时间配置有效期
		if duration != "" {
			parseDuration, err := time.ParseDuration(duration)
			if err != nil {
				return err
			}
			licenseDate = time.Now().Add(parseDuration)
		}
		formatter := fmt.Sprintf("%s&%s", constants.UUID,
			strconv.FormatInt(licenseDate.Unix(), 10))

		s, err := encrypt.Encrypt(formatter, constants.SECRET)
		if err != nil {
			logger.Error(err)
		}
		exist := myfile.CheckFileIsExist("license")
		if exist {
			err := os.Remove("license")
			if err != nil {
				logger.Warn("Remove file error")
			}
		}
		err = os.WriteFile("license", []byte(s), 0666)
		if err != nil {
			logger.Error(err)
		}
		logger.Info(s)
		return err
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		logger.CfgConsoleLogger(debug, false)
	})
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug logger")
	rootCmd.Flags().IntVar(&month, "month", 3, "license month")
	rootCmd.Flags().StringVar(&duration, "duration", "1m", "license second")
}
