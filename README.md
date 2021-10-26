<h1 align="center">Easy Admin</h1>
<div align="center">
后台系统快速脚手架，包含登录注册，以及相应的认证
</div>

项目组成
----
- 前端：基于[Ant Design Pro](https://pro.ant.design/)实现的，修改了一下登录注册相关页面和逻辑
- 后端：使用gin框架写的一个简单的restful

功能说明
----
- 注册

这个比较简单，提供用户名、密码、邮箱三个资源，前端使用`sha1`加密，后端使用[bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)加密

- 登录

登录成功之后，会写一份cookie，后续请求都会带上cookie，具体的逻辑可以看这里[代码](easy-api-template/util/cookie.go)

- 前端拦截器

简单登录状态，未登录或者登录失效跳转到登录页面，具体逻辑在这里[代码](easy-admin-template/src/utils/request.js)

- 后端拦截器

简单写了一个中间件，具体逻辑在这里[代码](easy-api-template/middleware/auth.go)

总览
----
图片1
图片2

环境和依赖
----

> 前端项目(easy-admin-template)

可以直接看AntDesignPro的文档，[地址](https://github.com/vueComponent/ant-design-vue-pro/blob/master/README.zh-CN.md)
启动完成之后，接着启动后端服务

> 后端项目(easy-api-template)

- 安装依赖
```shell
go的安装依赖方式，不会的自己google
```

- 启动本地mysql
```shell
启动完成之后，修改`conf/config/dev.ini`中mysql的配置
生产环境请修改`conf/config/prod.ini`
```

- 本地运行
```shell
go run main.go --env dev 本地
```

- 生产运行
```shell
1.sh build.sh 构建出二进制文件
2.然后使用dockerfile构建镜像
3.剩下的自己玩(不会容器的就先入门下)
```


其他说明
----
- 本地环境，后端配置文件是`conf/config/dev.ini`，前端配置文件是`.env.development`
- 生产环境，后端配置文件是`conf/config/dev.ini`，前端配置文件是`.env`
- 前端所有配置的接口地址`easy-admin-template/src/api/apis`，这个目录下，按照文件名来区分是哪个模块就行


