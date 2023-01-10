// Code generated by protoc-gen-natsrpc. DO NOT EDIT.
// versions:
// - protoc-gen-natsrpc v0.5.0
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	natsrpc "github.com/byebyebruce/natsrpc"
	nats_go "github.com/nats-io/nats.go"
	proto "google.golang.org/protobuf/proto"
)

var _ = new(context.Context)
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = natsrpc.Version
var _ = nats_go.Version

// 节点服务器注册发现服务
type DiscoveryNATSRPCServer interface {
	// 服务器启动注册获取appid
	RegisterServer(ctx context.Context, req *RegisterServerReq) (*RegisterServerRsp, error)
	// 服务器关闭取消注册
	CancelServer(ctx context.Context, req *CancelServerReq) (*NullMsg, error)
	// 服务器在线心跳保持
	KeepaliveServer(ctx context.Context, req *KeepaliveServerReq) (*NullMsg, error)
	// 获取负载最小的服务器的appid
	GetServerAppId(ctx context.Context, req *GetServerAppIdReq) (*GetServerAppIdRsp, error)
	// 获取区服密钥信息
	GetRegionEc2B(ctx context.Context, req *NullMsg) (*RegionEc2B, error)
	// 获取负载最小的网关服务器的地址和端口
	GetGateServerAddr(ctx context.Context, req *GetGateServerAddrReq) (*GateServerAddr, error)
	// 获取全部网关服务器信息列表
	GetAllGateServerInfoList(ctx context.Context, req *NullMsg) (*GateServerInfoList, error)
	// 获取主游戏服务器的appid
	GetMainGameServerAppId(ctx context.Context, req *NullMsg) (*GetMainGameServerAppIdRsp, error)
}

// RegisterDiscoveryNATSRPCServer register Discovery service
func RegisterDiscoveryNATSRPCServer(server *natsrpc.Server, s DiscoveryNATSRPCServer, opts ...natsrpc.ServiceOption) (natsrpc.IService, error) {
	return server.Register("hk4e.node.api.Discovery", s, opts...)
}

// 节点服务器注册发现服务
type DiscoveryNATSRPCClient interface {
	// 服务器启动注册获取appid
	RegisterServer(ctx context.Context, req *RegisterServerReq, opt ...natsrpc.CallOption) (*RegisterServerRsp, error)
	// 服务器关闭取消注册
	CancelServer(ctx context.Context, req *CancelServerReq, opt ...natsrpc.CallOption) (*NullMsg, error)
	// 服务器在线心跳保持
	KeepaliveServer(ctx context.Context, req *KeepaliveServerReq, opt ...natsrpc.CallOption) (*NullMsg, error)
	// 获取负载最小的服务器的appid
	GetServerAppId(ctx context.Context, req *GetServerAppIdReq, opt ...natsrpc.CallOption) (*GetServerAppIdRsp, error)
	// 获取区服密钥信息
	GetRegionEc2B(ctx context.Context, req *NullMsg, opt ...natsrpc.CallOption) (*RegionEc2B, error)
	// 获取负载最小的网关服务器的地址和端口
	GetGateServerAddr(ctx context.Context, req *GetGateServerAddrReq, opt ...natsrpc.CallOption) (*GateServerAddr, error)
	// 获取全部网关服务器信息列表
	GetAllGateServerInfoList(ctx context.Context, req *NullMsg, opt ...natsrpc.CallOption) (*GateServerInfoList, error)
	// 获取主游戏服务器的appid
	GetMainGameServerAppId(ctx context.Context, req *NullMsg, opt ...natsrpc.CallOption) (*GetMainGameServerAppIdRsp, error)
}

type _DiscoveryNATSRPCClient struct {
	c *natsrpc.Client
}

// NewDiscoveryNATSRPCClient
func NewDiscoveryNATSRPCClient(enc *nats_go.EncodedConn, opts ...natsrpc.ClientOption) (DiscoveryNATSRPCClient, error) {
	c, err := natsrpc.NewClient(enc, "hk4e.node.api.Discovery", opts...)
	if err != nil {
		return nil, err
	}
	ret := &_DiscoveryNATSRPCClient{
		c: c,
	}
	return ret, nil
}
func (c *_DiscoveryNATSRPCClient) RegisterServer(ctx context.Context, req *RegisterServerReq, opt ...natsrpc.CallOption) (*RegisterServerRsp, error) {
	rep := &RegisterServerRsp{}
	err := c.c.Request(ctx, "RegisterServer", req, rep, opt...)
	return rep, err
}
func (c *_DiscoveryNATSRPCClient) CancelServer(ctx context.Context, req *CancelServerReq, opt ...natsrpc.CallOption) (*NullMsg, error) {
	rep := &NullMsg{}
	err := c.c.Request(ctx, "CancelServer", req, rep, opt...)
	return rep, err
}
func (c *_DiscoveryNATSRPCClient) KeepaliveServer(ctx context.Context, req *KeepaliveServerReq, opt ...natsrpc.CallOption) (*NullMsg, error) {
	rep := &NullMsg{}
	err := c.c.Request(ctx, "KeepaliveServer", req, rep, opt...)
	return rep, err
}
func (c *_DiscoveryNATSRPCClient) GetServerAppId(ctx context.Context, req *GetServerAppIdReq, opt ...natsrpc.CallOption) (*GetServerAppIdRsp, error) {
	rep := &GetServerAppIdRsp{}
	err := c.c.Request(ctx, "GetServerAppId", req, rep, opt...)
	return rep, err
}
func (c *_DiscoveryNATSRPCClient) GetRegionEc2B(ctx context.Context, req *NullMsg, opt ...natsrpc.CallOption) (*RegionEc2B, error) {
	rep := &RegionEc2B{}
	err := c.c.Request(ctx, "GetRegionEc2B", req, rep, opt...)
	return rep, err
}
func (c *_DiscoveryNATSRPCClient) GetGateServerAddr(ctx context.Context, req *GetGateServerAddrReq, opt ...natsrpc.CallOption) (*GateServerAddr, error) {
	rep := &GateServerAddr{}
	err := c.c.Request(ctx, "GetGateServerAddr", req, rep, opt...)
	return rep, err
}
func (c *_DiscoveryNATSRPCClient) GetAllGateServerInfoList(ctx context.Context, req *NullMsg, opt ...natsrpc.CallOption) (*GateServerInfoList, error) {
	rep := &GateServerInfoList{}
	err := c.c.Request(ctx, "GetAllGateServerInfoList", req, rep, opt...)
	return rep, err
}
func (c *_DiscoveryNATSRPCClient) GetMainGameServerAppId(ctx context.Context, req *NullMsg, opt ...natsrpc.CallOption) (*GetMainGameServerAppIdRsp, error) {
	rep := &GetMainGameServerAppIdRsp{}
	err := c.c.Request(ctx, "GetMainGameServerAppId", req, rep, opt...)
	return rep, err
}
