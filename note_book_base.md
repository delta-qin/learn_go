# 看文档
`go doc builtin.delete`

[studygolang.com/pkgdoc](https://studygolang.com/pkgdoc)

# 环境搭建

安装之后配置GOPATH

查看`go env`


# fmt
fmt 实现了类似C语言的 printf 和 scanf的格式化 I/O。

## 向外输出
- Printf，格式化输出。想要输出变量值需要占位符，
- Sprintf("%s %s", name, world), 不打印，直接返回给新变量
- Println，想输出变量直接逗号隔开
- Print，想输出变量直接逗号隔开

```go
// Print 简单打印
fmt.Print("hello")
// Println 有换行符的打印
fmt.Println()

// Printf 是唯一一个可以使用 占位符 的输出
fmt.Printf("hello, %s", "deltaqin")


fmt.Printf("%d\n", n)
fmt.Printf("%b\n", n)
fmt.Printf("%o\n", n)
fmt.Printf("%x\n", n) // 16进制，使用a-f
fmt.Printf("%X\n", n) // 16进制，使用A-F
fmt.Printf("%U\n", n) // unicode U+1234 Unicode码
fmt.Printf("%q\n", n) // go语法字符字面值 65-》 A， 整数-》字符

fmt.Printf("%b\n", n) // 二进制指数的科学计数法 -123456p-78
fmt.Printf("%f\n", n) // 有小数部分，无指数部分 123.456
fmt.Printf("%9f\n", n) // 宽度9，精度默认，前面空格
fmt.Printf("%.2f\n", n) // 精度2
fmt.Printf("%9.2f\n", n) // 宽度9，精度2，前面空格
fmt.Printf("%9.f\n", n) // 宽度9，精度0
fmt.Printf("%F\n", n) // 同上
fmt.Printf("%e\n", n) // 科学计数法 -1234.456e+78
fmt.Printf("%E\n", n) // 科学计数法 -1234.456E+78
fmt.Printf("%g\n", n) // 实际而定  %e %f
fmt.Printf("%G\n", n) // 实际而定  %E %F

fmt.Printf("%c\n", n)
fmt.Printf("%s\n", n)
fmt.Printf("%5s\n", n) // 5个字符，不足前面补空格
fmt.Printf("%-5s\n", n) // 5个字符，不足后面补空格
fmt.Printf("%05s\n", n) // 5个字符，不足前面补0
fmt.Printf("%q\n", n) // 安全转义
fmt.Printf("%x\n", n) // 每个字节使用两字符16进制来表示
fmt.Printf("%X\n", n)

fmt.Printf("%p\n", n) // 指针

fmt.Printf("%v\n", n) // 万能，打印默认输出
fmt.Printf("%+v\n", n) // 万能，输出结构体的时候回加字段名
fmt.Printf("#%v\n", n) // 万能，如果是string按照类型加一个“”，打印更加详细，比 %v

fmt.Printf("%T\n", n)
fmt.Printf("%%\n", n) // 这个打印的是 % ,格式不是 \%
fmt.Printf("%t\n", n) // true 或 false
```

## 向内输入
fmt 下面有 

- fmt.Scan  
- fmt.Scanln 自带换行符
- fmt.Scanf 格式化，手动换行符

```go
var (
	name string
	age  int
)

// 没有引号
fmt.Scan(&name, &age)
fmt.Println(name, age)
fmt.Scanln(&name, &age)
fmt.Println(name, age)
fmt.Scanf("%s %d\n", &name, &age)
fmt.Println(name, age)
```

# 变量
- 全局变量
  - 可以定义不使用。
  - 必须以关键字开头
  - - var name string
- 局部变量
  - 定义必须使用，否则报错
  - 可以省略关键字，直接使用 :=
  - var name string
    - name = "de"
  - var name = "de"
  - name := "de"
- 匿名变量（哑元变量）
  - 必须接收但是不使用的时候

注意文件的命名，不可以直接使用_test结尾的命名，会报错

推荐使用驼峰

声明变量可以批量

全局变量

```go
// var name string
// var age int
// var flag bool

var (
	name string
	age  int
	flag bool
)
```

局部变量 

`var name string`

`var name = "deltaqin"`

函数内简短声明 `:=`

匿名变量 `_` 

```go 
var v1 int32
v1 = 1

var v2 int32 = 1

var v3 = int32(1)

// 强制类型转换
v4 := int32(1)
```

# 常量

## const
批量声明，某一行之后没有赋值，默认就和上一行一致

```go
// const name = "qzt"
// const age = 12
// const flag = false

const (
	name = "qzt"
	age  = 12
	flag = false
)

const (
	n1 = "qzt"
    n2
    n3  
)
```

## iota

go语言常量计数器，在常量表达式里面使用

iota **在 const 关键字出现的时候将被重置为0**。
**const 每新增一行常量声明将使 iota 计数一次**，

是 const 语句块里面的行索引。简化定义在枚举的时候比较有用

```go
const(
    n1 = iota // 0
    n2        // 1
    n3        // 2
)

// 换行就会+1
const(
    m1 = iota // 0
    m2          //1
    _
    m3          //3
)

// 一个体内只会初始化一次0，再次定义也是叠加
const (
	c1 = iota // 0
	c2 = 100
	c3 = iota // 2 新增一行就会加1
	c4        // 3
)

// 一行内定义多个
// 只有换行才会+1
const (
	a1, a2 = iota + 1, iota + 2 // 1,2
	b1, b2 = iota + 1, iota + 2 // 2,3
)

// 定义数量级
const (
	_ = iota
	// 1左移10位相当于乘了2^10，1000000000=1024
	KB = 1 << (10 * iota)
	// 左移20位,2^20 = 2^10*2^10 = 1024*1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)
```

# 基本数据类型

## 整型
- 有符号
  - int8：-128 127
  - int16：-32768 32767
  - int32： -2147483648 2147483647
  - int64：-9223372036854775808 9223372036854775807
- 无符号
  - uint8：byte 0 255
  - uint16：short 0 65535
  - uint32 0 4294967295
  - uint64：long 0 18446744073709551615

## 特殊整型
- uint：32系统就是uint32,64位就是uint64
- int：上面一致
- uintptr：无符号整型，用于存放一个指针，变量内存地址，16进制表示，0x...

## 八进制 十六进制

```go
var a1 = 100
fmt.Printf("%d\n", a1) //100
fmt.Printf("%b\n", a1) //1100100
fmt.Printf("%o\n", a1) //144
fmt.Printf("%x\n", a1) // 64

a2 := 017
fmt.Printf("%d\n", a2) // 15

a3 := 0x123
fmt.Printf("%d\n", a3) // 291

a4 := int8(1)
fmt.Printf("%d\n", a4) // 1

// 查看变量类型
fmt.Printf("%T\n", a1)
fmt.Printf("%T\n", a4) //int8
```

## 浮点数

```go
// math.MaxFloat32
f1 := 1.2333 // 默认64，32需要显式说明
fmt.Printf("%T\n", f1) //float64
f2 := float32(1.223)
fmt.Printf("%T\n", f2) //float32 

// 不能直接赋值，即使是小转大
```

## 复数
- complex64
  - 实部和虚部是32位
- complex128
  - 实部和虚部是64位

```go
	var c1 complex64
	c1 = 1 + 2i

	fmt.Println(c1)

	var c2 complex64
	c2 = 2 + 2i
	fmt.Println(c2)
```

## 布尔值bool
默认是false
整型不能强bool
不能参与数值运算

```go
bool1 := false
fmt.Printf("%T\n", bool1) // bool

var bool2 bool
fmt.Printf("%T : %v", bool2, bool2) //bool : false
```




## 字符串string
go的字符串是以原生数据类型出现的，和其他的数据一样。使用UTF-8编码。

**只可以时候用"",使用''是字符，和java是一样的**

### 转义符

- `\r`回车
- `\n`换行
- `\t` 制表
```go
	s1 := "\"你\\好'啊\""
	fmt.Println(s1)
```

### 多行字符串
```go
	s2 := `hello 
		啊这
	`
	fmt.Println(s2)
```

### 相关操作
- len(str)
- + 或 fmt.Sprintf
- strings.Split
- strings.Contains
- strings.HasPrefix, strings.HasSuffix
- strings.Index(), strings.LastIndex()
- strings.Join(a[] string, sep string)


```go
	// 相关操作
	s3 := "/Users/qinzetao/workspace/go"
	s4 := strings.Split(s3, "/")
	fmt.Println(s4)
	fmt.Println(strings.Join(s4, "?"))

	fmt.Println("len", len(s3))

	fmt.Println(strings.Contains(s3, "Users"))

	fmt.Println(strings.HasPrefix(s3, "/U"))
	fmt.Println(strings.HasSuffix(s3, "go"))
	fmt.Println(strings.Index(s3, "e"))
	fmt.Println(strings.LastIndex(s3, "go"))
```

## 字符 byte和rune
字符串与字符，打印类型可以看到，string类型和int32类型

- byte：uint8，代表一个 ASCII 码字符，使用''括起来
- rune：代表一个 UTF-8 字符，实际是int32，3个字节来表示一个字符，3*8=24.前面还有几位表示rune类型

`'H'`默认是int32，只有使用`byte('H')`显式转换之后才是uint8

**字符串底层是一个byte数组，可以直接和`[]byte`转换，也可以直接和`[]rune`，区别在于转换之后一个不可以直接复制ASCII之外的，一个可以**

字符串不可修改,想修改必须重新分配内存，复制字节数组。

一个rune字符由多个byte组成

```go
s := "deltaqin"

fmt.Println(len(s))

// byte
for i := 0; i < len(s); i++ {
	fmt.Printf("%v(%c)", s[i], s[i])
}
fmt.Println()

// rune
for _, r := range s {
	fmt.Printf("%v(%c)", r, r)
}

// 100(d)101(e)108(l)116(t)97(a)113(q)105(i)110(n)228(ä)189(½)160( )229(å)165(¥)189(½)
// 100(d)101(e)108(l)116(t)97(a)113(q)105(i)110(n)20320(你)22909(好)

```

```go
// 转换案例
// 转换为byte切片
byteS1 := []byte(s)
byteS1[0]='a'
fmt.Println(string(byteS1))

// 转换为 rune 切片
byteS2 := []rune(s)
// 注意字符不可以使用""
byteS2[0] = '秦'
fmt.Println(string(byteS2))
```

## 类型转换
```go
var a, b = 3,4
var c int

c = int(math.Sqrt(float64(a*a + b*b)))
fmt.Println(c)
```

# 流程控制

## if
条件没有括号

```go
age := 20
if age > 18 {
	fmt.Println(1)
} else if age < 10 {
	fmt.Println(0)
} else {
	fmt.Println(2)
}
```

使用局部变量
```go
// 特殊写法
// 局部变量，作用域，用完即销毁，后面获取不到
if age1:=20; age1>19 {
	fmt.Println(99)
} else {
	fmt.Println(22)
}
```


## for
break，continue不对if有效

```go
for i := 0; i < 10; i++ {
	fmt.Println(i)
}

// 1
i:= 5
for ; i < 10; i++ {
	fmt.Println(i)
}

// 2
for i<10 {
	fmt.Println(i)
	i++
}

// 死循环
for {
	fmt.Println("1")
}

// range
s1 := "deltaqin你好"
for i,v := range s1{
	// 直接打印是ASCII
	// fmt.Println(i,v)

	fmt.Printf("%d,%c\n" ,i,v)
}

// 0,d
// 1,e
// 2,l
// 3,t
// 4,a
// 5,q
// 6,i
// 7,n
// 8,你
// 11,好
// 注意最后的一个汉字是三个字符
```

## switch

```go
i := 3

switch i {
case 1:
	fmt.Println(1)
case 2:
	fmt.Println(2)
case 3:
	fmt.Println(3)
default:
	fmt.Println("over")
}
```

```go
// 局部变量，多枚举case
switch i:=3; i {
case 1:
	fmt.Println(1)
case 2,5,6:
	fmt.Println(2)
case 3:
	fmt.Println(3)
default:
	fmt.Println("over")
}
```

```go
// 条件判断作为case，switch后不需要写任何变量
age := 13
switch {
	case age > 20:
		// 打印
	case age < 18:
		// 打印
	default:
		// 打印
}
```

```go
// 
age := 13
switch {
	case age > 20:
		// 打印
		// fallthrogh会执行满足条件的当前case和下个case
		// 不建议
		fallthrough
	case age < 18:
		// 打印
	default:
		// 打印
}
```

## goto--不建议

```go
for i:=1; i<10 ; i++ {
	if i == 2 {
		goto breakTag
	}
}

breakTag:
	//语句
```

**跳出多层for循环**
```go
for i:=1; i<10; i++ {
	flag := true
	for j:=1; j<i; j++ {
		if i == 5 {
			flag = false
			break
		}
	}
	if !flag{
		break
	}
}
```

# 运算符
## 算数运算符
`+-*/%`

`a++` 是一个单独的语句，不能在`=`右边赋值


## 关系运算符
`== >= <= != `
强类型，不同类型不可以比较，`int32` 和 `int8` 不可以比较

## 逻辑运算符
`&& || !`


## 位运算
对整数在内存里面的二进制位进行运算
`& | ^` 与或异或

`<< >>`左移n位就是乘以2的n次方，右移n位就是除以2的n次方。左移高位丢弃，低位补0，右移直接移动即可

```go
// 位运算的时候移动之后的结果不要超出数据类型的容量
var i = int8(1)
i<<10
// 打印出来就是0，只取后8位，
// 10000000000
```

### 二进制的作用
- 加速运算
- IP
- 权限：Linux（自己，用户组，其他用户）或者自己设计
  
## 赋值运算符
操作之后再赋值

`x = x + 1`

`= += -= *= /= 5= <<= >>= &= |= ^=`


# 复杂数据类型

## 数组
声明的时候就确定，内容可以修改，大小不能变

`var name [3]int`

相当于一个容器，长度是数组类型的一部分，一旦定义就不能改变了。`[5]int`和 `[10]int` 是不同的类型

不同类型不可以赋值

### 初始化
不初始化默认都是 `0`,`false`,`""`

```go
// 初始化1
	var arr1 [5]int
	arr1 = [5]int{1, 0, 2, 2, 3}
	fmt.Println(arr1)

	// 初始化2
	arr2 := [5]int{1, 1, 1, 1, 1}
	fmt.Println(arr2)

	// 初始化3：不指定个数初始化
	arr3 := [...]int{2, 2, 2, 2, 2}
	fmt.Println(arr3)

	// 初始化4：按照位置初始化
	arr4 := [5]int{0:1, 1:2, 4:6}
	fmt.Println(arr4)

	// 初始化5：按照位置初始化
	arr5 := [...]int{0:1, 1:2, 90:6}
	fmt.Println(arr5)
```

### 数组遍历

```go
// 遍历1
arr5 := [...]string{"1","2","3"}
for i:=0; i<len(arr5);i++ {
	fmt.Println(arr5[i])
}

// 遍历2
for i, v :=  range arr5{
	fmt.Println(i,v)
}
```

### 多维数组
一定要赋值的时候也要指定值的类型，坑

```go
// 二维数组
var arr6 [3][2]int
arr6 = [3][2]int{
	[2]int{1,2},
	[2]int{3,2},
	[2]int{5,2},
}

for _, val1 := range arr6{
	// fmt.Println(val1)
	for _,val2 := range val1 {
		fmt.Println(val2)
	}
}
```

### 数组是值类型，赋值之后是复制，不是引用
**不是引用类型，赋值之后是复制，不是引用**

```go
// 数组是值类型，不同于引用类型
arr7 := [2]int{1, 2}
arr8 := arr7
arr8[0] = 100
fmt.Println(arr7)
// [1 2]
// [100 2]
fmt.Println(arr8)
```

## 切片
数组看做是容器，数组长度固定，类型固定。切片是引用类型，切片圈住连续的内存，对应的是其上面的数组，**看做切片是在底层数组上面的滑动窗口。**

**切片底层是一个数组**，直接创建的话go底层先创建一个数组，再将这个数组切片之后返回。和自己使用数组切的原理是一样的

### 直接定义切片

```go
// 1. 定义切片：不指定个数
var s1 []int
var s2 []string
fmt.Println(s1, s2)
fmt.Println(s1 == nil) // true 没有开辟内存空间
fmt.Println(s2 == nil)

// 初始化切片
s1 = []int{1, 2, 3}
s2 = []string{"1", "2", "3"}
fmt.Println(s1, s2)
fmt.Println(s1 == nil) // false
fmt.Println(s2 == nil)

// 长度和容量
fmt.Printf("len(s1):%d, cap(s1):%d\n", len(s1), cap(s1))
fmt.Printf("len(s2):%d, cap(s2):%d\n", len(s2), cap(s2))
```

### 从数组获取切片
```go
// 2. 数组得到切片
a1 := [...]int{1, 2, 3, 4, 5, 6}
s3 := a1[0:4] // [,)
fmt.Println(s3)
s4 := a1[:4]
fmt.Println(s4)
s5 := a1[4:]
fmt.Println(s5)
s6 := a1[:]
fmt.Println(s6)
// 切片容量指的是底层数组的容量
// 数组容量是从起始切的位置一直到最后。
fmt.Printf("len(s3):%d, cap(s3):%d\n", len(s3), cap(s3))
// len(s3):4, cap(s3):6
fmt.Printf("len(s5):%d, cap(s5):%d\n", len(s5), cap(s5))
// len(s5):2, cap(s5):2
```

### 理解“切片是在数组上面的滑动窗口”
切片是在数组上面的滑动窗口，引用类型
对切片再切片也只是改变窗口的大小

修改底层数组的值，切片的值自然就变了，切片没有自己的值，只是一个窗口，圈出来数组上的某一个区域

```go
// 切片是在数组上面的滑动窗口，引用类型
// 对切片再切片也只是改变窗口的大小
s7 := s5[1:]
a1[5] = 100
fmt.Println("s7",s7,"s5",s5) // s7 [100] s5 [5 100]
```

### `make([]T,size,cap)`动态创建切片
T是切片元素类型，后者作用显然。注意写类型的时候有[]
```go
// 3. make函数创建切片
s8 := make([]int, 5, 10)
fmt.Println(s8) // [0 0 0 0 0]
fmt.Println(len(s8)) //5
fmt.Println(cap(s8)) //10
// 显然，内部存储空间已经分配了10个
```

**注意初始make里面定义的 size 都会有默认值，后续使用append追加的时候是在这些之后追加的，不是从0位置开始**

### 切片赋值拷贝
拷贝之后，两个切片会共享一个底层数组，改变切片内容（不是指大小）就是在改变底层数组元素。两者都会受到影响

```go
s9 := []int{1, 2, 3}
s10 := s9
fmt.Println(s9, s10)
// [1 2 3] [1 2 3]
s9[0] = 100
fmt.Println(s9, s10)
// [100 2 3] [100 2 3]
```

### 切片遍历
和数组一样

### 注意事项
#### 切片之间不可以比较
- 不能用 == 里比较两个切片是否含有相同的元素，唯一合法的比较操作是和 nil 比较
- nil值的切片没有底层数组，没有开辟内存空间，切片长度和容量都是0。
- 长度和容量是0的切片不一定 == nil，
- 可以直接使用 `len(s1)` 来判断是否为空


```go
var s1 []int	// s1 == nil
s2 := []int{}	// s2 != nil
s3 := make([]int, 0, 0)	// s3 != nil
```

### 切片常用方法

#### append：为切片动态添加元素。
- 底层数组不能容纳的时候，会自动按照一定策略扩容，此时切片指向的底层数组就会更换，扩容往往发生在 append 函数调用时
- **调用append必须使用之前的切片来接收返回值**


```go
// append方法为切片添加元素
var s11 []int
for i := 0; i < 10; i++ {
	// 调用append必须使用之前的切片来接收返回值
	s11 = append(s11, i)
	fmt.Print("value：", s11, " len:", len(s11), " cap:", cap(s11))
	fmt.Printf(" ptr:%p\n", s11)
}
```

结果：
```
value：[0] len:1 cap:1 ptr:0xc00001c200
value：[0 1] len:2 cap:2 ptr:0xc00001c210
value：[0 1 2] len:3 cap:4 ptr:0xc000018220
value：[0 1 2 3] len:4 cap:4 ptr:0xc000018220
value：[0 1 2 3 4] len:5 cap:8 ptr:0xc000016140
value：[0 1 2 3 4 5] len:6 cap:8 ptr:0xc000016140
value：[0 1 2 3 4 5 6] len:7 cap:8 ptr:0xc000016140
value：[0 1 2 3 4 5 6 7] len:8 cap:8 ptr:0xc000016140
value：[0 1 2 3 4 5 6 7 8] len:9 cap:16 ptr:0xc00001a100
value：[0 1 2 3 4 5 6 7 8 9] len:10 cap:16 ptr:0xc00001a100
```

> 小tips:

在append里面，可以将一个数组使用...方式来代替一个个取出来赋值

```go
s1 := []int{}
ss := []string{"1","2","3","4"}
s1 = append(s1, ss...) // ...表示拆开
```

#### append 扩容策略：

- 1 新申请的 cap 大于原来的2倍，最终就是新申请的 cap 大小
- 2 如果旧切片长度小于1024，最终是旧的 2 倍
- 3 旧切片长度大于1024，最终容量从旧的开始循环增加原来的 1/4，直到大于等于新申请的 cap 大小。
- 4 最终 cap 溢出，最终就是新申请的 cap 大小

源码：

`/usr/local/go/src/runtime/slice.go`行：144-163

```go
newcap := old.cap
doublecap := newcap + newcap
if cap > doublecap {
	newcap = cap
} else {
	if old.len < 1024 {
		newcap = doublecap
	} else {
		// Check 0 < newcap to detect overflow
		// and prevent an infinite loop.
		for 0 < newcap && newcap < cap {
			newcap += newcap / 4
		}
		// Set newcap to the requested cap when
		// the newcap calculation overflowed.
		if newcap <= 0 {
			newcap = cap
		}
	}
}
```

int 和 string 对于类型的不同处理也不一样。

#### copy：复制切片
复制数据，而不是简单的指针复制。`copy(new, old)`

**注意需要把空的初始化长度计算好,size和cap都要预留好，否则可能copy不全	**

```go
// copy
s12 := []int{1,2,3,4,5}
s13 := make([]int, 5, 5)
copy(s13,s12)
fmt.Println(s12)
fmt.Println(s13)
s13[0] = 100
fmt.Println(s12)
// [1 2 3 4 5]
fmt.Println(s13)
// [100 2 3 4 5]
```

#### 删除元素
切片没有提供删除方法

```go
s1.append(s1[:index], s1[index+1:]...)
```

```go
// 删除
a2 := [...]int{1,2,3} // 数组，习惯数组的这个写法
s14 := a2[:]		// 切片
fmt.Println(s14)
// 删除第2个
s14 = append(s14[:1], s14[2:]...)
fmt.Println(s14) // [1 3]
fmt.Println(a2) // [1 3 3]
// 注意删除之后的底层数组位置没有变，只是东西变了
// 针对切片的修改全部都会反映到底层数组上面
```

理解底层数组最后的样子。删除的时候后面的顶上来，其实最后的位置没有被动，还是保留最后一位的元素，自己还被复制了一份放在自己的前面。


## map
map 类型的变量默认初始值是 nil ，需要使用 make(map, [cap]) 来分配内存。

**注意区分 slice 和 map 的 make 函数的区别。一个是三个参数，一个是两个参数。定义size大小，决定了可以使用 = 赋值几次，超出窗口的需要使用append来添加**

取值后不存在返回类型对应的默认值

### 基本使用

```go
fmt.Println("========1 初始化不分配内存得到nil=======")
var s map[string]string
fmt.Println(s==nil)
s = make(map[string]string, 10) 
s["name"] = "deltaqin"
fmt.Println(s)

fmt.Println("========2 直接分配内存获取map=======")
scores := make(map[string]int, 8)
scores["de"] = 90
scores["lta"] = 100
fmt.Println(scores)
fmt.Println(scores["de"])
fmt.Printf("%T\n", scores)

fmt.Println("========3 新建map时候直接赋值=======")
s1 := map[string]string {
	"name": "deltqin",
	"pass": "123",
}
fmt.Println(s1)

fmt.Println("========4 打印map里面的值=======")
s2 := make(map[string]string, 10)
val, ok := s2["de"]
if !ok {
	fmt.Println("查无此key")
} else {
	fmt.Println(val)
}
fmt.Println(s2)
```


### 遍历
```go
b1 := make(map[string]string , 10)
b1["a"] = "1"
b1["b"] = "2"
// 只k
for k := range b1{
	fmt.Println(k)
} 
// 只v
for _,v := range b1 {
	fmt.Println(v)
}
```

### 删除
删除，为空不进行操作

```go
	delete(b1, "a")
	fmt.Println(b1)
```


### 按照指定顺序遍历map 
map 结果是无序的，内部结构是hash。可以借助切片对 key 排序之后再for 这个切片从 map 取值

```go
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	scores := make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scores[key] = value
	}

	fmt.Println(scores)

	s := make([]string, 0, 200)
	for k := range scores {
		s = append(s, k)
	}

	fmt.Println(s)

	sort.Strings(s)

	for _, key := range s {
		fmt.Println(key, scores[key])
	}

}

```

### 元素是map类型的切片 & 值为切片类型的map
- 切片和map使用的时候一定要初始化，
- 可以使用make，也可以直接赋值初始化
```go
// 切片和map使用的时候一定要初始化，
// 可以使用make，也可以直接赋值初始化

// 值为map的切片
// 注意使用make分配切片内存空间的时候，size一定要指定好，
// 否则size为0，直接使用 = 赋值会失败，只能使用append，因为切片窗口是0
s1 := make([]map[string]int, 10, 10)
s1[0] = make(map[string]int,1)
s1[0]["deltaqin"] = 20
fmt.Println(s1)

// 值为切片的map
// map 在分配map的时候只需要cap，不要多写参数
m1 := make(map[string][]int,10)
m1["deltaqin"] = []int{1,2,3}
fmt.Println(m1)

// 输出
// [map[deltaqin:20] map[] map[] map[] map[] map[] map[] map[] map[] map[]]
// map[deltaqin:[1 2 3]]
```

# 指针
## & *
**go语言不存在程序员可以直接操作的指针，只有取地址符**

`&` 取地址
`*` 根据地址取值

打印出来的东西基本和C语言一致，
eg:
- *int ：说明是int类型的指针，

```go
n := 18
p := &n
t := *p

fmt.Println(p)
fmt.Println(t)
fmt.Printf("%T\n", p)
// 0xc00001c0b0
// 18
// *int
```

## new make
new函数申请一个内存地址，开辟内存空间，返回的是对应内存的指针。new是给基本数据类型申请空间的，一般很少用

```go
// new 返回的是空间的地址
var p1 *int
fmt.Println(p1)
p1 = new(int)
fmt.Println(p1)
fmt.Println(*p1)
*p1 = 100
fmt.Println(*p1)
```

**make函数申请一个内存地址，只用于 slice map channel 的内存创建，返回的是内容本身而不是内容的指针，因为本身就是引用类型，没必要返回指针的指针**

**注意make 在slice和map的区别，参数个数不一样，赋值错误将会导致后续添加值的时候报错，清楚每一个参数的作用，**
- slice：类型，size，cap
  - size 若为0，添加元素的时候只能使用append，而不能使用=
- map：类型，cap



# 函数
函数传参全是拷贝，和C语言一样

在一个命名的函数里面，不可以声明再次声明命名函数
## 函数定义--参数写法--返回值写法
- 参数
  - 无参数
  - 指定类型和名字
  - 省略类型（相同的时候）
  - 可变参数
- 返回值
  - 无参数
  - 只指定类型
  - 指定类型和返回值
  - 多返回值的省略类型
  - 可变参数

```go
func f2(x int, y int) {
	fmt.Println("没有返回值")
}

func f3() {
	fmt.Println("没有参数没有返回值")
}

func f4() int {
	fmt.Println("没有参数有返回值")
	ret := 3
	return ret
}

func f5(x int, y int) (ret int) {
	fmt.Println("命名的返回值，相当于在函数里面声明了一个变量")
	ret = x + y
	return // 命名返回值，return后面可以省略
}

func f6() (int, string) {
	fmt.Println("多个返回值")
	return 1, "deltaqin"
}

func f7(x, y int, z, m string) int {
	fmt.Println("参数类型简写，连续多个参数类型一致，可以将最后一个之前的参数类型去掉")
	return x + y
}

func f8(x string, y ...int) {
	fmt.Println("可变长参数，必须放在最后")
	fmt.Println(x)
	fmt.Println(y)
	// y是切片类型 []int
}

func main() {

	f2(1, 2)
	f3()
	r4 := f4()
	r5 := f5(1, 2)
	r6, r61 := f6()
	r7 := f7(1, 2, "1", "2")
	f8("1",1,2,3)
	fmt.Println(r4, r5, r6, r61, r7)

}
```


## 函数作用域
- 全局作用域
- 函数作用域
  - 先在函数内部找，找不到外部找
  - 函数内部的变量，在外部是访问不到的
- 代码块作用域

## defer语句
**当前函数即将返回的时候再执行defer后面的。（用于最后的资源释放：关闭数据库连接等等，可以有多个）**

defer 语句会将其后面跟随的语句延迟执行，在函数返回之后，延迟的语句按照defer定义的逆序执行，先被defer的语句最后被执行，最后defer的最先被执行。

**发生在返回值赋值之后，真正返回指令执行之前。**go语言中的return不是原子操作，是包含两步的：`返回值 = x`以及`RET 指令`。defer执行的时机是在二者之间的。

- 返回值 = x
- 运行defer
- RET 指令

注意上面的执行时机，在返回值之后对返回值进行修改是不会有任何效果的。

```go
fmt.Println("start")
defer fmt.Println(1)
defer fmt.Println(2)
defer fmt.Println(3)
fmt.Println("end")

// start
// end
// 3
// 2
// 1
```

面试题1

```go
func f1() int {
	x := 5
	defer func()  {
		x++
	}()
	return x  // 5 x是函数自己定义的变量，不是返回变量
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 6 改变的就是返回值x，defer自己没有声明变量
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 5 没有改变y
}

func f4() (x int) {
	defer func(x int)  {
		x++
	}(x)
	return 5  // 5 函数传参是拷贝，不是原来的x了
}

func f5() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5 // 6 立即执行函数传递进去的是指针
}

func main() {
	r1 := f1()
	r2 := f2()
	r3 := f3()
	r4 := f4()
	fmt.Println(r1, r2, r3, r4)
}
```

面试题2

defer里面如果调用了函数，会先执行，得到的结果作为参数在压入栈的时候放进去

defer压入的时候参数都是确切的状态，必须变成实际的值而不是变量，后面的变化不会影响到已经压入栈的函数的参数

```go
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 1
	// 执行到这里会先将参数里面的 calc("10", a, b)=2 算出来
	// 再把calc整体压入栈里面，并不是到时候一起算
	// 压入的时候参数是几就是几，后面的值不会影响到这里，所以a是1不是0
	defer calc("1", a, calc("10", a, b)) // 1 1 2
	a = 0
	// 执行到这里会先将参数里面的 calc("20", a, b)=1 算出来
	// 再把calc整体压入栈里面
	defer calc("2", a, calc("20", a, b)) // "2" 0 1 
	b = 1
	fmt.Println(111)
	// 10 1 1 2
	// 20 0 1 1
	// 111
	// 2 0 1 1
	// 1 1 2 3
}
```

## 高阶函数--函数类型和变量
参数，返回值，函数都是变量

函数可以作为参数和返回值

```go
package main

import "fmt"

func f1() {
	fmt.Println("1")
}

func f2() int {
	fmt.Println("2")
	return 10
}

func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func f4(x, y int) int {
	return x + y
}

func temp(a, b int) int {
	return a + b
}

func f5(x func() int) func(int,int) int {
	return temp
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)

	b := f2
	fmt.Printf("%T\n",b)

	f3(f2)

	d := f4
	fmt.Printf("%T\n",d)

	// 注意传递参数的时候，函数的类型要一致（包括参数以及返回值的个数和类型）
	e := f5(f2)
	fmt.Printf("%T\n", e)

	
	// func()
	// func() int
	// 2
	// 10
	// func(int, int) int
	// func(int, int) int
}
```

## 匿名函数
**一般在函数内部使用，因为命名函数内部不能再定义有名字的函数**

- **只使用一次，立即执行函数**
- **赋值给一个变量**


```go
// 1 匿名函数赋值给变量
var anonymous = func(x, y int) {
	fmt.Println(x, y)
}
anonymous(1,2)

// 2 立即执行函数
func (x,y int)  {
	fmt.Println(x, y)
}(1,2)
```

## 闭包 -- 装饰者模式
闭包（一个实体） = 函数 + 引用环境

一般在使用外界接口的时候，自己不能直接调用传递相应的参数，可以使用

参数不一致的时候，包装。需要的类型作参数，需要的参数类型作为返回的类型

**简单理解：就是装饰者模式，不修改原有函数的基础上加上自己像加的东西**

```go
func f1(f func()) {
	fmt.Println("f1")
	f()
}

func f2(x, y int) {
	fmt.Println("f2")
	fmt.Println(x, y)
}

func f3(f func(int, int), m, n int) func() {
	// 闭包：一个函数可以引用他外面的变量，这里的f和mn都是环境的变量
	// 最后main调用的时候看起来是使用了 temp ，实际上就是装饰者模式，
	// 在原有的方法基础上定义自己的相关属性
	temp := func() {
		f(m, n)
	}
	return temp
}

func main() {
	// 匿名函数的使用：闭包
	// 立即执行，相当于调用了 f2
	f3(f2, 1, 2)()
}
```

```go
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpg := makeSuffix(".jpg")
	png := makeSuffix(".png")

	name1 := jpg("1")
	name2 := png("2")

	fmt.Println(name1)
	fmt.Println(name2)
}
```

## 内置函数
- close：关闭channel
- len：求长度，string array slice map channel
- new：分配值类型的内存，返回的是指针，eg:int, struct
- make：分配引用类型的内存，返回的是值, eg: chan, map, slice
- append：追加元素到数组、slice中
- panic/recover：做错误处理

### panic/recover
panic 需要和 defer 配合使用，因为 panic 报错之后不会执行后面的代码，需要将处理异常代码放在defer里面

panic 相当于 java 里面的 throw， 后续代码不会执行

- recover 必须搭配defer
- recover 必须写在 panic 之前

defer里面一般使用 recover() 获取错误，解决之后继续后面的代码，这里defer就相当于 catch ，注意 try 的范围相当于是这一整个函数，只会跳过这一步执行下面的函数，但是当前产生异常的函数后面的代码不会执行了

```go
func f1() {
	// 如果是在defer后面使用匿名函数，一定要执行，不要忘记立即执行
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	panic("严重错误")
	fmt.Println("这里不会执行")
}

func main() {
	f1()
	fmt.Println("这里会执行")
}
```


# 结构体

## 自定义类型和类型别名
基于内置类型构造自己的类型

类型别名编译之后不存在

```go
// 自定义类型，编译之后依旧是自定义类型
var myInt int
// 类型别名，编译之后替换为int，编写的时候更加清晰一些
	// 代码的作用，错误代码等等
var myint = int
```

注意：rune 就是 int32 的类型别名。

## 结构体定义

```go
// 丰富维度的数据
// 自己造类型
type st1 struct {
	name  string
	age   int
	gen   string
	hobby []string
}

func main() {
	var de st1
	de.name = "de"
	de.age = 12
	de.gen = "男"
	de.hobby = []string{"篮球", "足球"}

	fmt.Println(de)
	fmt.Printf("%T\n", de)
	fmt.Println(de.age)
// 	{de 12 男 [篮球 足球]}
// main.st1
// 12
}
```

## 匿名结构体
临时使用，不浪费全局空间 

```go
func main() {
	// 直接使用变量声明
	var s struct{
		x string
		y int 
	}{"de", 12,}

	s.x = "de"
	fmt.Printf("%T\n", s)
	fmt.Println(s)
}
```

## 结构体初始化

```go 
type person struct {
	name string
	age  int
}

// kv初始化
var p3 = person{
	name: "deltaqin",
	age: 16,
}
// 按照值列表顺序初始化
p4 := person {
	"deltaqin",
	12,
}
```

参照切片和map初始化

```go
s1 := []int {1,2,3,4}
m1 := map[string]int {
	"stu1": 1,
	"stu2": 2,
}
```

## 指针类型的结构体
由于go不可以修改指针，所以改变指针执行的内存地址上的值的时候可以不加 `*`

```go
type st1 struct {
	name string
	age  int
}

func f(x st1) {
	x.name = "qin1"
}

func f1(x *st1) {
	(*x).name = "qin1"
}

func f2(x *st1) {
	x.name = "qin2"
}

func main() {
	var p st1
	p.name = "qin"
	p.age = 1
	f(p)
	fmt.Println(p)
	f1(&p)
	fmt.Println(p)
	f2(&p)
	fmt.Println(p)
}
```

```go
	var p st1
	fmt.Printf("%p\n", &p) // 0xc0000044a0

	// p1 是结构体指针
	var p1 = new(st1)
	fmt.Printf("%x\n", p1) // &{ 0}
	fmt.Printf("%x\n", &p1) // c000006030
	fmt.Printf("%p\n", &p1) // 0xc000006030
	fmt.Printf("%#v\n", p1) // &main.st1{name:"", age:0}
	fmt.Printf("%T\n", p1) // *main.st1
	// 可以直接对结构体指针使用 . 来获取属性
	p1.name = "deltaqin"
	p1.age = 12
	fmt.Printf("%#v", p1) // &main.st1{name:"deltaqin", age:12}

```

### 在内存上是连续的地址
```go
	p2 := st1{
		name: string("deltaqin"),
		age: int8(20),
	}

	fmt.Printf("%p\n", &(p2.name))
	fmt.Printf("%p\n", &(p2.age))
	// 0xc0000045c0，16位
	// 0xc0000045d0，16位
	// 还有内存对齐，在必要的时候
```

## 构造函数
- 可以返回的是指针
- 可以直接返回值，结构体本身是值类型不是引用类型，返回的时候就是值拷贝

- 开头使用new


```go
type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func newPerson2(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
	// {deltaqin 18}
	// 0xc000004500
}
```

## 补充：标识符
变量名，函数名，类名，方法名。**首字母大写，就是对外部包可见的。**

`fmt.Println()`

## 方法和接收者
**方法是作用于特定类型的函数**,

接收者就是调用该方法的类型

`func (接收者名字 接收者类型) 函数名字(参数名 参数类型) (返回参数) {}`

接收者名字最好使用接收者类型的首字母，而不是 this 和 self

```go
type dog struct {
	name string
}

func newDog(name string ) *dog {
	return &dog{
		name: name,
	}
}
// 和函数的区别就是函数名前面需要制定具体的类型变量，
// 这个只是值拷贝，不是指针拷贝
// 多用类型名首字母小写
func (d dog) wang() {
	fmt.Println("wangwang")
}

func main() {
	// 方法是作用于特定类型的函数
	d1 := newDog("de")
	d1.wang()
}
```


## 值接收和指针接收的区别
- 值接收不会改变类的属性，是值传递，结构体是值类型，是复制

- **指针接收就可以改变类的属性**


```go
func newDog(name string ) *dog {
	return &dog{
		name: name,
	}
}

func (d dog) wang() {
	// fmt.Println("wangwang")
	d.name = "1"
}

func (d *dog) wang1() {
	// fmt.Println("wangwang")
	d.name = "2"
}

func main() {
	// 方法是作用于特定类型的函数
	d1 := newDog("de")
	d1.wang()
	fmt.Println(d1)
	d1.wang1()
	fmt.Println(d1)

}
```

## 任意类型添加方法
只能给自己的包里面的类添加方法，想要给别的包需要自定义类型

```go 
type myInt int

func (i myInt) my()  {
	fmt.Println("直接操作int不可以，不是自己包定义的")
}

func main() {
	m := myInt(10)
	m.my()
	fmt.Println(m)
}
```

## 结构体匿名字段
相同类型只能写一个，字段少且简单的场景

```go
type stu struct {
	string
	int
}

func main() {
	stu1 := stu{
		"de",
		1,
	}
	fmt.Println(stu1)
}
```

## 结构体嵌套
匿名结构体嵌套在里面，使用 `.` 作为属性查询的时候，现在自己的属性，之后在匿名结构体属性里面查找

匿名嵌套结构体，可以省略中间字段的名字

匿名嵌套结构体的字段冲突解决，直接写全名

```go
type company struct{
	name string
	id int64
}

type address struct{
	name string
	id int64
}

type student struct{
	name string
	id int64
	address
	company
}

func main() {
	stu := student{
		name: "de",
		id : 1, 
		address: address{
			name: "de",
			id: 1,
		},
	}
	fmt.Println(stu.address.name)
}
```

## 结构体继承

```go
type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s move", a.name)
}

type dog struct {
	feet uint8
	animal
}

func (d dog) wang()  {
	fmt.Printf("%s wang\n", d.name)
}
func main() {
	d := dog{
		feet:4,
		animal: animal{
			name: "de",
		},
	}
	d.wang()
	// 继承匿名 animal，方法就全有了，变相实现继承
	d.move()
}
```

## 结构体和json
- Name不大写外界无法使用
- 反序列化需要传递指针


```go 
import (
	"encoding/json"
	"fmt"
)

type person struct {
	// Name不大写外界无法使用
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "de",
		Age:  16,
	}

	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	fmt.Printf("%v\n", string(b))

	// 字符串就是字节数组的切片
	str:= `{"name":"de", "age":12}`
	var p2 person
	// 反序列化需要传递指针
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%#v\n", p2)

}

```

# 接口
接口是一种类型，不关心变量是什么类型，只关心可以调用什么方法

基本数据类型，结构体类型，

- 面向接口编程：只关心方法，实现方法就可以了，**解耦**
- 面向对象编程：关心属性和方法

## 接口定义与实现（面向接口编程）
**实现了接口的所有方法，这个变量就实现了这个接口，可以叫做这个接口类型的变量**

```go
type 接口名 interface {
	方法名(参数)(返回值)
}
```

```go
// 只规定方法
type speaker interface {
	// 实现了speak方法的变量就是该接口类型
	speak()
}

type dog struct{}
type cat struct{}

func (d dog) speak() {
	fmt.Println("wang")
}

func (c cat)speak()  {
	fmt.Println("miao")
}

func attack(s speaker)  {
	s.speak()
}

func main() {
	var d dog
	var c cat
	attack(d)
	attack(c)
}
```

## 接口原理
接口保存的时候分为 两部分：**值的类型和值本身**，实现了存储不同类型和不同的值

实现类直接赋值给接口类型的变量的时候，接口变量变成了具体的实现类（可以理解为多态）

```go
type animal interface {
	eat()
}

type cat struct{}
type dog struct{}

func (c cat) eat() {
	fmt.Println("fish")
}

func (d dog) eat() {
	fmt.Println("骨头")
}

func main() {
	var a animal
	a = cat{}
	a.eat()
	a = dog{}
	a.eat()
}
```

## 值接收者和指针接收者
值接收者：可以接收指针和值类型

```go
type animal interface {
	eat()
}

type cat struct{}
type dog struct{}

func (c cat) eat() {
	fmt.Println("fish")
}

func (d dog) eat() {
	fmt.Println("骨头")
}

func main() {
	var a1 animal
	c1 := cat{}
	a1 = c1
	a1.eat()
}
```

指针接收者：可以接收指针类型

```go
type animal interface {
	eat()
}

type cat struct{}
type dog struct{}

func (c *cat) eat() {
	fmt.Println("fish")
}

func (d *dog) eat() {
	fmt.Println("骨头")
}

func main() {
	var a1 animal
	c1 := cat{}
	// 注意接口接收的类型应该和实现的方法里面定义的需要的类型一致，所以这里使用取地址符
	a1 = &c1
	a1.eat()
}
```

## 接口和类型的关系
**同一结构体可以实现多个接口**

接口也可以嵌套，直接将接口名字写在接口定义里面即可

将接口看做是一个规章制度，不是一个具体的对象拥有什么样的属性


## 空接口类型
如果一个结构体实现了一个接口所有的方法，那就是实现了这个接口

如果一个接口什么方法都没有，任何类型都实现了这个接口

`type a interface{}`

`interface{}` 没必要起名字

### 空接口作为函数的参数
不管什么类型都可以传递给函数，就可以将函数的参数设置为空接口 `...interface{}`

```go
func show(a interface{}){
	fmt.Println("type:%T value:%v\n", a, a)
}
```
### 空接口作为 map 的值
使用空接口实现可以保存任意值的字典

```go
func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["1"] = "1"
	m1["2"] = 2
	m1["3"] = true
	m1["4"] = [...]string{"1","2","3"}
	fmt.Println(m1)
}
```

## 类型断言
空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢？

### 接口值
一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值。

```go
var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)
w = nil
```

想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：`x.(T)`
- x：表示类型为interface{}的变量
- T：表示断言x可能是的类型。

该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。

```go
func main() {
	var x interface{}
	x = "de"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
```

上面的示例中如果要断言多次就需要写多个if判断，这个时候我们可以使用switch语句来实现：

```go
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
```

因为空接口可以存储任意类型值的特点，所以空接口在Go语言中的使用十分广泛。

关于接口需要注意的是，只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗。

# 包
只有main包才可以编译为可执行文件



# 并发编程

```go

```


# vscode 设置 snippets

```json
{
	// Place your snippets for go here. Each snippet is defined under a snippet name and has a prefix, body and 
	// description. The prefix is what is used to trigger the snippet and the body will be expanded and inserted. Possible variables are:
	// $1, $2 for tab stops, $0 for the final cursor position, and ${1:label}, ${2:another} for placeholders. Placeholders with the 
	// same ids are connected.
	// Example:
	// "Print to console": {
	// 	"prefix": "log",
	// 	"body": [
	// 		"console.log('$1');",
	// 		"$2"
	// 	],
	// 	"description": "Log output to console"
	// }
	"pln": {
		"prefix": "pln",
		"body": "fmt.Println($0)",
		"description": "fmt.Println()"
	},
	"plf": {
		"prefix": "plf",
		"body": "fmt.Printf(\"$0\\n\",)",
		"description": "fmt.Printf()"
	}


}
```





