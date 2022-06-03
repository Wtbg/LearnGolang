# 基于go echo框架登陆的简单实现及学习路径

## 引言

进入项目组之后才发现会的太少，方向也不完全和实用的对上，因此

<img src="file:///home/ld/.config/marktext/images/2022-06-02-18-21-33-image.png" title="" alt="" width="203">

学习一mentor&网络上收集的样例 & 相关知识的了解，路径如下

## package import and go mod

#### package

> - package由多个文件组成
> 
> - 最好package和目录名相同
> 
> - 每个子目录只能有一个package
> 
> - 以绝对路径GOPATH寻址

TRY：创建一个多package的项目，调用变量和函数

#### go mod

> 在工作区使用 `go mod init`启用go mod，前提条件是`go env`中`GO111MODULE`处于on/auto状态*可以使用export GO111MODULE修改，下载新版go默认开启*

go mod会在编译程序的时候自动添加依赖，更改工作区从GOPATH到go mod工作区，防止包含产生的冲突。

#### 类型学习

##### struct（和C类似，但好多不同

```go
type User struct{
    Username string
    SecretPassward string
}
var teacherZ User
```

但是似乎没有使用 -> 访问成员的用法，即使是指针，同样用 ‘ . ’ 访问

在返回值中返回结构体指针比返回结构体开销小

**返回一个结构体再改变会编译不过，因为并不确定返回的值是一个可被改变的左值**

在go中都是值传递

**小写的是不导出的成员，目前不能在包外访问/修改，**



如以下例子（来自[Go圣经节选]([结构体 · Go语言圣经](https://books.studygolang.com/gopl-zh/ch4/ch4-04.html))）

```go
func EmployeeByID(id int) *Employee { 
    //find the struct
    return &thestruct
}

fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"

id := dilbert.ID
EmployeeByID(id).Salary = 0 // fired for... no real reason
```

**在S类型结构体中不能有S类型，但是可以有S*类型**



**结构体简短赋值：使用成员名和相应的值，如：**

```go
auther := User{Username: "zhgg"}
```

##### 特性：面向对象匿名成员

```go
type Point struct {
    X, Y int
}

type Circle struct {
    Point
    Radius int
}

type Wheel struct {
    Circle
    Spokes int
}
```

> 匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针

可以直接访问叶子属性

```go
var w Wheel
w.X = 8            // equivalent to w.Circle.Point.X = 8
w.Y = 8            // equivalent to w.Circle.Point.Y = 8
w.Radius = 5       // equivalent to w.Circle.Radius = 5
w.Spokes = 20
```

匿名对象的名字可以视为类型名，但是这样不能用简短赋值了，总的来说是语法糖，但是能很好的结合面向对象

##### JSON

JavaScript对象表示法（JSON）是一种用于发送和接收结构化信息的标准协议。

JSON是对JavaScript中各种类型的值——字符串、数字、布尔值和对象——Unicode本文编码。它可以用有效可读的方式表示第三章的基础数据类型和本章的数组、slice、结构体和map等聚合数据类型。

```go
boolean         true
number          -273.15
string          "She said \"Hello, BF\""
array           ["gold", "silver", "bronze"] //JSON的数组
object          {"year": 1980,
                 "event": "archery",
                 "medals": ["gold", "silver", "bronze"]}//JSON的对象
```

**JSON marshaling**


