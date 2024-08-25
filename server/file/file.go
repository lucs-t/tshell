package file

import (
	"encoding/json"
	"os"

	"github.com/lucs-t/tshell/models"
	"github.com/lucs-t/tshell/utils"
)

var sshInfoFile = "config/" + utils.SshInfoName
var sshConfigFile = "config/" + utils.SshConfigName

// WriteSSHinfoFiles 将新的 SSHInfo 写入 sshinfo.json 文件
func WriteSSHinfoFiles(info models.SSHInfo) error {
	// 检查文件是否存在
	err := CheckFile(sshInfoFile)
	if err != nil {
		return utils.Errorf("failed to check SSH info file: %w", err)
	}

	// 读取现有的 SSH 信息
	infos, err := ReadSSHinfoFiles()
	if err != nil {
		return utils.Errorf("failed to read SSH info files: %w", err)
	}

	// 将新信息追加到列表中
	infos = append(infos, info)

	// 打开文件以写入数据
	file, err := os.OpenFile(sshInfoFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return utils.Errorf("failed to open SSH info file: %w", err)
	}
	defer file.Close()

	// 将数据编码为 JSON 并写入文件
	err = json.NewEncoder(file).Encode(infos)
	if err != nil {
		return utils.Errorf("failed to encode SSH info: %w", err)
	}

	return nil
}

// ReadSSHinfoFiles 从 sshinfo.json 文件中读取 SSH 信息
func ReadSSHinfoFiles() ([]models.SSHInfo, error) {
	// 检查文件是否存在
	err := CheckFile(sshInfoFile)
	if err != nil {
		return nil, utils.Errorf("failed to check SSH info file: %w", err)
	}

	// 打开文件以读取数据
	file, err := os.Open(sshInfoFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.SSHInfo{}, nil
		}
		return nil, utils.Errorf("failed to open SSH info file: %w", err)
	}
	defer file.Close()

	// 解码 JSON 数据
	var infos []models.SSHInfo
	err = json.NewDecoder(file).Decode(&infos)
	if err != nil {
		return nil, utils.Errorf("failed to decode SSH info: %w", err)
	}

	return infos, nil
}

// WriteSSHconfigFiles 将新的 SSHConfig 写入 sshconfig.json 文件
func WriteSSHconfigFiles(config models.SSHConfig) error {
	// 检查文件是否存在
	err := CheckFile(sshConfigFile)
	if err != nil {
		return utils.Errorf("failed to check SSH config file: %w", err)
	}

	// 打开文件以写入数据
	file, err := os.OpenFile(sshConfigFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return utils.Errorf("failed to open SSH config file: %w", err)
	}
	defer file.Close()

	// 将数据编码为 JSON 并写入文件
	err = json.NewEncoder(file).Encode(config)
	if err != nil {
		return utils.Errorf("failed to encode SSH config: %w", err)
	}

	return nil
}

// ReadSSHconfigFiles 从 sshconfig.json 文件中读取 SSH 配置信息
func ReadSSHconfigFiles() (models.SSHConfig, error) {
	// 检查文件是否存在
	err := CheckFile(sshConfigFile)
	if err != nil {
		return models.SSHConfig{}, utils.Errorf("failed to check SSH config file: %w", err)
	}

	// 打开文件以读取数据
	file, err := os.Open(sshConfigFile)
	if err != nil {
		if os.IsNotExist(err) {
			return models.SSHConfig{}, nil
		}
		return models.SSHConfig{}, utils.Errorf("failed to open SSH config file: %w", err)
	}
	defer file.Close()

	// 解码 JSON 数据
	var config models.SSHConfig
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return models.SSHConfig{}, utils.Errorf("failed to decode SSH config: %w", err)
	}

	return config, nil
}

// CheckFile 检查路径是否存在，如果不存在则创建它
func CheckFile(path string) error {
	// 检查 config 目录是否存在
	if _, err := os.Stat("config"); os.IsNotExist(err) {
		err := os.Mkdir("config", os.ModePerm)
		if err != nil {
			return utils.Errorf("failed to create config directory: %w", err)
		}
	}else if err != nil {
		return utils.Errorf("failed to check config directory: %w", err)
	}

	return nil
}