docker build -t aprianfirlanda/user-service:v1 .
echo $PASSWORD_DOCKER_HUB | docker login -u aprianfirlanda --password-stdin
docker push aprianfirlanda/user-service:v1
