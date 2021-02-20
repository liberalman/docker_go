package container

import (
	// system package
	"encoding/json"
	"fmt"
	"io/ioutil"
    "math/rand"
	"os"
	"path"
	"strconv"
	"strings"
    "text/tabwriter"
	"time"

	// third package
	"github.com/sirupsen/logrus"

	// item package
	"docker_go/common"
)

type ContainerInfo struct {
	Pid        string   `json:"pid"`     // 容器的init进程在宿主机上的PID
	Id         string   `json:"id"`      // 容器ID
	Command    string   `json:"command"` // 容器内init进程的运行命令
	Name       string   `json:"name"`    //
	CreateTime string   `json:"createTime"`
	Status     string   `json:"status"`
	Volume     string   `json:"volume"`      // 容器的数据卷
	PortMaping []string `json:"portmapping"` // 端口映射
}

func GenContainerID(length int) string {
    letterBytes := "0123456789"
    letterLen := len(letterBytes)
    rand.Seed(time.Now().UnixNano()) // 如果不加随机种子，每次遍历获取都是重复的数据
    buf := make([]byte, length)
    for i := range buf {
        buf[i] = letterBytes[rand.Intn(letterLen)]
    }
    return string(buf)
}

// 记录容器信息
func RecordContainerInfo(containerPID int, cmdArray []string, containerName, containerID string) error {
	info := &ContainerInfo{
		Pid:        strconv.Itoa(containerPID),
		Id:         containerID,
		Command:    strings.Join(cmdArray, ""),
		Name:       containerName,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		Status:     common.Running,
	}

	dir := path.Join(common.DefaultContainerInfoPath, containerName) // 生成目录路径
	_, err := os.Stat(dir) // 获取目录状态
	if err != nil && os.IsNotExist(err) {
		// 目录不存在，则新建
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			logrus.Errorf("mkdir container dir: %s, err: %v", dir, err)
			return nil
		}
	}

	fileName := fmt.Sprintf("%s/%s", dir, common.ContainerInfoFileName) // 生成config.json文件路径
	file, err := os.Create(fileName) // 不管之前存不存在，都新建文件
	if err != nil {
		logrus.Errorf("create config.json, fileName: %s, err: %v", fileName, err)
		return err
	}

	bs, _ := json.Marshal(info) // 转换结构图为json字符串
	_, err = file.WriteString(string(bs)) // 写json串入到文件
	if nil != err {
		logrus.Errorf("write config.json, fileName: %s, err: %v", fileName, err)
		return err
	}

	return nil
}

func DeleteContainerInfo(containerName string) {
    dir := path.Join(common.DefaultContainerInfoPath, containerName)
    err := os.RemoveAll(dir)
    if err != nil {
        logrus.Errorf("remove container info, err: %v", err)
    }
}

func getContainerInfo(containerName string) (*ContainerInfo, error) {
    filePath := path.Join(common.DefaultContainerInfoPath, containerName, common.ContainerInfoFileName)
    bs, err := ioutil.ReadFile(filePath)
    if nil != err {
        logrus.Errorf("read file, path: %s, err: %v", filePath, err)
        return nil, err
    }
    info := &ContainerInfo{}
    err = json.Unmarshal(bs, info)

    return info, err
}

func ListContainerInfo() {
    files, err := ioutil.ReadDir(common.DefaultContainerInfoPath) // 读取根路径下各个容器目录名称
    if err != nil {
        logrus.Errorf("read info dir, err: %v", err)
    }

    var infos []*ContainerInfo
    for _, file := range files {
        info, err := getContainerInfo(file.Name()) // 通过文件名获取其下config.json内容
        if err != nil {
            logrus.Errorf("get container info, name: %s, err: %v", file.Name(), err)
            continue
        }
        infos = append(infos, info)
    }

    // 打印到标准输出（屏幕）上
    w := tabwriter.NewWriter(os.Stdout, 12, 1, 2, ' ', 0)
    _, _ = fmt.Fprintf(w, "ID\tNAME\tPID\tSTATUS\tCOMMAND\tCREATED\n")
    for _, info := range infos {
        _, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t\n",
            info.Id, info.Name, info.Pid, info.Status, info.Command, info.CreateTime)
    }

    // 刷新标准输出流缓存区，将容器列表打印出来
    if err := w.Flush(); nil != err {
        logrus.Errorf("flush info, err: %v", err)
    }
}

