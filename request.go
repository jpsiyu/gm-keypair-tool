package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const url = "http://localhost:8801"

type Res33 struct {
	Id     int32       `json:"id"`
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
}

type Req33 struct {
	JsonRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

func Post33(data *Req33) (error, *Res33) {
	dataJson, _ := json.Marshal(data)
	dataBytes := bytes.NewBuffer(dataJson)
	req, _ := http.NewRequest("POST", url, dataBytes)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return err, nil
	}

	defer rsp.Body.Close()
	body, _ := ioutil.ReadAll(rsp.Body)
	//fmt.Println(string(body))

	var res33 Res33
	json.Unmarshal(body, &res33)
	return nil, &res33
}

func Version() (error, *Res33) {
	data := Req33{"2.0", "Chain33.Version", make([]interface{}, 0)}
	err, res33 := Post33(&data)
	return err, res33
}

func GetWalletStatus() (error, *Res33) {
	data := Req33{"2.0", "Chain33.GetWalletStatus", make([]interface{}, 0)}
	err, res33 := Post33(&data)
	return err, res33
}

func GenSeed() (error, *Res33) {
	param := map[string]int32{}
	data := Req33{"2.0", "Chain33.GenSeed", []interface{}{param}}
	err, res33 := Post33(&data)
	return err, res33
}

func SaveSeed(seed, passwd string) (error, *Res33) {
	param := map[string]string{}
	param["seed"] = seed
	param["passwd"] = passwd
	data := Req33{"2.0", "Chain33.SaveSeed", []interface{}{param}}
	err, res33 := Post33(&data)
	return err, res33
}

func UnLock(passwd string) (error, *Res33) {
	param := map[string]interface{}{"timeout": 0, "passwd": passwd}
	data := Req33{"2.0", "Chain33.UnLock", []interface{}{param}}
	err, res33 := Post33(&data)
	return err, res33
}

func NewAccount(label string) (error, *Res33) {
	param := map[string]interface{}{"label": label}
	data := Req33{"2.0", "Chain33.NewAccount", []interface{}{param}}
	err, res33 := Post33(&data)
	return err, res33
}

func DumpPrivkey(addr string) (error, *Res33) {
	param := map[string]interface{}{"data": addr}
	data := Req33{"2.0", "Chain33.DumpPrivkey", []interface{}{param}}
	err, res33 := Post33(&data)
	return err, res33
}
