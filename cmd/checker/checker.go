package main

import "udesk_license/cmd/checker/cmd"

func main() {
	cmd.Execute()
	//logger.Info("check license !!!")
	//l := license.NewLicense()
	//go l.CronCheckLicense()
	//
	//logger.Info("start nginx !!!")
	//err := exec.Cmd("nginx", "-g daemon off;")
	//if err != nil {
	//	logger.Info("start nginx error: %v", err)
	//}
}
