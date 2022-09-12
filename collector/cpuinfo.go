package collector

import (
	"strings"
)

var Cpufilepath="/cpuinfo"
const CPUSubsystem = "cpuinfo"

func init(){
	Registry(CPUSubsystem,GetCpuinfo())
}

type Cpuinfo struct {
}

func GetCpuinfo()Collectors{
	return Cpuinfo{}
}

func (c Cpuinfo)GetData(dirpath string)map[string][]string{
	cpuinfo:=make(map[string][]string)
	filepath:=dirpath+Cpufilepath
	fileHandle:=ReadFile(filepath)
	delSpace:=strings.ReplaceAll(fileHandle," ","")
	delLine:=strings.Split(delSpace,"\n")
	for _,v:=range delLine{
		args:=strings.Split(v,":")
		if len(args)==2&&args[1]!=""{
			cpuinfo[args[0]]=args[1:]
		}
	}
	return cpuinfo
}
