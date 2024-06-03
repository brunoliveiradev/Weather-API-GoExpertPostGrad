### Languages: [Português 🇧🇷](#clima-por-cep-weather-api) | [English 🇨🇦](#weather-by-zipcode-api)

---

# Clima por CEP (Weather API)

Este projeto implementa uma API em Golang que recebe um CEP válido e retorna as temperaturas da localidade
correspondente em Celsius, Fahrenheit e Kelvin. A API faz parte do desafio técnico do curso de Pós-Graduação em
Engenharia de Software [GoExpert](https://goexpert.fullcycle.com.br/pos-goexpert/).

## ⚙️ Configuração

Você precisará das seguintes tecnologias abaixo:

- [Docker](https://docs.docker.com/get-docker/) 🐳
- [Docker Compose](https://docs.docker.com/compose/install/) 🐳
- [GIT](https://git-scm.com/downloads)
- [Postman ☄️](https://www.postman.com/downloads/) ou [VS Code](https://code.visualstudio.com/download) com a
  extensão [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) instalada.
- [Google Cloud SDK](https://cloud.google.com/sdk/docs/install#mac) (opcional, para testes no Google Cloud Run)

## 🚀 Iniciando

1. Clone o repositório e entre no diretório do projeto.
   ```sh
   git clone https://github.com/brunoliveiradev/Weather-API-GoExpertPostGrad.git
   cd Weather-API-GoExpertPostGrad
   ```

2. Execute o comando abaixo na pasta raiz do projeto para iniciar o ambiente de desenvolvimento:
   ```sh
   docker compose up --build -d
   ```

   Para parar os serviços:
   ```sh
   docker compose down
   ```

3. A API estará disponível em `http://localhost:8080` 🚀.

## 🧪 Testes

1. Para testar localmente você pode usar os cURLS de exemplo dentro do arquivo `api/get_temperature.http` exemplo:
    ```sh
    curl http://localhost:8080/01001000
    ```
2. Resposta esperada:
   ```json
    {
        "temp_C": 28.5,
        "temp_F": 83.3,
        "temp_K": 301.5
    }
    ```

## 🛠️ Deploy e Teste no Google Cloud Run

A aplicação já está disponível no Google Cloud Run no seguinte
URL: `https://weather-by-zipcode-api-4w3swr2nrq-rj.a.run.app`

1. Para testar no Google Cloud Run, você pode usar o Postman e o cURL de exemplo dentro do
   arquivo `api/get_temperature.http`
   exemplo:
2. Cep Válido deve retornar `200` e o JSON com as temperaturas:
    ```sh
    curl https://weather-by-zipcode-api-4w3swr2nrq-rj.a.run.app/01001000
    ```
3. Cep não encontrado deve retornar `404` com `can not find zipcode`:
    ```sh
    curl https://weather-by-zipcode-api-4w3swr2nrq-rj.a.run.app/89216369
    ```
4. CEP inválido deve retornar `422` com `invalid zipcode`:
    ```sh
    curl https://weather-by-zipcode-api-4w3swr2nrq-rj.a.run.app/123456
    ```

---

# Weather By Zipcode API

This project implements an API in Golang that receives a valid zip code and returns the temperatures of the
corresponding
location in Celsius, Fahrenheit and Kelvin. The API is part of the technical challenge of the Postgraduate course in
Software Engineering [GoExpert](https://goexpert.fullcycle.com.br/pos-goexpert/).

## ⚙️ Configuration

You will need the following technologies below:

- [Docker](https://docs.docker.com/get-docker/) 🐳
- [Docker Compose](https://docs.docker.com/compose/install/) 🐳
- [GIT](https://git-scm.com/downloads)
- [Postman ☄️](https://www.postman.com/downloads/) ou [VS Code](https://code.visualstudio.com/download) with
  the [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) extension installed.
- [Google Cloud SDK](https://cloud.google.com/sdk/docs/install#mac) (optional, for tests on Google Cloud Run)

## 🚀 Getting Started

1. Clone the repository and navigate to the project directory.
   ```sh
   git clone https://github.com/brunoliveiradev/Weather-API-GoExpertPostGrad.git
   cd Weather-API-GoExpertPostGrad
   ```
2. Run the command below in the project root folder to start the development environment:
   ```sh
    docker compose up --build -d
    ```
   To stop the services:
    ```sh
    docker compose down
    ```
3. The API will be available at `http://localhost:8080` 🚀.
4. To test locally you can use the example cURLS inside the `api/get_temperature.http` file example:
    ```sh
    curl http://localhost:8080/01001000
    ```
Expected response:
```json
    {
        "temp_C": 28.5,
        "temp_F": 83.3,
        "temp_K": 301.5
    }
```

## 🛠️ Deploy and Test on Google Cloud Run

The application is already available on Google Cloud Run at the following
URL: `https://weather-by-zipcode-api-4w3swr2nrq-rj.a.run.app`

1. To test on Google Cloud Run, you can use Postman and the example cURL inside the `api/get_temperature.http` file
   example:
2. Valid Zipcode should return `200` and the JSON with the temperatures:
    ```sh
    curl https://weather-by-zipcode-api-4w3swr2nrq-rj.a.run.app/01001000
    ```
3. Zipcode not found should return `404` with `can not find zipcode`:
    ```sh
    curl https://weather-by-zipcode-api-4w3swr2nrq-rj.a.run.app/89216369
    ```
4. Invalid Zipcode should return `422` with `invalid zipcode`:
    ```sh
    curl https://weather-by-zipcode-api-4w3swr2nrq-rj.a.run.app/123456
    ```
      
