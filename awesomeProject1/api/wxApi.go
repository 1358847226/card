package api

import (
	"awesomeProject1/models"
	"bytes"
	"encoding/json"
	"io"
	_ "io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetAccessToken() string {
	client := &http.Client{Timeout: 5 * time.Second}
	resp2,_ := client.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxdf131fbe9fc9887f&secret=659e5540cf1c36bdadd5f2e08fd877eb",)
	defer resp2.Body.Close()
	var buffer2 [512]byte
	result2 := bytes.NewBuffer(nil)
	for {
		n, err := resp2.Body.Read(buffer2[0:])
		result2.Write(buffer2[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	var resultData2 = result2.String()
	log.Println("res2",resultData2)
	comma := strings.Index(resultData2, ":")
	d := strings.Index(resultData2,",")
	res2 := resultData2[comma+2 :d-1]
	return res2
}

func SubscribeMessage(res2 string, data models.SubscribeMessage) string{
	client := &http.Client{Timeout: 5 * time.Second}
	b ,_ := json.Marshal(data)

	log.Println("b",data)
	body := bytes.NewBuffer(b)
	log.Println("body",body)
	resp, _ := client.Post("https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token="+res2,"content-type:application/json;charset=UTF-8",
		body)
	resp.Header.Set("Content-Type", "application/json; encoding=utf-8")
	defer resp.Body.Close()


	var buffer2 [512]byte
	result2 := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer2[0:])
		result2.Write(buffer2[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	var resultData2 = result2.String()
	log.Println("res2",resultData2)

	return resultData2
}