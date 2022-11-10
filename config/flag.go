package config

import (
	"flag"
	"github.com/gookit/color"
	"strconv"
)

var Url string
var Pathfile string
var Threads int
var Timeout int
var Recursion bool
var Rcode string
var Urlfile string
var Outfile string
var Requestmode string
var Neglect string
var Proxy string
var Sitetype string


func init() {
	//加载配置文件
	configs := InitConfig("./default/default.ini")
	intThreads, _ := strconv.Atoi(configs["Threads"])
	intTimeout, _ := strconv.Atoi(configs["Timeout"])

	flag.StringVar(&Sitetype, "m","","根据指定类型进行扫描，可设置php,asp,aspx,jsp")
	flag.StringVar(&Url, "u","","指定url")
	flag.StringVar(&Urlfile, "uf","","指定url列表")
	flag.StringVar(&Pathfile, "f",configs["Pathfile"],"指定目录字典")
	flag.StringVar(&Rcode, "i",configs["Rcode"],"筛选指定状态码,示例：200,403,404,500或者200-400")
	flag.StringVar(&Neglect, "ei",configs["Neglect"],"忽略指定状态码,示例：200,403,404,500或者200-400")
	flag.IntVar(&Threads,"T",intThreads,"设置线程，默认30")
	flag.IntVar(&Timeout,"t",intTimeout,"设置延时时间，默认5s")
	flag.StringVar(&Outfile, "o","","保存扫描结果,默认输出日期+地址")
	flag.StringVar(&Requestmode, "R",configs["Requestmode"],"指定G->Get扫描还是H->Head扫描")
	flag.BoolVar(&Recursion,"r",false,"进行递归扫描")
	flag.StringVar(&Proxy, "p","","proxy，可设置http代理或socks5代理，socks5://admin:corun@x.x.x.x:1080")
	flag.Parse()


	logo := `

 ██████████   █████ ███████████    █████████    █████████    █████████   ██████   █████
░░███░░░░███ ░░███ ░░███░░░░░███  ███░░░░░███  ███░░░░░███  ███░░░░░███ ░░██████ ░░███ 
 ░███   ░░███ ░███  ░███    ░███ ░███    ░░░  ███     ░░░  ░███    ░███  ░███░███ ░███ 
 ░███    ░███ ░███  ░██████████  ░░█████████ ░███          ░███████████  ░███░░███░███ 
 ░███    ░███ ░███  ░███░░░░░███  ░░░░░░░░███░███          ░███░░░░░███  ░███ ░░██████ 
 ░███    ███  ░███  ░███    ░███  ███    ░███░░███     ███ ░███    ░███  ░███  ░░█████ 
 ██████████   █████ █████   █████░░█████████  ░░█████████  █████   █████ █████  ░░█████
░░░░░░░░░░   ░░░░░ ░░░░░   ░░░░░  ░░░░░░░░░    ░░░░░░░░░  ░░░░░   ░░░░░ ░░░░░    ░░░░░ 

[+] code by Corun V1.4.1
[+] https://github.com/corunb/Dirscan
`
	color.HiGreen.Println(logo)


}


