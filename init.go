package fingerplus

import "net/http"

var header = http.Header{
	"Cache-Control": {"no-cache"},
	"Content-Type":  {"application/x-www-form-urlencoded"},
}

type Response struct {
	Result bool `json:"Result"`
}

type Fingerplus struct {
	ServerURL string
}

func New() *Fingerplus {
	return &Fingerplus{}
}

const (
	getDeviceInfo = "/dev/info"
	setTime       = "/dev/settime"

	getAllScanLog = "/scanlog/all"
	getNewScanLog = "/scanlog/new"
	deleteScanLog = "/scanlog/del"

	getAllUser    = "/user/all"
	setUser       = "/user/set"
	deleteUser    = "/user/delall"
	deleteUserPIN = "/user/del"
)
