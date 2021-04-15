package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func main() {
	fmt.Println("Cron Starting ...")

	// 实例化cron
	cr := cron.New()

	// 设置cron任务
	cr.AddFunc("* * * * * *", func() {
		fmt.Println("Run Hello...")
	})

	// 运行cron
	cr.Start()

	// 定时器10s执行一次
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}

}
