### バージョン（前提条件）

1. Go バージョン: `1.20.1`
2. swag インストール

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```
3. DockerDesktop


# I. Projectの構成

---
```
project
│   README.md
│   Makefile
│   ...
│   main.go
│       
└───pkg
    ├── config
    │       ├── config.go
    │       └── config.yml
    ├── controller
    │       ├── app.go
    │       ├── auth.go
    │       └── registry.go
    ├── domain
    │       └── model
    │       └── auth
    │               ├── access_token.go
    │               └── auth.go
    ├── repository
    │       ├── auth.go
    │       └── repository.go
    ├── router
    │       └── router.go
    ├── usecase
    │       └── auth.go
    └── util
            ├── common.go 
            ├── language.go 
            └── logger.go 
```
---

# II.  Swager 自動生成方法

```sh
    swag init
```

# III.  Docker Build
#### 1. Setup database on Docker

```sh
    sh ./bin/init_db.sh
```

###### Check connection
```sh
    docker exec -it postgres_container psql -U postgres

    // Then 
    // to get version
    SELECT VERSION(); 
    
    // To get all table
    \dt 
```

#### 2. 起動する

```sh
    docker compose up -d
```

### IV. リンク参照

1. Swagger: `http://localhost:8080/swagger/index.html#/`
2. Swager Json: `http://localhost:8080/swagger/doc.json`