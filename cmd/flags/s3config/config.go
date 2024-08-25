package s3config

import (
	"flag"
	"os"
	"strings"

	"github.com/lucs-t/tshell/utils"
)

type S3Config struct {
	FlagSets   map[string]*flag.FlagSet
	Config     map[string]string
	ak         string
	sk         string
	region     string
	bucket    string
	updatePath string
	endpoint  string
}

func NewS3Config() *S3Config {
	var s = S3Config{
		FlagSets: make(map[string]*flag.FlagSet),
		Config:   make(map[string]string),
	}
	s.add()
	return &s
}

func (s *S3Config) GetData() map[string]string {
	return s.Config
}

func (s *S3Config) Parse(arg string) error {
	if set, ok := s.FlagSets[arg]; ok {
		set.Parse(os.Args[3:])
		if arg == "configAdd" {
			if s.ak != "" {
				s.Config[utils.AK] = s.ak
			}else{
				return utils.Errorf("access key is required")
			}
			if s.sk != "" {
				s.Config[utils.SK] = s.sk
			}else{
				return utils.Errorf("access secret is required")
			}
			if s.region != "" {
				s.Config[utils.Region] = s.region
			}else{
				return utils.Errorf("region is required")
			}
			if s.endpoint != "" {
				s.Config[utils.Endpoint] = s.endpoint
			}else{
				return utils.Errorf("endpoint is required")
			}
			if s.updatePath != "" {
				ss := strings.Split(s.updatePath, ":")
				if len(ss) != 2 {
					return utils.Errorf("update path format is wrong")
				}
				s.Config[utils.Bucket] = ss[0]
				s.Config[utils.UpdatePath] = ss[1]
			}else{
				return utils.Errorf("update path is required")
			}
		}
	}
	return nil
}

func (s *S3Config) add() {
	s.FlagSets["configAdd"] = flag.NewFlagSet("add", flag.ExitOnError)
	s.FlagSets["configAdd"].Usage = utils.ConfigAddUsage(s.FlagSets["configAdd"])
	s.FlagSets["configAdd"].StringVar(&s.ak, "k", "", "access key")
	s.FlagSets["configAdd"].StringVar(&s.sk, "s", "", "access secret")
	s.FlagSets["configAdd"].StringVar(&s.region, "r", "", "region")
	s.FlagSets["configAdd"].StringVar(&s.updatePath, "p", "", "format: bucket:updatePath")
	s.FlagSets["configAdd"].StringVar(&s.endpoint, "e", "", "endpoint")
}
