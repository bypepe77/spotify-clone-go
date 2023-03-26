<div align="center">
 <h1>Spotify clone api with go</h1>
    <span><strong>Spotify clone </strong>api for go it's a clone of spotify only for learning purpose.</span><br />
</div>

## Prerequisites to execute:

Clone or fork this repository. 

Create a .env with this data: 


```
PORT="8080"
GIN_MODE="debug"
MYSQL_USERNAME= "your_username"
MYSQL_PASSWORD= "your_password"
MYSQL_HOST= "your_host"
MYSQL_PORT= "your_mysql_port"
MYSQL_DATABASE= "your_database_name"
APP_NAME = "your_app_name
HOST = "Were your app will be running, example: localhost"
```

You can run a local databse running this command: 

```bash
 docker-compose up -d 
```

## How to run the app:
To run this app you have to execute the following command: 

```bash
 go run cmd/main.go
```


