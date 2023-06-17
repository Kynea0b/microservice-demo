apply:
	source .env; helm upgrade gacha . --install \
	--set auth.username=${DB_GACHA_USERNAME} \
	--set auth.password=${DB_GACHA_PASSWORD} \
	-f values.yaml

dbpass:
	kubectl get secret --namespace default gacha-mysql -o jsonpath="{.data.mysql-root-password}" | base64 --decode; echo
