package utils

import (
	"fmt"
	"net"
	"os"
)

func CheckFile(path string) bool {
	return CheckDirOrFileExist(path)
}

func CheckDir(path string) bool {
	return CheckDirOrFileExist(path)
}

// 判断目录/文件是否存在
func CheckDirOrFileExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return os.IsExist(err)
	}
	return true
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
