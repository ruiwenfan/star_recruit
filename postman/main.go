package main

// 这个文件下是为了测试http服务
import (
	"log"
	"net/http"
	"star"
)

func main() {
	http.HandleFunc("/createPost", star.HttpCreatePost)
	http.HandleFunc("/parseStudentId", star.HttpStudentRequest)
	http.HandleFunc("/searchPostAndUser", star.HttpSearchPostAndUser)
	log.Println("Defaulting to port 1234")
	http.ListenAndServe("127.0.0.1:1234", nil)
}
