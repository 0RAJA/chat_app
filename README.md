使用go语言实现的基于websocket的服务端

需求文档 https://ubqf66qdnq.feishu.cn/docx/doxcnef0UsiFJzmNTbBaPr85Kfc

错误码以及ws事件说明 https://ubqf66qdnq.feishu.cn/docx/doxcndbrG2BzzYCQXWxBCuP9bqg

请在`config/app`下新建`private.yml`配置文件,格式为:

```yaml
Postgresql: # Postgresql配置
  DriverName: postgresql # 驱动名
  SourceName: "postgres://root:123456@chat_postgres_zr:5432/chat?sslmode=disable&pool_max_conns=10"
Redis: # Redis配置
  Address: "chat_redis_62:6379"
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

开发者本地测试步骤:

1. 确保go环境go1.17,docker版本20.10.17
2. `make init` 安装项目依赖的工具
3. `make docker_net` 创建docker网络
4. 根据自己的项目地址进行`postgres_zr_init`和`redis_init`初始化数据库
5. `make build` 构建项目
6. `make docker_run` 启动项目

服务器docker-compose部署:

1. 修改`docker-compose.yml`中挂载卷中的绝对路径为本地项目路径
2. `docker-compose up -d` 后台启动项目
3. 可以修改`docker-compose.yml`配置文件来更改项目的部署配置(可以用自己打包的镜像)

功能(通过RESTFUL API接口和socket.io实现):

1. 一个用户(email唯一)拥有多个账号，并进行切换，同时支持多设备同时在线，消息进行同步推送
2. 账号之间可以进行好友申请，好友申请后需要等待对方同意，同意后成为好友
3. 好友之间可以发送普通消息或文本，消息可以被撤回，可以pin，可以置顶，可以查看已读未读，可以根据关键字搜索消息。
4. 群和好友在主页可以设置pin状态，是否显示，免打扰状态 ，群和好友可以设置备注，备注可以被搜索
5. 通过ws处理消息通信和已读操作，同时对各类操作或通知进行主动推送

优点:

1. 可以支持多设备同时在线，消息进行同步推送
2. 接口文档清晰，方便开发者调用
3. 代码结构规范，便于维护和拓展
4. 使用docker进行部署，支持多种环境部署

缺点:

1. 功能仍不完善，仍有改善空间
2. 数据库表设计仍存在不合理之处，需要后期进行调整，例如已读回执的存储需要进行调整
3. 服务需要进行拆分，例如IM服务可以单独部署，与主服务之间通过RPC进行通信

