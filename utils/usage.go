package utils

import (
	"flag"
	"fmt"
)

func FlagUsage() {
	fmt.Println("Usage: tshell [command]")
	fmt.Println("Description:")
	fmt.Println("  A simple tool to manage SSH connections")
	fmt.Println("Commands:")
	fmt.Println("  cnn     Connect to a SSH connection")
	fmt.Println("  add     Add a new SSH connection")
	fmt.Println("  remove  Remove a SSH connection")
	fmt.Println("  show    Show all SSH connection")
	fmt.Println("  s3      Synchronize ssh connection information to s3")
}


func AddUsage(f *flag.FlagSet) func() {
	return func() {
		fmt.Println("Usage: tshell add [options]")
		fmt.Println("Description:")
		fmt.Println("  Add a new SSH connection to the list\n")
		fmt.Println("Example:")
		fmt.Println("  # Add a new SSH connection with password")
		fmt.Println(fmt.Sprintf("  tshell add -u user:host:port -p 123456")+"\n")
		fmt.Println("  # Add a new SSH connection with key")
		fmt.Println(fmt.Sprintf("  tshell add -u user:host:port -k /path/to/key")+"\n")
		fmt.Println("Options:")
		f.PrintDefaults()
	}
}

func RemoveUsage(f *flag.FlagSet) func() {
	return func() {
		fmt.Println("Usage: tshell remove [options]")
		fmt.Println("Description:")
		fmt.Println("  Remove a SSH connection from the list\n")
		fmt.Println("Example:")
		fmt.Println("  # Remove a SSH connection")
		fmt.Println("  tshell remove -n sshName\n")
		fmt.Println("  # Remove all SSH connection")
		fmt.Println("  tshell remove --all\n")
		fmt.Println("Options:")
		f.PrintDefaults()
	}
}

func ConfigUsage() {
	fmt.Println("Usage: tshell config [options]")
	fmt.Println("Description:")
	fmt.Println("  Configure the tshell")
	fmt.Println("Options:")
	fmt.Println("  add    Add ssh configuration")
	fmt.Println("  remove Remove ssh configuration")
	fmt.Println("  info   Show ssh configuration info")
}

func ConfigAddUsage(f *flag.FlagSet) func() {
	return func() {
		fmt.Println("Usage: tshell config add [options]")
		fmt.Println("Description:")
		fmt.Println("  Add a new configuration to the list\n")
		fmt.Println("Example:")
		fmt.Println("  # Add a new configuration")
		fmt.Println("  tshell config add -ak accessKey -sk accessSecret -region region -path bucket:updatePath\n")
		fmt.Println("Options:")
		f.PrintDefaults()
	}
}