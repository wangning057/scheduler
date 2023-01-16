package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/go-redis/redis/v9"
	c "github.com/wangning057/scheduler/executeServiceClient"
	"github.com/wangning057/scheduler/service/task"
	"google.golang.org/grpc"
)

// 初始化一个redis客户端工具，以备使用
var rdb *redis.Client

func init() {
	fmt.Println("init in main.go")
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// 实现 ExecutorServiceServer 接口
type executeServiceServer struct {
	task.UnimplementedExecuteServiceServer
}

/*
	Execute 函数功能：

0.从ninja客户端接收到任务 （不需要了）
1.设置任务在redis中的状态为ready
2.将任务发给executor执行
3.将执行结果发送给ninja客户端
*/
func (e *executeServiceServer) Execute(ctx context.Context, t *task.ExecutionTask) (*task.ExecuteResult, error) {
	// TODO
	// 1.设置任务在redis中的状态为"ready"
	taskId := t.GetTaskId()
	err := rdb.Set(ctx, taskId, "ready", 0).Err()
	if err != nil {
		log.Fatalf("设置任务%+v在redis中的状态为ready失败", taskId)
	}
	// 2.将任务发给executor执行

	// BUG:下面这两个应该是开两个goroutine同时执行，，用channel拿返回值，然后返回非nil的那一个res，而不是一个一个执行
	res1, err1 := c.Client1.Execute(ctx, t)
	if err1 != nil {
		log.Fatalln("res1, err1 := c.Client1.Execute(ctx, t)失败", err1)
	}
	res2, err2 := c.Client2.Execute(ctx, t)
	if err2 != nil {
		log.Fatalln("res2, err2 := c.Client2.Execute(ctx, t)失败", err2)
	}

	//返回res1和res2中非nil的那个
	if res1 != nil {
		return res1, nil
	} else {
		return res2, nil
	}
}

func main() {

	//相对于ninja的gRPC服务端
	server := grpc.NewServer()
	task.RegisterExecuteServiceServer(server, &executeServiceServer{})
	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("从ninja客户端接收到任务的服务监听端口失败", err)
	}
	_ = server.Serve(listener)

	//相对于executor的gRPC客户端

}
