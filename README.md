# gon

[![emojicom](https://img.shields.io/badge/emojicom-%F0%9F%90%9B%20%F0%9F%86%95%20%F0%9F%92%AF%20%F0%9F%91%AE%20%F0%9F%86%98%20%F0%9F%92%A4-%23fff)](http://neni.dev/emojicom) [![CI](https://github.com/nenitf/gon/actions/workflows/ci.yml/badge.svg)](https://github.com/nenitf/gon/actions/workflows/ci.yml)

![Meme bob esponsa cansado](./bob-cansado.jpg)

CLI que reexecuta um comando no terminal com <kbd>enter</kbd> ao invés de <kbd>↑</kbd><kbd>enter</kbd>

## Utilização

1. Baixe o [executável](https://github.com/nenitf/gon/releases/latest)
2. Invoque-o antes do comando (ex: ``./gon -c date "+%H:%M:%S %d/%m/%y"``)
3. Aperte <kbd>enter</kbd> para repeti-lo 

> Também pode executar o binário se estiver no `PATH` do sistema, ex: ``gon phpunit``

> Atualização para linux: `curl -L https://github.com/nenitf/gon/releases/latest/download/gon_linux_amd64.tar.gz | tar -xvz`

## Desenvolvimento

### Ambiente

Para executar e testar é necessário subir o ambiente

```sh
make up
```

> Pare com ``make down``

### Execução

```sh
docker-compose exec app go run main.go echo "funciona!"
```

### Linting

```sh
make fmt
```

### Build e atualização local no Linux

```sh
make selfbuild
```

## Alternativas

### Shell script
```sh
while echo "exemplo"; read line; do true; done;

# sugestão de função para `.bashrc`
# wrl echo "exemplo"
#wrl(){
#	while $@; read line; do true; done;
#}
```
