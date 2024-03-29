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

In case you want to execute a local database using docker modify **docker-compose.yml** with your data. 

```
 version: '3'

services:
  db:
    image: mysql
    command: --default-authentication-plugin=mysql_native_password
    restart: always
    environment:
      MYSQL_DATABASE: exampleDatabase
      MYSQL_USER: exampleUser
      MYSQL_PASSWORD: examplePassword
      MYSQL_ROOT_PASSWORD: exampleRootPassword
    ports:
      - "3306:3306"
    volumes:
      - ./mysql_data:/var/lib/mysql
```

Now to run your local database using docker use the following command: 

```bash
 docker-compose up -d 
```


## How to run the app:
To run this app you have to execute the following command: 

```bash
 go run cmd/main.go
```


## Contributions:
[![Stargazers repo roster for @USERNAME/REPO_NAME](https://reporoster.com/stars/bypepe77/spotify-clone-go)](https://github.com/bypepe77/spotify-clone-go/stargazers)
[![Forkers repo roster for @USERNAME/REPO_NAME](https://reporoster.com/forks/bypepe77/spotify-clone-go)](https://github.com/bypepe77/spotify-clone-go/network/members)

