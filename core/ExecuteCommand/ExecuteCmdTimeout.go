package ExecuteCommand

import (
	"agent/common"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strconv"
	"syscall"
	"time"
)

type CmdRequest struct {
	Interpreter string
	Path        string
	User        string
	Content     string
	Param       string
	TimeOut     int64
}

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

func (cmdRequest *CmdRequest) ExecuteCMDTimeOut() CmdResult {
	var cmdResult CmdResult
	cmdResult.Code = 0
	cmdResult.Msg = ""

	// 1、检查脚本执行路径
	if res := common.CheckDirAndCreate(cmdRequest.Path); !res {
		cmdResult.Code = 1
		cmdResult.Msg = "要执行的目录不存在"
		return cmdResult
	}

	// 2、将脚本内容写入文件
	scriptPath := path.Join("/tmp/", common.RandString(9))
	defer os.RemoveAll(scriptPath)

	if err := ioutil.WriteFile(scriptPath, []byte(cmdRequest.Content), 0755); err != nil {
		cmdResult.Code = 2
		cmdResult.Msg = "脚本落地失败"
		return cmdResult
	}

	// 3、开启脚本执行
	cmd := exec.Command(cmdRequest.Interpreter, "-c", scriptPath, cmdRequest.Param)
	// 切换用户
	if cmdRequest.User != "root" {
		user, err := user.Lookup(cmdRequest.User)
		if err == nil {
			uid, _ := strconv.Atoi(user.Uid)
			gid, _ := strconv.Atoi(user.Gid)

			cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
			cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
		}
	}

	// 切换脚本执行目录
	cmd.Dir = cmdRequest.Path
	cmdResult.StartTime = time.Now().Format("2006-01-02 15:04:05")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	time.AfterFunc(time.Duration(cmdRequest.TimeOut)*time.Second, func() { syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL) })
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
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	cmdResult.Pid = cmd.ProcessState.Pid()
	cmdResult.Stdout = stdout.String()
	cmdResult.Stderr = stderr.String()
	cmdResult.FinishTime = time.Now().Format("2006-01-02 15:04:05")
	return cmdResult
}
