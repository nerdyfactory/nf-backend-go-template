# KED-api-server

## Installation

Make sure you have [**go**](https://golang.org/doc/install) and [**docker**](https://www.docker.com/products/docker-desktop) installed.

### Docker

1. download oracle xe install file from https://download.oracle.com/otn-pub/otn_software/db-express/oracle-database-xe-18c-1.0-1.x86_64.rpm
2. Place downloaded file to `docker/db`
3. Make sure the install file name is `oracle-database-xe-18c-1.0-1.x86_64.rpm`
4. build image

```
$ docker-compose build
```

5. run docker (This will take some time at the first time)

```
$ docker-compose up
```

or running in background

```
$ docker-compose up -d
```

6. create DB and import data

```
$ cd docker/db
$ ./init.sh
```

### Go

1. Install dependencies.

```
$ go mod vendor
```

2. Run test

```
$ make test
```

or

```
$ go test ./internal/routes -run YourCaseName # run single case
```

3. Run server

```
$ make start
```
