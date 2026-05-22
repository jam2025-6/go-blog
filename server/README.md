### server

├── api (api层)
├── assets (静态资源包)
├── config (配置包)
├── core (核心文件)
├── flag (flag命令)
├── global (全局对象)
├── initialize (初始化)
├── log (日志文件)
├── middleware (中间件层)
├── model (模型层)
│ ├── appTypes (自定义类型)
│ ├── database (mysql结构体)
│ ├── elasticsearch (es结构体)
│ ├── other (其他结构体)
│ ├── request (入参结构体)
│ └── response (出参结构体)
├── router (路由层)
├── service (service层)
├── task (定时任务包)
├── uploads (文件上传目录)
└── utils (工具包)
├── hotSearch (热搜接口封装)
└── upload (oss接口封装)

api 层负责接收请求和返回响应（"接客"）
service 层负责实现具体的业务逻辑（"干活"）。
