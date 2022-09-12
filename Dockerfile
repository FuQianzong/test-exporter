FROM centos

RUN mkdir test-exporter && makdir -p /data/proc

ADD ./test-exporter  /test-exporter/test-exporter
 
 RUN chmod +x /test-exporter/test-exporter

 ENTRYPOINT [ "/test-exporter/test-exporter" ]