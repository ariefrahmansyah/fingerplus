package fingerplus

import (
	"encoding/json"
	"net/url"
)

type DeviceInfo struct {
	Result bool           `json:"Result"`
	Data   DeviceInfoData `json:"DEVINFO"`
}

type DeviceInfoData struct {
	Jam            string `json:"Jam"`
	Admin          string `json:"Admin"`
	User           string `json:"User"`
	FP             string `json:"FP"`
	CARD           string `json:"CARD"`
	PWD            string `json:"PWD"`
	AllOperasional string `json:"All Operasional"`
	AllPresensi    string `json:"All Presensi"`
	NewOperasional string `json:"New Operasional"`
	NewPresensi    string `json:"New Presensi"`
}

func (fp *Fingerplus) GetDeviceInfo(serialNumber string) (DeviceInfo, error) {
	params := url.Values{
		"sn": {serialNumber},
	}

	postman := NewPostman()
	postman.Header = header
	postman.Method = MethodPost
	postman.URL = fp.ServerURL + getDeviceInfo
	postman.Params = params

	_, body, err := postman.Send()
	if err != nil {
		return DeviceInfo{}, err
	}

	deviceInfo := DeviceInfo{}

	err = json.Unmarshal(body, &deviceInfo)
	if err != nil {
		return DeviceInfo{}, err
	}

	return deviceInfo, nil
}

func (fp *Fingerplus) SetTime(serialNumber string) (Response, error) {
	params := url.Values{
		"sn": {serialNumber},
	}

	postman := NewPostman()
	postman.Header = header
	postman.Method = MethodPost
	postman.URL = fp.ServerURL + setTime
	postman.Params = params

	_, body, err := postman.Send()
	if err != nil {
		return Response{}, err
	}

	response := Response{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return Response{}, err
	}

	return response, nil
}
