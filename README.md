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

## 接口设计

### SuperAdmin部分

#### 获得所有商户信息

- Method : HTTP GET
- Request URL: /somewhere/stores
- Response:

```json
`{``  ``"token"``:``"xxxx"``,``  ``"error_code"``:``0``,``  ``"error_msg"``:``"xx"``,``  ``"request_id"``:``"xxx"``}`
```



#### 增加某个商户

- Method : HTTP POST
- Request URL: /somewhere/stores
- Request: raw json

```json
`{``  ``"error_code"``:``0``,``  ``"error_msg"``:``"xx"``,``  ``"request_id"``:``"xxx"``}`
```

- Response:

```json
`{``  ``"error_code"``:``0``,``  ``"error_msg"``:``"xx"``,``  ``"request_id"``:``"xxx"``}`
```

