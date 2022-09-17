#!/usr/bin/python3
from prometheus_client import start_http_server,Gauge
from prometheus_client.core import CollectorRegistry
import os
import time

dirpath="/proc"
cpucores=["user","nice","system","idle","iowait","irrq","softirq","steal","guest","guest_nice"]
statfilepath=dirpath+"/stat"
cpuinfofilepath=dirpath+"/cpuinfo"
meminfofilepath=dirpath+"/meminfo"
ip=os.popen("hostname -i").read().strip('\n')
namespace="test"
cpuinfosubsystem="cpuinfo"
meminfosubsystem="meminfo"
statsubsystem="stat"
REGISTRY = CollectorRegistry(auto_describe=False)
gauges={}

def initstatregistry(filepath,subsystem):
        f=open(filepath,"r")
        for line in f.readlines():
            s=del_special(line)
            args=s.split( )
            #print(args)
            metricsregistry(args,subsystem)
        f.close()

def initinforegistry(filepath,subsystem):
        f=open(filepath,"r")
        for line in f.readlines():
            s=del_special(line)
            delspace=s.replace(" ","")
            args=delspace.split(":")
            #print(args)
            metricsregistry(args,subsystem)
        f.close()

def readstatfile(filepath,subsystem):
        f=open(filepath,"r")
        for line in f.readlines():
            s=del_special(line)
            args=s.split( )
            #print(args)
            metricsbuild(args,subsystem)
        f.close()

def readinfofile(filepath,subsystem):
        f=open(filepath,"r")
        for line in f.readlines():
            s=del_special(line)
            delspace=s.replace(" ","")
            args=delspace.split(":")
            #print(args)
            metricsbuild(args,subsystem)
        f.close()

def is_number(s):
    try:
        float(s)
        return True
    except ValueError:
        pass
    return False



def del_special(s):
    strip=s.strip()
    a=strip.replace("(","")
    b=strip.replace(")","")
    delr=strip.replace('\r',"")
    delt=delr.replace('\t',"")
    deln=delt.replace('\n',"")
    return deln

def metricsregistry(args,subsystem):
    metricsname=namespace+"_"+subsystem+"_"+args[0]
    description=subsystem+" field "+args[0]
    if len(args)==2 and is_number(args[1]):
        try:
            #print(metricsname,args)
            infobuild=Gauge(metricsname,description,['instance'],registry=REGISTRY)
            gauges[metricsname]=infobuild
        except ValueError:
            pass
    if len(args)==11:
        statbuild=Gauge(metricsname,description,['instance','mode'],registry=REGISTRY)
        gauges[metricsname]=statbuild

def metricsbuild(args,subsystem):
    metricsname=namespace+"_"+subsystem+"_"+args[0]
    if len(args)==2 and is_number(args[1]):
        #print(metricsname)
        gauges[metricsname].labels({'instance':ip}).set(float(args[1]))
        print(gauges[metricsname],ip,args[1])
    if len(args)==11:
        #print(metricsname)
        argsplite=args[1:]
        for i in range(10):
            gauges[metricsname].labels({'instance':ip},{'mode':cpucores[i]}).set(float(argsplite[i]))
            print(gauges[metricsname],ip,cpucores[i],args[i])

if __name__ == "__main__":
    initstatregistry(statfilepath,statsubsystem)
    initinforegistry(cpuinfofilepath,cpuinfosubsystem)
    initinforegistry(meminfofilepath,meminfosubsystem)
    #print(gauges)
 #暴露端口
    start_http_server(9032)
 #不断传入数据
    while True:
        readstatfile(statfilepath,statsubsystem)
        readinfofile(cpuinfofilepath,cpuinfosubsystem)
        readinfofile(meminfofilepath,meminfosubsystem)
        time.sleep(30)
