package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// clearScreen 用于清除终端屏幕
var clearScreen = make(map[string]func())

func init() {
	clearScreen["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clearScreen["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clearScreen["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	// 动画帧
	dogFrames := []string{
		`
    __
o-''|\_____/)
 \_/|_)     )
    \  __  /
    (_/ (_/
`,
		`
    __
o-''|\____/)
 \_/|_)    )
    \  __ /
    (_/ (_/
`,
		`
    __
o-''|\___/)
 \_/|_)   )
    \  __
    (_/ (_
`,
		`
    __
o-''|\__/)
 \_/|_)  )
    \  _
    (_/ (
`,
		`
    __
o-''|\_/)
 \_/|_) )
    \
    (_/
`,
		`
    __
o-''|\/)
 \_/|_)
    
    (_
`,
	}

	// 动画循环
	for {
		for i := 0; i < len(dogFrames); i++ {
			callClear()
			fmt.Println("看，一只可爱的小狗在摇尾巴！ (按 Ctrl+C 退出)")
			fmt.Println(dogFrames[i])
			time.Sleep(150 * time.Millisecond)
		}
		for i := len(dogFrames) - 2; i > 0; i-- {
			callClear()
			fmt.Println("看，一只可爱的小狗在摇尾巴！ (按 Ctrl+C 退出)")
			fmt.Println(dogFrames[i])
			time.Sleep(150 * time.Millisecond)
		}
	}
}

// callClear 根据操作系统调用相应的清屏函数
func callClear() {
	value, ok := clearScreen[runtime.GOOS]
	if ok {
		value()
	} else { // 对于未知的操作系统
		panic("你的平台不受支持！无法清除屏幕。")
	}
}
