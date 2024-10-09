apply create namespace centralized auth
```shell
kubectl apply -f centralized-auth/centralized-auth-namespace.yaml
```

apply create deployment keycloak
```shell
kubectl apply -f centralized-auth/keycloak-deployment.yaml
```

apply create service keycloak
```shell
kubectl apply -f centralized-auth/keycloak-service.yaml
```

check keycloak from apisix container
```shell
kubectl -n apisix exec -it $(kubectl get pods -n apisix -l app.kubernetes.io/name=apisix -o name) -- curl http://keyclaok.centralized-auth:8080/realms/master
```

add route for keycloak on apisix
```shell
kubectl apply -f centralized-auth/keycloak-route.yaml
```

check route keycloak from apisix gateway
```shell
kubectl -n apisix exec -it $(kubectl get pods -n apisix -l app.kubernetes.io/name=apisix -o name) -- curl http://127.0.0.1:9080/keycloak/realms/master -H "Host: 192.168.100.105"
```
