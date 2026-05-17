run:
	set -a && . ./.env && set +a && go run ./cmd/api

build:
	go build -o bin/web ./cmd/api
	source .env
	./bin/web

debug:
	go build -o bin/web-debug -gcflags="all=-N -l" ./cmd/api
	source .env && ./bin/web-debug

create-migration:
	migrate create -seq -ext=.sql -dir=./migrations $(mig-name)

up-migrate:
	# set -a && . ./.env && set +a && 
	source .env && \
	migrate -path=./migrations -database=$${PARMAAN_PATR_DB_DSN} up
