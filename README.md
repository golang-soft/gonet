# go-server
gonet 游戏服务器架构，mmo架构，包含数学库(box,matrix,point2d,point3d),[Recast Navigation寻路模块](https://blog.csdn.net/mango9126/article/details/79390543)，
a星寻路模块。

分布式雪花uuid,ai行为树，ai状态机，[excel导出配置](https://github.com/bobohume/gonet/tree/master/tool/data),raft同步模块，分片raft同步模块，hashring分布式一致性算法。

gonet核心思想是actor模式,消息驱动,采用无锁队列mailbox替换channel阻塞模型

微服务，微服务之间使用分布式消息队列
支持客户端调试
    Client 支持常规的socket服务器
    Clientwebsocket  支持golang版本的websocket服务器
    clientwebsocketG 支持gorilla 版本的websocket服务器

# 交流

# 分布式集群
##### center中央服务器，提供服务器注册，服务器id分配等等服务
##### config 配置文件
##### account账号服务，提供注册账号，登录校验，集群服务
##### netgate网关服务，对外连接，消息防火墙，对内消息转发，集群服务
##### world世界服务，所有逻辑，集群服务
##### worlddb db服务器，数据的层出与缓存，数据库的支持
##### login 登陆服，网关负载以及a，b版本切换
##### gm gm服务器
##### grpcserver grpc服务器
##### pay 支付服务器
##### proxy 代理服务器
##### zone 地图分区服务器
##### 第三方中间件：etcd 1分布式服发现，注册启动成功的服务， etcd 2 用于注册玩家，redis分布式缓存，nats分布式消息队列。
##### cmessage 客户端协议， 支持protobuf
##### smessage 服务器内部协议， 支持protobuf 
##### xconf 统一配置中心，暂时不支持，没有实现
# 服务器架构如下：
![image](框架.jpg)

