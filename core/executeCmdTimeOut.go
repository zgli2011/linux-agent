package core

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"linux-agent/common"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strconv"
	"syscall"
	"time"
)

func (cmdInfo *CmdInfo) ExecuteCMDTimeOut() (CmdResult, error) {
	var cmdResult CmdResult
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

	// 3、开启脚本执行
	cmdCTX, cancel := context.WithTimeout(context.Background(), time.Second * time.Duration(cmdInfo.ScriptTimeOut))
	defer cancel()

	cmd := exec.CommandContext(cmdCTX, cmdInfo.Interpreter, scriptPath, cmdInfo.ExecuteScriptParam)
	// 切换用户
	if cmdInfo.ExecuteUser != "root" {
		user, err :=user.Lookup(cmdInfo.ExecuteUser)
		if err == nil {
			uid, _ := strconv.Atoi(user.Uid)
			gid, _ := strconv.Atoi(user.Gid)

			cmd.SysProcAttr = &syscall.SysProcAttr{}
			cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
		}
	}
	// 切换脚本执行目录
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



	if cmdCTX.Err() == context.DeadlineExceeded {
		cmdResult.exitCode = -2
		fmt.Println("超时了")
	}else if cmdCTX.Err() == nil{
		cmdResult.exitCode = cmd.ProcessState.ExitCode()
	}

	return cmdResult, nil
}
