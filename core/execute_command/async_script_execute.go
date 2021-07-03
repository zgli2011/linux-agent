package execute_command

import (
	"fmt"
	"time"
)

type CMDInfo struct {
	Interpreter string
	Path        string
	User        string
	Content     string
	Param       string
	TimeOut     int64
}

type DataContainer struct {
	Queue chan interface{}
}

var async_queue *DataContainer

func NewDataContainer(max_queue_len int) {
	async_queue.Queue = make(chan interface{}, max_queue_len)
}

func GetQueue() *DataContainer {
	return async_queue
}

//非阻塞push
func (dc *DataContainer) Push(data interface{}) bool {
	click := time.After(10 * time.Millisecond)
	select {
	case dc.Queue <- data:
		return true
	case <-click:
		return false
	}
}

//非阻塞pop
func (dc *DataContainer) Pop() (data interface{}) {
	click := time.After(10 * time.Millisecond)
	select {
	case data = <-dc.Queue:
		return data
	case <-click:
		return nil
	}
}

func ScriptExecuteQueue() {
	for {
		select {
		case data := <-GetQueue().Queue:
			async_script_execute_queue(data)
		}
	}
}

func async_script_execute_queue(data interface{}) {
	fmt.Println("异步任务执行:")
	fmt.Println(data)
}
