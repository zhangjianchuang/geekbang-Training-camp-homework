# eshop

极客时间GO进阶训练营毕业项目。简单的在线商城，使用微服务架构。

- cart： 购物车
- catalog： 商品目录
- catalog.admin：商品目录管理
- order： 订单
- shop： 项目BFF，提供对用于，商品目录，用户，购物车，订单的管理
- user： 用户
# 总结

1. 微服务架构
- 微服务网关
- BFF
- 微服务划分
- 服务发现
- 多集群，多租户
2. Golang错误处理
- error接口
- error用法：
    - 1. Sentinel Error
    - 2. ErrorType
    - 3. Opaque Error
- Wrap Error
3. 并发
- Goroutines
- sync包
- Channel
- context
4. 工程化
- 目录
    - cmd
    - api
    - internal
    - pkg
    - configs
- DI
- API设计
    - API命令
    - 错误码定义
- 配置项
5. 微服务治理
- 服务隔离
- 超时
- 过载保护
- 限流
- 降级
- 重试
- 负载均衡
6. 分布式缓存
- 缓存
    - redis
    - memcache
- 数据一致性
7. 分布式事务
8. Kafka
- Topic
- Partition
- Broker
- Producer
- Consumer
