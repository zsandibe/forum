build:
	docker build -t forum .
run-img:
	docker run --name=forum -p 8081:8081 --rm -d forum 
run:
	go run cmd/main.go
stop:
	docker stop forum