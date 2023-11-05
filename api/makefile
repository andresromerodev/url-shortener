.PHONY: generate
generate:
	go run github.com/steebchen/prisma-client-go generate

.PHONY: db-push
db-push:
	go run github.com/steebchen/prisma-client-go db push

.PHONY: db-pull
db-pull:
	go run github.com/steebchen/prisma-client-go db pull

.PHONY: migrate-dev
migrate-dev:
	go run github.com/steebchen/prisma-client-go migrate dev

.PHONY: migrate-deploy
migrate-deploy:
	go run github.com/steebchen/prisma-client-go migrate deploy
