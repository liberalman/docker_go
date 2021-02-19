package common

const (
    RootPath   = "/home/socho/"     // /root/
    MntPath    = "/home/socho/mnt/" // /root/mnt/
    WriteLayer = "writeLayer"
)

const (
    Running = "running"
    Stop    = "stopped"
    Exit    = "exited"
)

const (
    DefaultContainerInfoPath = "/var/run/docker_go/"
    ContainerInfoFileName    = "config.json"
    ContainerLogFileName     = "container.log"
)

const (
    EnvExecPid = "docker_pid"
    EnvExecCmd = "docker_cmd"
)

const (
    DefaultNetworkPath   = "/var/run/docker_go/network/network/"
    DefaultAllocatorPath = "/var/run/docker_go/network/ipam/subnet.json"
)
