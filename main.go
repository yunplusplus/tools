package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	base64 "github.com/yunplusplus/tools/io.yunplusplus.encoding"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Query struct {
	Pagenum  int `json:"pagenum"`
	Pagesize int `json:"pagesize"`
}

func (q Query) toBytes() (body *bytes.Buffer, err error) {
	b, err := json.Marshal(q)
	log.Print(string(b))
	if err != nil {
		return &bytes.Buffer{}, err
	}
	return bytes.NewBuffer(b), err
}

func syncRoutine() {
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		fmt.Println(i)
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}
func Request() {
	query := Query{Pagenum: 1, Pagesize: 10}
	body, err := query.toBytes()
	if err != nil {
		return
	}
	resp, err := http.Post("http://www.chnsys-law.com/onlineLitigation/listEtmsCourt", "application/json", body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read failed:", err)
		return
	}
	log.Println("content:", string(content))

	syncRoutine()
}
func main() {
	var c = base64.Encode64("赵云涛")
	log.Println(c)
	c, _ = base64.Decode64(c)
	log.Println(c)
}
