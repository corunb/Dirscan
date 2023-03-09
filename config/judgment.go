package config

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

// FindUrl 识别http/https,并探存活
func FindUrl(Turl string)  bool{
	var result bool
	u, err := url.Parse(Turl)
	if err != nil {
		fmt.Println(err)
	}

	if strings.ToLower(u.Scheme) != "http" && strings.ToLower(u.Scheme) != "https" {
		fmt.Println("[!] 请输入正确的url!")
		result = false
	}else {
		//fmt.Println("正确的url:",Turl)
		resp := Request(Turl)

		if resp != nil {
			respCode := resp.StatusCode   //状态码

			if respCode <200 && respCode >=501  {
				//fmt.Println("[!] 目标url无法访问")
				result = false
			}else {
				//fmt.Println("respCode",respCode)
				result = true
			}
		}else {
			//fmt.Println("resp:",resp)
			//fmt.Println("[!] 目标url是无法访问")
			result = false
		}

	}

	return result
}

// Thundred 对200的页面进行识别判断
func Thundred(Turl string) bool{
	_ , body := JumpUrl(Turl)
	var str bool
	//对页面的长度、内容2个方面进行判断，越趋近3的网站，误报越小
	result :=  Lenbody(len(body)) +Resbodya(string(body))
	if result < 2 {
		str = false
	}else {
		str = true
	}
	return str
}

// Redirect 判断识别302页面内容
func Redirect(Turl string) bool{
	respCode , body := JumpUrl(Turl)
	var str bool
	//对页面的状态码、长度、内容3个方面进行判断，越趋近4的网站，误报越小
	result := ResCode(respCode) + Lenbody(len(body)) +Resbody(string(body))
	if result < 2 {
		str = false
	}else {
		str = true
	}
	//fmt.Println("str",str)
	return str
}

// Thundreds 判断400-500
func Thundreds(Turl string) bool{
	_ , body := JumpUrl(Turl)
	var str bool
	//对页面的长度、内容2个方面进行判断，越趋近3的网站，误报越小
	result :=  Lenbody(len(body)) +Resbodyb(string(body))
	if result < 2 {
		str = false
	}else {
		str = true
	}
	return str
}

// Fhundreds 判断500的页面
func Fhundreds(Turl string) bool{
	_ , body := JumpUrl(Turl)
	var str bool
	//对页面的长度进行判断，越趋近2的网站，误报越小
	result :=   Lenbody(len(body))
	if result < 1 {
		str = false
	}else {
		str = true
	}
	return str
}

func JumpUrl(Turl string)  (int,[]byte){
	resp := GETRequest(Turl)
	var respCode int
	var bodys  = []byte(nil)
	if resp != nil {
		respCode = resp.StatusCode //状态码
		bodys, _ = ioutil.ReadAll(resp.Body)
	}

	return respCode, bodys
}

func ResCode(respCode int) int {
	var str int
	if respCode == 200 {
		str = 1
	}else {
		str = 0
	}
	return str
}

func Lenbody(body int) int {
	var str int
	if body >200 && body <1000 {
		str =1
	}else if body >1000{
		str =2
	}else {
		str =0
	}
	return str
}

func Resbody(body string) int {
	var str int
	if strings.ContainsAny(body,"页面不存在")  {
		str = 0
	}else {
		str = 1
	}
	//fmt.Println("a",str)
	return str
}

func Resbodya(body string) int {
	var str int
	if strings.ContainsAny(body,"重新登录")  {
		str = 0
	}else {
		str = 1
	}
	//fmt.Println("a",str)
	return str
}

func Resbodyb(body string) int {
	var str int
	if strings.ContainsAny(body,"403 Forbidden")  {
		str = 0
	}else {
		str = 1
	}
	//fmt.Println("a",str)
	return str
}
