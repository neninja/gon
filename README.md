# gon

[![emojicom](https://img.shields.io/badge/emojicom-%F0%9F%90%9B%20%F0%9F%86%95%20%F0%9F%92%AF%20%F0%9F%91%AE%20%F0%9F%86%98%20%F0%9F%92%A4-%23fff)](http://neni.dev/emojicom) [![CI](https://github.com/nenitf/gon/actions/workflows/ci.yml/badge.svg)](https://github.com/nenitf/gon/actions/workflows/ci.yml)

![Meme bob esponsa cansado](./bob-cansado.jpg)

CLI que reexecuta um comando no termimnal com <kbd>enter</kbd> ao invés de <kbd>↑</kbd><kbd>enter</kbd>

## Utilização

1. Baixe o [executável](https://github.com/nenitf/gon/releases/latest)
2. Invoque-o antes do comando (ex: ``./gon composer test``)
3. Aperte <kbd>enter</kbd> para repeti-lo 

> Também pode executar o binário se estiver no `PATH` do sistema, ex: ``gon composer tdd``

> Com `-c` é possível limpar o terminal a cada execução, ex: ``gon -c date "+%H:%M:%S %d/%m/%y"``

## Desenvolvimento

### Setup inicial

```sh
docker-compose up -d
#docker-compose exec app go mod tidy
```

> Pare com ``docker-compose down``

### Execução

```sh
docker-compose exec app go run main.go echo "funciona!"
```

### Linting

```sh
docker-compose exec app go fmt ./...
```

### Build e atualização local

- Linux

```sh
docker-compose exec app go build -ldflags="-X 'main.Version=`git describe --tags --abbrev=0`'" && mv -f gon ~/.local/bin/gon
```
