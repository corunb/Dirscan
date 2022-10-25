package config

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Bar struct {
	mu      sync.Mutex
	graph   string    // 显示符号
	rate    string    // 进度条
	percent int       // 百分比
	current int       // 当前进度位置
	total   int       // 总进度
	start   time.Time // 开始时间
}

// NewBar 初始化
func (bar *Bar) NewBar(current, total int) *Bar {
	//bar := new(Bar)
	bar.current = current
	bar.total = total
	bar.start = time.Now()
	if bar.graph == "" {
		bar.graph = "█"
	}
	bar.percent = bar.getPercent()
	for i := 0; i < bar.percent; i += 2 {
		bar.rate += bar.graph //初始化进度条位置
	}
	return bar
}

//func NewBarWithGraph(start, total int, graph string) *Bar {
//	bar := NewBar(start, total)
//	bar.graph = graph
//	return bar
//}

//计算当前百分比
func (bar *Bar) getPercent() int {
	return int((float64(bar.current) / float64(bar.total)) * 100)
}

//获取当前花费时间
func (bar *Bar) getTime() (s string) {
	u := time.Now().Sub(bar.start).Seconds()
	h := int(u) / 3600
	m := int(u) % 3600 / 60
	if h > 0 {
		s += strconv.Itoa(h) + "h "
	}
	if h > 0 || m > 0 {
		s += strconv.Itoa(m) + "m "
	}
	s += strconv.Itoa(int(u)%60) + "s"
	return
}


//加载进度条
func (bar *Bar) load() {
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}
	//fmt.Printf("\r [%-40s]% 3d%%    %2s   %d/%d", bar.rate, bar.percent, bar.getTime(), bar.current, bar.total)
	fmt.Printf("\r  [%d/%d] %3d%%  ",bar.current, bar.total,bar.percent)
}

func (bar *Bar) Reset(current int) {
	bar.mu.Lock()
	defer bar.mu.Unlock()
	bar.current = current
	bar.load()

}

// Add 设置进度
func (bar *Bar) Add(i int) {
	bar.mu.Lock()
	defer bar.mu.Unlock()
	bar.current += i
	bar.load()
}

func (bar *Bar) Clear() {
	//_ = fmt.Sprintf("\r%s\r", " ")
	fmt.Printf( "\r ")
	// 去除换行符
	//str = strings.Replace(str, "\n", "", -1)
	//fmt.Printf( "\\033[5B\r")
}

func (bar *Bar) Close() {
	fmt.Println(" ")
}