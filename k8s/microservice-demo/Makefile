apply:
	source .env; helm upgrade microservice-demo-chart . --install \
	--set auth.username=${DB_USERNAME} \
	--set auth.password=${DB_PASSWORD} \
	-f values.yaml
