start:
	docker-compose up -d

protogen:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	cd pbapi && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative user-service.proto

mockery_install:
	go install github.com/vektra/mockery/v2@v2.24.0

test_with_db:
	docker-compose -f docker-compose.test.yaml up --build --abort-on-container-exit && \
    docker-compose -f docker-compose.test.yaml rm -fsv

push_to_dockerhub:
	docker build -t $(DOCKERHUB_USERNAME)/user-service:latest . && \
	docker push $(DOCKERHUB_USERNAME)/user-service:latest