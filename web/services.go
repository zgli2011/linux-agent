package web

import (
	"agent/core/execute_command"
	"time"
)

func ScriptSyncService(scripBody ScriptBody) *ScriptResponse {
	var scriptResponse = &ScriptResponse{
		Code:       -1, // 0:正常;1:脚本落地失败;2:切换目录失败;3:切换用户失败
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
	cmdRequest := &execute_command.CmdRequest{}
	cmdRequest.Interpreter = scripBody.Interpreter
	cmdRequest.Path = scripBody.Path
	cmdRequest.User = scripBody.User
	cmdRequest.Content = scripBody.Content
	cmdRequest.Param = scripBody.Param
	cmdRequest.TimeOut = scripBody.TimeOut
	cmdResult := cmdRequest.ExecuteScriptTimeOut()

	scriptResponse.Code = cmdResult.Code
	scriptResponse.Msg = cmdResult.Msg
	scriptResponse.Pid = cmdResult.Pid
	scriptResponse.ExitCode = cmdResult.ExitCode
	scriptResponse.Stdout = cmdResult.Stdout
	scriptResponse.Stderr = cmdResult.Stderr
	scriptResponse.StartTime = cmdResult.StartTime
	scriptResponse.FinishTime = cmdResult.FinishTime

	return scriptResponse
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

func FileTransferService(file_body FileBody) FileResponse {
	var fileResponse FileResponse
	fileResponse.BCommand = file_body.BCommand
	fileResponse.BCommandUser = file_body.BCommandUser
	fileResponse.ACommand = file_body.ACommand
	fileResponse.ACommand_user = file_body.ACommand_user
	fileResponse.IP = file_body.IP
	fileResponse.TaskID = file_body.TaskID

	// 执行前置动作，如果失败直接返回
	exit_code, stdout, stderr := execute_command.ExecuteShellTimeOut(file_body.ACommand_user, file_body.BCommand, 10)
	fileResponse.BCommandExitCore = exit_code
	fileResponse.BCommandStdout = stdout
	fileResponse.BCommandStderr = stderr
	if exit_code != 0 {
		// 退出
	}
	// 检查原来文件是否存在，如果不存在直接上传，如果文件存在，则看是否需要校验md5值，如果需要则校验md5，否则直接上传

	// 文件传输

	// 执行后置动作
}
