# somewhere

## 数据库设计

### stores

| name        | type     | meaning      |
| ----------- | -------- | ------------ |
| store_id    | ObjectId | 商户id       |
| store_name  | string   | 商户名称     |
| store_score | double   | 商户评分     |
| store_city  | string   | 商户所在城市 |

### users

| name            | type     | meaning                   |
| --------------- | -------- | ------------------------- |
| user_id         | ObjectId | 用户id                    |
| user_name       | string   | 用户名称                  |
| user_gender     | int      | 用户性别，0未知，1男，2女 |
| user_age        | int      | 用户年龄，0未知           |
| user_city       | string   | 用户城市                  |
| user_timestamp  | int      | 用户创建时间              |
| user_historysum | double   | 用户历史消费              |

### items

| name           | type     | meaning      |
| -------------- | -------- | ------------ |
| item_id        | ObjectId | 商品id       |
| store_id       | ObjectId | 商品的商铺id |
| item_name      | string   | 商品名称     |
| item_score     | double   | 商品评分     |
| item_price     | double   | 商品价格     |
| item_salecount | int      | 商品销量     |
| item_brand     | string   | 商品品牌     |
| item_timestamp | int      | 商品创建时间 |

### records

| name            | type     | meaning                   |
| --------------- | -------- | ------------------------- |
| record_id       | ObjectId | 记录id                    |
| user_id         | ObjectId | 用户id                    |
| store_id        | ObjectId | 商户id                    |
| item_id         | ObjectId | 商品id                    |
| is_trade        | int      | 交易信息，0没交易，1交易  |
| timestamp       | int      | 交易时间                  |
| store_city      | string   | 商户所在城市              |
| user_gender     | int      | 用户性别，0未知，1男，2女 |
| user_age        | int      | 用户年龄，0未知           |
| user_city       | string   | 用户城市                  |
| user_historysum | double   | 用户历史消费              |
| item_score      | double   | 商品评分                  |
| item_price      | double   | 商品价格                  |
| item_salecount  | int      | 商品销量                  |
| item_brand      | string   | 商品品牌                  |

### recommend

| name    | type   | meaning                     |
| ------- | ------ | --------------------------- |
| user_id | string | 用户id                      |
| list    | list   | 推荐物品列表，内容是item_id |

## 接口设计

### SuperAdmin部分



#### 获得所有商户信息

- Method : HTTP GET
- Request URL: /somewhere/stores
- Response:

```json
{
    "list": [
        {
            "store_id": "5ddd03339156005687913584",
            "store_name": "素云满湖",
            "store_level": 2.33,
            "store_city": "敦煌"
        }
    ],
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574821448.8081"
}
```



#### 增加某个商户

- Method : HTTP POST
- Request URL: /somewhere/stores
- Request: raw json

```json
{
    "store_name": "湖",
    "store_level": 2.33,
    "store_city":"敦煌"
}
```

- Response:

```json
{
    "store_id": "5ddddcf891560069a33cb544",
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574821112.1847"
}
```





#### 删除某个商户

- Method : HTTP DELETE
- Request URL: /somewhere/stores
- Request: form data

|   key    |          value           |
| :------: | :----------------------: |
| store_id | 5ddddcf891560069a33cb544 |

- Response:

```json
{
    "delete_success_num": 1,
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574823102.4059"
}
```



#### 更新某个商户

- Method : HTTP PUT
- Request URL: /somewhere/stores
- Request: raw json

```json
{
    "store_id": "5ddddcf891560069a33cb544",
    "store_name": "湖",
    "store_level": 2.33,
    "store_city":"敦煌"
}
```

- Response:

```json
{
    "update_sucess_num": 1,
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574824271.1847"
}
```



 #### 获得所有用户

- Method : HTTP GET
- Request URL: /somewhere/users
- Response:

```json
{
    "list": [
        {
            "user_id": "5ddcfb6425c5b119f58a17c4",
            "user_name": "素云满湖",
            "user_age": 22,
            "user_gender": 1,
            "user_city": "长春",
            "user_timestamp": 1574763364,
            "user_historysum": 2.33
        }
    ],
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574907357.7887"
}
```



