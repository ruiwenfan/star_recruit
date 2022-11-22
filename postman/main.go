package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/test", testHandler)
	//选择对应URL所处理的函数
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	http.ListenAndServe("127.0.0.1:1234", nil)
}

// 处理函数
func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	if r.Method == "POST" { //监听是否为POST方法
		b, err := ioutil.ReadAll(r.Body)
		//读取r中body的所有数据
		fmt.Print(string(b))
		//控制台打印数据
		if err != nil {
			log.Println("Data failed:", err)
		}
	}
}
