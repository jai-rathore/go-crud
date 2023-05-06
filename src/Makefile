install:
	go install github.com/swaggo/swag/cmd/swag@latest
	ls -al /go/bin

generate:
	swag init
	go generate .

flyway:
	cd utils; sudo sh ./create_flyway.sh ${SCRIPTNAME}

configWithSecrets:
	sudo npm i -g json-replace-secrets
	cd config ; json-replace-secrets -c config -s secrets

test-coverage:
	go test -v ./... -covermode=count -coverpkg=./... -coverprofile test-coverage/coverage.out
	go tool cover -html coverage/coverage.out -o test-coverage/coverage.html
	open test-coverage/coverage.html

generate-proto:
	protoc --proto_path=grpc/protos grpc/protos/*.proto --go_out=. --go-grpc_out=.