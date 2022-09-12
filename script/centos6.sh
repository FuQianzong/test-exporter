#/bin/bash
mkdir -p /usr/local/test-exporter
cp ./test-exporter /usr/local/test-exporter
cp ./service/test-exporter.sh /etc/init.d
chmod +x /etc/init.d/test-exporter.sh
chkconfig --add /etc/init.d/test-exporter.sh