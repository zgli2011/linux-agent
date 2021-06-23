package execute_command

import (
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func PutFile(users string, group string, file string, content string) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	} else {
		defer f.Close()
		// 写入文件
		n, _ := f.Seek(0, os.SEEK_END)
		_, err := f.WriteAt([]byte(content), n)
		if err != nil {
			return err
		}
		// 获取文件信息
		fileInfo, err := f.Stat()
		if err != nil {
			return err
		}
		// 找到用户的信息
		stat := fileInfo.Sys().(*syscall.Stat_t)
		user, err := user.Lookup(users)
		if err != nil {
			return err
		}
		// 找到用户id
		uid, err := strconv.Atoi(user.Uid)
		if err != nil {
			return err
		}
		// 找到用户组ID
		gid, err := strconv.Atoi(user.Gid)
		if err != nil {
			return err
		}
		//判断前后uid和gid是否一致， 如果不一致则修改为文件的属主信息
		if stat.Uid == uint32(uid) && stat.Gid == uint32(gid) {
			return os.Chown(file, uid, gid)
		}
		return nil
	}
}
