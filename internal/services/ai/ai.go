package ai

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sync"

	fastjson "github.com/goccy/go-json"
)

var (
	addr       string
	bufferPool = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 10240)
		},
	}
)

type RequestData struct {
	Prompt  string   `json:"prompt"`
	History []string `json:"history"`
}

type ResponseData struct {
	Response string   `json:"response"`
	History  []string `json:"history"`
}

func Init() {
	fmt.Println("[AI] Init")
	addr = os.Getenv("AI_HOST") + ":" + os.Getenv("AI_PORT")
}

func Request(requestData *RequestData) (*ResponseData, error) {
	// connect to this socket
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("[AI]", err)
		return nil, err
	}
	defer conn.Close()

	request, err := fastjson.Marshal(requestData)
	if err != nil {
		fmt.Println("[AI]", err)
		return nil, err
	}
	request = append(request, byte('\x00'))

	conn.Write(request)

	// listen for reply
	vbuffer := bufferPool.Get()
	buffer := vbuffer.([]byte)
	len, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("[AI]", err)
		bufferPool.Put(vbuffer)
		return nil, err
	}

	var data ResponseData
	err = json.Unmarshal(buffer[:len], &data)
	if err != nil {
		fmt.Println("[AI]", err)
		bufferPool.Put(vbuffer)
		return nil, err
	}

	bufferPool.Put(vbuffer)
	return &data, nil
}
