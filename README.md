# gon

[![emojicom](https://img.shields.io/badge/emojicom-%F0%9F%90%9B%20%F0%9F%86%95%20%F0%9F%92%AF%20%F0%9F%91%AE%20%F0%9F%86%98%20%F0%9F%92%A4-%23fff)](http://neni.dev/emojicom) [![CI](https://github.com/nenitf/gon/actions/workflows/ci.yml/badge.svg)](https://github.com/nenitf/gon/actions/workflows/ci.yml)

Reexecuta comando com <kbd>enter</kbd> ao invés de <kbd>↑</kbd><kbd>enter</kbd>. Criado para deixar *tdd* ainda mais divertido.

## Utilização

1. Baixe o [executável](https://github.com/nenitf/gon/releases/latest)
2. Invoque-o antes do comando (ex: ``./gon composer test``)

> Também pode executar o binário se estiver no `PATH` do sistema, ex: ``gon composer test``

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
