package pkg

import "strconv"

// Configure any server settings, ie if you want to host this somewhere
// You could change this to have the URL or a secret or something
type ServerConfig struct {
	debugLevel string
	port       int
}

func (cfg *ServerConfig) GetPort() string {
	return strconv.Itoa(cfg.port)
}
