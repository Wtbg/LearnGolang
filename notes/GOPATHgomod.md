### GOPATH

#### 工作区

工作区包含版本控制仓库

版本控制工具仓库包含很多包

包包含 *.go文件*

`src`内有Go文件

`bin`内有二进制文件

---

#### go mod

> Modules 是相关 Go 包的集合，是源代码交换和版本控制的单元。Go语言命令直接支持使用 Modules，包括记录和解析对其他模块的依赖性，Modules 替换旧的基于 GOPATH 的方法，来指定使用哪些源文件。

常用的`go mod`命令如下表所示：  

| 命令              | 作用                              |
| --------------- | ------------------------------- |
| go mod download | 下载依赖包到本地（默认为 GOPATH/pkg/mod 目录） |
| go mod edit     | 编辑 go.mod 文件                    |
| go mod graph    | 打印模块依赖图                         |
| go mod init     | 初始化当前文件夹，创建 go.mod 文件           |
| go mod tidy     | 增加缺少的包，删除无用的包                   |
| go mod vendor   | 将依赖复制到 vendor 目录下               |
| go mod verify   | 校验依赖                            |
| go mod why      | 解释为什么需要依赖                       |

> `go get` 
> 
> - 运行`go get -u`命令会将项目中的包升级到最新的次要版本或者修订版本；
> 
> - 运行`go get -u=patch`命令会将项目中的包升级到最新的修订版本；
> 
> - 运行`go get [包名]@[版本号]`命令会下载对应包的指定版本或者将对应包升级到指定的版本。

> 提示：`go get [包名]@[版本号]`命令中版本号可以是 x.y.z 的形式，例如 go get foo@v1.2.3，也可以是 git 上的分支或 tag，例如 go get foo@master，还可以是 git 提交时的哈希值，例如 go get foo@e3702bed2。

☝复制的，用的时候在写其他的

> 一旦go.mod和创建，go toolchain会在执行命令时修改和维护go.mod文件
> 
> 命令下详

##### 添加依赖
