package star

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type Student struct {
	Grade      string
	College    string
	Department string
	Class      string
	ClassNum   string
}

// 判断学号是否合格
func IsValid(studentID string) bool {
	ans, _ := regexp.MatchString(`^\d{12,13}$`, studentID)
	art, _ := regexp.MatchString(`^\d{11}ys$`, studentID)
	return ans || art
}

// 用此函数获得所需信息
func GetAnswer(studentID string) (*Student, error) {
	// 先判断是否是正确的学号格式
	if !IsValid(studentID) {
		fmt.Println("not valid studentID")
		return nil, errors.New("not valid studentID")
	}
	var re *regexp.Regexp
	// 本科生
	if len(studentID) == 13 {
		art, _ := regexp.MatchString(`^\d{11}ys$`, studentID)
		if art {
			// 艺术生
			re = regexp.MustCompile(`(\d{4})(\d{2})(\d{2})(\d{2})(.{3})`)
		} else {
			re = regexp.MustCompile(`(\d{4})(\d{2})(\d{2})(\d{2})(\d{3})`)
		}
	} else {
		// 研究生
		re = regexp.MustCompile(`(\d{4})(\d{2})(\d{2})(\d{2})(\d{2})`)
	}
	answer := re.FindStringSubmatch(studentID)
	fmt.Println("True,", strings.Join(answer[1:], ","))
	return &Student{Grade: answer[1], College: answer[2], Department: answer[3], Class: answer[4], ClassNum: answer[5]}, nil
}

// 注册为http服务
func HttpStudentRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	studentID := r.Form.Get("id")
	student, err := GetAnswer(studentID)
	if err != nil {
		// 学号不合乎规矩，直接填入400错误，不写入校验信息
		w.WriteHeader(400)
		w.Write([]byte("student id is not valid"))
		return
	}
	studentJson, err := json.Marshal(*student)
	if err != nil {
		fmt.Println("json failed err is ", err)
		// 服务器内部错误
		w.WriteHeader(500)
		// 将错误写入body
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(studentJson)
	w.WriteHeader(200)
}
