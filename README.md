# Nothing Here

Nginx Controller
```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.10.0/deploy/static/provider/cloud/deploy.yaml
```

Resources Port Fowarding
```bash
kubectl port-forward service/mysql 3306:3306
kubectl port-forward service/sentinel 5000:5000
```