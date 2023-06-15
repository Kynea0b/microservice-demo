apply:
	source .env; helm upgrade microservice-demo-chart . --install \
	--set auth.username=${DB_GACHA_USERNAME} \
	--set auth.password=${DB_GACHA_PASSWORD} \
	-f values.yaml
