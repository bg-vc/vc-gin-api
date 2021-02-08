## Server Run
```shell script
cd $GOPATH/src/vc-gin-api
go build
./admin.sh start
```

## 1. Account Login

**Url** : `http://127.0.0.1:8080/api/account/login`

**Method** : `POST`

**Params** : `name`, `password`

**Auth Required** : `No`

**Request** :

```shell script
curl -X POST -H "Content-Type:application/json" \
http://127.0.0.1:8080/api/account/login -d'{"name":"admin", "password":"admin123456"}'
```
**Response** :

```
{
    "code":0,
    "msg":"ok",
    "data":{
        "name":"admin",
        "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTI3NTA4ODIsImlhdCI6MTYxMjc0OTA4MiwiaWQiOjEsIm5iZiI6MTYxMjc0OTA4MiwidXNlcm5hbWUiOiJhZG1pbiJ9.MmmQMd9V596vCbZaXx758bNSKHJdD5_tsD-K-RU3RP8"
    }
}
```

## 2. Account Password Update

**Url** : `http://127.0.0.1:8080/api/account/updatePass`

**method** : `POST`

**Params** : `name`, `oldPwd`, `newPwd`

**Auth Required** : `No`

**request** :

```shell script
curl -X POST -H "Content-Type:application/json" \
http://127.0.0.1:8080/api/account/updatePass -d'{"name":"admin", "oldPwd":"admin123456", "newPwd":"admin666666"}'
```
**response** :

```
{
    "code":0,
    "msg":"ok"
}
```

## 3. Query Users

**Url** : `http://127.0.0.1:8080/api/user/common/queryAll`

**Method** : `GET`

**Params** : `start`, `limit`, `name`, `address`

**Auth Required** : `No`

**Request** :

```shell script
curl http://127.0.0.1:8080/api/user/common/queryAll
```

**Response**  :

```
{
    "code":0,
    "msg":"ok",
    "data":{
        "total":2,
        "rows":[
            {
                "id":1,
                "name":"vince",
                "age":24,
                "address":"beijing"
            },
            {
                "id":2,
                "name":"join",
                "age":25,
                "address":"shanghai"
            }
        ]
    }
}
```

## 4. Add User

**Url** : `http://127.0.0.1:8080/api/user/auth/add`

**Method** : `POST`

**Params** : `name`, `age`, `address`

**Auth Required** : `Yes`

**Request** :

```shell script
curl -X POST -H "Content-Type:application/json" \
http://127.0.0.1:8080/api/user/auth/add -d'{"name":"david", "age":28, "address":"shenzhen"}' \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Njk1NTc3NjIsImlhdCI6MTU2OTU1NzE2MiwiaWQiOjEsIm5iZiI6MTU2OTU1NzE2MiwidXNlcm5hbWUiOiJKYW1lcyJ9.-WTzVOqRe6-Ax3M4TpGIACjEqef8PqzcKArLF4_gmRE"
```

**Response**  :
```shell script
{
    "code":0,
    "msg":"ok",
    "data":null
}
```

## 4. Update User

**Url** : `http://127.0.0.1:8080/api/user/common/queryAll`

**Method** : `POST`

**Params** : `name`, `age`, `address`

**Auth Required** : `Yes`

**Request** :

```shell script
curl -X POST -H "Content-Type:application/json" \
http://127.0.0.1:8080/api/user/auth/update -d'{"name":"david", "age":30, "address":"beijing"}' \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Njk1NTc3NjIsImlhdCI6MTU2OTU1NzE2MiwiaWQiOjEsIm5iZiI6MTU2OTU1NzE2MiwidXNlcm5hbWUiOiJKYW1lcyJ9.-WTzVOqRe6-Ax3M4TpGIACjEqef8PqzcKArLF4_gmRE"
```

**Response**  :
```
{
    "code":0,
    "msg":"ok"
}
```