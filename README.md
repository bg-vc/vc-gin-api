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
http://127.0.0.1:8080/api/account/login -d'{"name":"James", "password":"admin123456"}'
```
**Response** :

```
{
    "code":0,
    "msg":"ok",
    "data":{
        "name":"James",
        "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Njk1NTU1NzksImlhdCI6MTU2OTU1NTI3OSwiaWQiOjEsIm5iZiI6MTU2OTU1NTI3OSwidXNlcm5hbWUiOiJKYW1lcyJ9.fsfgR494cmJiO7Jq0uRJqOdVGz9icZ5JnkHtWcbTiFs"
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
http://127.0.0.1:8080/api/account/updatePass -d'{"name":"admin", "oldPwd":"admin", "newPwd":"admin123456"}'
```
**response** :

```
{
    "code":0,
    "msg":"ok",
    "data":null
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

**Url** : `http://127.0.0.1:8080/api/user/common/queryAll`

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
    "msg":"ok",
    "data":null
}
```