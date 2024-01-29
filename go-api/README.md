# MySQL Docker Container

Exemplo simples de como construir e executar um contêiner Docker com o servidor MySQL 8.

## Construir a Imagem Docker

Para construir a imagem Docker, você pode executar o seguinte comando:

```sh
$ docker build -t mysql .
```

Este comando usa o Dockerfile neste repositório para criar uma imagem Docker baseada no servidor MySQL 8 oficial.

## Executar o Contêiner Docker

Para executar o contêiner Docker MySQL, você pode usar o seguinte comando:

```sh
$ docker run --detach -p 3306:3306 mysql
```

Isso iniciará o contêiner MySQL em segundo plano (-d), mapeando a porta 3306 do contêiner para a porta 3306 do host (-p 3306:3306).

## Informações de Conexão

Aqui estão as informações de conexão para acessar o banco de dados MySQL dentro do contêiner:

- URL de conexão: `mysql://localhost:3306/poc`
- Nome de usuário: `user`
- Senha: `dummypassword`

Você pode usar essas informações de conexão para se conectar ao banco de dados MySQL a partir de um cliente JDBC ou qualquer outra ferramenta de banco de dados que você preferir.

## Acessar o Banco de Dados MySQL via Linha de Comando

Você também pode acessar o banco de dados MySQL dentro do contêiner usando a linha de comando. Para fazer isso, use o seguinte comando:

```sh

$ docker exec -it <ID_DO_CONTAINER_MYSQL> mysql -u user -p
```

Substitua `<ID_DO_CONTAINER_MYSQL>` pelo ID do contêiner MySQL. Isso abrirá um prompt de comando MySQL interativo, onde você pode executar consultas SQL diretamente.