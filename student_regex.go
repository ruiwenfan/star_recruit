package star

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValid(studentID string) bool {
	ans, _ := regexp.MatchString(`^\d{12,13}$`, studentID)
	art, _ := regexp.MatchString(`^\d{11}ys$`, studentID)
	return ans || art
}

// 输入此函数的一定被检验过是合乎标准的学号
func GetAnswer(studentID string) {
	// 本科生
	var re *regexp.Regexp
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
}
