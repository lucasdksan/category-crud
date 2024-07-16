# Category CRUD

Simples CRUD usando GO.

# Descrição

Este projeto tem como objetivo simular um sistema CRUD (Create, Read, Update, Delete) para categorias aleatórias utilizando o framework Gin em Go. Todos os dados são armazenados em um slice na memória, o que significa que os dados persistem apenas enquanto o sistema estiver em execução.

# Objetivo

* Aprender e testar o framework Gin: Este projeto serve como um exercício prático para entender e utilizar o Gin, um framework web rápido e leve para Go.

* Fortalecer os conceitos de Go: Implementar um CRUD completo ajudará a consolidar conhecimentos essenciais sobre a linguagem Go, incluindo manipulação de slices, estruturas de dados, tratamento de erros, e manipulação de dados JSON.

* Simular um CRUD completo: Permitir a criação, leitura, atualização e exclusão de categorias em um ambiente simulado, semelhante a um banco de dados em memória.

# Funcionalidades

* Create: Adicionar novas categorias.

* Read: Listar todas as categorias.

* Update: Atualizar categorias existentes.

* Delete: Remover categorias por ID.

# Estrutura do Projeto

* Controladores: Gerenciam as requisições HTTP e chamam os casos de uso apropriados.

* Casos de Uso: Contêm a lógica de negócios para cada operação CRUD.

* Repositórios: Interface para o armazenamento dos dados, implementado com slices em memória.

* Entidades: Definem as estruturas de dados para as categorias.

# Referências

* [ChatGPT](https://chatgpt.com/)
* [Gin](https://gin-gonic.com/docs/)
