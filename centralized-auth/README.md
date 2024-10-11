apply create namespace centralized auth
```shell
kubectl apply -f centralized-auth/centralized-auth-namespace.yaml
```

apply create ingress for host centralized-auth.local to apisix
```shell
kubectl apply -f centralized-auth/centralized-auth-ingress.yaml
```

## create keycloak

apply create deployment keycloak
```shell
kubectl apply -f centralized-auth/keycloak/keycloak-deployment.yaml
```

apply create service keycloak
```shell
kubectl apply -f centralized-auth/keycloak/keycloak-service.yaml
```

check keycloak from apisix container
```shell
kubectl -n apisix exec -it $(kubectl get pods -n apisix -l app.kubernetes.io/name=apisix -o name) -- curl http://keycloak.centralized-auth:8080/keycloak/realms/master
```

add route for keycloak on apisix
```shell
kubectl apply -f centralized-auth/keycloak/keycloak-route.yaml
```

check route keycloak from apisix gateway
```shell
kubectl -n apisix exec -it $(kubectl get pods -n apisix -l app.kubernetes.io/name=apisix -o name) -- curl http://127.0.0.1:9080/keycloak/realms/master -H "Host: 192.168.100.105"
```


## create user server

apply create deployment user server
```shell
kubectl apply -f centralized-auth/user-server/user-server-deployment.yaml
```

apply create service user server
```shell
kubectl apply -f centralized-auth/user-server/user-server-service.yaml
```

check user server from apisix container
```shell
kubectl -n apisix exec -it $(kubectl get pods -n apisix -l app.kubernetes.io/name=apisix -o name) -- curl http://user-server.centralized-auth:3000/api/users/me
```

add route for user server on apisix
```shell
kubectl apply -f centralized-auth/user-server/user-server-route.yaml
```

check route user server from apisix gateway
```
kubectl -n apisix exec -it $(kubectl get pods -n apisix -l app.kubernetes.io/name=apisix -o name) -- curl http://127.0.0.1:9080/user-server/api/users/me -H "Host: centralized-auth.local"
```



## payment user server

apply create deployment payment server
```shell
kubectl apply -f centralized-auth/payment-server/payment-server-deployment.yaml
```

apply create service payment server
```shell
kubectl apply -f centralized-auth/payment-server/payment-server-service.yaml
```

check payment server from apisix container
```shell
kubectl -n apisix exec -it $(kubectl get pods -n apisix -l app.kubernetes.io/name=apisix -o name) -- curl http://payment-server.centralized-auth:3000/api/users/payments
```

add route for payment server on apisix
```shell
kubectl apply -f centralized-auth/payment-server/payment-server-route.yaml
```

check route payment server from apisix gateway
```
kubectl -n apisix exec -it $(kubectl get pods -n apisix -l app.kubernetes.io/name=apisix -o name) -- curl http://127.0.0.1:9080/payment-server/api/users/payments -H "Host: centralized-auth.local"
```
