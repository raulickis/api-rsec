# Para implantar a api-rsec no Kubernetes com Persistência em disco

```
# diretorio /api-rsec/infra
cd infra

kubectl apply -f https://raw.githubusercontent.com/kubernetes/kops/master/addons/metrics-server/v1.8.x.yaml

kubectl delete -f deploy-postgres.yml
kubectl delete configmap api-rsec-config 
kubectl delete -f deploy-api-rsec.yml

kubectl apply -f deploy-postgres.yml
kubectl create configmap api-rsec-config --from-env-file=../.env.kubernetes
kubectl apply -f deploy-api-rsec.yml

# diretorio /web-rsec/infra
# kubectl create configmap web-rsec-config --from-env-file=../.env.kubernetes
# kubectl apply -f deploy-web-rsec.yml


# Para testar no Kubernetes
kubectl port-forward services/api-rsec 9990:9990


```

