#!/bin/bash
#description: test service script
#chkconfig: 235 90 90

INITIATE_FILE="/usr/local/test-exporter/initiate.file"

start(){
  if [ -f ${INITIATE_FILE} ];then
    echo "服务已启动" > /dev/null
    exit 1
  else
    /usr/local/test-exporter 1 > /dev/null 2>&1 &
    touch ${INITIATE_FILE}
    exit 0
  fi
}
stop(){
  killall /usr/local/test-exporter
  rm -f ${INITIATE_FILE}
  echo "服务已关闭" > /dev/null
  exit 0
}
restart(){
  stop
  start
}

#服务管理:start,stop,status
case $1 in
   stop)
       stop
   ;;
   start)
       start
   ;;
   restart)
       restart
esac