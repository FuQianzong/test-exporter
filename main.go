package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"test1/collector"
	"test1/getargs"
)

func main() {
	args := getargs.GetArgs()
	log.Println("获取参数")
	endpoint := GetEndpoint(args.Endpoint)
	collector.Upload(args.DirPath)
	PromeMetricsHandle()
	log.Println("Listening on ", args.Endpoint)
	if err := endpoint.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func GetEndpoint(endpoint string) *http.Server {
	return &http.Server{Addr: ":" + endpoint}
}

//制作Prometheus格式的指标
func PromeMetricsHandle() {
	reg := prometheus.NewPedanticRegistry()
	c := collector.Get()
	reg.MustRegister(c)
	gatherers := prometheus.Gatherers{
		//prometheus.DefaultGatherer,（这个为默认指标，可启用，也可不用）
		reg,
	}
	h := promhttp.HandlerFor(gatherers, promhttp.HandlerOpts{ErrorHandling: promhttp.ContinueOnError})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) { h.ServeHTTP(w, r) })
	//http.Handle("/metrics", promhttp.Handler())
}
