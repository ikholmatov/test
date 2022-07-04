CURRENT_DIR=$(shell pwd)

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen:
	./scripts/gen-proto.sh	${CURRENT_DIR}
	ls genproto/*.pb.go | xargs -n1 -IX bash -c "sed -e '/bool/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"

migration-up:
	migrate -path migration -database 'postgres://venom:112233@localhost:5432/user_service?sslmode=disable' up

migration-down:
	migrate -path migration -database 'postgres://venom:112233@localhost:5432/user_service?sslmode=disable' down

migration-gen:
	migrate create -ext sql -dir migration -seq type_changes

