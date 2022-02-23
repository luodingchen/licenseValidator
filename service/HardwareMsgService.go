package service

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"io/ioutil"
	"net"
	"verifyLinux/models"
)

func GetHardwareMsg() models.HardwareInfo {
	var HardwareInfo models.HardwareInfo
	cpuInfo, _ := cpu.Info()
	HardwareInfo.Cpu = cpuInfo

	fs, _ := ioutil.ReadDir("/dev/disk/by-uuid")
	for _, v := range fs {
		HardwareInfo.Disk = append(HardwareInfo.Disk, v.Name())
	}

	hostInfo, _ := host.Info()
	HardwareInfo.Host = *hostInfo

	HardwareInfo.Net, _ = net.Interfaces()

	return HardwareInfo
}
