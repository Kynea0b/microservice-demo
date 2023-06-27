# kubernetes my Helm chart

## Helm

```
# install
brew install helm

# create
helm create microservice-demo

# 依存関係解決
helm dependencies build
```

## MySQL

```
# パスワード
kubectl get secret --namespace default gacha-mysql -o jsonpath="{.data.mysql-root-password}" | base64 --decode; echo

# 別のコンテナから繋いでみる
kubectl exec -it <Pod name> -- bash
apt update && apt install default-mysql-core
mysql -uroot -p -h gacha-mysql
```

## apply

一括デプロイ

```
make apply
```

## delete

一括削除

```
make delete
```