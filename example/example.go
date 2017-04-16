package main

import (
	"log"

	"github.com/ariefrahmansyah/fingerplus"
)

const (
	serverURL     = "http://192.168.1.2:8080"
	serialNumber1 = "2251016030256"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fp := fingerplus.New()
	fp.ServerURL = serverURL

	devInfo, err := fp.GetDeviceInfo(serialNumber1)
	log.Printf("GetDeviceInfo: %+v Error: %+v\n", devInfo, err)

	setTime, err := fp.SetTime(serialNumber1)
	log.Printf("SetTime: %+v Error: %+v\n", setTime, err)

	allScanLog, err := fp.GetAllScanLog(serialNumber1)
	log.Printf("GetAllScanLog: %+v Error: %+v\n", allScanLog, err)

	newScanLog, err := fp.GetNewScanLog(serialNumber1)
	log.Printf("GetNewScanLog: %+v Error: %+v\n", newScanLog, err)

	delScanLog, err := fp.DeleteScanLog(serialNumber1)
	log.Printf("DeleteScanLog: %+v Error: %+v\n", delScanLog, err)

	allUser, err := fp.GetAllUser(serialNumber1)
	log.Printf("GetAllUser: %+v Error: %+v\n", allUser, err)
}
