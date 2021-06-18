package web

type ScriptBody struct {
	Interpreter string `json:"interpreter" binding:"required"`
	Path        string `json:"path" binding:"required"`
	User        string `json:"user" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Param       string `json:"param"`
	TimeOut     int64  `json:"timeout" binding:"required, gt=0, dive"`
	IP          string `json:"ip" binding:"required"`
	TaskID      string `json:"task_id" binding:"required"`
}

type ScriptSyncResponse struct {
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
