# Quick Start

1. Apply Redis
```bash
kubectl apply -f kubernetes/redis/
```

2. Waiting for redis all redis pod are started, then delete sentinel statefulset and apply again
```kubectl
kubectl delete statefulset sentinel
```
```kubectl
kubectl apply -f kubernetes/redis/
```

3. Apply MySQL
```kubectl
kubectl apply -f kubernetes/mysql/
```

4. Apply App (API)
```kubectl
kubectl apply -f kubernetes/app/
```

5. Apply Ingress
```kubectl
kubectl apply -f kubernetes/ingress.yml
```

6. Apply Nginx Controller
```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.10.0/deploy/static/provider/cloud/deploy.yaml
```

# Resources Port Fowarding
MySQL
```bash
kubectl port-forward service/mysql 3306:3306
```

Redis
```bash
kubectl port-forward service/sentinel 5000:5000
```