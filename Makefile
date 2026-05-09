run:
	set -a && . ./.env && set +a && go run ./cmd/api

create-migration:
	migrate create -seq -ext=.sql -dir=./migrations $(mig-name)

up-migrate:
	# set -a && . ./.env && set +a && 
	source .env && \
	migrate -path=./migrations -database=$${PARMAAN_PATR_DB_DSN} up
