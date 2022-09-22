docker pull appcelerator/etcd
docker run -d -p 2379:2379 -p 2380:2380 appcelerator/etcd --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379
@pause


