package network

import (
	"docker_go/container"
	"net"
)

// 网络
type Network struct {
	Name    string
	IpRange *net.IPNet
	Driver  string
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
