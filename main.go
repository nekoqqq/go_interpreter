package main

import (
	"fmt"
	"go_interpreter/repl"
	"log"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatalf("用户获取失败, err: %v", err)
	}

	fmt.Printf("尊敬的:%s, 欢迎使用本解释器!请输入任意命令\n", u.Username)
	repl.Serve(os.Stdin, os.Stdout)

}
