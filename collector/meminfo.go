package collector

import (
	"strings"
)

var Memfilepath="/meminfo"
const MemSubsystem = "meminfo"

func init(){
	Registry(MemSubsystem,GetMeminfo())
}

type Meminfo struct {
}

func GetMeminfo()Collectors{
	return Meminfo{}
}

func (m Meminfo)GetData(dirpath string)map[string][]string{
	meminfo:=make(map[string][]string)
	filepath:=dirpath+Memfilepath
	fileHandle:=ReadFile(filepath)
	delSpace:=strings.ReplaceAll(fileHandle," ","")
	delLine:=strings.Split(delSpace,"\n")
	//fmt.Println(delLine)
	for _,v:=range delLine{
		args:=strings.Split(v,":")
		if len(args)==2&&args[1]!=""{
			meminfo[args[0]]=args[1:]
		}
	}
	return meminfo
}
