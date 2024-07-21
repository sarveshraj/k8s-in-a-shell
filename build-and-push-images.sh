TAG=${1:-'0.1'}

docker build -t ghcr.io/sarveshraj/k8s-in-a-shell/server:$TAG ./deployment
docker push ghcr.io/sarveshraj/k8s-in-a-shell/server:$TAG

docker build -t ghcr.io/sarveshraj/k8s-in-a-shell/worker:$TAG ./cronjob
docker push ghcr.io/sarveshraj/k8s-in-a-shell/worker:$TAG

docker build -t ghcr.io/sarveshraj/k8s-in-a-shell/frontend:$TAG ./service
docker push ghcr.io/sarveshraj/k8s-in-a-shell/frontend:$TAG