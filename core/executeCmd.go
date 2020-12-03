package core

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"linux-agent/common"
	"os"
	"os/exec"
	"path"
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

	args := ""
	if cmdInfo.ExecuteUser != "root" {
		args = "su - " + cmdInfo.ExecuteUser
	}

	arg := []string{args, "-c", "cd", cmdInfo.ExecutePath, "&&", cmdInfo.Interpreter, scriptPath, cmdInfo.ExecuteScriptParam}
	cmd := exec.CommandContext(cmdCTX, "sh", arg...)
	fmt.Println(cmd.String())
	output, _ := cmd.CombinedOutput()
	fmt.Println(output)
	//stdin, _ := cmd.StdinPipe()
	//defer stdin.Close()
	//
	//stdout, _ := cmd.StdoutPipe()
	//defer stdout.Close()
	//
	//stderr, _ := cmd.StderrPipe()
	//defer stderr.Close()
	//cmd.Start()
	//go func() {
	//	s := bufio.NewScanner(stdout)
	//	for s.Scan() {
	//		fmt.Println("程序输出:" + s.Text())
	//	}
	//}()




	//if err := cmd.Run(); err != nil {
	//	fmt.Println(stdout.String())
	//	return cmdResult, err
	//}
	//fmt.Println("---")
	//fmt.Println(stdout.String())
	//fmt.Println("---")
	//fmt.Println(stderr.String())
	//fmt.Println("---")
	//fmt.Println(cmd.ProcessState.ExitCode())
	//
	//
	//cmdResult.exitCode = cmd.ProcessState.ExitCode()
	//cmdResult.stdout = stdout.String()
	//cmdResult.stderr = stderr.String()

	return cmdResult, nil
}
