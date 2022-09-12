这是个简单获取Linux的/proc目录下stat,cpuinfo,meminfo数据并根据个人喜好制作成Prometheus指标格式的代码，正在学习中。
以下所有操作都为centos7的root用户，执行命令时的位置都为main.go所在位置

使用方法:

1.本地运行
将代码拉取到本地Linux环境，下载好依赖，直接执行"go run main.go"就好，目前可以自定义的参数只有端口和文件目录，
参数信息查看执行"go run main.go -h"

2.打包二进制文件
执行"make build",也可以执行:"go build -o test-exporter ./main.go"，一样的效果，命令在Makefile中封装了而已

3.以Linux的service服务运行
需要执行过第2步打包二进制文件，以下用centos6和centos7为例子，centos7可以用cetos6的方式部署service服务，但反过来不行
centos6执行"./script/centos6.sh",启动服务:"service test-exporter.sh start",关闭服务:"service test-exporter.sh stop",重启服务:"service test-exporter.sh restart"
centos7执行"./script/centos7.sh",启动服务:"systemctl start test-exporter.service"，关闭服务:"systemctl stop test-exporter.service"

4.以docker方式运行
如果你已经安装了docker的话
先执行"make build && make images"打包镜像，然后执行:"docker run -d -p 9031:9031 test-exporter:0.0.1",
检查镜像是否运行起来:"docker ps |grep test-exporter",如果已经运行，执行:"curl localhost:9031/metrics"看看是否有指标

5.容器化部署:
如果你已经完成k8s所有组件的安装，还有执行了第四步的"make build && make images"
如果你没有部署Prometheus-operator，那test-exporter-serviceMonitor.yaml你可以不用部署
在main.go的目下执行"kubectl apply -f ./yaml/*",文件不更改的话是获取部署节点的stat、cpuinfo、meminfo文件信息
删除挂载卷的话是获取container(容器)信息

前4步都可以执行:"curl 127.0.0.1:9031/metrics"来查看指标