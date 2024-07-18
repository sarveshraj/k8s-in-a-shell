docker build -t server:0.1 ./deployment
docker tag server:0.1 ghcr.io/sarveshraj/learnk8s/server:0.1
docker push ghcr.io/sarveshraj/learnk8s/server:0.1

docker build -t worker:0.1 ./cronjob
docker tag worker:0.1 ghcr.io/sarveshraj/learnk8s/worker:0.1
docker push ghcr.io/sarveshraj/learnk8s/worker:0.1

docker build -t frontend:0.1 ./service
docker tag frontend:0.1 ghcr.io/sarveshraj/learnk8s/frontend:0.1
docker push ghcr.io/sarveshraj/learnk8s/frontend:0.1