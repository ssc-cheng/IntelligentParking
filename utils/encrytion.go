package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)


// 随机生成数的校验码,count是指定的生成数的位数
func Num(count int) string {
	numeric := [10]byte{1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var random strings.Builder
	for i := 0; i < count; i++ {
		fmt.Fprintf(&random, "%d", numeric[ rand.Intn(r) ])
	}
	return random.String()
}


/*
生成随机字符串
*/
func GetRandomString(lens int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}



//打码：身份证、姓名、手机
func EncryptField(input string) (output string) {
	output = input
	regsfz := regexp.MustCompile(`^(^[1-9]\d{5}(18|19|([23]\d))\d{2}(((0[13578]|10|12)(0[1-9]|[12]\d|3[01]))|((0[469]|11)(0[1-9]|[12]\d|30))|(02(0[1-9]|[12]\d)))\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}(((0[13578]|10|12)(0[1-9]|[12]\d|3[01]))|((0[469]|11)(0[1-9]|[12]\d|30))|(02(0[1-9]|[12]\d)))\d{3}$)`)
	regxm := regexp.MustCompile("[\u4e00-\u9fa5]")
	regdhhm := regexp.MustCompile(`^1([38]\d|5[0-35-9]|7[3678])\d{8}$`)
	if result := string(regsfz.Find([]byte(string(input)))); len(result) > 0 {
		output = result[0:5] + `******` + result[11:]
		return

	} else if result := regxm.FindAll([]byte(input), -1); len(result) > 1 {
		var m = []byte{}
		if len(result) == 2 {
			m = append(m, []byte("*")...)
			m = append(m, result[1]...)
			output = string(m)
			return
		} else if len(result) > 2 {
			m = append(m, result[0]...)
			m = append(m, []byte("*")...)
			for i := 2; i < len(result); i++ {
				m = append(m, result[i]...)
			}
			output = string(m)
			return
		} else {
			output = string(result[0])
			return
		}
	} else if result := string(regdhhm.Find([]byte(string(input)))); len(result) > 0 {
		m := result[0:5] + `***` + result[8:]
		output = m
		return
	} else {
		output = input
		return
	}
	return
}
