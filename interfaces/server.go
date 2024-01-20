package interfaces

import (
	"crypto/rsa"

	"znci.dev/lamp-v2/utils"
)

type Server interface {
	GetConfig() utils.ServerConfig
	GetPlayerCount() int
	GetKeyPair() *rsa.PrivateKey
}
