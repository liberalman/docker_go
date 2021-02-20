package tests

import (
	// system package
    "fmt"
	"testing"

	// item package
	"docker_go/container"
)

func TestRecordContainerInfo(t *testing.T) {
	// 记录容器信息
	err := container.RecordContainerInfo(1234, []string{"ls"}, "mycontainer", "5678")
	if nil != err {
		t.Errorf("record container info, err: %v", err)
	}
}

func TestDeleteContainerInfo(t *testing.T) {
	container.DeleteContainerInfo("mycontainer")
}

func TestGenContainerID(t *testing.T) {
    fmt.Println(container.GenContainerID(10))
}

func TestListContainerInfo(t *testing.T) {
	container.ListContainerInfo()
}

func TestStopContainer(t *testing.T) {
	container.StopContainer("mycontainer")
}

func TestRemoveContainer(t *testing.T) {
	container.RemoveContainer("mycontainer")
}

