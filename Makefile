start:
	cd server && docker-compose up -d

apply:
	kubectl apply -f k8s/service.yaml
	kubectl apply -f k8s/config.yaml
	kubectl apply -f k8s/deployment.yaml

delete:
	kubectl delete -f k8s/service.yaml
	kubectl delete -f k8s/config.yaml
	kubectl delete -f k8s/deployment.yaml