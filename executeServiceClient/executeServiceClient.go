package executeServiceClient

import (
	"log"

	"github.com/wangning057/scheduler/service/execute"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	exeIp1 string = "172.17.0.45:8003"   //待修改
	exeIp2 string = "172.17.0.21:8003" //待修改
)

var (
	Client1 execute.ExecuteServiceClient
	Client2 execute.ExecuteServiceClient
)

func init() {
	// 连接executor1
	conn1, err1 := grpc.Dial(exeIp1, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err1 != nil {
		log.Fatalf("did not connect: %v", err1)
	}
	// defer conn1.Close()
	Client1 = execute.NewExecuteServiceClient(conn1)

	// 连接executor2
	conn2, err2 := grpc.Dial(exeIp2, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err2 != nil {
		log.Fatalf("did not connect: %v", err2)
	}
	// defer conn2.Close()
	Client2 = execute.NewExecuteServiceClient(conn2)
}
