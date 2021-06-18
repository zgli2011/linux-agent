package web

import (
	"agent/core/ExecuteCommand"
	"time"
)

func ScriptSyncService(scripBody ScriptBody) *ScriptSyncResponse {
	var scriptSyncResponse = &ScriptSyncResponse{
		Code:       -1, // 0:正常;1:创建脚本失败;2:切换目录失败;3:切换用户失败
		ExitCode:   -1, // -1:未执行; >=0表示脚本的返回码
		Stdout:     "",
		Stderr:     "",
		IP:         scripBody.IP,
		TaskID:     scripBody.TaskID,
		Pid:        0,
		Msg:        "",
		StartTime:  time.Now().Format("2006-01-02 15:04:05"),
		FinishTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	cmdRequest := &ExecuteCommand.CmdRequest{}
	cmdRequest.Interpreter = scripBody.Interpreter
	cmdRequest.Path = scripBody.Path
	cmdRequest.User = scripBody.User
	cmdRequest.Content = scripBody.Content
	cmdRequest.Param = scripBody.Param
	cmdRequest.TimeOut = scripBody.TimeOut
	cmdResult := cmdRequest.ExecuteCMDTimeOut()

	scriptSyncResponse.Code = cmdResult.Code
	scriptSyncResponse.Msg = cmdResult.Msg
	scriptSyncResponse.Pid = cmdResult.Pid
	scriptSyncResponse.ExitCode = cmdResult.ExitCode
	scriptSyncResponse.Stdout = cmdResult.Stdout
	scriptSyncResponse.Stderr = cmdResult.Stderr
	scriptSyncResponse.StartTime = cmdResult.StartTime
	scriptSyncResponse.FinishTime = cmdResult.FinishTime

	return scriptSyncResponse
}

type ScriptAsyncResponse struct {
	IP   string `json:"ip"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func ScriptAsyncService(scripBody ScriptBody) ScriptAsyncResponse {
	var scriptAsyncResponse ScriptAsyncResponse
	// 将任务塞到任务队列

	scriptAsyncResponse.Code = 0
	scriptAsyncResponse.IP = scripBody.IP
	scriptAsyncResponse.Msg = ""
	return scriptAsyncResponse
}
