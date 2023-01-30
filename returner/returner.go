package returner

import (
	"errors"
	"log"

	"github.com/wangning057/scheduler/service/task"
)

// 每一个task都有自己的一个channel
var resultMap = make(map[string]chan *task.ExecuteResult, 1024)

func InitChan(task_id string) chan *task.ExecuteResult {
	ch := make(chan *task.ExecuteResult, 1)
	resultMap[task_id] = ch
	return ch
}

func SetRes(task_id string, res *task.ExecuteResult) error {
	if ch, ok := resultMap[task_id]; ok {
		ch <- res
		return nil
	} else {
		return errors.New("找不到task_id:" + task_id + "对应的channel")
	}
}

// func GetRes(task_id string) *task.ExecuteResult {
// 	resChan := resultMap[task_id]
// 	res := <-resChan
// 	delete(resultMap, task_id)
// 	close(resChan)
// 	return res
// }

func GetRes(task_id string) *task.ExecuteResult {
	resChan := resultMap[task_id]
	if resChan == nil {
		log.Printf("resChan == nil，无法取到任务id=%v的res\n", task_id)
	}

	for {
		select {
		case res := <-resChan:
			delete(resultMap, task_id)
			close(resChan)
			return res
		}
	}
}
