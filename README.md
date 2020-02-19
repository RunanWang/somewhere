## somewhere

这是一个毕业设计的项目，是一个简易的推荐系统。功能包括推荐展示、用户行为收集以及推荐模型的训练。

## 技术选型

### 算法

- keras 学习框架

### 后端

- go gin 框架
- MongoDB
- Redis

### 前端

- vue框架

## 启动

- 后端

    go run main.go

- 前端

    npm install --registry=https://registry.npm.taobao.org

    npm run dev

    浏览器访问 [http://localhost:8036](http://localhost:8036)

- 示例

    示例 [搭建在阿里云上](http://182.92.196.182)
    

## 文件组织

- 文档部分

        在./doc

- 算法部分

       目前只有demo，处理的是movielens数据集，在./algo中:
       ./algo/data 数据处理
       ./algo/src 几种常见的模型实验

- 后端部分
        
        ./main.go 启动文件
        ./app gin框架启动
        ./conf config.toml文件
        ./config 配置源码
        ./db 数据库库相关代码
        ./err 错误处理相关
        ./handler 处理函数相关
        ./middleware 跨域、鉴权中间件
        ./model 数据库模型
        ./msg 前后端传递的消息
        ./service 后端提供的服务
        ./vendor 外部调用的包

- 前端部分

        ./vue-admin

- docker

## 效果展示
