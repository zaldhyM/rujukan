**konfigurasi api**
``` 
cd api/ 
go mod init rujukan
```

**install go fiber**

```
go get -u github.com/gin-gonic/gin
```

**install load env**

``` 
go get -u github.com/joho/godotenv
```


**Koneksi db**

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

**Konfigurasi docker run countener**

```
docker compose up -d
```

**docker untuk cek mysql**

```
docker ps
```

**docker masuk ke terminal mysql**

```
docker exec -it mysql_server mysql -u root -p
```

**docker compose (stop dan menghapus countener)**

```
docker compose down
docker compose down -v
```

**docker di hentikan tampa menhapus**

```
docker compose stop
```

