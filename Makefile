up:
	@docker-compose up -d

down:
	@docker-compose down

fmt:
	@docker-compose exec app go fmt ./...

selfbuild:
	@docker-compose exec app go build -ldflags="-X 'main.Version=`git describe --tags --abbrev=0`-local'"
	@mv -f gon ~/.local/bin/gon
