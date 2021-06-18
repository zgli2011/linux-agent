package ExecuteCommand

// type CmdResultList struct {
// 	CmdResultList map[string]*ScriptSyncResponse
// 	lock          sync.RWMutex
// }

// func (cmdResultList *CmdResultList) GetResult(taskId string) *ScriptSyncResponse {
// 	return cmdResultList.CmdResultList[taskId]
// }

// func (cmdResultList *CmdResultList) RecordResult(taskId string, taskResult *ScriptSyncResponse) bool {
// 	cmdResultList.lock.Lock()
// 	cmdResultList.CmdResultList[taskId] = taskResult
// 	cmdResultList.lock.Unlock()
// 	return true
// }

// func (cmdResultList *CmdResultList) DeleteResult(taskId string) bool {
// 	cmdResultList.lock.Lock()
// 	delete(cmdResultList.CmdResultList, taskId)
// 	cmdResultList.lock.Unlock()
// 	return true
// }
