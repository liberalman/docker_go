# docker_go
使用golang写一个docker应用出来

unit test
```sh
$ go test ./tests/ -v -test.run GenContainerID

$ go test ./tests/ -v -test.run RecordContainerInfo
$ cat /var/run/docker_go/mycontainer/config.json

$ go test ./tests/ -v -test.run NewWorkSpace

$ go test ./tests/ -v -test.run ListContainerInfo

$ go test ./tests/ -v -test.run StopContainer

$ go test ./tests/ -v -test.run RemoveContainer

```

设置CentOS支持aufs

查看是否支持

cat /proc/filesystems

安装aufs
```sh
cd /etc/yum.repo.d
# 下载文件
wget https://yum.spaceduck.org/kernel-ml-aufs/kernel-ml-aufs.repo
# 安装
yum install kernel-ml-aufs
# 修改内核启动
vim /etc/default/grub
## 修改参数
GRUB_DEFAULT=0

# 重新生成grub.cfg
grub2-mkconfig -o /boot/grub2/grub.cfg

# 重启计算机
reboot
```

配置busybox
```sh
# 下载 busybox
docker pull busybox
# 运行
docker run -d busybox top -b
# 导出
docker export -o busybox.tar (容器ID)
# 解压到 /root文件夹下
cd /root
mkdir busybox
tar -xvf busybox.tar -C busybox/
```

使用指南
```sh
# 编译
go build .

# 启动一个容器, busybox为镜像名，存放在 /root/busybox.tar
./go-docker run -ti --name test busybox sh

# 后台启动
./go-docker run -d --name test busybox sh

# 挂载文件
./go-docker run -d -v /root/test:/test --name test busybox sh

# 进入容器
./go-docker exec test sh

# 查看容器日志
./go-docker logs test

# 查看容器列表
./go-docker ps

# 停止容器
./go-docker stop test

# 删除容器
./go-docker rm test
```
