// goguide
package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"reflect"
	"runtime"
)

func main() {
	//randFunc()

	//capitalFunc()

	//fmt.Println(euclid(45, 18))

	//a, b := swap("previous", "later")
	//fmt.Printf("Former: previous, later; Now: %v, %v \n", a, b)

	//varFunc()

	//typesFunc()

	//forFunc()

	//fmt.Println(ifFunc(3, 2, 10))
	//fmt.Println(ifFunc(3, 3, 20))

	//fmt.Println("sqrtFunc: ", sqrtFunc(3))
	//fmt.Println("math.Sqrt: ", math.Sqrt(3))

	//switchFunc("1")
	//switchFunc("2")
	//switchFunc("3")
	//switchFunc("4")
	//switchFunc("5")
	//switchFunc("6")
	//switchFunc("7")

	//deferFunc()

	//pointerFunc()

	//var a, b, c, d = Point{4, 3}, Point{0, 3}, Point{4, 0}, Point{-1, -1}
	//fmt.Println(structFunc(a, b, c))
	//fmt.Println(structFunc(d, b, c))
	//p := &a
	//fmt.Println(p, ":", *p, " : ", a, " : ", p.X, " : ", p.Y)
	////还可以通过Name:value的形式进行部分初始化
	//e := Point{X: 233}
	//f := Point{Y: 666}
	//fmt.Println(e, f)

	//arrayFunc()

	//fmt.Println(example(3))

	sliceFunc()
}

func randFunc() {

	//包名与导入路径最后一个目录一致，例如rand
	//这还是随机数么。。。每次运行都一样啊。。。
	//继续学习
	fmt.Println(rand.Intn(5))
	fmt.Println(rand.Intn(5))
	fmt.Println(rand.Intn(5))
	fmt.Println(rand.Intn(5))
	fmt.Println(rand.Intn(5))
}
func capitalFunc() {
	//首字母大写包外可访问，相当于其他语言的 public
	fmt.Println(math.Pi)
	//fmt.Println(math.pi)
}
func euclid(a, b int) int {
	//函数形参是变量名前类型后，列表后是返回值类型
	//声明变量时同一类型可以在一行声明，只留最后一个type
	//   x, y int
	if b == 0 {
		return a
	}
	return euclid(b, a%b)
}

func swap(a, b string) (string, string) {
	//go可以有多个返回值
	return b, a
}

var outer int

func varFunc() {
	//var可以定义在包级别和函数级别
	//这里看出，int值默认是0
	var local int
	//多个变量赋值，可以省略类型
	var v1, v2, v3 = 233, 2.33, false
	fmt.Println(outer, local)
	if !v3 {
		fmt.Println("v1 / v2 = ", float64(v1)/v2)
	}
	//短声明变量
	//函数中可以用:=来赋值，代替var。
	//函数外每个语句必须以'var','func'开始，:=也不能在函数外使用
	//可以用reflect包的Typeof获取变量类型；或者用%T来得到类型
	sh := 1000000000000000
	fmt.Printf("%T \n", sh)
	fmt.Println(reflect.TypeOf(sh))
}

func typesFunc() {
	//Go 的基本类型有Basic types

	//bool

	//string

	//int  int8  int16  int32  int64
	//uint uint8 uint16 uint32 uint64 uintptr

	//byte uint8 的别名

	//rune int32 的别名
	//     代表一个Unicode码

	//float32 float64

	//complex64 complex128
	var (
		boo   bool       = false
		maxIn uint64     = 1<<64 - 1
		comp  complex128 = cmplx.Sqrt(12 - 33i)
	)
	fmt.Println(!boo)
	fmt.Println(maxIn)
	fmt.Println(comp)
	//变量默认为零值，数值为 0， bool 为 false， 字符串为 “”
	var i int
	var b bool
	var f float64
	var s string
	fmt.Println(i, " ", b, " ", f, " ", s)
	fmt.Printf("%v %v %v %q \n", i, b, f, s)

	//类型转换
	//不同类型变量赋值时需要显式转换
	var v1, v2 = 1.732, 3
	var r float64 = (math.Sqrt(float64(v2)))
	fmt.Println(v1, " = ", r, " ? ", v1 == r)
	//const 关键字可以声明常量 常量不能用:=来定义
	const Haha = "I am angry !"
	//Haha = "excited~"
	fmt.Println(Haha)

	//数值常量
	const (
		big   = 1 << 100
		small = big >> 110
	)
	fmt.Println("small int ", needInt(small))
	fmt.Println("big float ", needFloat(big))
	fmt.Println("small float ", needFloat(small))
}
func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func forFunc() {
	//go中的for是没有（）的，但是{}是必须有的
	//下面是一个十阶的fabonacci
	var a, b = 1, 1
	for i := 0; i < 10; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
	//for中的前置后置语句可空
	//for就可以相当于别的语言中的while
	for {
		b <<= 2
		if b > 500 {
			break
		}
	}
	fmt.Println(b)
}

