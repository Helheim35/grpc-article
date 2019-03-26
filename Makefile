IPATH = car

all: clean vendor proto

clean:
	rm -Rf vendor ; rm $(IPATH)/*.pb.go

proto:
	protoc -I=$(IPATH) --go_out=plugins=grpc:$(IPATH) ./$(IPATH)/*.proto

vendor:
	go mod vendor

runclient:
	go run client/main.go

runserver:
	go run server/main.go