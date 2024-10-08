# Installation APISIX

install APISIX ingress with kubernetes manifest file.

create namespace

```shell
kubectl create ns apisix
```

apply etcd

```shell
kubectl -n apisix apply -f apisix/etcd.yaml
```

check health etcd

```shell
kubectl -n apisix exec -it etcd-0 -- etcdctl endpoint health
```

apply config

```shell
kubectl -n apisix apply -f apisix/config.yaml
```

apply apisix deployment

```shell
kubectl -n apisix apply -f apisix/apisix-deployment.yaml
```

check readiness apisix deployment

```shell
kubectl get pods -n apisix -l app.kubernetes.io/name=apisix -o name
```

check apisix deploy correctly
```shell
kubectl -n apisix exec -it $(kubectl get pods -n apisix -l app.kubernetes.io/name=apisix -o name) -- curl http://127.0.0.1:9080
```

# httpbin for testing

install httpbin
```shell
kubectl create ns demo
kubectl -n demo run httpbin --image-pull-policy=IfNotPresent --image kennethreitz/httpbin --port 80
kubectl -n demo expose pod httpbin --port 80
```

try access httpbin from apisix pod
```shell
kubectl -n apisix exec -it $(kubectl get pods -n apisix -l app.kubernetes.io/name=apisix -o name) -- curl http://httpbin.demo/get
```

create route apisix into httpbin
```shell
kubectl -n apisix exec -it $(kubectl get pods -n apisix -l app.kubernetes.io/name=apisix -o name) -- curl "http://127.0.0.1:9180/apisix/admin/routes/1" -H "X-API-KEY: edd1c9f034335f136f87ad84b625c8f1" -X PUT -d '
{
  "uri": "/*",
  "host": "httpbin.org",
  "upstream": {
    "type": "roundrobin",
    "nodes": {
      "httpbin.demo:80": 1
    }
  }
}'
```

test apisix route to httpbim
```shell
kubectl -n apisix exec -it $(kubectl get pods -n apisix -l app.kubernetes.io/name=apisix -o name) -- curl "http://127.0.0.1:9080/get" -H 'Host: httpbin.org'
```

# Installing APISIX Ingress controller

create service account and cluster role
```shell
git clone https://github.com/apache/apisix-ingress-controller.git --depth 1
cd apisix-ingress-controller/
kubectl apply -k samples/deploy/rbac/apisix_view_clusterrole.yaml # apply cluster role
kubectl -n apisix create serviceaccount apisix-ingress-controller # create service account
# bind cluster role and service account
kubectl create clusterrolebinding apisix-viewer --clusterrole=apisix-view-clusterrole --serviceaccount=apisix:apisix-ingress-controller
```

install crd

```shell
kubectl apply -k samples/deploy/crd
```

apply apisix config

```shell
kubectl -n apisix apply -f apisix/apisix-config.yaml
```

apply ingress service

```shell
kubectl -n apisix apply -f apisix/ingress-service.yaml
```

remove existing route created before for httpbin
```shell
kubectl -n apisix exec -it $(kubectl get pods -n apisix -l app.kubernetes.io/name=apisix -o name) -- curl "http://127.0.0.1:9180/apisix/admin/routes/1" -X DELETE -H "X-API-KEY: edd1c9f034335f136f87ad84b625c8f1"
```

apply ingress deployment

```shell
kubectl -n apisix apply -f apisix/ingress-deployment.yaml
```

apply apisixroute http bin

```shell
kubectl -n demo apply -f demo/httpbin-route.yaml
```

test apisixroute
```shell
kubectl -n apisix exec -it $(kubectl get pods -n apisix -l app.kubernetes.io/name=apisix -o name) -- curl "http://127.0.0.1:9080/get" -H "Host: local.httpbin.org"
```
