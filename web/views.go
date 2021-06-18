package web

import (
	"agent/common"
	"agent/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 同步任务
func ScriptSyncView(c *gin.Context) {
	var request_data ScriptBody

	if err := c.BindJSON(&request_data); err != nil {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(200, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	scriptResponse := ScriptSyncService(request_data)
	c.JSON(http.StatusOK, scriptResponse)
}

func ScriptAsyncView(c *gin.Context) {
	request_data := ScriptBody{}
	c.BindJSON(&request_data)
	scriptResponse := ScriptAsyncService(request_data)
	c.JSON(200, scriptResponse)
}

func FileTransferView(c *gin.Context) {

	c.JSON(200, gin.H{
		"status": "posted",
	})
}

func CurrentConfigView(c *gin.Context) {
	// config := c.MustGet("config")
	config := config.GetConfiguration()
	common.Log.Info(config)
	c.JSON(200, config)
}

func TerminateScriptAsyncView(c *gin.Context) {

}

func CreateCrontabView(c *gin.Context) {

}

func DeleteCrontabView(c *gin.Context) {

}

func ModifyCrontabView(c *gin.Context) {

}

func ListCrontabView(c *gin.Context) {

}
