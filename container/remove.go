package container

import (
	// system package

	// third package
	"github.com/sirupsen/logrus"

	// item package
	"docker_go/common"
)

// 删除容器
func RemoveContainer(containerName string) {
	info, err := getContainerInfo(containerName) // 读取容器信息
	if nil != err {
		logrus.Errorf("get container info, err: %v", err)
		return
	}

    // 不能删除正在运行的容器，只能删除停止状态的容器
	if info.Status != common.Stop {
        logrus.Errorf("can't remove running container")
        return
    }

    DeleteContainerInfo(containerName)
}

