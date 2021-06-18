package utils

import (
	"fmt"
	"net"
	"os"
)

// 判断目录是否存在
func CheckDir(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}

// 判断目录是否存在， 不存在则创建目录
func CheckDirAndCreate(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return true
		} else {
			return false
		}
	}
}

func CheckFile(file string) bool {
	return true
}

func GetLocalIP4() []string {
	var ipList []string
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	for _, address := range addrList {
		if ip, ok := address.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				ipList = append(ipList, ip.IP.String())
			}

		}
	}
	return ipList
}