func ifFunc(x, n, lim float64) float64 {
	//if的条件判断之前可以执行一个简单语句
	//else中还可以用if简单语句中的变量
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println("here v equals ", v)
		return lim
	}
}

func sqrtFunc(x float64) float64 {
	//练习，牛顿法计算开方
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func switchFunc(start string) {
	//switch也可以在条件前加简单语句
	//如果case语句中无fallthrough，那么默认执行到case末尾
	//switch无条件时，相当于switch true，这种形式可以在case 后添加条件，
	//用一种更清晰的方式编写if-then-else
	switch os := runtime.GOOS; start {
	case "1":
		fmt.Print(" 1: ", os)
	case "2":
		fmt.Print(" 2: ", os)
	case "3":
		fmt.Print(" 3: ", os)
	case "4":
		fmt.Print(" 4: ", os)
		fallthrough
	case "5":
		fmt.Print(" 5: ", os)
		fallthrough
	case "6":
		fmt.Print(" 6: ", os)
		fallthrough
	default:
		fmt.Print(" Default: ", os)
	}
	fmt.Println()
}

func deferFunc() {
	//defer语句的函数会延迟到上层函数返回前执行，但是参数会立刻生成
	//defer延迟的函数调用压入栈中，按照LIFO顺序调用函数
	a := 2333
	defer fmt.Println("defer ", a)
	a += a
	fmt.Println("common ", a)
	for i := 0; i < 5; i++ {
		defer fmt.Println("loop ", i+1, ", value : ", (i+1)*a)
	}
}

func pointerFunc() {
	//和C中的指针很像，*T表示T类型的指针，*p表示p指针指向的值，&v生成变量v的指针
	//但是不同的是go中的指针不可以指针运算
	var p *string
	str := "pointers~"
	p = &str
	fmt.Println("str: ", *p, " &str: ", p)
	//无指针运算
	//fmt.Println(" *(&str+2): ", *(p + 2))
}

type Point struct {
	X int
	Y int
}

func powerOfDist(x, y Point) int {
	return (x.X-y.X)*(x.X-y.X) + (x.Y-y.Y)*(x.Y-y.Y)
}
func structFunc(x, y, z Point) bool {
	//判断三个点能否够构成直角三角形
	var a, b, c = powerOfDist(x, y), powerOfDist(y, z), powerOfDist(x, z)
	if a+c == b || a+b == c || b+c == a {
		return true
	} else {
		return false
	}
}
func arrayFunc() {
	//[n]T表示n个T类型元素的数组
	//数组不可以改变大小
	var a [5]string
	for i := 0; i < 5; i++ {
		a[i] = "一夫振臂万夫雄!"
	}
	fmt.Println(a)
	fmt.Println(a[3])
	l := len(a[4])
	for i := 0; i < l; i++ {
		//字符串用下标访问得到的是byte
		fmt.Print(i, " : ", a[4][i], ", ")
	}
	fmt.Println()
	for i, ch := range a[4] {
		//字符串用range访问得到的是Unicode编码的rune
		fmt.Print(i, " : ", ch, ", ")
	}
	fmt.Println()
}

//这是许式伟书中的例子，说把return完全放到if else中会报错
//然而并没有。。。是版本的问题，在Go1.0中是会出错的，1.1之后就可以编译了
func example(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}
func sliceFunc() {
	a := []int{1, 7, 3, 2, 0, 5, 0, 8, 0, 7}
	//和Python中的切片规则类似，左闭右开区间
	s := a[3:]
	fmt.Println(a)
	fmt.Println(s)
	//make可以指定类型、长度、容量
	m := make([]int, 5, 10)
	//copy可以将第二个参数内容复制到第一个，以短的那个为基准
	copy(m, s)
	fmt.Println(m, len(m), cap(m))
	//append方法，第一个是要追加的目的切片，剩余参数数可选，也可以是另一个切片，后头需要加“...”
	m = append(m, a...)
	fmt.Println(m)
	temp := a
	for l := len(a); l > 0; l-- {
		temp = temp[:l]
		fmt.Println(temp)
	}
}
