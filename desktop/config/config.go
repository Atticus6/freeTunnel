package config

import (
	"log"
	"os"
	"path/filepath"
)

var StoreDir string

func init() {
	homeDir, err2 := os.UserHomeDir()
	if err2 != nil {
		log.Fatal("无法获取用户目录:", err2)
	}
	StoreDir = filepath.Join(homeDir, ".freeTunnel")

}
