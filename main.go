package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/go-redis/redis/v9"
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
type executorServiceServer struct {
	task.UnimplementedExecutorServiceServer
}

/*
	Execute 函数功能：

0.从ninja客户端接收到任务 （不需要了）
1.设置任务在redis中的状态为ready
2.将任务发给executor执行
3.将执行结果发送给ninja客户端
*/
func (e *executorServiceServer) Execute(ctx context.Context, in *task.ExecutionTask) (*task.ExecuteResult, error) {
	// TODO
	// 1.设置任务在redis中的状态为"ready"
	action_id := in.GetActionId()
	err := rdb.Set(ctx, action_id, "ready", 0).Err()
	if err != nil {
		log.Fatalf("设置任务%+v在redis中的状态为ready失败", action_id)
	}
	// 2.将任务发给executor执行
	// TODO:需要先写与executor的连接部分的客户端代码

}

func main() {

	//相对于ninja的gRPC服务端
	server := grpc.NewServer()
	task.RegisterExecutorServiceServer(server, &executorServiceServer{})
	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("从ninja客户端接收到任务的服务监听端口失败", err)
	}
	_ = server.Serve(listener)

	//相对于executor的gRPC客户端

}
