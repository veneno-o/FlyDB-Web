# FlyDB-Web
Build the official website of FlyDB



## 接口文档

## 1.Third party login 

- 三方登录
- This application adopts OAuth 2.0 to support third-party login. Users can log in through GitHub or Twitter.

- 本应用采用OAuth 2.0支持第三方登录。 用户可以通过GitHub或Twitter登录。

**Example:**

```
GET /auth/github 
GET /auth/twitter 
```

**Authorization successful return in JSON format**:

```
{
    "status": "success",
    "message": "login success",
    "data": {
        "token": "<token>"
    }
}
```



## 2.DELETE string data 

- 删除数据

**Request example:**

```
DELETE /api/v1/kv/string/delete

Request header:
Content-Type: application/json
Authorization: Bearer <token>

Request body:
{
    "key": "key_example"
}
```

**Request params**

| Param name | Type   | Required | Description |
| ---------- | ------ | -------- | ----------- |
| key        | string | yes      | key         |

**Successful return example:**

```
{
    "code": "200",
    "message": "data delete success",
    "data": {
        "key": "key_example",
        "value": "value_example"
    }
}
```

**Failed return example:**

```
{
    "code": "400",
    "message": "data delete failed"
}
```



## 3.Get Keys list

**Request example:**

```
GET /api/v1/kv/string/keys

Request header:
Content-Type: application/json
Authorization: Bearer <token>

Request body:
{
    "prefix": "key_example"
}
```

**Request params**

| Param name | Type   | Required | Description |
| ---------- | ------ | -------- | ----------- |
| prefix     | string | yes      | key prefix  |

**Successful return example:**

```
{
    "code": "200",
    "message": "data get success",
    "data": {
        "keys": [
            "key_example1",
            "key_example2"
        ]
    }
}
```

**Failed return example:**

```
{
    "code": "400",
    "message": "data get failed"
}
```

## 4.PUT list data

**Request example:**

```
POST /api/v1/kv/list/put

Request header:
Content-Type: application/json
Authorization: Bearer <token>

Request body:
{
    "key": "key_example",
    "value": [
        "value_example1",
        "value_example2"
    ]
}
```

**Request params**

| Param name | Type   | Required | Description |
| ---------- | ------ | -------- | ----------- |
| key        | string | yes      | key         |
| value      | list   | yes      | value       |

**Successful return example:**

```
{
    "code": "200",
    "message": "data write success",
    "data": {
        "key": "key_example",
        "value": [
            "value_example1",
            "value_example2"
        ]
    }
}
```

**Failed return example:**

```
{
    "code": "400",
    "message": "data write failed"
}
```



## 5.GET list data

**Request example:**

```
GET /api/v1/kv/list/get

Request header:
Content-Type: application/json
Authorization: Bearer <token>

Request body:
{
    "key": "key_example"
}
```

**Request params**

| Param name | Type   | Required | Description |
| ---------- | ------ | -------- | ----------- |
| key        | string | yes      | key         |

**Successful return example:**

```
{
    "code": "200",
    "message": "data get success",
    "data": {
        "key": "key_example",
        "value": [
            "value_example1",
            "value_example2"
        ]
    }
}
```

**Failed return example:**

```
{
    "code": "400",
    "message": "data get failed"
}
```

