package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

//启动Go程序的方法是在终端中导航到包含main包的目录，然后运行以下命令：
//
//go run t.go
//
//这将编译并运行Go程序，并在终端中输出"Hello, World!"。

//另外，你也可以先编译程序，然后运行生成的可执行文件：
//
//go build t.go
//./t   （在Windows上是t.exe）
//
//这两种方法都可以用来启动Go程序。

//确保你已经安装了Go编程语言，并且正确配置了环境变量GOPATH和GOROOT。
//你可以在终端中运行"go version"来检查Go是否已正确安装。
