<h1 align="center">CRUD API SERVICE</h1>


<p align="center">
 <img alt="Golang" src="https://img.shields.io/badge/Golang-4285f4?style=for-the-badge&logo=go&labelColor=black">

 <img alt="PostgreSQL" src="https://img.shields.io/badge/PostgreSQL-%23212121?style=for-the-badge&logo=postgreSQL&labelColor=black">

 <img alt="Docker" src="https://img.shields.io/badge/Docker-%231d63ed?style=for-the-badge&logo=docker&labelColor=black">
</p>
<hr>

![minishell](https://i.imgur.com/jJxqbjv.png)


## ABOUT ##

This is a service written in Go that provides a user-friendly interface for managing users and books via a REST API. The service supports user registration and authentication using JWT tokens, ensuring security and efficiency. The service provides the ability to create, update, delete and obtain information about books, access to which is limited to authenticated users.

Swagger documentation is provided for the convenience of users and developers. Interaction with the audit-log-server is carried out via gRPC to write activity logs to the MongoDB database.

The service uses PostgreSQL to store information about users, books, and refresh tokens, and MongoDB to store logs. Both databases run in Docker containers for ease of deployment and scalability.

Final technology stack: Go, Gin, REST API, gRPC, PostgreSQL, MongoDB, Docker, JWT. This service is an excellent example of modern, secure and scalable web development.

<hr>

## ⚡ STARTING ##

```bash
# Clone this project
$ git clone https://github.com/andy-ahmedov/crud_service

# Up database
$ make up

# Create Tables
$ make create_table

# Create Swagger documentation
$ make swag

# Run the project
$ make run

```

## 🔨 STACK ##
  <a href="https://git-scm.com/" target="_blank"> <img src="https://raw.githubusercontent.com/rahul-jha98/github_readme_icons/main/language_and_tools/square/git-scm/git-scm.svg" align="left" alt="git" height='50px'/> </a>
  <a href="https://go.dev/" target="_blank"> <img src="https://github.com/andy-ahmedov/andy-ahmedov/blob/main/files/Go-Logo_Yellow.png?raw=true" alt="golang" align="left" height='50px'/> </a>
  <a href="https://www.gnu.org/software/bash/" target="_blank"> <img src="https://raw.githubusercontent.com/andy-ahmedov/README_icons/5a0bd0723991e5d95e0eb90ce4e544345b69e05b/language_and_tools/square/bash/bash.svg" alt="bash" align="left" height='50px'/> </a>
  <a href="https://www.docker.com/" target="_blank"> <img src="https://raw.githubusercontent.com/andy-ahmedov/README_icons/5a0bd0723991e5d95e0eb90ce4e544345b69e05b/language_and_tools/square/docker/docker.svg" alt="docker" align="left" height='50px'/> </a>
  <a href="https://www.openapis.org/" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/6910f0503efdd315c8f9b858234310c06e04d9c0/icons/openapi/openapi-plain-wordmark.svg" align="left" alt="openapi" height='50px'/> </a>
  <a href="https://www.postgresql.org/" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/6910f0503efdd315c8f9b858234310c06e04d9c0/icons/postgresql/postgresql-plain-wordmark.svg" align="left" alt="postgres" height='50px'/> </a>
  <a href="https://jwt.io/" target="_blank"> <img src="https://store-images.s-microsoft.com/image/apps.2250.63b92ddd-7108-4be3-8fb4-8c8137f54bcf.2427e723-5f43-491f-9b30-fd1cfe84a62f.55ccd47a-09f3-4319-abd5-7eb798ed76d1?mode=scale&h=100&q=90&w=100" align="left" alt="jwt" height='45px'/> </a>
  <a href="https://grpc.io/" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/6910f0503efdd315c8f9b858234310c06e04d9c0/icons/grpc/grpc-original.svg" align="left" alt="grpc" height='50px'/> </a>
  <a href="https://www.json.org/json-en.html" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/6910f0503efdd315c8f9b858234310c06e04d9c0/icons/json/json-plain.svg" align="left" alt="json" height='50px'/> </a>
  <a href="https://restfulapi.net/" target="_blank"> <img src="https://repository-images.githubusercontent.com/436490321/0e58e265-ef30-40b1-b439-8b302b06ec42" align="left" alt="rest-api" height='50px'/> </a>
  <a href="https://swagger.io/tools/swagger-ui/?ref=the-best-recipe" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/6910f0503efdd315c8f9b858234310c06e04d9c0/icons/swagger/swagger-original-wordmark.svg" align="left" alt="swagger" height='50px'/> </a>
  <a href="https://code.visualstudio.com/" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/6910f0503efdd315c8f9b858234310c06e04d9c0/icons/visualstudio/visualstudio-plain-wordmark.svg" align="left" alt="vscode" height='50px'/> </a>

