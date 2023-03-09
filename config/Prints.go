package config

import (
	"fmt"
	"github.com/gookit/color"
	"strconv"
)

// GetPrint  结果显示
func GetPrint(respCode int,Bodylen string ,url string,path string,Rurl string) {

	if respCode >= 200 && respCode < 300 && Thundred(url + path) == true {
		Urlpath := url+path
		color.Green.Printf("\r%v     [%v]     %v    \t- %v \n", Time,respCode,Bodylen,Urlpath )
		Write(Time+"     "+"["+strconv.Itoa(respCode)+"]"+" ["+ Bodylen +"]    "+Urlpath,url)
	} else if respCode >= 300 && respCode < 400  && Redirect(Rurl) == true{
		color.Yellow.Printf("\r%v     [%v]     %v    \t- %v --> %v \n", Time,respCode, Bodylen, path, Rurl     )
		Write(Time+"     "+"["+strconv.Itoa(respCode)+"]"+"     "+ path +"     "+ Rurl,url)
	} else if respCode >= 400 && respCode < 500  && Thundreds(url + path) == true{
		if respCode == 403 {
			color.Red.Printf("\r%v     [%v]     %v    \t- %v \n", Time,respCode, Bodylen, path     )
			Write(Time+"     "+"["+strconv.Itoa(respCode)+"]"+"     "+path,url)
		}else {
			color.Blue.Printf("\r%v     [%v]     %v    \t- %v \n", Time,respCode, Bodylen, path     )
		}
	} else if respCode >= 500 && respCode < 600 && Fhundreds(url + path) == true {
		color.Cyan.Printf("\r%v     [%v]     %v    \t- %v \n", Time, respCode, Bodylen, path)
	}
	//} else {
	//	fmt.Printf( "\r%v     [%v]     %v    \t- %v \n", Time,respCode, Bodylen, path)
	//}



}

// HeadPrint 结果显示
func HeadPrint(respCode int ,url string,path string,Rurl string) {

	if respCode >= 200 && respCode < 300 && Thundred(url + path) == true{
		Urlpath := url + path
		color.Green.Printf("\r%v  [%v] - %v \n", Time, respCode, Urlpath)
		Write(Time+"     "+"["+strconv.Itoa(respCode)+"]"+"     "+Urlpath, url)
	} else if respCode >= 300 && respCode < 400 && Redirect(Rurl) == true{
		color.Yellow.Printf("\r%v  [%v] - %v  --> %v \n", Time, respCode, path, Rurl)
		Write(Time+"     "+"["+strconv.Itoa(respCode)+"]"+"     "+path+"     "+Rurl, url)
	} else if respCode >= 400 && respCode < 500 && Thundreds(url + path) == true{
		if respCode == 403 {
			color.Red.Printf("\r%v  [%v] - %v  \n", Time, respCode, path)
			Write(Time+"     "+"["+strconv.Itoa(respCode)+"]"+"     "+path, url)
		}
		color.Blue.Printf("\r%v  [%v] - %v  \n", Time, respCode, path)
	} else if respCode >= 500 && respCode < 600 && Fhundreds(url + path) == true {
		color.Cyan.Printf("\r%v  [%v] - %v  \n", Time, respCode, path)
	}
	//} else {
	//	fmt.Printf("\r%v  [%v] - %v  \n", Time, respCode, path)
	//}
}


// Tishi 提示
func Tishi() {
	////提示
	fmt.Println("----------------------------------------")
	color.Red.Printf("[+] 字典:%v \n",Pathfile)
	color.Red.Printf("[+] 字典数:%v \n",len(Typeselection()))
	color.Green.Printf("[+] 线程:%v \n",Threads)
	color.Green.Printf("[+] 目标:%v \n",Url)
	fmt.Println("----------------------------------------")
}
// Tishis 提示
func Tishis() {
	////提示
	fmt.Println("----------------------------------------")
	color.Red.Printf("[+] 字典:%v \n",Pathfile)
	color.Red.Printf("[+] 字典数:%v \n",len(Typeselection()))
	color.Green.Printf("[+] 线程:%v \n",Threads)
	color.Green.Printf("[+] 目标数量:%v \n",len(ReadFile(Urlfile)))
	fmt.Println("----------------------------------------")
}