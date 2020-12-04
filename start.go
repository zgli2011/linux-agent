package main

import (
	"fmt"
	"linux-agent/core"
)
func main() {
	//orig := "http://c.biancheng.net/golang/"
	//key := "12345678123456781234567812345678"
	//fmt.Println("原文：", orig)
	//encryptCode := cryption.AesEncrypt(orig, key)
	//fmt.Println("密文：", encryptCode)
	//decryptCode := cryption.AesDecrypt(encryptCode, key)
	//fmt.Println("解密结果：", decryptCode)


	cmd := &core.CmdInfo{}
	cmd.Interpreter = "/bin/bash"
	cmd.ExecuteUser = "asher"
	cmd.ExecutePath = "/tmp"
	cmd.ExecuteScript = "ifconfig && echo $1 $2 && whoami"
	cmd.ExecuteScriptParam = "p1 p2"
	cmd.ScriptTimeOut = 10
	fmt.Println(cmd.ExecuteCMD())
}