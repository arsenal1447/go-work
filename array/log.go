package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// log.Println("这是日志")
	// v := "很普通的"
	// log.Printf("这是一条%s日志\n", v)
	// log.Fatalln("这是一条会触发fatal日志")
	// log.Panicln("这是一条会触发 panic 日志")

	//logger会打印每条日志信息的日期、时间，默认输出到系统的标准错误。
	// Fatal系列函数会在写入日志信息后调用os.Exit(1)。
	// Panic系列函数会在写入日志信息后panic。

	//log标准库中的Flags函数会返回标准logger的输出配置，
	//而SetFlags函数用来设置标准logger的输出配置。

	// 控制输出日志信息的细节，不能控制输出的顺序和格式。

	// log标准库提供了如下的flag选项，它们是一系列定义好的常量。
	// // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	// Ldate         = 1 << iota     // 日期：2009/01/23
	// Ltime                         // 时间：01:23:23
	// Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
	// Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
	// Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
	// LUTC                          // 使用UTC时间
	// LstdFlags     = Ldate | Ltime // 标准logger的初始值

	//2021/04/16 17:48:17.720643 /Users/xxxxx/Desktop/mywork/gowork/test/log.go:20:
	// log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	// log.Println("\n 这是一条很普通的日志")

	// log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	// log.Println("\n 这是一条很普通的日志")

	// log标准库中还提供了关于日志信息前缀的两个方法：
	// func Prefix() string
	// func SetPrefix(prefix string)
	// 其中Prefix函数用来查看标准logger的输出前缀，
	// SetPrefix函数用来设置输出前缀。
	//[小王子]2021/04/16 17:50:23.477034 /Users/zhouxinxin/Desktop/mywork/gowork/test/log.go:39:
	// log.SetPrefix("[小王子]")
	// log.Println("\n 这是一条很普通的日志")

	//配置日志输出位置
	// func SetOutput(w io.Writer)

	// 例如，下面的代码会把日志输出到同目录下的xx.log文件中。
	// logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	// if err != nil {
	// 	fmt.Println("open log file failed ,err:", err)
	// 	return
	// }

	// log.SetOutput(logFile)
	// log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	// log.Println("\n 这是一条很普通的日志")
	// log.SetPrefix("[小王子]")
	// log.Println("\n 这是一条很普通的日志")

	// log标准库中还提供了一个创建新logger对象的构造函数–New，
	// 支持我们创建自己的logger示例。New函数的签名如下：
	// func New(out io.Writer, prefix string, flag int) *Logger
	// New创建一个Logger对象。其中，参数out设置日志信息写入的目的地。
	// 参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）。

	logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ldate|log.Ltime)

	//运行结果：<New>2021/04/16 19:14:24 log.go:75: 这是自定义的logger记录的日志
	logger.Println("这是自定义的logger记录的日志")

}

// 如果你要使用标准的logger，我们通常会把上面的配置操作写到init函数中。
func init() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("open log file failed ,err:", err)
		return
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("\n 这是一条很普通的日志")
	log.SetPrefix("[小王子]")
	log.Println("\n 这是一条很普通的日志")
}
