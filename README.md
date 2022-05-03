# gon

[![emojicom](https://img.shields.io/badge/emojicom-%F0%9F%90%9B%20%F0%9F%86%95%20%F0%9F%92%AF%20%F0%9F%91%AE%20%F0%9F%86%98%20%F0%9F%92%A4-%23fff)](http://neni.dev/emojicom) [![CI](https://github.com/nenitf/gon/actions/workflows/ci.yml/badge.svg)](https://github.com/nenitf/gon/actions/workflows/ci.yml)

Reexecuta comando com <kbd>enter</kbd> ao invés de <kbd>↑</kbd><kbd>enter</kbd>. Útil para reexecutar scripts de teste para aplicar *tdd*.

Para utilizar, baixe o [executável](https://github.com/nenitf/gon/releases/latest) e o invoque antes do comando (ex: ``./gon composer test``)

## Desenvolvimento

```sh
docker-compose up -d
```

> Pare com ``docker-compose down``

## Linting

```sh
docker-compose exec app go fmt ./...
```
