package params

import (
	"errors"
	"flag"
	"fmt"
)

type Params struct {
	IpAddr     string
	Port       string
	Timeout    int
	FileConfig string
}

var (
	FileConfig  = flag.String("fileConfig", "", "define the file with config")
	Port        = flag.String("port", "8080", "Define port")
	Timeout     = flag.Int("timeout", 10, "Define timeout")
	HelpMessage = `This is the help message, please use flags properly:
	--help: for this help message
	--fileConfig=fileName: the path to the config file
	--port: the port number
	--timeout: the defined timeout for server`
	ErrNoConfigFile = errors.New("No config file provided")
)

func Init() (*Params, error) {
	flags := Params{}
	flag.Parse()
	fmt.Println(*FileConfig)
	if *FileConfig == "" {
		return nil, ErrNoConfigFile
	}
	flags.Port = *Port
	flags.Timeout = *Timeout
	flags.FileConfig = *FileConfig
	return &flags, nil
}
