package scripts

import (
	"fmt"
	"time"

	"github.com/junmocsq/bookstore/api/models/user"
	"github.com/junmocsq/bookstore/api/tools"
	"github.com/robfig/cron/v3"
)

func init() {

	// 创建一个定时任务调度器
	c := cron.New()

	// 添加定时任务，每一段时间执行一次
	_, err := c.AddFunc("*/10 * * * *", func() {
		// 在这里执行您的任务代码
		fmt.Println("定时任务执行时间:", time.Now())

		// TODO
		for i := 0; i < 20; i++ {
			go func() {
				user.NewUser().GetByAccount(tools.CreateRandomString(10))
			}()
		}
	})

	if err != nil {
		fmt.Println("添加定时任务失败:", err)
		return
	}

	// 启动定时任务调度器
	c.Start()

}
