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
	"path"
	"strings"
	"time"
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

	if cmdInfo.ExecuteUser == "" {
		cmdInfo.ExecuteUser = "root"
	}

	name, args := "", ""
	if cmdInfo.ExecuteUser != "root" {
		argsList := []string{"su -", cmdInfo.ExecuteUser, "-c", "'", "cd", cmdInfo.ExecutePath, "&&", cmdInfo.Interpreter}
		args = strings.Join(argsList, " ")
	}else{
		argsList := []string{"cd", cmdInfo.ExecutePath, "&&", cmdInfo.Interpreter}
		args = strings.Join(argsList, " ")
	}

	arg := []string{args, scriptPath, cmdInfo.ExecuteScriptParam, "'"}
	fmt.Println(name, arg)
	cmd := exec.CommandContext(cmdCTX, "sh", arg...)
	fmt.Println(cmd.String())
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		fmt.Println(stdout.String())
		fmt.Println(stderr.String())
		return cmdResult, err
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println(stdout.String())
		fmt.Println(stderr.String())
		return cmdResult, err
	}
	fmt.Println(string(stdout.Bytes()))
	fmt.Println(string(stderr.Bytes()))
	fmt.Println(cmd.ProcessState.ExitCode())


	cmdResult.exitCode = cmd.ProcessState.ExitCode()
	cmdResult.stdout = stdout.String()
	cmdResult.stderr = stderr.String()

	return cmdResult, nil
}
