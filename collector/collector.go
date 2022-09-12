package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"sync"
)

var (
	f=make(map[string]Collectors)
	Dirpath string
	CPUField=[10]string{"user","nice","system","idle","iowait","irrq","softirq","steal","guest","guest_nice"}
)

const (
	Namespace="Test"
	CpucoreLenth = 10
)

func Registry(statSubsystem string,collectors Collectors){
	f[statSubsystem]=collectors
}

type Collector struct {}

func Get() Collector{
	return Collector{}
}

func (c Collector) Describe(ch chan<- *prometheus.Desc) {}

func (c Collector) Collect(ch chan<- prometheus.Metric) {
	Update(ch)
}

func Upload(dirPath string){
	Dirpath=dirPath
}

type Collectors interface {
	GetData(dirpath string)map[string][]string
}

func ReadFile(filepath string) string{
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Println(err)
	}
	f:=string(file)
	a:=strings.ReplaceAll(f,"(","_")
	b:=strings.ReplaceAll(a,")","")
	//delSpace:=strings.ReplaceAll(b," ","")
	delTag:=strings.ReplaceAll(b,"\t","")
	delR:=strings.ReplaceAll(delTag,"\r","")
	delKb:=strings.ReplaceAll(delR,"kB","")
	return delKb
}



func Update(ch chan<- prometheus.Metric){
	log.Println("读取数据")
	wg:=sync.WaitGroup{}
	for k,v:=range f{
		wg.Add(1)
		go func(k string,v Collectors){
			data:=v.GetData(Dirpath)
			DataHandle(data,k,ch)
			wg.Done()
		}(k,v)
	}
	wg.Wait()
	log.Println("指标制作完成")
}

func DataHandle(data map[string][]string,subsystem string,ch chan<- prometheus.Metric){
	wg:=sync.WaitGroup{}
	for k,v:=range data{
		wg.Add(1)
		go func(k string,v []string){
			if len(v)==1{
				DataMetricsCreate(k,v,ch,subsystem)
			}
			if len(v)>1&&subsystem==StatSubsystem{
				StatMetricsCreate(k,v,ch,subsystem)
			}
			wg.Done()
		}(k,v)
	}
	wg.Wait()
}

func DataMetricsCreate(k string,v []string,ch chan<- prometheus.Metric,subsystem string){
	value,err:=strconv.ParseFloat(v[0],64)
	if err==nil{
		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				prometheus.BuildFQName(Namespace, subsystem, k),
				fmt.Sprintf("%s information field %s", subsystem,k),
				nil,
				nil,
			),
			prometheus.GaugeValue,
			value,
		)
	}
}
func StatMetricsCreate(key string,data []string,ch chan<- prometheus.Metric,subsystem string){
	realdata:=DelNullValue(data)
	wg:=sync.WaitGroup{}
	if len(realdata)==CpucoreLenth {
		for i, v := range realdata {
			wg.Add(1)
			go func(i int,v string){
				value, err := strconv.ParseFloat(v, 64)
				if err == nil {
					ch <- prometheus.MustNewConstMetric(
						prometheus.NewDesc(
							prometheus.BuildFQName(Namespace, subsystem, key),
							fmt.Sprintf("%s information field %s", subsystem,key),
							[]string{"mod"},
							nil,
						),
						prometheus.GaugeValue,
						value,
						CPUField[i],
					)
				}
				wg.Done()
			}(i,v)
		}
		wg.Wait()
	}
}

func DelNullValue(data []string)[]string{
	realdata:=[]string{}
	for _,v:=range data{
		if v!=""{
			realdata=append(realdata,v)
		}
	}
	return realdata
}