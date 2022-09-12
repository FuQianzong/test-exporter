package getargs

import (
	"flag"
	"os"
)

type Args struct {
	DirPath string
	Endpoint string
}

//获取命令行参数
func GetArgs() *Args{
	dirPathExegesis:="A directory that gets CPU and memory information.Container information is read by default"
	endpointExegesis:="Service port,default port 9031"
	dirPath := flag.String("dp", "/proc", dirPathExegesis)
	endpoint := flag.String("ep", "9031", endpointExegesis)
	help := flag.Bool("h", false, "help")
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	return &Args{
		DirPath: *dirPath,
		Endpoint: *endpoint,
	}
}