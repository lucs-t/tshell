package main

import (
	"fmt"
	"os"

	"github.com/lucs-t/tshell/cmd"
	"github.com/lucs-t/tshell/utils"
)

func main() {
	flagMgr := cmd.NewFlagManager()
	if len(os.Args) < 2 {
		utils.FlagUsage()
		return
	}
	var m = map[string]string{}
	switch os.Args[1] {
	case "add":
		err := flagMgr.Parse("add")
		if err != nil {
			return
		}
		m = flagMgr.GetData()
	case "remove":
		err := flagMgr.Parse("remove")
		if err != nil {
			return
		}
		m = flagMgr.GetData()
	case "show":
		//todo
		fmt.Println("show")
		return
	case "s3":
		if len(os.Args) < 3 {
			utils.ConfigUsage()
			return
		}
		if os.Args[2] == "add" {
			err := flagMgr.Parse("configAdd")
			if err != nil {
				return
			}
			m = flagMgr.GetData()
		}else if os.Args[2] == "remove" {
			// todo
			fmt.Println("s3 remove")
		}else if os.Args[2] == "info" {
			// todo
			fmt.Println("s3 info")
		}else if os.Args[2] == "-h" || os.Args[2] == "--help" {
			utils.ConfigUsage()
			return
		}else{
			fmt.Println("Error: Invalid Argument")
		}
	case "-h","--help":
		utils.FlagUsage()
		return
	default:
		utils.FlagUsage()
		return
	}
	fmt.Println(m)
}
