package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	sc := []constant.ServerConfig{
		{
			IpAddr: "172.30.0.113",
			Port:   8849,
		},
	}

	cc := constant.ClientConfig{
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogRollingConfig:    &lumberjack.Logger{MaxSize: 10},
		LogLevel:            "debug",
		AppendToStdout:      true,
	}

	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	clientList := []naming_client.INamingClient{}
	for i := 0; i < 3; i++ {
		tmp, err := clients.NewNamingClient(
			vo.NacosClientParam{
				ClientConfig:  &cc,
				ServerConfigs: sc,
			},
		)
		if err != nil {
			panic(err)
		}
		clientList = append(clientList, tmp)
	}
	for _, namingClient := range clientList {
		info, err := namingClient.GetAllServicesInfo(vo.GetAllServiceInfoParam{})
		fmt.Printf("%v\n", info)
		if err != nil {
			panic(err)
		}
	}
	info, err := client.GetAllServicesInfo(vo.GetAllServiceInfoParam{})
	fmt.Printf("%v\n", info)
	if err != nil {
		panic(err)
	}
	client.CloseClient()
	_ = http.ListenAndServe("0.0.0.0:7567", nil)
}
