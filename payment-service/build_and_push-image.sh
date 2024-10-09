docker build -t aprianfirlanda/payment-service:v1 .
echo $PASSWORD_DOCKER_HUB | docker login -u aprianfirlanda --password-stdin
docker push aprianfirlanda/payment-service:v1
