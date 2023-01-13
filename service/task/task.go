package task

import (
	context "context"
)

var ExecutorService = &executorService{}

// 实现 ExecutorServiceServer 接口
type executorService struct {
}

/* Execute 函数功能：
1.从ninja客户端接收到任务
2.改变任务在redis中的状态为"ready"
3.将任务发给executor
4.接收executor的执行结果
5.将执行结果发送给ninja客户端 
*/
func (e *executorService) Execute(context.Context, *Command) (*ExecuteResult, error) {
	// TODO
}
func (e *executorService) mustEmbedUnimplementedExecutorServiceServer() {}
