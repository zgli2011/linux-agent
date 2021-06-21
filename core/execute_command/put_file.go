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
		n, _ := f.Seek(0, os.SEEK_END)
		if _, err := f.WriteAt([]byte(content), n); err == nil {
			if fileInfo, err := f.Stat(); err == nil {
				stat := fileInfo.Sys().(*syscall.Stat_t)
				stat.
				if user, err := user.Lookup(users); err == nil {
					if uid, err := strconv.Atoi(user.Uid); err == nil {
						if gid, err := strconv.Atoi(user.Gid); err == nil {
							if stat.Uid == uid && stat.Gid == gid {
								return os.Chown(file, uid, gid)
							}
						}
					}
				}
			}
		}
		return err
	}
}
