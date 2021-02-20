# docker_go
使用golang写一个docker应用出来

unit test
```sh
$ go test ./tests/ -v -test.run GenContainerID

$ go test ./tests/ -v -test.run RecordContainerInfo
$ cat /var/run/docker_go/mycontainer/config.json

$ go test ./tests/ -v -test.run ListContainerInfo

$ go test ./tests/ -v -test.run StopContainer

$ go test ./tests/ -v -test.run RemoveContainer

```

