package network

import (
	"docker_go/container"
	"github.com/vishvananda/netlink"
	"net"
)

// 网络
type Network struct {
	Name    string
	IpRange *net.IPNet
	Driver  string
}

// 网络端点
type Endpoint struct {
	ID          string           `json:"id"`
	Device      netlink.Veth     `json:"dev"`
	IPAddress   net.IP           `json:"ip"`
	MacAddress  net.HardwareAddr `json:"mac"`
	Network     *Network
	PortMapping []string
}

// 网络驱动接口
type NetworkDriver interface {
	// 驱动名
	Name() string
	// 创建网络
	Create(subnet string, name string)
	// 删除网络
	Delete(network Network) error
	// 连接容器网络端点到网络
	Connect(network *Network, endpoint *Endpoint) error
	// 从网络上移除容器网络端点
	Disconnect(network Network, endpoint *Endpoint) error
}

// 初始化网络驱动
func Init() error {
	return nil
}

func Connect(networkName string, containerInfo *container.ContainerInfo) error {
	return nil
}

// 创建网络
func CreateNetwork(driver, subnet, name string) error {
	return nil
}
