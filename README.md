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
go get -u gorm.io/driver/sqlite
```

**Konfigurasi docker run countener**

```
docker compose up -d
```

**docker untuk cek mysql**

```
docker ps
```

**docker compose stop (stop dan menghapus countener)**

```
docker compose down
```

**docker di hentikan tampa menhapus**

```
docker compose stop
```

**perbaruan configurasi docker**

```
docker compose down -v
```
