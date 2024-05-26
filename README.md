# RESTFUL API básica en Go 🐭

RESTFUL API básica desarrollada con [Go](https://go.dev/) y [Gin](https://gin-gonic.com/) siguiendo los pasos de la [guía oficial para desarrollar una RESTFUL API](https://go.dev/doc/tutorial/web-service-gin) básica en este lenguaje

## Endpoints
|HTTP method|Description|Path|
|-----------|-----------|----|
|POST|Save new album|http://localhost:8080/albums|
|GET|Obtain all albums|http://localhost:8080/albums|
|GET|Obtain album by id|http://localhost:8080/albums/:id|