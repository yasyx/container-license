package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"udesk_license/pkg/license"
	"udesk_license/pkg/utils/exec"
	"udesk_license/pkg/utils/logger"
)

var (
	debug   bool
	appCmd  string
	appArgs string
)

var exampleRun = `
check the nginx container license
	checker --cmd=nginx --args="-g daemon off;"
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "checker",
	Short:   "checker is a command-line to check the app license ",
	Example: exampleRun,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.Info("check license !!!")
		l := license.NewLicense()
		go l.CronCheckLicense()

		logger.Info("start nginx !!!")
		err := exec.Cmd(appCmd, appArgs)
		if err != nil {
			logger.Info("start nginx error: %v", err)
		}
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
	rootCmd.Flags().StringVar(&appCmd, "cmd", "nginx1", "app exec cmd")
	rootCmd.Flags().StringVar(&appArgs, "args", "-g daemon off;", "app exec args")
}
