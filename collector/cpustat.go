package collector

import (
	"strings"
)
//D:/appdata/gowork/src/test1/file
var Statfilepath="/stat"
const StatSubsystem="cpustat"

func init(){
	Registry(StatSubsystem,GetCpustat())
}

type Cpustat struct {
}

func GetCpustat()Collectors{
	return Cpustat{}
}

func (c Cpustat)GetData(dirpath string)map[string][]string{
	filepath:=dirpath+Statfilepath
	cpustat:=make(map[string][]string)
	fileHandle:=ReadFile(filepath)
	delLine:=strings.Split(fileHandle,"\n")
	for _,v:=range delLine{
		args:=strings.Split(v," ")
		if len(args)>=2{
			cpustat[args[0]]=args[1:]
		}
	}
	return cpustat
}
