package core

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
