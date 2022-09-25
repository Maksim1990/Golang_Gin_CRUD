# Golang Gin Framework Template (+ MySQL setup) 

Dockerized Go Gin-gonic framework initial setup with MySQL.

### Start Gin framework manual command
``docker-compose up -d``


#### (Optional) Gin is already started automatically by command in docker-compose file. Incase need manual start of Gin then use following commands:
``docker exec -it gin_app sh``

``gin -i -all run main.go``

### After it navigate to `http://127.0.0.1:9091/ping`