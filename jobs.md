# 1. Authorization
This application adopts OAuth 2.0 to support third-party login. Users can log in through GitHub or Twitter.
**Example:**
```bash
GET /auth/github 
GET /auth/twitter 
```
**Authorization successful return in JSON format**:
```json
{
    "status": "success",
    "message": "login success",
    "data": {
        "token": "<token>"
    }
}
```
# 2. Data operation
## 2.1 PUT string data
**Request example:**
```
POST /api/v1/kv/string/put

Request header:
Content-Type: application/json
Authorization: Bearer <token>

Request body:
{
    "key": "key_example",
    "value": "value_example"
}
```
**Request params**

| Param name | Type | Required | Description |
| :----: | :----: | :----: | :----: |
| key | string | yes | key |
| value | string | yes | value |

**Successful return example:**
```json
{
    "code": "200",
    "message": "data write success",
    "data": {
        "key": "key_example",
        "value": "value_example"
    }
}
```
**Failed return example:**
```json
{
    "code": "400",
    "message": "data write failed"
}
```
## 2.2 GET string data
**Request example:**
```
GET /api/v1/kv/string/get

Request header:
Content-Type: application/json
Authorization: Bearer <token>

Request body:
{
    "key": "key_example"
}
```
**Request params**

| Param name | Type | Required | Description |
| :----: | :----: | :----: | :----: |
| key | string | yes | key |


**Successful return example:**
```json
{
    "code": "200",
    "message": "data get success",
    "data": {
        "key": "key_example",
        "value": "value_example"
    }
}
```
**Failed return example:**
```json
{
    "code": "400",
    "message": "data get failed"
}
```
## 2.3 DELETE string data
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

| Param name | Type | Required | Description |
| :----: | :----: | :----: | :----: |
| key | string | yes | key |


**Successful return example:**
```json
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
```json
{
    "code": "400",
    "message": "data delete failed"
}
```

## 2.4 Get Keys list
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

| Param name | Type | Required | Description |
| :----: | :----: | :----: | :----: |
| prefix | string | yes | key prefix |

**Successful return example:**
```json
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
```json
{
    "code": "400",
    "message": "data get failed"
}
```

## 2.5 PUT list data
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

| Param name | Type | Required | Description |
| :----: | :----: | :----: | :----: |
| key | string | yes | key |
| value | list | yes | value |

**Successful return example:**
```json
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
```json
{
    "code": "400",
    "message": "data write failed"
}
```

## 2.6 GET list data
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

| Param name | Type | Required | Description |
| :----: | :----: | :----: | :----: |
| key | string | yes | key |

**Successful return example:**
```json
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
```json
{
    "code": "400",
    "message": "data get failed"
}
```