#### 增加某个用户

- Method : HTTP POST
- Request URL: /somewhere/users
- Request: raw json

```json
{
    "user_name": "素满湖",
    "user_gender": 1,
    "user_age":22,
    "user_city": "长春",
    "user_historysum": 2.33
}
```

- Response:

```json
{
    "user_id": "5ddf2e449156000de1e03acd",
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574907460.1847"
}
```



#### 更改某个用户

- Method : HTTP PUT
- Request URL: /somewhere/users
- Request: raw json

```json
{
	"user_id":"5ddf2e449156000de1e03acd",
    "user_name": "素满湖",
    "user_gender": 1,
    "user_age":22,
    "user_city": "长春",
    "user_historysum": 2.33
}
```

- Response:

```json
{
    "update_sucess_num": "5ddf2e449156000de1e03acd",
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574907588.2081"
}
```



#### 删除某个用户

- Method : HTTP DELETE
- Request URL: /somewhere/users
- Request: form data

|   key   |          value           |
| :-----: | :----------------------: |
| user_id | 5ddcfab325c5b119f58a1786 |

- Response:

```json
{
    "delete_success_num": "5ddcfab325c5b119f58a1786",
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574907666.4425"
}
```



#### 获得所有商品

- Method : HTTP GET
- Request URL: /somewhere/products
- Response:

```json
{
    "list": [
        {
            "item_id": "5ddcd1f691560037116545cd",
            "store_id": "5ddb7289e2ea5cecbe4605b6",
            "item_name": "ssdasxa",
            "item_price": 2.3,
            "item_score": 4.5523423,
            "item_salecount": 345252,
            "item_brand": "wahaha",
            "item_timestamp": 1574752758
        }
    ],
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574907755.2540"
}
```



#### 增加某个商品

- Method : HTTP POST
- Request URL: /somewhere/products
- Request: raw json

```json
{
    "store_id": "5ddb7289e2ea5cecbe4605b6",
    "item_name": "矿水",
    "item_price": 2.3,
    "item_score": 4.5523423,
    "item_salecount": 345252,
    "item_brand": "wahaha"
}
```

- Response:

```json
{
    "item_id": "5ddf300b9156000de1e03acf",
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574907915.694"
}
```



#### 删除某个商品

- Method : HTTP DELETE
- Request URL: /somewhere/products
- Request: form data

|   key   |          value           |
| :-----: | :----------------------: |
| item_id | 5ddddcf891560069a33cb544 |

- Response:

```json
{
    "delete_success_num": 1,
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574823102.4059"
}
```



#### 更改某个商品

- Method : HTTP PUT
- Request URL: /somewhere/products
- Request: raw json

```json
{
	"item_id":"5ddcdb249156004ffaa78b27",
    "store_id": "5ddb7289e2ea5cecbe4605b6",
    "item_name": "矿水",
    "item_price": 2.3,
    "item_score": 4.5523423,
    "item_salecount": 345252,
    "item_brand": "wahaha"
}
```

- Response:

```json
{
    "update_sucess_num": 0,
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574908012.8162"
}
```



#### 获得所有记录

- Method : HTTP GET
- Request URL: /somewhere/records
- Response:

```json
{
    "list": [
        {
            "record_id": "5ddb7289e2ea5cecbe4605b6",
            "user_id": "",
            "item_id": "",
            "is_trade": 1,
            "timestamp": 1574662793
        }
    ],
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574907853.456"
}
```





### 商户部分



#### 更改自己商户的信息

- Method : HTTP PUT
- Request URL: /somewhere/users
- Request: raw json

```json
{
	"user_id":"5ddf2e449156000de1e03acd",
    "user_name": "素满湖",
    "user_gender": 1,
    "user_age":22,
    "user_city": "长春",
    "user_historysum": 2.33
}
```

- Response:

```json
{
    "update_sucess_num": "5ddf2e449156000de1e03acd",
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574907588.2081"
}
```



