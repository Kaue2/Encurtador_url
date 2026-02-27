# 🔗 Encurtador de URL

Um serviço rápido e eficiente de encurtamento de URLs construído em **Go**, utilizando **PostgreSQL** para persistência de dados e **Redis** para um sistema de cache de alta performance.



## 🚀 O Projeto

Este projeto tem como objetivo receber uma URL longa e retornar uma versão curta e amigável. Para garantir um tempo de resposta baixíssimo e reduzir a carga no banco de dados principal, o sistema utiliza o Redis como camada de cache para as URLs acessadas com mais frequência.

### 🛠️ Tecnologias Utilizadas

* **Linguagem:** [Go (Golang)](https://go.dev/)
* **Banco de Dados:** [PostgreSQL](https://www.postgresql.org/)
* **Cache:** [Redis](https://redis.io/)
* **Containers:** [Docker](https://www.docker.com/) (Recomendado para facilitar a execução)

---

## ⚙️ Pré-requisitos

Antes de começar, você vai precisar ter instalado em sua máquina as seguintes ferramentas:
* [Git](https://git-scm.com)
* [Go](https://go.dev/doc/install) (versão 1.20+)
* [PostgreSQL](https://www.postgresql.org/download/) e [Redis](https://redis.io/download) (opcional docker tem as imagens)
* [Docker](https://docs.docker.com/get-docker/) e `docker-compose` para rodar os bancos em containers.

---

## 🏃‍♂️ Como executar o projeto

### 1. Clonando o repositório

```bash
git clone https://github.com/Kaue2/Encurtador_url.git
cd Encurtador_url
```
