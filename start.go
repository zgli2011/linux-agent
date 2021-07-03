package main

import (
	"agent/core/execute_command"
	"agent/web"

	"github.com/gin-gonic/gin"
)

func gin_web() {
	router := gin.Default()
	// router.Use(task_info())
	router.POST("/agent/current_config", web.CurrentConfigView)        // 查询当前配置
	router.POST("/agent/script/sync", web.ScriptSyncView)              // 同步任务
	router.POST("/agent/script/async", web.ScriptAsyncView)            // 异步任务
	router.POST("/agent/file/transfer", web.FileTransferView)          // 文件传输
	router.DELETE("/agent/script/async", web.TerminateScriptAsyncView) // 拦截异步任务
	// router.POST("/agent/crontab", web.CreateCrontabView)               // 新建本地crontab任务
	// router.DELETE("/agent/crontab", web.DeleteCrontabView)             // 删除本地crontab任务
	// router.PUT("/agent/crontab", web.ModifyCrontabView)                // 修改本地crontab任务
	// router.GET("/agent/crontab", web.ListCrontabView)                  // 查看本地crontab任务
	router.Run(":3000")

}

// // 定时向proxy代理上报存活状态
// func agent_report() {

// }

// 定时任务用于处理异步任务
// func clean_task_result() {
// 	c := cron.New()
// 	c.AddFunc("@every 5s", func() { fmt.Println("Every 5 second, starting an hour thirty from now") })
// 	c.Start()
// 	defer c.Stop()
// 	select {}
// }

func main() {
	// 启动web项目
	gin_web()
	// 初始化队列
	execute_command.NewDataContainer(100)
	// 启动执行队列
	go execute_command.ScriptExecuteQueue()
	// a := queue.NewDataContainer(1000)
	// b := a.Pop()
	// fmt.Println(b)
	// go script_execute_queue()
	// go task_queue()
	// go agent_report()
}
