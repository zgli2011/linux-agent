package core

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"linux-agent/common"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strconv"
	"syscall"
)

type CmdInfo struct {
	Interpreter        string
	ExecuteUser        string
	ExecutePath        string
	ExecuteScript      string
	ExecuteScriptParam string
	ScriptTimeOut      int64
}

type CmdResult struct {
	exitCode int
	stdout string
	stderr string
}

func (cmdInfo *CmdInfo) ExecuteCMD() (CmdResult, error) {
	var cmdResult CmdResult
	cmdResult.exitCode = -1
	// 1、检查脚本执行路径
	if common.CheckDir(cmdInfo.ExecutePath) == false{
		return cmdResult, errors.New("要执行的目录不存在")
	}

	// 2、将脚本内容写入文件
	scriptPath := path.Join("/tmp/", common.RandString(9))
	defer os.RemoveAll(scriptPath)

	if err := ioutil.WriteFile(scriptPath, []byte(cmdInfo.ExecuteScript), 0755); err!=nil{
		return cmdResult, err
	}

	cmd := exec.Command(cmdInfo.Interpreter, scriptPath, cmdInfo.ExecuteScriptParam)

	// 切换用户
	if cmdInfo.ExecuteUser != "root" {
		user, err :=user.Lookup(cmdInfo.ExecuteUser)
		if err == nil {
			log.Printf("uid=%s,gid=%s", user.Uid, user.Gid)
			uid, _ := strconv.Atoi(user.Uid)
			gid, _ := strconv.Atoi(user.Gid)

			cmd.SysProcAttr = &syscall.SysProcAttr{}
			cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
		}
	}

	cmd.Dir = cmdInfo.ExecutePath

	fmt.Println(cmd.String())
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}
	cmdResult.stdout = string(buf.Bytes())

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
	}
	cmdResult.stderr = string(buf.Bytes())

	cmdResult.exitCode = cmd.ProcessState.ExitCode()
	return cmdResult, nil
}
