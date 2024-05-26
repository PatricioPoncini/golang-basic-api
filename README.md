# RESTFUL API básica en Go 🐭

RESTFUL API básica desarrollada con [Go](https://go.dev/) y [Gin](https://gin-gonic.com/) siguiendo los pasos de la [guía oficial para desarrollar una RESTFUL API](https://go.dev/doc/tutorial/web-service-gin) básica en este lenguaje

## Clonar el repositorio
1. Primero se debe clonar el repositorio en la ubicación que se desee utilizando `git clone https://github.com/PatricioPoncini/golang-basic-api.git`.
2. Una vez clonado, ejecutamos `cd golang-basic-api` para entrar al proyecto.
3. Ejecutamos `go mod tidy` para instalar dependencias.
4. Ya se puede levantar localmente ejecutando `go run .` y se le puede pegar a los endpoints. La API estará escuchando en el puerto `8080`

## Endpoints
|HTTP method|Description|Path|
|-----------|-----------|----|
|POST|Save new album|http://localhost:8080/albums|
|GET|Obtain all albums|http://localhost:8080/albums|
|GET|Obtain album by id|http://localhost:8080/albums/:id|