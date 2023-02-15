package main

import (
	"Dirscan/config"
	"fmt"
	"github.com/gookit/color"
	"time"
)




func main() {
	start := time.Now()


	if config.Url != "" && config.Urlfile == ""{
		//进行单url扫描
		config.Tishi()
		Turl := config.Urll(config.Url)
		if config.FindUrl(Turl) == true {
			config.Scans(Turl)
		}

	}else if config.Urlfile != "" && config.Url == ""{
		config.Tishis()
		config.Scanes()
	}else {
		fmt.Println("请输入-h查看帮助！")
	}



	end := time.Since(start)

	color.HiGreen.Printf("\n运行时间为: %v\n", end)
}
