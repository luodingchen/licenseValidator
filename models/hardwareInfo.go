package models

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"net"
)

type HardwareInfo struct {
	Cpu  []cpu.InfoStat  `json:"cpu"`
	Disk []string        `json:"disk"`
	Host host.InfoStat   `json:"host"`
	Net  []net.Interface `json:"net"`
}

type HardwareJson struct {
	Cpu  string `json:"cpu"`
	Disk string `json:"disk"`
	Host string `json:"host"`
	Net  string `json:"net"`
}
