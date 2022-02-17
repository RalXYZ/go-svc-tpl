# Golang Service Template

Golang back-end service template. Using this template, you can get started with back-end projects quickly.  

|  Web Framework   |     ORM      |   Database Driver    | Configuration Manager |   Log Manager   |  API Documentation  |
|:----------------:|:------------:|:--------------------:|:---------------------:|:---------------:|:-------------------:|
| labstack/echo/v4 | gorm.io/gorm | gorm.io/driver/mysql |      spf13/viper      | sirupsen/logrus | swaggo/echo-swagger |

## Configuration

1. Create database and user in MySQL, and grant privileges to user.  

```sql
CREATE DATABASE buz;
CREATE USER 'foo'@'localhost' IDENTIFIED BY 'bar';
GRANT ALL PRIVILEGES ON buz.* TO 'foo'@'localhost';
```

2. Create a file named `conf.yaml` in the project's root directory.  

```yaml
sql:
  username: foo
  password: bar
  db_name: buz
```

3. To generate API doc, you need to execute the following commands: 

```shell
go install github.com/swaggo/swag/cmd/swag@latest  # download swag, a doc generator
swag init  # docs will be written into `docs/` directory
```

4. Compile and run this project. You will find API doc in `http://localhost:1323/api/doc/index.html` on default.  
5. Change the go module name of this project to whatever you like.  

## Usage

1. To create tables, add structs in `model/`. `model/foo.go` is an example of such struct. Then, add the struct to `DB.AutoMigrate` in `model/init.go`. 
2. Add business logics in `app/controller/`, then add router in `app/routers.go`. 
3. Write API doc alongside with the controller function. `app/controller/foo.go` is an example.  
