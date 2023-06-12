package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"udesk_license/pkg/constants"
	"udesk_license/pkg/utils/encrypt"
	myfile "udesk_license/pkg/utils/file"
	"udesk_license/pkg/utils/logger"
)

var months int64 = 3

func main() {
	if len(os.Args) > 1 {
		logger.Info("months:", os.Args[1])
		months, _ = strconv.ParseInt(os.Args[1], 10, 64)
	}
	formatter := fmt.Sprintf("%s&%s", constants.UUID,
		strconv.FormatInt(time.Now().AddDate(0, int(months), 0).Unix(), 10))

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
		return
	}
	logger.Info(s)
}
