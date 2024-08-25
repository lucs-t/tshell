package main

import (
	"fmt"
	"os"

	"github.com/lucs-t/tshell/cmd"
	"github.com/lucs-t/tshell/cmd/utils"
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
	case "config":
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
			fmt.Println("config remove")
		}else if os.Args[2] == "info" {
			// todo
			fmt.Println("config info")
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
