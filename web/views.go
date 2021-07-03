package web

import (
	"agent/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 同步任务
func ScriptSyncView(c *gin.Context) {
	var request_data ScriptBody
	// 接收并绑定请求参数， 如果出现异常则直接返回
	if err := c.BindJSON(&request_data); err != nil {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	// 执行命令
	scriptResponse := ScriptSyncService(request_data)
	c.JSON(http.StatusOK, scriptResponse)
}

// TODO: 异步任务
func ScriptAsyncView(c *gin.Context) {
	var request_data ScriptBody
	// 接收并绑定请求参数， 如果出现异常则直接返回
	if err := c.BindJSON(&request_data); err != nil {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	scriptResponse := ScriptAsyncService(request_data)
	c.JSON(http.StatusOK, scriptResponse)
}

// TODO: 文件传输
func FileTransferView(c *gin.Context) {
	var request_data FileBody
	// 接收并绑定请求参数， 如果出现异常则直接返回
	if err := c.BindJSON(&request_data); err != nil {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	fileTransferResponse := FileTransferService(request_data)
	c.JSON(http.StatusOK, fileTransferResponse)
}

// 获取当前应用的加载的配置信息
func CurrentConfigView(c *gin.Context) {
	config := config.GetConfiguration()
	c.JSON(http.StatusOK, config)
}

// TODO: 关闭异步任务
func TerminateScriptAsyncView(c *gin.Context) {

}

// TODO: 创建crontab任务
func CreateCrontabView(c *gin.Context) {

}

// TODO: 删除crontab任务
func DeleteCrontabView(c *gin.Context) {

}

// TODO: 查询crontab任务列表
func ListCrontabView(c *gin.Context) {

}
