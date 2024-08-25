package file

import (
	"encoding/json"
	"log"
	"os"

	"github.com/lucs-t/tshell/models"
)

var sshInfoFile = "config/sshinfo.json"
var sshConfigFile = "config/sshconfig.json"

func WriteSSHinfoFiles(info []models.SSHInfo) error{
	err := CheckDir()
	if err != nil {
		return err
	}
	if _,err := os.Stat(sshInfoFile); os.IsNotExist(err) {
		_,err := os.Create(sshInfoFile)
		if err != nil {
			log.Println("Error: ",err)
			return err
		}
	}else if err != nil {
		log.Println("Error: ",err)
		return err
	}
	file,err := os.OpenFile(sshInfoFile,os.O_RDWR,os.ModePerm)
	if err != nil {
		log.Println("Error: ",err)
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(info)
	if err != nil {
		log.Println("Error: ",err)
		return err
	}
	return nil
}

func ReadSSHinfoFiles() ([]models.SSHInfo,error){
	err := CheckDir()
	if err != nil {
		return nil,err
	}
	if _,err := os.Stat(sshInfoFile); os.IsNotExist(err) {
		_,err := os.Create(sshInfoFile)
		if err != nil {
			log.Println("Error: ",err)
			return nil,err
		}
	}else if err != nil {
		log.Println("Error: ",err)
		return nil,err
	}
	file,err := os.Open(sshInfoFile)
	if err != nil {
		log.Println("Error: ",err)
		return nil,err
	}
	defer file.Close()
	var info []models.SSHInfo
	err = json.NewDecoder(file).Decode(&info)
	if err != nil {
		log.Println("Error: ",err)
		return nil,err
	}
	return info,nil
}

func Checfile(path string) error{
	if _,err := os.Stat("config"); os.IsNotExist(err) {
		err := os.Mkdir("config",os.ModePerm)
		if err != nil {
			log.Println("Error: ",err)
			return err
		}
	}
	if _,err := os.Stat(sshInfoFile); os.IsNotExist(err) {
		_,err := os.Create(sshInfoFile)
		if err != nil {
			log.Println("Error: ",err)
			return err
		}
	}else if err != nil {
		log.Println("Error: ",err)
		return err
	}
	return nil
}