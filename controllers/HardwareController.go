package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"verifyLinux/models"
	"verifyLinux/service"
)

const aesKeyString = "4IPQW2YFCEWPMBA7"

func GenerateHardwareMsg(c *gin.Context) {
	var HardwareInfo models.HardwareInfo

	HardwareInfo = service.GetHardwareMsg()

	cpuBytes, _ := json.Marshal(HardwareInfo.Cpu)
	diskBytes, _ := json.Marshal(HardwareInfo.Disk)
	hostBytes, _ := json.Marshal(HardwareInfo.Host)
	netBytes, _ := json.Marshal(HardwareInfo.Net)

	encryptedCpu, err := service.AesEncrypt(cpuBytes, aesKeyString)
	if err != nil {
		panic(err)
	}

	encryptedDisk, err := service.AesEncrypt(diskBytes, aesKeyString)
	if err != nil {
		panic(err)
	}

	encryptedHost, err := service.AesEncrypt(hostBytes, aesKeyString)
	if err != nil {
		panic(err)
	}

	encryptedNet, err := service.AesEncrypt(netBytes, aesKeyString)
	if err != nil {
		panic(err)
	}

	var HardwareJson = models.HardwareJson{
		Cpu:  encryptedCpu,
		Disk: encryptedDisk,
		Host: encryptedHost,
		Net:  encryptedNet,
	}

	HardwareJsonBytes, err := json.Marshal(HardwareJson)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("files/HardwareMsg.json", HardwareJsonBytes, 0666)
	if err != nil {
		panic(err)
	}
	c.Header("Content-LicenseType", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"HardwareMsg.json")
	c.Header("Content-Transfer-Encoding", "binary")
	c.File("files/HardwareMsg.json")
}
