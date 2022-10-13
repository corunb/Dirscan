package main

import (
	"Dirscan/config"
	"fmt"
	"github.com/gookit/color"
	"time"
)




func main() {
	start := time.Now()


	if config.Url == "" && config.Urlfile == ""{
		fmt.Println("请输入-h查看帮助！")
	}else if config.Url != "" {
		//进行单url扫描
		config.Tishi()
		url := config.Urll(config.Url)
		config.Scans(url)
	}else if config.Urlfile != "" {
		config.Tishis()
		config.Scanes()
	}



	end := time.Since(start)
	color.Cyan.Println("运行时间为:", end)
}
