# Consul Traffic Management

## Learn by example

## Run front-end web service

```
git clone https://github.com/servicemeshbook/hello-echo.git

cd hello-echo

kubectl -n consul apply -f hello-echo.yaml
```

## Run API service

### Create v1 and v2 pod with api-service

```
cd rest
kubectl -n consul apply -f api.yaml
```

## Run front end web service

Open browser and run http://localhost:30145

You will notice that the frontend web-service calls api service which is pointing to api-v1 pod.

The `POD_IP` and `POD_NAME` are shown so that we know which POD is returning the results.

```
===============================================
Requested path    is: /
The HOST IP       is: 192.168.142.101
The POD IP        is: 192.168.230.236
The POD NAME      is: web
The POD NAMESPACE is: consul

===============================================
"Accept": ["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3"]
"Accept-Encoding": ["gzip, deflate, br"]
"Accept-Language": ["en-US,en;q=0.9"]
"Connection": ["keep-alive"]
"Upgrade-Insecure-Requests": ["1"]
"User-Agent": ["Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36"]
===============================================
Host       = "localhost:30145"
RemoteAddr = "192.168.142.101:50014"
===============================================
Response starts from http://api-service.consul.svc.cluster.local:8080
===============================================
Requested path    is: /
The HOST IP       is: 192.168.142.101
The POD IP        is: 192.168.230.206
The POD NAME      is: api-v1
The POD NAMESPACE is: consul

===============================================
"User-Agent": ["Go-http-client/1.1"]
"Accept-Encoding": ["gzip"]
===============================================
Host       = "api-service.consul.svc.cluster.local:8080"
RemoteAddr = "192.168.142.101:38750"
===============================================
Response starts from http://httpbin.org/post
{
  "args": {}, 
  "data": "{\"POD_IP\":\"192.168.230.206\",\"POD_NAME\":\"api-v1\"}", 
  "files": {}, 
  "form": {}, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Content-Length": "48", 
    "Content-Type": "application/json", 
    "Host": "httpbin.org", 
    "User-Agent": "Go-http-client/1.1"
  }, 
  "json": {
    "POD_IP": "192.168.230.206", 
    "POD_NAME": "api-v1"
  }, 
  "origin": "1.2.3.4, 1.2.3.4", 
  "url": "https://httpbin.org/post"
}
Response ends from http://httpbin.org/post
Response ends from http://api-service.consul.svc.cluster.local:8080
```
