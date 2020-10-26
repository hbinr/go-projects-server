# go-projects-server

gin-dev-kit 分支主要做了gin脚手架开发


## 项目目录解释

### main.go
程序入口文件
### /internal
这是 Go 包的一个特性，放在该包中的代码，表明只希望项目内部使用，是项目或库私有的，其他项目或库不能使用。请注意，不限于顶层internal目录，internal在项目树的任何级别上都可以有多个目录。
可以选择向内部包中添加一些额外的结构，以分隔共享和非共享内部代码。它不是必需的（尤其是对于较小的项目），但是最好有视觉提示来显示包的用途。实际应用程序代码可以进入/internal/app目录（例如/internal/app/myapp），而这些应用程序共享的代码可以进入/internal/pkg目录（例如/internal/pkg/myprivlib）。

### /pkg
该包可以和 internal 对应，是公开的。一般来说，放在该包的代码应该和具体业务无关，方便本项目和其他项目重用。当你决定将代码放入该包时，你应该对其负责，因为别人很可能使用它。

如果应用程序项目很小，并且嵌套的额外层次不会增加太多价值（除非您真的想要，请不要使用它。当它变得足够大并且您的根目录变得非常复杂时（特别是如果您有很多非Go应用程序组件），请考虑一下。

### /api

1. 负责获取和反解析前端发来的数据
2. 做基本的逻辑判断,传入的数据是否正确和合法
3. 将获取的数据传给服务层
4. 把结果返回给前端
5. 可能还开始和结束还包含日志记录

这一层将作为表现者。决定数据如何呈现。任何传递类型都可以作为是 REST API， 或者是 HTML 文件，或者是 gRPC

这一层将接收来自用户的输入， 并清理数据然后传递给用例层。

对于我的示例项目， 我使用 REST API 作为表现方式。客户端将通过网络调用资源节点， 表现层将获取到输入或请求，然后将它传递给用例层。

该层依赖于service层。

### /router
供外部访问,路由最好是分组设计,这样有比较好的逻辑边界,并且同一组的路由在一起,看代码可以一目了然。一般理由的分组按照模型分组即可。


### /service
这层将会扮演业务流程处理器的角色。任何流程都将在这里处理。该层将决定哪个仓库层被使用。并且负责提供数据给服务以便交付。处理数据进行计算或者在这里完成任何事。

用例层将接收来自传递层的所有经过处理的输入，然后将处理的输入存储到数据库中， 或者从数据库中获取数据等。

用例层将依赖于仓库层。

### /dao
数据访问层 是服务层获取数据的接口包

存放所有的数据库处理器，查询，创建或插入数据库的处理器将存放在这一层，该层仅对数据库执行 CRUD 操作。 该层没有业务流程。只有操作数据库的普通函数。

这层也负责选择应用中将要使用什么样的数据库。 可以是 Mysql， MongoDB， MariaDB，Postgresql，无论使用哪种数据库，都要在这层决定。

如果使用 ORM， 这层将控制输入，并与 ORM 服务对接。

如果调用微服务， 也将在这层进行处理。创建 HTTP 请求去请求其他服务并清理数据，这层必须完全充当仓库。 处理所有的数据输入，输出，并且没有特定的逻辑交互。

该层（dao）将依赖于连接数据库 或其他微服务（如果存在的话）

### /model
数据仓库层,把数据库和redis和其他存储都放在这个包下

与实体（ Entities ）一样， 模型会在每一层中使用，在这一层中将存储对象的结构和它的方法。例如： Article， Student， Book。
```go
import "time"

type Article struct {
    ID        int64     `json:"id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    UpdatedAt time.Time `json:"updated_at"`
    CreatedAt time.Time `json:"created_at"`
}
```
所以**实体**或者**模型**将会被存放在这一层

### middleware
中间件包，比如jwt，cors等

## web application 目录 (本项目为前后端开发，无此层)
### /web
Web应用程序特定的组件：静态Web资产，服务器端模板和SPA。

## common application目录
### /config
配置文件模板或默认配置。
### /util
工具包,主要是一些小工具的包，比如MD5，time格式化，json序列化等


### /script
存放 build、install、analysis 等操作脚本。这些脚本使得项目根目录的 Makefile 很简洁。

### /build
该目录用于存放打包和持续集成相关脚本。将云（AMI），容器（Docker），操作系统（deb，rpm，pkg）软件包配置和脚本放在/build/package目录中。
将CI（travis，circle，drone）配置和脚本放在/build/ci目录中。请注意，某些配置项工具（例如Travis CI）对于其配置文件的位置非常挑剔。尝试将配置文件放在/build/ci目录中，将它们链接到CI工具期望它们的位置（如果可能）。


### /test
一般用来存放除单元测试、基准测试之外的测试，比如集成测试、测试数据等。

## 其他目录
## /doc
设计和用户文档（除了godoc生成的文档之外）。


## 不应该拥有的目录
/src


参考：

https://studygolang.com/articles/26941?fr=sidebar

https://www.jianshu.com/p/fced6e751ada

https://studygolang.com/articles/12909

https://www.jianshu.com/p/d2ccf971ded2