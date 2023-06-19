apply:
	helm secrets upgrade gacha . --install \
	-f values.yaml \
	-f secrets/secrets.enc.yaml

delete:
	helm uninstall gacha
	kubectl delete pvc data-gacha-mysql-0

dbpass:
	kubectl get secret --namespace default gacha-mysql -o jsonpath="{.data.mysql-root-password}" | base64 --decode; echo

encrypt:
	sops --encrypt --age age158hhkp5kn6aeauhqkaqf2rqcqvjxan0mkuz30qkcr4g3sf9skpzqkm07ad secrets/secrets.yaml > secrets/secrets.enc.yaml