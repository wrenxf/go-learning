# Go 后端学习路线

**阶段 1：Go 基础**

- 吃透基础语法，重点掌握 slice、map、接口、错误处理；
- 精通并发：goroutine、channel、sync 同步原语；
- 理解核心原理：GMP 调度、GC 三色标记。
- 熟练 go mod，会简单 pprof 排查问题。

**阶段 2：业务 Web 开发**

- Gin 框架：中间件、参数解析，弄懂底层基数树路由；
- 存储：GORM + MySQL CRUD，能手写 SQL；Redis 基础使用；
- 基础网络：TCP、HTTP 基础，了解 net 原生网络库；
- 工具：Apifox 接口调试、数据库可视化工具；学习常用设计模式。

**阶段 3：微服务 & 分布式实战**

- gRPC、etcd 服务发现，上手 go-zero；落地一个微服务项目；
- 中间件：消息队列，解决分布式幂等、重试、缓存策略，提升 QPS；
- 可观测：Prometheus+Grafana、链路追踪；
- 运维部署：Linux 基础、Docker，搭建简易 CI/CD 流水线。

# 学习进度

| 功能点         | 状态       |
| -------------- | -------------- |
| Go基础语法     | 完成     |
| MySQL基础 | 完成 |
| Gin 框架 | 完成 |
| GORM | 完成 |
| Redis | 进行中 |

# 学习笔记

[1.Go语言基础](./GoBasics/Go_study.md)

[2.MySQL](./MySQL.md)

[3.Gin框架](./gin_study/Gin_study.md)

[4.GORM](./GORM.md)

[5.数据序列化格式](./YAML.md)

[6.GMP](./GMP.md)
