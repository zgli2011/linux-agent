package web

type ScriptBody struct {
	Interpreter string `json:"interpreter" binding:"required"`
	Path        string `json:"path" binding:"required"`
	User        string `json:"user" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Param       string `json:"param"`
	TimeOut     int64  `json:"timeout" validate:"required||integer=1,300"`
	IP          string `json:"ip" binding:"required"`
	TaskID      string `json:"task_id" binding:"required"`
}

type ScriptResponse struct {
	Pid        int    `json:"pid"`
	ExitCode   int    `json:"exit_code"`
	Stdout     string `json:"stdout"`
	Stderr     string `json:"stderr"`
	IP         string `json:"ip"`
	TaskID     string `json:"task_id"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
	StartTime  string `json:"start_time"`
	FinishTime string `json:"finish_time"`
}

type FileBody struct {
	BCommand      string `json:"bcommand"`
	BCommandUser  string `json:"bcommand_user"`
	ACommand      string `json:"acommand"`
	ACommand_user string `json:"acommand_user"`
	IP            string `json:"ip" binding:"required"`
	TaskID        int64  `json:"task_id" binding:"required"`
	FileList      []*struct {
		FileID   int    `json:"file_id" binding:"required"`
		Path     string `json:"path" binding:"required"`
		User     string `json:"user" binding:"required"`
		Group    string `json:"group"`
		Content  string `json:"content" binding:"required"`
		Name     string `json:"name" binding:"required"`
		MD5Check bool   `json:"md5_check" binding:"required"`
		MD5      string `json:"md5"`
	} `json:"file_list" binding:"required"`
}

type FileResponse struct {
	BCommand         string `json:"bcommand"`
	BCommandUser     string `json:"bcommand_user"`
	BCommandExitCore int    `json:"bcommand_exist_code"`
	BCommandStdout   string `json:"bcommand_stdout"`
	BCommandStderr   string `json:"bcommand_stderr"`
	ACommand         string `json:"acommand"`
	ACommand_user    string `json:"acommand_user"`
	ACommandExitCore int    `json:"acommand_exist_code"`
	ACommandStdout   string `json:"acommand_stdout"`
	ACommandStderr   string `json:"acommand_stderr"`
	IP               string `json:"ip"`
	TaskID           int64  `json:"task_id"`
	FileList         []File
}

type File struct {
	FileID  int    `json:"file_id"`
	Path    string `json:"path"`
	User    string `json:"user"`
	Group   string `json:"group"`
	Content string `json:"content"`
	Name    string `json:"name"`
	Result  bool   `json:"result"`
	Msg     string `json:"msg"`
}
