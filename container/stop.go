package container

import (
	// system package
	"encoding/json"
	"io/ioutil"
	"path"
	"strconv"
	"syscall"

	// third package
	"github.com/sirupsen/logrus"

	// item package
	"docker_go/common"
)

func StopContainer(containerName string) {
	info, err := getContainerInfo(containerName) // 读取容器信息
	if nil != err {
		logrus.Errorf("get container info, err: %v", err)
		return
	}

	if info.Pid != "" {
		pid, _ := strconv.Atoi(info.Pid) // 这是容器的init进程在宿主机上的Pid

		// 杀死宿主机进程
		if err := syscall.Kill(pid, syscall.SIGTERM); nil != err { // 调用kill的系统掉用杀死进程
			logrus.Errorf("stop container, pid: %d, err: %v", pid, err)
			return
		}

		// 修改容器状态
		info.Status = common.Stop
		info.Pid = ""
		bs, _ := json.Marshal(info)
		fileName := path.Join(common.DefaultContainerInfoPath, containerName, common.ContainerInfoFileName)
		if err := ioutil.WriteFile(fileName, bs, 0622); nil != err { // 回写
			logrus.Errorf("write container config.json, err: %v", err)
		}
	}
}
