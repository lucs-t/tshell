package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	var pwdInfo, keyInfo string

	// 定义 -p 和 -k 选项
	addCmd.StringVar(&pwdInfo, "p", "", "user:password:host:port for ssh connection")
	addCmd.StringVar(&keyInfo, "k", "", "user:keyPath:host:port for ssh connection")

	// 检查是否提供了 "add" 子命令
	if len(os.Args) < 2 {
		fmt.Println("Expected 'add' subcommand")
		os.Exit(1)
	}

	// 解析 "add" 子命令
	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
	default:
		fmt.Println("Expected 'add' subcommand")
		os.Exit(1)
	}

	// 处理 -p 选项
	if pwdInfo != "" {
		parts := strings.Split(pwdInfo, ":")
		if len(parts) == 4 {
			user := parts[0]
			password := parts[1]
			host := parts[2]
			port := parts[3]
			fmt.Printf("Adding SSH connection with password:\nUser: %s\nPassword: %s\nHost: %s\nPort: %s\n", user, password, host, port)
		} else {
			fmt.Println("Invalid format for -p. Expected format: user:password:host:port")
		}
	}

	// 处理 -k 选项
	if keyInfo != "" {
		parts := strings.Split(keyInfo, ":")
		if len(parts) == 4 {
			user := parts[0]
			keyPath := parts[1]
			host := parts[2]
			port := parts[3]
			fmt.Printf("Adding SSH connection with key:\nUser: %s\nKey Path: %s\nHost: %s\nPort: %s\n", user, keyPath, host, port)
		} else {
			fmt.Println("Invalid format for -k. Expected format: user:keyPath:host:port")
		}
	}

	// 如果没有提供任何参数
	if pwdInfo == "" && keyInfo == "" {
		fmt.Println("Please specify either -p or -k with correct format.")
	}
}
