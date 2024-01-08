export ORACLE_PASSWORD=<your password>

run_local:
	go run cmd/main.go

install_deps:
	go get ./...
	go mod download
	go mod tidy
	go mod vendor

build_local:
	go build cmd/main.go

test_local:
	go test ./... -p=1 -count=1

test_local_with_coverage:
	go test ./... -p=1 -count=1 -coverprofile=cover.out

generate_swagger:
	swag init

start_db:
	docker run -d -p 1521:1521 -e ORACLE_PASSWORD=$(ORACLE_PASSWORD) -v oracle-volume:/opt/oracle/oradata gvenzl/oracle-xe
