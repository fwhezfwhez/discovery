## discovery
Use etcd to realize service discovery.

## note
Make sure your etcd run in 0.0.0.0 not localhost.When docker, it will run like
`docker run -itd -p 2379:2379 -p 2380:2380 --rm -e ETCD_ADVERTISE_CLIENT_URLS=http://0.0.0.0:2379 -e ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379 etcd:3.3`

etcd docker 3+ can be refer to `fwhezfwhez/etcd:3.3`

## core api
server side:

#### Register server
```go
	ser, e := discovery.NewServiceReg([]string{"localhost:2379"}, 5)
 	if e!=nil {
 		panic(e)
 	}
 	if e := ser.PutService("/user-login/node/111/", "10.0.1.1:8081"); e != nil {
 		panic(e)
 	}
```

#### Assume server break accidentally
```go
    	ser, e := discovery.NewServiceReg([]string{"localhost:2379"}, 5)
     	if e!=nil {
     		panic(e)
     	}
     	if e := ser.PutService("/user-login/node/111/", "10.0.1.1:8081"); e != nil {
     		panic(e)
     	}
     	time.Sleep(40 * time.Second)
     	os.Exit(1)
```
On the case, client will lose this server list value
