# Order Stream

## Descrição

Este projeto é uma aplicação de streaming de pedidos que utiliza Kafka para comunicação entre produtores e consumidores. A aplicação inclui um CRUD de produtos, um produtor de pedidos via API que envia mensagens para o Kafka, e um consumidor que lê as mensagens do Kafka e as envia para o MongoDB.

## Funcionalidades

- CRUD de Produtos
- Produtor de Pedidos via API
- Envio de Pedidos para Kafka
- Consumidor de Pedidos do Kafka
- Armazenamento de Pedidos no MongoDB

## Estrutura do Projeto

- `internal/models`: Contém as definições das estruturas de dados, como `Order`, `Item` e `Customer`.
- `internal/repositories/mongo`: Contém a lógica de interação com o MongoDB.
- `internal/stream`: Contém a lógica de interação com o Kafka, incluindo o produtor e o consumidor.

## Configuração

### Pré-requisitos

- Docker
- Docker Compose

### Instalação

1. Clone o repositório:

    ```sh
    git clone https://github.com/seu-usuario/orderstreamrest.git
    cd orderstreamrest
    ```

2. Inicie os serviços com Docker Compose:

    ```sh
    docker-compose up -d
    ```

3. Acesse a aplicação em `http://localhost:8080`.

## Uso

### CRUD de Produtos

- **Criar Produto**: `POST /products`
- **Listar Produtos**: `GET /products`
- **Atualizar Produto**: `PUT /products/{id}`
- **Deletar Produto**: `DELETE /products/{id}`

### Produtor de Pedidos

- **Criar Pedido**: `POST /orders`

### Consumidor de Pedidos

O consumidor é iniciado automaticamente e lê as mensagens do Kafka, enviando-as para o MongoDB.

## Tecnologias Utilizadas

- Go
- Kafka
- MongoDB
- Docker
- Docker Compose

## Contribuição

1. Faça um fork do projeto.
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`).
3. Commit suas mudanças (`git commit -am 'Adiciona nova feature'`).
4. Faça o push para a branch (`git push origin feature/nova-feature`).
5. Abra um Pull Request.

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo LICENSE para mais detalhes.
