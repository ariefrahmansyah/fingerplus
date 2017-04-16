package fingerplus

import (
	"encoding/json"
	"net/url"
)

type User struct {
	Result bool       `json:"Result"`
	Data   []UserData `json:"Data"`
}

type UserData struct {
	PIN       string         `json:"PIN"`
	Name      string         `json:"Name"`
	RFID      string         `json:"RFID"`
	Password  string         `json:"Password"`
	Privilege int            `json:"Privilege"`
	Template  []TemplateData `json:"Template`
}

type TemplateData struct {
	PIN      string `json:"pin"`
	IDX      int    `json:"idx"`
	AlgVer   int    `json:"alg_ver"`
	Template string `json:"template"`
}

func (fp *Fingerplus) GetAllUser(serialNumber string) (User, error) {
	params := url.Values{
		"sn": {serialNumber},
	}

	postman := NewPostman()
	postman.Header = header
	postman.Method = MethodPost
	postman.URL = fp.ServerURL + getAllUser
	postman.Params = params

	_, body, err := postman.Send()
	if err != nil {
		return User{}, err
	}

	user := User{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// TODO
func (fp *Fingerplus) SetUser(serialNumber string) error {
	return nil
}

// TODO
func (fp *Fingerplus) DeleteAllUser(serialNumber string) error {
	return nil
}

// TODO
func (fp *Fingerplus) DeleteUserPIN(serialNumber string) error {
	return nil
}
