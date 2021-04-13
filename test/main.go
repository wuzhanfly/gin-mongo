package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type S struct {
	Id string `json:"_id"`
}

func main() {
	Database := new(S)
	res := Post_1("https://api.fgas.io/api/v1/getFil",`{"type": "32G"}`,"application/json")
	json.Unmarshal(res, &Database)
	fmt.Println(11,Database.Id)
	
}

//url:请求地址		data:POST请求提交的数据		contentType:请求体格式，如：application/json
//content:请求返回的内容
func Post_1(url string, data string, contentType string) []byte {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	fmt.Println(url)
	req.Header.Add("content-type", contentType)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(result))
	return result
}





