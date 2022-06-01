# go基础

### go命名的作用域

- 函数内：函数内

- 函数外：当前包内+大写字母开头包外（非私有
  
  *go命名风格简短*

### go声明定义

#### 变量

**var 变量名 类型 = 表达式**

> 省略类型字面类型
> 
> 省略表达式自动默认值  

**go个性承接函数返回值：**

```go
var f, err = os.Open(name)
if err != nil {
    return err
}
//use f
f.Close()
```

**简短声明**（只能函数内）

> 变量名 := 表达式

*简短声明能赋值存在的变量，但不能但不能单纯赋值，必须包含新声明的变量*

*但是如果是外部词法域中声明的，不等于赋值而是重新声明变量*

#### 指针

如果

```go
var x int
```

&x类型为*int //和C相似

> go语言中返回局部变量的地址是安全的，但每次地址不同，变量“逃逸“
> 
> 在go中*p++只改变变量的值，与C不同，类似引用调用
> 
> slice/map/chan也会创建别名

#### new函数

返回类型指针

```go
p := new(int)
*p = 2
fmt.Println(*p) // 2

func new(type)*type{
    var i type
    return &i
}
```

**new不是关键字可以作为变量名/别的类型**

#### go垃圾回收

> 避开完整的技术细节，基本的实现思路是，从每个包级的变量和每个当前运行函数的每一个局部变量开始，通过指针或引用的访问路径遍历，是否可以找到该变量。如果不存在这样的访问路径，那么说明该变量是不可达的，也就是说它是否存在并不会影响程序后续的计算结果。

#### 赋值

> 在go语言中 v++是语句
> 
> 因此不能使用
> 
> ```go
> x = v++
> ```

> 可元组赋值，但数量必须一样，不会取前几个

#### 可赋值性

对于非常量，类型必须完全相同

#### 新声明类型

type 类型名 底层类型 //一般在包一级

```go
package main

type i1 int
type i2 int

func main() {
	var var1 i1 = 1
	var var2 i2 = 2
	var2 = var1
	//cannot use var1 (variable of type i1) 
    //as i2 value in assignmentcompilerIncompatibleAssign
    //不同类型之间不能不能操作
}
```

#### 声明类型的方法

```go
//声明typename类型的方法String()
func (c typename) String() string {
    return fmt.Sprintf(~~~)
}
```

```go
package main

import (
	"fmt"
)

type i1 int
type i2 int

func (i i1) String() string {
	return fmt.Sprintf("the value is %d", i)
}
func main() {
	var i i1 = 1
	fmt.Println(i)
	fmt.Printf("%d", i)
}
//the value is 1
//1
```

#### 导入包

一个包由一个或多个.go文件组合，利用目标路径导入包

> GOPATH是 Go语言中使用的一个环境变量，它使用绝对路径提供项目的工作目录。

```shell
go env
GO111MODULE="on"
GOARCH="amd64"
GOBIN=""
GOCACHE="/home/ld/.cache/go-build"
GOENV="/home/ld/.config/go/env"
GOEXE=""
GOEXPERIMENT=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOINSECURE=""
GOMODCACHE="/home/ld/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/home/ld/go"
GOPRIVATE=""
GOPROXY="https://goproxy.cn,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GOVCS=""
GOVERSION="go1.18.1"
GCCGO="gccgo"
GOAMD64="v1"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD="/dev/null"
GOWORK=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build1543317169=/tmp/go-build -gno-record-gcc-switches"
```

> - 第 1 行，执行 go env 指令，将输出当前 Go 开发包的环境变量状态。
> - 第 2 行，GOARCH 表示目标处理器架构。
> - 第 3 行，GOBIN 表示编译器和链接器的安装位置。
> - 第 7 行，GOOS 表示目标操作系统。
> - 第 8 行，GOPATH 表示当前工作目录。
> - 第 10 行，GOROOT 表示 Go 开发包的安装目录。

- *在GOPATH指定路径下代码会保存在$GOPATH/src*

- *二进制文件在$GOPATH/bin*

> 建立工作目录GOPATH：
> 
> ```shell
> export GOPATH='pwd'
> ```

> 建立源码目录：
> 
> ```shell
> mkdir -p src/proname  
> ```

[GOPATH网站]([Go语言GOPATH详解（Go语言工作目录）](http://c.biancheng.net/view/88.html))




