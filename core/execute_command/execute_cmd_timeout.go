package execute_command

import (
	"agent/log"
	"agent/utils"

	"bytes"
	"io/ioutil"

	"os"
	"os/exec"
	"os/user"
	"path"
	"strconv"
	"syscall"
	"time"
)

type CmdResult struct {
	Pid        int
	Code       int
	ExitCode   int
	Stdout     string
	Stderr     string
	StartTime  string
	FinishTime string
	Msg        string
}

type ExecuteCMD struct {
	Interpreter string
	Path        string
	User        string
	Content     string
	Param       string
	TimeOut     int64
}

func (executeCMD *ExecuteCMD) ExecuteScriptTimeOut() CmdResult {
	var cmdResult CmdResult
	cmdResult.Code = 0
	cmdResult.Msg = ""

	// 1、检查脚本执行路径
	if res := utils.CheckDirAndCreate(executeCMD.Path); !res {
		cmdResult.Code = 1
		cmdResult.Msg = "要执行的目录不存在"
		return cmdResult
	}

	// 2、将脚本内容写入文件
	scriptPath := path.Join("/tmp/", utils.RandString(9))
	defer os.RemoveAll(scriptPath)

	if err := ioutil.WriteFile(scriptPath, []byte(executeCMD.Content), 0755); err != nil {
		cmdResult.Code = 2
		cmdResult.Msg = "脚本落地失败"
		return cmdResult
	}

	// 3、开启脚本执行
	cmd := exec.Command(executeCMD.Interpreter, "-c", scriptPath, executeCMD.Param)
	// 切换用户
	if executeCMD.User != "root" {
		user, err := user.Lookup(executeCMD.User)
		if err == nil {
			uid, _ := strconv.Atoi(user.Uid)
			gid, _ := strconv.Atoi(user.Gid)

			cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
			cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
		}
	}

	// 切换脚本执行目录
	cmd.Dir = executeCMD.Path
	cmdResult.StartTime = time.Now().Format("2006-01-02 15:04:05")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	log.Log.Info(time.Now().Format("2006-01-02 15:04:05"))
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	time.AfterFunc(time.Duration(executeCMD.TimeOut)*time.Second, func() { syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL) })
	if err := cmd.Run(); err != nil {
		if err.Error() == "signal: killed" {
			cmdResult.Code = 3
			cmdResult.Msg = "脚本执行超时"
			cmdResult.ExitCode = 130
		} else {
			cmdResult.ExitCode = cmd.ProcessState.ExitCode()
		}
	} else {
		cmdResult.ExitCode = cmd.ProcessState.ExitCode()
	}
	log.Log.Info(time.Now().Format("2006-01-02 15:04:05"))
	cmdResult.Pid = cmd.ProcessState.Pid()
	cmdResult.Stdout = stdout.String()
	cmdResult.Stderr = stderr.String()
	cmdResult.FinishTime = time.Now().Format("2006-01-02 15:04:05")
	return cmdResult
}

func ExecuteShellTimeOut(cmd_user string, command string, timeout int) (int, string, string) {
	// 3、开启脚本执行
	cmd := exec.Command("/bin/bash", "-c", command)
	// 切换用户
	if cmd_user != "root" {
		user, err := user.Lookup(cmd_user)
		if err == nil {
			uid, _ := strconv.Atoi(user.Uid)
			gid, _ := strconv.Atoi(user.Gid)

			cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
			cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
		}
	}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	time.AfterFunc(time.Duration(timeout)*time.Second, func() { syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL) })
	exit_code := -1
	if err := cmd.Run(); err != nil {
		if err.Error() == "signal: killed" {
			exit_code = 130
		} else {
			exit_code = cmd.ProcessState.ExitCode()
		}
	} else {
		exit_code = cmd.ProcessState.ExitCode()
	}
	rstdout := stdout.String()
	rstderr := stderr.String()
	return exit_code, rstdout, rstderr
}
