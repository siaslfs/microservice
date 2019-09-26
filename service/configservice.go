package service

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"microservice/dto"
	app "microservice/pb"
	"strconv"
	"strings"
)

const PORT = "8028"

func initGrpc() (configClient app.WaiterClient, err error) {
	c, err := credentials.NewClientTLSFromFile("server.pem", dto.Config.Grpc.ServerName)
	if err != nil {
		log.Fatalf("tlsClient.GetTLSCredentials err: %v", err)
	}
	//grpc的连接地址
	grpcUrl := fmt.Sprintf("%s", dto.Config.Grpc.Host)

	conn, err := grpc.Dial(grpcUrl, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("config.Dial err: %v", err)
	}
	defer conn.Close()
	configClient = app.NewWaiterClient(conn) // 模拟请求数据
	return configClient, err
}

func Get(param app.Params) (*app.Result, error) {
	//初始化grpc
	configClient, _ := initGrpc()
	//传递的key格式为AppId:EnId:KV
	arr := strings.Split(param.Keys, ":")
	if len(arr) != 3 {
		return nil, errors.New("keys is undefined")
	}
	EnId, err := strconv.ParseInt(arr[1], 10, 64)
	if err != nil {
		return nil, errors.New("keys not contains incorrect enid")
	}
	//内置系统参数，环境变量
	evnMap := map[int64]string{
		1: "1",
		2: "2",
		3: "3",
	}
	if _, ok := evnMap[EnId]; ok {
		// 调用gRPC接口
		tr, err := configClient.GetAppConfig(context.Background(), &param)
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("服务端响应: %s", tr.Data)
		return tr, err
	} else {
		return nil, errors.New("enid is null")
	}
}
