package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
	"time"
	"verifyLinux/models"
	"verifyLinux/service"
)

func VerifyLicense(c *gin.Context) {
	var res = NewResultMsg(c)
	productCode := c.Query("product_code")

	fs, err := ioutil.ReadFile("files/license")
	if err != nil {
		res.Error(err.Error())
		return
	}
	decrypted, _ := service.AesDecrypt(string(fs), aesKeyString)

	var license models.License
	err = yaml.Unmarshal(decrypted, &license)
	if err != nil {
		res.Error(err.Error())
		return
	}
	configBytes := []byte(license.ConfigJson)

	err = json.Unmarshal([]byte(license.ConfigJson), &license.Config)
	if err != nil {
		res.Error(err.Error())
		return
	}

	publicKey, err := service.DecodePublicKeyString(license.LicensePublicKey)
	if err != nil {
		res.Error(err.Error())
		return
	}

	err = service.Verify(publicKey, configBytes, license.LicenseSignature)
	if err != nil {
		res.Error(err.Error())
		return
	}
	deadlineTime, err := time.Parse("2006-01-02 15:04:05.000", license.Config.Deadline)
	if deadlineTime.Before(time.Now()) {
		res.Error("Time 已过期")
		return
	}

	if service.TimeTamperProofService(license.Config.StartTime, license.Config.Deadline) != license.Config.HardwareList[0].Host.Uptime {
		res.Error("Time 已被篡改")
		return
	}

	fmt.Println("Time 校验成功")

	var localHardwareInfo models.HardwareInfo
	localHardwareInfo = service.GetHardwareMsg()
	verifyNumber := 0

	for _, licenseHardwareInfo := range license.Config.HardwareList {
		if reflect.DeepEqual(licenseHardwareInfo.Cpu, localHardwareInfo.Cpu) {
			fmt.Println("Cpu 校验成功")
			verifyNumber += 1
		}
		if reflect.DeepEqual(licenseHardwareInfo.Disk, localHardwareInfo.Disk) {
			fmt.Println("Disk 校验成功")
			verifyNumber += 1
		}
		if licenseHardwareInfo.Host.Hostname == localHardwareInfo.Host.Hostname {
			fmt.Println("Host 校验成功")
			verifyNumber += 1
		}
		if reflect.DeepEqual(licenseHardwareInfo.Net, localHardwareInfo.Net) {
			fmt.Println("Net 校验成功")
			verifyNumber += 1
		}
		if verifyNumber < 4 {
			verifyNumber = 0
		} else {
			break
		}
	}
	if verifyNumber == 4 {
		fmt.Println("Hardware 校验完毕")
	} else {
		res.Error("Hardware 校验失败")
		return
	}
	for _, product := range license.Config.ProductList {
		if productCode == product.ProductCode {
			fmt.Println("Product 校验成功")
			fmt.Println("License 校验完毕")
			res.Success("License 校验成功", product.ProductFuncList)
			return
		}
	}
	res.Error("License 校验失败")
}

func UploadLicense(c *gin.Context) {
	var res = NewResultMsg(c)
	file, err := c.FormFile("file")
	if err != nil {
		res.Error(err.Error())
		return
	}
	err = c.SaveUploadedFile(file, "files/license")
	if err != nil {
		res.Error(err.Error())
		return
	}
	res.Success("Upload success", nil)
}
