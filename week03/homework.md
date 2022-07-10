# 运行结果

![](../img/2022-07-10-18-48-24-image.png)

# 设计

1. 获取cancel上下文
2. 获取group
3. 使用group启动StartServer goroutine和销毁的goroutine
4. 等待所有goroutine完成，有异常将异常打印


