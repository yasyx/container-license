package license

import (
	"github.com/robfig/cron/v3"
	"github.com/yasyx/container-license/pkg/constants"
	"github.com/yasyx/container-license/pkg/utils/encrypt"
	"github.com/yasyx/container-license/pkg/utils/logger"
	"os"
	"strconv"
	"strings"
	"time"
)

type License struct {
	lastCheckTime int64  // 上次检查时间,防止修改系统时间 Unix时间戳
	uuid          string // 程序的UUID 防License复制
	expiryTime    int64  // 过期时间，Unix时间戳
}

func NewLicense() *License {
	err, uuid, expiryTime := decryptLicense()
	if err != nil {
		logger.Error("授权文件解析失败")
		os.Exit(1)
	}
	return &License{lastCheckTime: time.Now().Unix(), uuid: uuid, expiryTime: expiryTime}
}

func decryptLicense() (error, string, int64) {
	file, err := os.ReadFile("/app/license")
	if err != nil {
		return err, "", 0
	}
	str := string(file)
	decrypt, err := encrypt.Decrypt(str, constants.SECRET)
	if err != nil {
		return err, "", 0
	}
	// uuid&expiryTime
	split := strings.Split(decrypt, "&")
	uuid := split[0]
	expiryTimeStr := split[1]
	expiryTime, err := strconv.ParseInt(expiryTimeStr, 10, 64)
	logger.Info("uuid:%v,expiryTime:%v", uuid, time.Unix(expiryTime, 0).Add(time.Hour*8).Format("2006-01-02 15:04:05"))
	return nil, uuid, expiryTime
}

func (l *License) CronCheckLicense() {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("*/5 * * * * ?", func() {
		license := l.checkLicense()
		if !license {
			os.Exit(1)
		}
	})
	if err != nil {
		os.Exit(1)
	}
	// 开启
	c.Start()
	defer c.Stop()
	select {}
}

func (l *License) checkLicense() bool {
	// 重新读取license 文件
	err, s, i := decryptLicense()
	if err != nil {
		os.Exit(1)
	}
	l.uuid = s
	l.expiryTime = i

	var now = time.Now().Unix()
	// 校验license 复制
	if l.uuid != constants.UUID {
		logger.Info("UUID不匹配 ")
		return false
	}
	// 校验系统时间是否被修改过
	if l.lastCheckTime < now {
		l.lastCheckTime = now
	} else {
		logger.Info("系统时间修改过 ")
		return false
	}
	// 校验过期时间
	if l.expiryTime < now {
		logger.Info("授权已过期: 授权到期时间为:%v", time.Unix(l.expiryTime, 0).Add(time.Hour*8).Format("2006-01-02 15:04:05"))
		return false
	}
	return true
}
