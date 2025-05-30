package task

import (
	"github.com/robfig/cron/v3"
	"time"

	"go-cloud-disk/utils/logger"
)

var Cron *cron.Cron

type jobFunc func() error

// Run runing job and print result that job executed
func Run(jobName string, job jobFunc) {
	// caculate job executed time
	from := time.Now().UnixNano()
	err := job()
	to := time.Now().UnixNano()
	if err != nil {
		logger.Log().Error("%s error: %dms\n err:%v", jobName, (to-from)/int64(time.Millisecond), err)
	} else {
		logger.Log().Info("%s success: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	}
}

// CronJob start cron job
func CronJob() {
	if Cron == nil {
		Cron = cron.New()
	}

	// every day restart dailyrank in 0:0:0
	if _, err := Cron.AddFunc("@daily", func() { Run("restart daily rank", RestartDailyRank) }); err != nil {
		logger.Log().Error("set restart daily rank func err", err)
	}
	// every day delete last day file in 1:0:0
	if _, err := Cron.AddFunc("0 1 * * *", func() { Run("delete last day file", DeleteLastDayFile) }); err != nil {
		logger.Log().Error("set delete last day file func err", err)
	}
	Cron.Start()
}
