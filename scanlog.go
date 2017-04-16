package fingerplus

import (
	"encoding/json"
	"net/url"
)

type ScanLog struct {
	Result bool          `json:"Result"`
	Data   []ScanLogData `json:"Data"`
}

type ScanLogData struct {
	SN         string `json:"SN"`
	ScanDate   string `json:"ScanDate"`
	PIN        string `json:"PIN"`
	VerifyMode int    `json:"VerifyMode"`
	IOMode     int    `json:"IOMode"`
	WorkCode   int    `json:"WorkCode"`
}

func (fp *Fingerplus) GetAllScanLog(serialNumber string) (ScanLog, error) {
	params := url.Values{
		"sn": {serialNumber},
	}

	postman := NewPostman()
	postman.Header = header
	postman.Method = MethodPost
	postman.URL = fp.ServerURL + getAllScanLog
	postman.Params = params

	_, body, err := postman.Send()
	if err != nil {
		return ScanLog{}, err
	}

	scanLog := ScanLog{}

	err = json.Unmarshal(body, &scanLog)
	if err != nil {
		return ScanLog{}, err
	}

	return scanLog, nil
}

func (fp *Fingerplus) GetNewScanLog(serialNumber string) (ScanLog, error) {
	params := url.Values{
		"sn": {serialNumber},
	}

	postman := NewPostman()
	postman.Header = header
	postman.Method = MethodPost
	postman.URL = fp.ServerURL + getNewScanLog
	postman.Params = params

	_, body, err := postman.Send()
	if err != nil {
		return ScanLog{}, err
	}

	scanLog := ScanLog{}

	err = json.Unmarshal(body, &scanLog)
	if err != nil {
		return ScanLog{}, err
	}

	return scanLog, nil
}

func (fp *Fingerplus) DeleteScanLog(serialNumber string) (Response, error) {
	params := url.Values{
		"sn": {serialNumber},
	}

	postman := NewPostman()
	postman.Header = header
	postman.Method = MethodPost
	postman.URL = fp.ServerURL + deleteScanLog
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
