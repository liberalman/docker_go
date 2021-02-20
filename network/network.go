package network

import "docker_go/container"

// 初始化网络驱动
func Init() error {
    return nil
}

func Connect(networkName string, containerInfo *container.ContainerInfo) error {
    return nil
}
