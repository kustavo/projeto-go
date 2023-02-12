# Exemplo Go

Exempo de uma API em Go utilizando Redis, JWT...

## Execução docker-compose

`docker-compose docker-compose up -d`

## k6

Medir tempo gasto pelas requisições http.

- Instalação: <https://k6.io/docs/getting-started/installation>

- Execução: `k6 run k6.js`

## pmap

Medir memória usada pelo processo

- Obter PID: `pidof <nome-arquivo-execucao>`

- Ver consumo memória: `pmap -x <pid>`

## Referências

- <https://medium.com/@putuprema/spring-boot-vs-asp-net-core-a-showdown-1d38b89c6c2d>
- <https://github.com/putuprema/spring-vs-aspnetcore>
