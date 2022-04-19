consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul -node=n1 -bind=127.0.0.1 -ui -rejoin -config-dir=/etc/consul.d/ -client 0.0.0.0
micro run getArea/

 
