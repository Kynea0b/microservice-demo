apply:
	source .env; helm upgrade gacha . --install \
	-f values.yaml

dbpass:
	kubectl get secret --namespace default gacha-mysql -o jsonpath="{.data.mysql-root-password}" | base64 --decode; echo
