package Queue

import (
	"sync"
)

type MessageQueue struct {
	queueList []interface{}
	lock sync.RWMutex
}

// 初始化队列
func (queue *MessageQueue) New() *MessageQueue {
	return queue
}

// 队列添加数据
func (queue *MessageQueue) Enqueue(msg interface{}) {
	queue.lock.Lock()
	queue.queueList = append(queue.queueList, msg)
	queue.lock.Unlock()
}

// 队列读取数据
func (queue *MessageQueue) Dequeue() *interface{}{
	queue.lock.Lock()
	item := queue.queueList[0]
	queue.queueList = queue.queueList[1:len(queue.queueList)]
	queue.lock.Unlock()
	return &item
}

// 获取队列的第一个元素，不移除
func (queue *MessageQueue) TestRead() interface{} {
	queue.lock.Lock()
	item := queue.queueList[0]
	queue.lock.Unlock()
	return &item
}

// 判空
func (queue *MessageQueue) IsEmpty() bool {
	return len(queue.queueList) == 0
}

// 获取队列的长度
func (queue *MessageQueue) Size() int {
	return len(queue.queueList)
}
