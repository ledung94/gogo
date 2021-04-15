huong dan chay postgres trong docker

1. docker pull postgres

docker run --name my-postgres -d -p 2345:5432 -e POSTGRES_PASSWORD=ai_khoc_noi_dau_nay postgres

docker run --name postgres -d -p 2345:5432 -e POSTGRES_PASSWORD="11235813" postgres

3. $ docker exec -it [my-postgres] bash

4. psql -U postgres

Connecting to a database

PostgreSQL can have multiple databases inside of it. In order to connect to one, you need to run the following command


\c <database_name>
Viewing the tables

To check all the tables that exist inside a database, run
CREATE DATABASE mytestdb;
\d
To check the details of a particular table, run

\d+ <table_name>
Exiting the container

To exit the container, run the following command

\q


5. GORM libs

batch insert only use in gorm v2
	"gorm.io/gorm"
  	"gorm.io/driver/postgres"
In this version, dont have close func. So we must create Generic database interface sql.DB -> close
gorm v1 is 
    "github.com/jinzhu/gorm"
    "github.com/jinzhu/gorm/dialects/postgres"