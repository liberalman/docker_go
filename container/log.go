package container

import (
	// system package
	"fmt"
	"io/ioutil"
	"os"
	"path"

	// third package
	"github.com/sirupsen/logrus"

	// item package
	"docker_go/common"
)

func LookContainerLog(containerName string) {
	logFileName := path.Join(common.DefaultContainerInfoPath, containerName, common.ContainerLogFileName)
	file, err := os.Open(logFileName)
	if nil != err {
		logrus.Errorf("open log file, path: %s, err: %v", logFileName, err)
		return
	}

	bs, _ := ioutil.ReadAll(file)
	if err != nil {
		logrus.Errorf("read log file, err: %v", err)
	}

	_, _ = fmt.Fprint(os.Stdout, string(bs))
}
