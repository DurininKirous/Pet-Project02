APP_NAME=bigscanner

build:
	go build -o $(APP_NAME) ./cmd/app/main.go

test:
	go test -v ./...

run:
	go run ./cmd/app/main.go --dir . --json

docker:
	docker build -t $(APP_NAME):latest .

docker-run:
	docker run --rm -v $(PWD):/data $(APP_NAME):latest --dir /data --json

docker-push:
	docker login -u $(DOCKERHUB_USERNAME) -p $(DOCKERHUB_TOKEN)
	docker tag $(APP_NAME) $(DOCKERHUB_USERNAME)/$(APP_NAME):latest
	docker push $(DOCKERHUB_USERNAME)/$(APP_NAME):latest

clean:
	rm -f $(APP_NAME)

