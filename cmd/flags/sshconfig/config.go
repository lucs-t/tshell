package sshconfig

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/lucs-t/tshell/cmd/utils"
)

type SSHConfig struct {
	FlagSets   map[string]*flag.FlagSet
	Config     map[string]string
	ak         string
	sk         string
	region     string
	updatePath string
}

func NewSSHConfig() *SSHConfig {
	var s = SSHConfig{
		FlagSets: make(map[string]*flag.FlagSet),
		Config:   make(map[string]string),
	}
	s.add()
	return &s
}

func (s *SSHConfig) GetData() map[string]string {
	return s.Config
}

func (s *SSHConfig) Parse(arg string) error {
	if set, ok := s.FlagSets[arg]; ok {
		set.Parse(os.Args[3:])
		if arg == "configAdd" {
			if s.ak != "" {
				s.Config[utils.AK] = s.ak
			}else{
				fmt.Println("Error: access key is required")
				return errors.New("access key is required")
			}
			if s.sk != "" {
				s.Config[utils.SK] = s.sk
			}else{
				fmt.Println("Error: access secret is required")
				return errors.New("access secret is required")
			}
			if s.region != "" {
				s.Config[utils.Region] = s.region
			}else{
				fmt.Println("Error: region is required")
				return errors.New("region is required")
			}
			if s.updatePath != "" {
				s.Config[utils.UpdatePath] = s.updatePath
			}else{
				fmt.Println("Error: bucket:filerpath is required")
				return errors.New("bucket:filerpath is required")
			}
		}
	}
	return nil
}

func (s *SSHConfig) add() {
	s.FlagSets["configAdd"] = flag.NewFlagSet("add", flag.ExitOnError)
	s.FlagSets["configAdd"].Usage = utils.ConfigAddUsage(s.FlagSets["configAdd"])
	s.FlagSets["configAdd"].StringVar(&s.ak, "k", "", "access key")
	s.FlagSets["configAdd"].StringVar(&s.sk, "s", "", "access secret")
	s.FlagSets["configAdd"].StringVar(&s.region, "r", "", "region")
	s.FlagSets["configAdd"].StringVar(&s.updatePath, "p", "", "format: bucket:updatePath")
}
