# Consul Traffic Management

## Learn by example

## Deploy front-end web service

```
git clone https://github.com/servicemeshbook/hello-echo.git

cd hello-echo

kubectl -n consul apply -f web.yaml
```

### Run web service

```
$ curl http://localhost:30145 ; echo
===============================================
Request time   : 2019-09-23 13:51:54.444014201 +0000 UTC
Requested path : /
Host IP        : 192.168.142.101
Pod IP         : 192.168.230.216
Pod Name       : web-deployment-86459c6966-2clhv
Pod Namespace  : consul
Host           : localhost:30145
RemoteAddr     : 192.168.142.101:50126
===============================================
no healthy upstream
```

Notice `no healthy upstream` since we have not yet deployed api services.


## Deploy backend API service

### Create v1 and v2 pod with api-service

```
kubectl -n consul apply -f api-v1.yaml

kubectl -n consul apply -f api-v2.yaml

```

### Run api-v1 on node port 30146.

```
$ curl http://localhost:30146
===============================================
Request time   : 2019-09-23 13:55:06.636072035 +0000 UTC
Requested path : /
Host IP        : 192.168.142.101
Pod IP         : 192.168.230.246
Pod Name       : api-v1-7fcf5d98d4-tgqrk
Pod Namespace  : consul
Host           : localhost:30146
RemoteAddr     : 192.168.142.101:34394
===============================================
```

### Run api-v2 service on node port 30147.

```
$ curl http://localhost:30147
===============================================
Request time   : 2019-09-23 13:55:12.075860863 +0000 UTC
Requested path : /
Host IP        : 192.168.142.101
Pod IP         : 192.168.230.205
Pod Name       : api-v2-5d64d5f8ff-zlcp6
Pod Namespace  : consul
Host           : localhost:30147
RemoteAddr     : 192.168.142.101:58432
===============================================
```

## Run front end web service

Run web service on node port 30145 again.

You will notice that the frontend web-service calls api service which is pointing to api-v1 pod.

```
===============================================
Request time   : 2019-09-21 01:25:29.844609478 +0000 UTC
Requested path : /
Host IP        : 192.168.142.101
Pod IP         : 192.168.230.202
Pod Name       : web-86459c6966-2clhv
Pod Namespace  : consul
Host           : localhost:30145
RemoteAddr     : 192.168.142.101:47332
===============================================
Request time : 2019-09-23 14:11:06.091295669 +0000 UTC
Requested path : /
Host IP : 192.168.142.101
Pod IP : 192.168.230.205
Pod Name : api-v1-7fcf5d98d4-tgqrk
Pod Namespace : consul
Host : localhost:8081
RemoteAddr : 127.0.0.1:57152
```

Repeat same again and you will notice that the upstream service called is api-v2. This happens in a round-robin fashion.

## Canary Deployment

First create a service-spiltter that defines the service subsets.

Create `service-split.hcl` as shown below.

```
kind = "service-splitter",
name = "api"

splits = [
  {
    weight = 99,
    service_subset = "v1"
  },
  {
    weight = 1,
    service_subset = "v2"
  }
]
```

### Create service split

```
consul config write service-split.hcl
```

Run same curl command 200 times and notice that api-v2 is called 1% of the time.

```
$ curl -s http://localhost:30145?[1-200] | grep "Pod Name.*api"

<< Removed >>
Pod Name : api-v1-7fcf5d98d4-tgqrk
Pod Name : api-v1-7fcf5d98d4-tgqrk
Pod Name : api-v2-5d64d5f8ff-zlcp6
Pod Name : api-v1-7fcf5d98d4-tgqrk
<< Removed >>
```

## Shift Traffic Permanently

If it is determined that 100% of traffic should now be shifted to v2 of api after testing it successfully, the weight can be defined to 100% to the subset v2.  

Create file `service-split-100.hcl`

```
kind = "service-splitter",
name = "api"

splits = [
  {
    weight = 0,
    service_subset = "v1"
  },
  {
    weight = 100
    service_subset = "v2"
  }
]
```
### Create service split

```
consul config write service-split-100.hcl
```

Repeat same curl command 10 times.

```
$ curl -s http://localhost:30145?[1-10] | grep "Pod Name.*api"

Pod Name : api-v2-5d64d5f8ff-zlcp6
Pod Name : api-v2-5d64d5f8ff-zlcp6
Pod Name : api-v2-5d64d5f8ff-zlcp6
Pod Name : api-v2-5d64d5f8ff-zlcp6
Pod Name : api-v2-5d64d5f8ff-zlcp6
Pod Name : api-v2-5d64d5f8ff-zlcp6
Pod Name : api-v2-5d64d5f8ff-zlcp6
Pod Name : api-v2-5d64d5f8ff-zlcp6
Pod Name : api-v2-5d64d5f8ff-zlcp6
Pod Name : api-v2-5d64d5f8ff-zlcp6
```

