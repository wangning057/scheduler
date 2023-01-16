package executeServiceClient

import (
	"log"

	"github.com/wangning057/scheduler/service/task"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	exeIp1 string = "192.168.0.0" //待修改
	exeIp2 string = "192.168.0.0" //待修改
)

var (
	Client1 task.ExecuteServiceClient
	Client2 task.ExecuteServiceClient
)

func init() {
	// 连接executor1
	conn1, err1 := grpc.Dial(exeIp1, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err1 != nil {
		log.Fatalf("did not connect: %v", err1)
	}
	defer conn1.Close()
	Client1 = task.NewExecuteServiceClient(conn1)

	// 连接executor2
	conn2, err2 := grpc.Dial(exeIp2, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err2 != nil {
		log.Fatalf("did not connect: %v", err2)
	}
	defer conn2.Close()
	Client2 = task.NewExecuteServiceClient(conn2)
}