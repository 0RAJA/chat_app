使用go语言实现的基于websocket的服务端
开发者本地测试步骤:

1. 确保go环境go1.17,docker版本20.10.17
2. `make init` 安装项目依赖的工具
3. `make docker_net` 创建docker网络
4. 根据自己的项目地址进行`postgres_zr_init`和`redis_init`初始化数据库
5. 通过`make migrate_up`迁移数据库表
6. 之后可以通过`make run`启动服务

请在`config/app`下新建`private.yml`配置文件,格式为:

```yaml
Postgresql: # Postgresql配置
  DriverName: postgresql # 驱动名
  SourceName: "postgres://root:123456@localhost:5432/chat?sslmode=disable&pool_max_conns=10"
Redis: # Redis配置
  Address: "localhost:6379"
  DB: 0
  Password: 123456
  PoolSize: 100 #连接池
  CacheTime: 1h
Token: # Token配置
  Key: "12345678123456781234567812345678" # 32位以上字符串，用于加密token
  UserTokenDuration: 24h # 用户token有效期
  AccountTokenDuration: 24h # 账号token有效期
  AuthorizationKey: Authorization # 授权头密钥
  AuthorizationType: bearer # 承载前缀
Email: # 邮件配置
  Host: smtp.qq.com
  Port: 465
  UserName: "***@qq.com" # 发送邮箱
  Password: "***" # 发送邮箱密钥
  IsSSL: true
  From: "***@qq.com" # 发送邮箱
  To: # 接收邮箱
    - "***@qq.com"
AliyunOSS: # OSS配置
  Endpoint: "https://oss-cn-hangzhou.aliyuncs.com"
  AccessKeyID: "***"
  AccessKeySecret: "***"
  BucketName: "***"
  BucketUrl: "***"
  BasePath: "chat/"
```
