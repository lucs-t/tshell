package cmd

import (
	sshinfo "github.com/lucs-t/tshell/cmd/Flags/sshinfo"
	"github.com/lucs-t/tshell/cmd/flags/s3config"
)

type Flag interface {
	Parse(string)error
	GetData() map[string]string
}

var _  Flag = &sshinfo.SSHInfo{}
var _  Flag = &s3config.S3Config{}

type FlagManager struct {
	Flags []Flag 
}
func NewFlagManager() *FlagManager {
	flags := []Flag{}
	flags = append(flags, sshinfo.NewSSHInfo())
	flags = append(flags, s3config.NewS3Config())
	return &FlagManager{Flags: flags}
}

func (f *FlagManager) Parse(arg string) error{
	for _, flag := range f.Flags {
		err := flag.Parse(arg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *FlagManager) GetData() map[string]string {
	data := make(map[string]string)
	for _, flag := range f.Flags {
		for k, v := range flag.GetData() {
			data[k] = v
		}
	}
	return data
}
