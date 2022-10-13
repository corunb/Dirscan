package config

import (
	"bufio"
	"fmt"
	"github.com/gookit/color"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

////颜色红绿黄青蓝
//var red = color.FgRed.Render
//var green = color.FgGreen.Render
//var Yellow = color.Yellow.Render
//var Cyan = color.HiCyan.Render
//var Blue = color.Blue.Render


var Time = time.Now().Format("2006/01/02 15:04:05")

// Codel code 输入范围转换为数组
func Codel(code string) []int{
	var codes []int
	codeArr := strings.Split(strings.Trim(code, ","), ",")
	for _, v := range codeArr {
		codeArr2 := strings.Split(strings.Trim(v, "-"), "-")
		startPort, err := filterCode(codeArr2[0])
		if err != nil {
			fmt.Println(err)
		}

		codes = append(codes, startPort)
		if len(codeArr2) > 1 {

			endPort, _ := filterCode(codeArr2[1])
			if endPort > startPort {
				for i := 1; i <= endPort-startPort; i++ {
					codes = append(codes, startPort+i)
				}
			}
		}
	}
	//去重复
	codes = arrayUnique(codes)
	return codes
	//fmt.Println(ports)
}

//转换strng为int
func filterCode(str string) (int, error) {
	code, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	//if port < 1 || port > 600 {
	//	return 0, errors.New("端口号范围超出")
	//}
	return code, nil
}

//去重
func arrayUnique(arr []int) []int {
	var newArr []int
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}




// Storage 转换大小
func Storage(num int) string{
	a := num / 1024
	var b int
	var c string
	switch {

	case a == 0:
		//fmt.Print(num,"B")
		c = strconv.Itoa(num) +"B"
	case a >=1 && a < 1024:
		//fmt.Print(a,"KB")
		c = strconv.Itoa(a) +"KB"
	case a >=1024 && a < 1048576:
		b = a / 1024
		//fmt.Print(b,"MB")
		c = strconv.Itoa(b) +"MB"
	case a >=1048576 && a < 1073741824:
		b = (a/1024) / 1024
		//fmt.Print(b,"GB")
		c = strconv.Itoa(b) +"GB"
	}
	return c
}


// Urll   url后面加去除/
func Urll(url string) string {

	a := url[len(url)-1:]
	if a == "/" {
		url = strings.TrimRight(url,a)
	}
	//fmt.Println(url)
	return url
}

// IsPath 判断字典是目录还是文件
func IsPath(pathname string) bool{
	//如果文件名和文件前缀相同，则是目录
	filenameall := path.Base(pathname)	 //获取不包含目录的文件名
	filesuffix := path.Ext(pathname)	//获取文件后缀
	fileprefix := filenameall[0:len(filenameall) - len(filesuffix)] //文件前缀


	if fileprefix == filenameall{ //文件名和文件前缀相同为目录，否则为文件
		//fmt.Println(pathname,"这是目录")
		return true
	}else {
		//fmt.Println(pathname," 文件")
		return false
	}
}

// Urladdress url地址获取
func Urladdress(url string) string{
	a1 := strings.Split(url, "//")[1]
	a2 := strings.Split(a1, "/")[0]
	a3 := strings.Replace(a2, ".", "_", -1)
	a4 := strings.Replace(a3, ":", "_", -1)
	return a4
	//fmt.Println(a4)

}

// CodeIstrue codeIstrue 状态码输入检测
func CodeIstrue(intcode []int) {
	for _,v := range intcode {
		if v <100 || v> 600 {
			fmt.Println("状态码输入错误，请输入正确状态码")
			os.Exit(0)
		}
	}
}

// InitConfig 读取配置文件
func InitConfig(path string) map[string]string {
	config := make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config
}

// Tishi 提示
func Tishi() {
	////提示
	fmt.Println("----------------------------------------")
	color.Red.Printf("[+] 字典:%v \n",Pathfile)
	color.Red.Printf("[+] 字典数:%v \n",len(ReadFile(Pathfile)))
	color.Green.Printf("[+] 线程:%v \n",Threads)
	color.Green.Printf("[+] 目标:%v \n",Url)
	fmt.Println("----------------------------------------")
}
// Tishis 提示
func Tishis() {
	////提示
	fmt.Println("----------------------------------------")
	color.Red.Printf("[+] 字典:%v \n",Pathfile)
	color.Red.Printf("[+] 字典数:%v \n",len(ReadFile(Pathfile)))
	color.Green.Printf("[+] 线程:%v \n",Threads)
	color.Green.Printf("[+] 目标数量:%v \n",len(ReadFile(Urlfile)))
	fmt.Println("----------------------------------------")
}

//求交集
func intersect(slice1, slice2 []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}
//求差集 slice1-并集
func difference(slice1, slice2 []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
	inter := intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

// GetPrint  结果显示
func GetPrint(respCode int,Bodylen string ,url string,path string) {
	if respCode >= 200 && respCode < 300 {
		Urlpath :=url+path
		color.Green.Printf("%v\t[%v]\t%v\t - %v \n", Time,respCode,Bodylen,Urlpath)
		Write(Time+"\t"+"["+strconv.Itoa(respCode)+"]"+"\t"+Urlpath,url)
	} else if respCode >= 300 && respCode < 400 {
		color.Yellow.Printf("%v\t[%v]\t%v\t - %v %v \n", Time,respCode, Bodylen, path, Redirect)
		Write(Time+"\t"+"["+strconv.Itoa(respCode)+"]"+"\t"+ path +"\t"+ Redirect,url)
	} else if respCode >= 400 && respCode < 500 {
		if respCode == 403 {
			color.Red.Printf("%v\t[%v]\t%v\t - %v \n", Time,respCode, Bodylen, path)
			Write(Time+"\t"+"["+strconv.Itoa(respCode)+"]"+"\t"+path,url)
		}else {
			color.Blue.Printf("%v\t[%v]\t%v\t - %v \n", Time,respCode, Bodylen, path)
		}

	} else if respCode >= 500 && respCode < 600 {
		color.Cyan.Printf("%v\t[%v]\t%v\t - %v \n", Time,respCode, Bodylen, path)
	} else {
		fmt.Printf("%v\t[%v]\t%v\t - %v \n", Time,respCode, Bodylen, path)
	}
}

// HeadPrint 结果显示
func HeadPrint(respCode int ,url string,path string) {
	if respCode >= 200 && respCode < 300 {
		Urlpath := url + path
		color.Green.Printf("%v - [%v] - %v \n", Time, respCode, Urlpath)
		Write(Time+"\t"+"["+strconv.Itoa(respCode)+"]"+"\t"+Urlpath, url)
	} else if respCode >= 300 && respCode < 400 {
		//Urlpath := url + path
		color.Yellow.Printf("%v- [%v] - %v  %v \n", Time, respCode, path, Redirect)
		Write(Time+"\t"+"["+strconv.Itoa(respCode)+"]"+"\t"+path+"\t"+Redirect, url)
	} else if respCode >= 400 && respCode < 500 {
		if respCode == 403 {
			color.Red.Printf("%v- [%v] - %v  \n", Time, respCode, path)
			Write(Time+"\t"+"["+strconv.Itoa(respCode)+"]"+"\t"+path, url)
		}
		color.Blue.Printf("%v- [%v] - %v  \n", Time, respCode, path)
	} else if respCode >= 500 && respCode < 600 {
		color.Cyan.Printf("%v- [%v] - %v  \n", Time, respCode, path)
	} else {
		fmt.Printf("%v- [%v] - %v  \n", Time, respCode, path)
	}
}


// RemoveRepByLoop 通过两重循环过滤重复元素数组去重
func RemoveRepByLoop(slc []string) []string {
	var result []string // 存放结果
	for i := range slc{
		flag := true
		for j := range result{
			if slc[i] == result[j] {
				flag = false  // 存在重复元素，标识为false
				break
			}
		}
		if flag {  // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// Recursionchoose 递归扫描的数据存储
func Recursionchoose(respCode int ,url string,path string) {
	if (respCode == 200 || respCode == 301  || respCode == 302 || respCode == 403) && IsPath(url+path) == true  {
		//if  {
		//color.Green.Println(path, "标记")

		Urlpath := Urll(url+path)
		BiaoJi = append(BiaoJi, Urlpath)
		//fmt.Println(urlChan)
		//Recursionscan(url+path)
		//}
	}
}