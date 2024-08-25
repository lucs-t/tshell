package sshinfo

import (
	"flag"
	"fmt"

	"os"
	"strconv"
	"strings"

	"github.com/lucs-t/tshell/utils"
)

type SSHInfo struct {
	FlagSets map[string]*flag.FlagSet
	Info map[string]string
	sshCnn string
	password string
	keyPath string
	sshName string
	removeAll bool
}

func NewSSHInfo() (*SSHInfo) {
	var s = SSHInfo{
		FlagSets: make(map[string]*flag.FlagSet),
		Info: make(map[string]string),
	}
	s.add()
	s.remove()
	return &s
}

func (s *SSHInfo) GetData() map[string]string {
	return s.Info
}


func (s *SSHInfo) Parse(arg string) error{
	if set,ok := s.FlagSets[arg]; ok {
		set.Parse(os.Args[2:])
		if s.sshName != "" {
			s.Info[utils.SshInfoName] = s.sshName
		}
		if s.keyPath != "" {
			s.Info[utils.KeyPath] = s.keyPath
		}
		if s.password != "" {
			s.Info[utils.Password] = s.password
		}
		if s.sshCnn != "" {
			var isvalid bool
			info := strings.Split(s.sshCnn, ":")
			if len(info) == 3 {
				s.Info[utils.User] = info[0]
				s.Info[utils.Host] = info[1]
				s.Info[utils.Port] = info[2]
				if s.Info[utils.Port] == "" {
					s.Info[utils.Port] = "22"
				}
				_,err := strconv.Atoi(s.Info[utils.Port])
				if err != nil {
					return utils.Errorf("invalid port\n%s", err.Error())
				}
				isvalid = true
			}
			if len(info) == 2 {
				s.Info[utils.User] = info[0]
				s.Info[utils.Host] = info[1]
				s.Info[utils.Port] = "22"
				isvalid = true
			}
			if !isvalid {
				return utils.Errorf("invalid argument,format: user:host:port")
			}
			if s.sshName == "" {
				s.Info[utils.SshName] = s.Info[utils.Host]
			}
		}
		if s.removeAll {
			s.Info[utils.RemoveAll] = fmt.Sprintf("%t", s.removeAll)
		}
		if arg == "add" {
			if s.Info[utils.User] == "" {
				return utils.Errorf("user is required")
			}
			if s.Info[utils.Host] == "" {
				return utils.Errorf("host is required")
			}
		}
		if arg == "remove" {
			if s.Info[utils.SshName] == "" && !s.removeAll {
				return utils.Errorf("-n or --all is required")
			}
		}
	}
	return nil
}


func (s *SSHInfo) add() {
	s.FlagSets["add"] = flag.NewFlagSet("add", flag.ExitOnError)
	s.FlagSets["add"].Usage = utils.AddUsage(s.FlagSets["add"])
	s.FlagSets["add"].StringVar(&s.sshCnn, "u", "", fmt.Sprintf("%s for ssh connection", utils.SshUrl))
	s.FlagSets["add"].StringVar(&s.password, "p", "", "password for ssh connection")
	s.FlagSets["add"].StringVar(&s.keyPath, "k", "", "key path for ssh connection")
	s.FlagSets["add"].StringVar(&s.sshName, "n", "", "ssh connection name,if not set,use host as name")
}

func (s *SSHInfo) remove() {
	s.FlagSets["remove"] = flag.NewFlagSet("remove", flag.ExitOnError)
	s.FlagSets["remove"].Usage = utils.RemoveUsage(s.FlagSets["remove"])
	s.FlagSets["remove"].StringVar(&s.sshName, "n", "", "ssh connection name")
	s.FlagSets["remove"].BoolVar(&s.removeAll,"all", false, "remove all ssh connection")
}

