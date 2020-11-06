# G-commerce

## Estudo da linguagem GO(desenvolvendo um e-commerce básico)

__Tecnologias utilizadas__

| Nome | Descrição | versão | fonte | status|
|---|---|:---:|---|:---:|
| GO            | Linguagem de programação  | 1.15.3    | [GO](https://golang.org)| Em uso|
| PostgreSQL    | Banco de dados relacional | 13.0      | [PostgreSQL](https://www.postgresql.org/)| Em uso|
|Podman¹        | Mecanismo para uso de containers  | 2.1.1 | [Podman](https://podman.io/)| Em uso|
|Podman-compose²        | Execuçao de scripts para containers  | 2.1.1 | [Podman-compose](https://github.com/containers/podman-compose)| Em uso|
|PGadmin        | Adminnistração para platarformas PostgreSQL| 4    | [PGadmin](https://www.pgadmin.org/)| Em uso|
| KeyCloak      | Gerenciamente de identificação e acesso| 11.0.3| [Keycloak](https://www.keycloak.org/)| Em breve|

Observações:

1. O Podman pode ser substituido pelo Docker.
2. O Docker pode ser substituido pelo Docker-compose.

## Tarefas e casos de uso
- [x] Estruturação inicial(arquitetura de software)
- [ ] internacionalização(i18n)
- [ ] integrar com banco de dados postgre
- [ ] API REST-FULL
- [ ] Login
- [ ] Produto
- [ ] Estoque
- [ ] Venda
- [ ] Carrinho

## Exectar projeto:

- Clonar este repositório.
- Ter as tecnologias compativeis.
- Iniciar o banco de dados.
- executa o camando na raiz do projeto: 

>go run main.go

#### Em atualização ...