#### 在自己的店铺增加某个商品

- Method : HTTP POST
- Request URL: /somewhere/products
- Request: raw json

```json
{
    "store_id": "5ddb7289e2ea5cecbe4605b6",
    "item_name": "矿水",
    "item_price": 2.3,
    "item_score": 4.5523423,
    "item_salecount": 345252,
    "item_brand": "wahaha"
}
```

- Response:

```json
{
    "item_id": "5ddf300b9156000de1e03acf",
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574907915.694"
}
```



#### 在自己的店铺删除某个商品

- Method : HTTP DELETE
- Request URL: /somewhere/products
- Request: form data

|   key   |          value           |
| :-----: | :----------------------: |
| item_id | 5ddddcf891560069a33cb544 |

- Response:

```json
{
    "delete_success_num": 1,
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574823102.4059"
}
```



#### 在自己的店铺更改某个商品

- Method : HTTP PUT
- Request URL: /somewhere/products
- Request: raw json

```json
{
	"item_id":"5ddcdb249156004ffaa78b27",
    "store_id": "5ddb7289e2ea5cecbe4605b6",
    "item_name": "矿水",
    "item_price": 2.3,
    "item_score": 4.5523423,
    "item_salecount": 345252,
    "item_brand": "wahaha"
}
```

- Response:

```json
{
    "update_sucess_num": 0,
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574908012.8162"
}
```



#### 获得自己所有的商品

- Method : HTTP GET
- Request URL: /somewhere/products
- Request: form data


|   key    |          value           |
| :------: | :----------------------: |
| store_id | 5ddb7289e2ea5cecbe4605b6 |

- Response:

```json
{
    "list": [
        {
            "item_id": "5ddcd1f691560037116545cd",
            "store_id": "5ddb7289e2ea5cecbe4605b6",
            "item_name": "ssdasxa",
            "item_price": 2.3,
            "item_score": 4.5523423,
            "item_salecount": 345252,
            "item_brand": "wahaha",
            "item_timestamp": 1574752758
        }
    ],
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574907755.2540"
}
```



#### 查看自己某个商品的所有记录

- Method : HTTP GET
- Request URL: /somewhere/records
- Request: form data


|   key   |          value           |
| :-----: | :----------------------: |
| item_id | 5ddf2e449156000de1e03acd |

- Response:

```json
{
    "list": [
        {
            "record_id": "5ddb7289e2ea5cecbe4605b6",
            "user_id": "",
            "item_id": "",
            "is_trade": 1,
            "timestamp": 1574662793
        }
    ],
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574907853.456"
}
```





### 用户部分



#### 更新自己的信息

- Method : HTTP PUT
- Request URL: /somewhere/users
- Request: raw json

```json
{
	"user_id":"5ddf2e449156000de1e03acd",
    "user_name": "素满湖",
    "user_gender": 1,
    "user_age":22,
    "user_city": "长春",
    "user_historysum": 2.33
}
```

- Response:

```json
{
    "update_sucess_num": "5ddf2e449156000de1e03acd",
    "error_code": 0,
    "error_msg": "",
    "request_id": "1574907588.2081"
}
```



#### 查看推荐的列表

- Method : HTTP GET
- Request URL: /somewhere/recommend
- Request: form data


|   key   |          value           |
| :-----: | :----------------------: |
| user_id | 5ddf2e449156000de1e03acd |

- Response:

```json
{
    "list": [
        {
            "item_id": "5ddcd1f691560037116545cd",
            "store_id": "5ddb7289e2ea5cecbe4605b6",
            "item_name": "ssdasxa",
            "item_price": 2.3,
            "item_score": 4.5523423,
            "item_salecount": 345252,
            "item_brand": "wahaha",
            "item_timestamp": 1574752758
        }
    ],
    "error_code": 0,
    "error_msg": "",
    "request_id": "1575625479.8081"
}
```





#### 查看某个具体的商品

#### 购买某个商品

#### 对购买的商品作出评价

