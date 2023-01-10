package app

import (
	"context"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hk4e/common/config"
	"hk4e/common/mq"
	"hk4e/common/rpc"
	"hk4e/gate/net"
	"hk4e/node/api"
	"hk4e/pkg/logger"
)

var APPID string

func Run(ctx context.Context, configFile string) error {
	config.InitConfig(configFile)

	// natsrpc client
	client, err := rpc.NewClient()
	if err != nil {
		return err
	}

	// 注册到节点服务器
	rsp, err := client.Discovery.RegisterServer(context.TODO(), &api.RegisterServerReq{
		ServerType: api.GATE,
		GateServerAddr: &api.GateServerAddr{
			KcpAddr: config.CONF.Hk4e.KcpAddr,
			KcpPort: uint32(config.CONF.Hk4e.KcpPort),
			MqAddr:  config.CONF.Hk4e.GateTcpMqAddr,
			MqPort:  uint32(config.CONF.Hk4e.GateTcpMqPort),
		},
		Version: config.CONF.Hk4e.Version,
	})
	if err != nil {
		return err
	}
	APPID = rsp.GetAppId()
	go func() {
		ticker := time.NewTicker(time.Second * 15)
		for {
			<-ticker.C
			_, err := client.Discovery.KeepaliveServer(context.TODO(), &api.KeepaliveServerReq{
				ServerType: api.GATE,
				AppId:      APPID,
			})
			if err != nil {
				logger.Error("keepalive error: %v", err)
			}
		}
	}()
	defer func() {
		_, _ = client.Discovery.CancelServer(context.TODO(), &api.CancelServerReq{
			ServerType: api.GATE,
			AppId:      APPID,
		})
	}()

	logger.InitLogger("gate_" + APPID)
	logger.Warn("gate start, appid: %v", APPID)

	messageQueue := mq.NewMessageQueue(api.GATE, APPID, client)
	defer messageQueue.Close()

	connectManager := net.NewKcpConnectManager(messageQueue, client.Discovery)
	defer connectManager.Close()

	go func() {
		outputChan := connectManager.GetKcpEventOutputChan()
		for {
			<-outputChan
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case <-ctx.Done():
			return nil
		case s := <-c:
			logger.Warn("get a signal %s", s.String())
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				logger.Warn("gate exit, appid: %v", APPID)
				time.Sleep(time.Second)
				return nil
			case syscall.SIGHUP:
			default:
				return nil
			}
		}

	}
}
