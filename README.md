使用go语言实现的基于websocket的服务端
开发者本地测试步骤:

1. 确保go环境go1.17,docker版本20.10.17
2. `make init` 安装项目依赖的工具
3. `make docker_net` 创建docker网络
4. 根据自己的项目地址进行`postgres_zr_init`和`redis_init`初始化数据库
5. 通过`make migrate_up`迁移数据库表
6. 之后可以通过`make run`启动服务
