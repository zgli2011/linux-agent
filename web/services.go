package web

import (
	"agent/core/execute_command"
	"agent/utils"
	"path"
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
	// 初始化命令执行所需的所有信息
	cmd := &execute_command.ExecuteCMD{}
	cmd.Interpreter = scripBody.Interpreter
	cmd.Path = scripBody.Path
	cmd.User = scripBody.User
	cmd.Content = scripBody.Content
	cmd.Param = scripBody.Param
	cmd.TimeOut = scripBody.TimeOut
	// 执行脚本
	cmdResult := cmd.ExecuteScriptTimeOut()
	// 将命令执行结果转换成http返回的消息体
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
	cmd_info := &execute_command.CMDInfo{
		Interpreter: scripBody.Interpreter,
		Path:        scripBody.Path,
		User:        scripBody.User,
		Content:     scripBody.Content,
		Param:       scripBody.Param,
		TimeOut:     scripBody.TimeOut,
	}
	// 将任务塞到任务队列
	res := execute_command.GetQueue().Push(cmd_info)
	code := 0
	msg := ""
	if !res {
		code = 1
		msg = "任务加入执行队列失败"
	}
	return ScriptAsyncResponse{
		Code: code,
		IP:   scripBody.IP,
		Msg:  msg,
	}
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
	exit_code, stdout, stderr := execute_command.ExecuteShellTimeOut(file_body.BCommandUser, file_body.BCommand, 10)
	fileResponse.BCommandExitCore = exit_code
	fileResponse.BCommandStdout = stdout
	fileResponse.BCommandStderr = stderr
	if exit_code != 0 {
		// 退出
		for _, file := range file_body.FileList {
			var f File
			f.FileID = file.FileID
			f.Path = file.Path
			f.User = file.User
			f.Group = file.Group
			f.Content = file.Content
			f.Name = file.Name
			f.Result = false
			f.Msg = "undo"
			fileResponse.FileList = append(fileResponse.FileList, f)
		}
		fileResponse.ACommandExitCore = -1
		fileResponse.ACommandStdout = ""
		fileResponse.ACommandStderr = ""
		return fileResponse
	}

	// 文件传输
	flag := true
	for _, file := range file_body.FileList {
		var f File
		f.FileID = file.FileID
		f.Path = file.Path
		f.User = file.User
		f.Group = file.Group
		f.Content = file.Content
		f.Name = file.Name
		file_path := path.Join(file.Path, file.Name)
		if !flag { // 如果有文件放失败了，则不再执行下面的动作，只打标记
			f.Result = false
			f.Msg = "undo"
			fileResponse.FileList = append(fileResponse.FileList, f)
			continue
		}
		if file.MD5Check && utils.MD5Value(file.Content) != file.MD5 { // 如果开启了文件校验，但是校验不通过则直接退出
			f.Result = false
			f.Msg = "md5校验失败"
			flag = false
		} else { // 存放文件
			if err := execute_command.PutFile(file.User, file.Group, file_path, file.Content); err != nil {
				f.Result = false
				f.Msg = "文件上传失败:" + err.Error()
				flag = false
			} else {
				f.Result = true
				f.Msg = ""
			}
		}
		fileResponse.FileList = append(fileResponse.FileList, f)
	}

	// 执行后置动作
	exit_code, stdout, stderr = execute_command.ExecuteShellTimeOut(file_body.ACommand_user, file_body.ACommand, 10)
	fileResponse.ACommandExitCore = exit_code
	fileResponse.ACommandStdout = stdout
	fileResponse.ACommandStderr = stderr
	return fileResponse
}
