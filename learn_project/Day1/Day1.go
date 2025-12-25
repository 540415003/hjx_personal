// ==========================================
// Day 1: Go语言基础入门
// 主题：环境、基本语法、数据类型、变量、常量
// ==========================================

package main

import (
	"fmt"
	"math"
	"strings"
)

// ==========================================
// 1. Hello World - 你的第一个Go程序
// ==========================================
// 知识点：
// - package main: 程序入口包（类似Java的main类）
// - import: 导入包（类似Java的import）
// - func main(): 程序入口函数（类似Java的public static void main）
// - fmt: 格式化I/O包（类似Java的System.out）

func helloWorld() {
	fmt.Println("Hello, Go!")
	// Java对比: System.out.println("Hello, Java!");
}

// ==========================================
// 2. 变量声明 - Go有多种声明方式
// ==========================================

func variableDeclaration() {
	fmt.Println("\n========== 变量声明 ==========")

	// 方式1: var 变量名 类型 = 值
	var name string = "张三"
	fmt.Println("方式1:", name)

	// 方式2: var 变量名 = 值 (类型推断)
	var age = 25
	fmt.Println("方式2:", age)

	// 方式3: := 短声明 (最常用，只能在函数内使用)
	city := "北京"
	fmt.Println("方式3:", city)
	// Java对比: String city = "北京";

	// 方式4: 批量声明
	var (
		username = "admin"
		password = "123456"
		isActive = true
	)
	fmt.Printf("用户名: %s, 密码: %s, 激活: %v\n", username, password, isActive)

	// 方式5: 多变量同时声明
	var x, y, z int = 1, 2, 3
	fmt.Println("多变量:", x, y, z)

	a, b := "Hello", "World"
	fmt.Println("短声明多变量:", a, b)

	// 注意：Go中未使用的变量会导致编译错误！
	// 这与Java不同，Go更严格
}

// ==========================================
// 3. 基本数据类型
// ==========================================

func basicDataTypes() {
	fmt.Println("\n========== 基本数据类型 ==========")

	// 布尔类型
	var isTest bool = true
	fmt.Printf("布尔类型: %v (类型: %T)\n", isTest, isTest)
	// Java对比: boolean isTest = true;

	// 字符串类型 (UTF-8编码，支持中文)
	var message string = "Go语言测试"
	fmt.Printf("字符串: %s (长度: %d字节)\n", message, len(message))
	// Java对比: String message = "Go语言测试";

	// 整数类型 - Go有多种整数类型
	var num1 int = 42           // 根据平台自动选择32或64位
	var num2 int8 = 127         // -128 到 127
	var num3 int16 = 32767      // -32768 到 32767
	var num4 int32 = 2147483647 // 约-21亿到21亿
	var num5 int64 = 9223372036854775807
	fmt.Printf("int: %d, int8: %d, int16: %d, int32: %d, int64: %d\n",
		num1, num2, num3, num4, num5)
	// Java对比: int对应int32, long对应int64

	// 无符号整数
	var unum1 uint = 42
	var unum2 uint8 = 255 // 0 到 255
	fmt.Printf("uint: %d, uint8: %d\n", unum1, unum2)

	// 浮点数
	var price float32 = 99.99
	var pi float64 = 3.14159265359
	fmt.Printf("float32: %.2f, float64: %.10f\n", price, pi)
	// Java对比: float和double

	// 复数（Go特有，Java没有原生支持）
	var c1 complex64 = 1 + 2i
	var c2 complex128 = 3 + 4i
	fmt.Printf("复数: c1=%v, c2=%v\n", c1, c2)

	// byte类型 (uint8的别名，常用于处理ASCII字符)
	var b byte = 'A'
	fmt.Printf("byte: %c (数值: %d)\n", b, b)

	// rune类型 (int32的别名，用于Unicode字符)
	var r rune = '中'
	fmt.Printf("rune: %c (Unicode: %U)\n", r, r)
	// Java对比: char在Java中是16位，Go的rune是32位
}

// ==========================================
// 4. 零值 - Go的默认值
// ==========================================

func zeroValues() {
	fmt.Println("\n========== 零值（默认值） ==========")

	var i int
	var f float64
	var b bool
	var s string
	// Go会自动初始化为零值，不会像某些语言那样是未定义
	fmt.Printf("int零值: %d\n", i)      // 0
	fmt.Printf("float零值: %f\n", f)    // 0.0
	fmt.Printf("bool零值: %v\n", b)     // false
	fmt.Printf("string零值: '%s'\n", s) // ""（空字符串）
	// Java对比: 成员变量有默认值，局部变量必须初始化
}

// ==========================================
// 5. 类型转换 - Go需要显式转换
// ==========================================

func typeConversion() {
	fmt.Println("\n========== 类型转换 ==========")

	var i int = 42
	var f float64 = float64(i) // 必须显式转换
	var u uint = uint(f)
	fmt.Printf("int: %d -> float64: %.2f -> uint: %d\n", i, f, u)

	// Go不允许隐式类型转换！
	// var x int = 10
	// var y float64 = x  // 错误！必须显式转换
	var x int = 10
	var y float64 = float64(x) // 正确
	fmt.Printf("显式转换: %d -> %.1f\n", x, y)

	// 字符串和数字的转换需要使用strconv包（后面会学）
	// Java对比: Java允许某些隐式转换，Go更严格
}

// ==========================================
// 6. 常量 - const关键字
// ==========================================

func constants() {
	fmt.Println("\n========== 常量 ==========")

	// 单个常量
	const pi = 3.14159
	const apiKey = "your-api-key-here"
	fmt.Printf("常量: pi=%.5f, apiKey=%s\n", pi, apiKey)
	// Java对比: final double PI = 3.14159;

	// 批量常量
	const (
		StatusOK       = 200
		StatusNotFound = 404
		StatusError    = 500
	)
	fmt.Printf("HTTP状态码: OK=%d, NotFound=%d, Error=%d\n",
		StatusOK, StatusNotFound, StatusError)

	// 常量表达式
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
	)
	fmt.Printf("存储单位: 1GB = %d bytes\n", GB)
}

// ==========================================
// 7. iota - Go的枚举计数器
// ==========================================

func iotaExample() {
	fmt.Println("\n========== iota枚举 ==========")

	// iota: 在const块中，从0开始自动递增
	const (
		Sunday    = iota // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
	fmt.Printf("星期: Sunday=%d, Monday=%d, Friday=%d\n", Sunday, Monday, Friday)
	// Java对比: enum DayOfWeek { SUNDAY, MONDAY, ... }

	// iota的高级用法
	const (
		_  = iota             // 0 (使用_忽略)
		KB = 1 << (10 * iota) // 1 << 10 = 1024
		MB                    // 1 << 20 = 1048576
		GB                    // 1 << 30 = 1073741824
		TB                    // 1 << 40
	)
	fmt.Printf("存储单位: KB=%d, MB=%d, GB=%d, TB=%d\n", KB, MB, GB, TB)

	// 权限示例（位运算）
	const (
		ReadPermission    = 1 << iota // 1 << 0 = 1  (二进制: 001)
		WritePermission               // 1 << 1 = 2  (二进制: 010)
		ExecutePermission             // 1 << 2 = 4  (二进制: 100)
	)
	fmt.Printf("权限: Read=%d, Write=%d, Execute=%d\n",
		ReadPermission, WritePermission, ExecutePermission)
}

// ==========================================
// 8. 格式化输出 - fmt包的常用函数
// ==========================================

func formattedOutput() {
	fmt.Println("\n========== 格式化输出 ==========")

	name := "测试工程师"
	age := 28
	salary := 15000.50

	// Print系列：输出后不换行
	fmt.Print("这是Print，")
	fmt.Print("不会自动换行\n")

	// Println系列：输出后换行
	fmt.Println("这是Println，会自动换行")

	// Printf系列：格式化输出（最常用）
	fmt.Printf("姓名: %s, 年龄: %d, 工资: %.2f\n", name, age, salary)

	// 常用格式化动词
	var testNum = 255
	fmt.Printf("%%d 十进制: %d\n", testNum)
	fmt.Printf("%%b 二进制: %b\n", testNum)
	fmt.Printf("%%o 八进制: %o\n", testNum)
	fmt.Printf("%%x 十六进制(小写): %x\n", testNum)
	fmt.Printf("%%X 十六进制(大写): %X\n", testNum)

	var testFloat = 123.456
	fmt.Printf("%%f 默认浮点: %f\n", testFloat)
	fmt.Printf("%%.2f 保留2位小数: %.2f\n", testFloat)
	fmt.Printf("%%e 科学计数法: %e\n", testFloat)

	fmt.Printf("%%s 字符串: %s\n", "Go语言")
	fmt.Printf("%%v 默认格式: %v\n", struct{ Name string }{"测试"})
	fmt.Printf("%%+v 带字段名: %+v\n", struct{ Name string }{"测试"})
	fmt.Printf("%%T 类型: %T\n", 123)

	// Java对比: System.out.printf() 或 String.format()
}

// ==========================================
// 9. 数学运算
// ==========================================

func mathOperations() {
	fmt.Println("\n========== 数学运算 ==========")

	a, b := 10, 3

	// 基本运算
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d (整数除法)\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d (取余)\n", a, b, a%b)

	// 浮点除法
	var x, y float64 = 10, 3
	fmt.Printf("%.1f / %.1f = %.4f (浮点除法)\n", x, y, x/y)

	// 自增自减（注意：Go中只有后置，没有前置）
	count := 0
	count++ // 可以
	// ++count  // 错误！Go不支持前置++
	fmt.Printf("自增后: %d\n", count)
	// Java对比: 支持++i和i++

	// math包的常用函数
	fmt.Printf("绝对值: |%.1f| = %.1f\n", -5.5, math.Abs(-5.5))
	fmt.Printf("向上取整: %.1f -> %.0f\n", 5.3, math.Ceil(5.3))
	fmt.Printf("向下取整: %.1f -> %.0f\n", 5.8, math.Floor(5.8))
	fmt.Printf("四舍五入: %.1f -> %.0f\n", 5.5, math.Round(5.5))
	fmt.Printf("平方根: √16 = %.0f\n", math.Sqrt(16))
	fmt.Printf("幂运算: 2³ = %.0f\n", math.Pow(2, 3))
	fmt.Printf("最大值: max(10, 20) = %.0f\n", math.Max(10, 20))
	fmt.Printf("最小值: min(10, 20) = %.0f\n", math.Min(10, 20))
}

// ==========================================
// 10. Go与Java的主要区别总结
// ==========================================

func goVsJava() {
	fmt.Println("\n========== Go vs Java 主要区别 ==========")
	fmt.Println("1. 类型声明：Go是 'var name type'，Java是 'type name'")
	fmt.Println("2. 类型推断：Go的 := 可以省略类型")
	fmt.Println("3. 无隐式转换：Go必须显式类型转换")
	fmt.Println("4. 无异常机制：Go使用error返回值（后面会学）")
	fmt.Println("5. 无类：Go使用struct和方法")
	fmt.Println("6. 接口实现：Go是隐式实现，Java是显式implements")
	fmt.Println("7. 并发模型：Go用goroutine和channel，Java用Thread")
	fmt.Println("8. 编译速度：Go编译极快")
	fmt.Println("9. 依赖管理：Go用go.mod，Java用Maven/Gradle")
	fmt.Println("10. 未使用变量：Go编译报错，Java只是警告")
}

// ==========================================
// 主函数 - 程序入口
// ==========================================

func main() {
	fmt.Println("╔════════════════════════════════════╗")
	fmt.Println("║   Go语言 Day 1: 基础入门学习        ║")
	fmt.Println("╚════════════════════════════════════╝")

	// 依次运行各个示例
	helloWorld()
	variableDeclaration()
	basicDataTypes()
	zeroValues()
	typeConversion()
	constants()
	iotaExample()
	formattedOutput()
	mathOperations()
	goVsJava()

	fmt.Println("\n========== 练习题 ==========")
	exercise1()
	exercise2()
	exercise3()
	exercise4()

	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("🎉 恭喜！Day 1 学习完成！")
	fmt.Println(strings.Repeat("=", 40))
}

// ==========================================
// 📝 Day 1 练习题（在下面编写答案）
// ==========================================

/*
练习1：变量声明和基本运算
任务：声明以下变量并进行计算
- 你的姓名（字符串）
- 你的年龄（整数）
- 你的工资（浮点数）
- 计算你一年的总收入
- 输出格式化的结果
*/
func exercise1() {
	var name string = "hjx"
	age := 25
	salary := 18800.00
	yearlyIncome := salary * 12
	fmt.Printf("姓名: %s, 年龄: %d, 工资: %.2f, 年收入: %.2f\n", name, age, salary, yearlyIncome)
}

/*
练习2：常量和iota
任务：定义测试用例的状态常量
- 使用iota定义：Pending(待执行)=1, Running(执行中)=2, Passed(通过)=3, Failed(失败)=4
- 定义测试用例总数常量 = 100
- 输出所有常量的值
*/
func exercise2() {
	const (
		Pending = iota + 1
		Running
		Passed
		Failed
	)
	const TotalCases = 100
	fmt.Printf("测试用例状态定义: Pending=%d, Running=%d, Passed=%d, Failed=%d\n", Pending, Running, Passed, Failed)
	fmt.Printf("测试用例总数: %d\n", TotalCases)
}

/*
练习3：类型转换和计算
任务：测试用例通过率计算
- 总用例数: 150 (int)
- 通过用例数: 127 (int)
- 计算通过率（百分比，保留2位小数）
- 提示：需要转换为float64进行计算
*/
func exercise3() {
	totalCases := 150
	passedCases := 127
	passRate := float64(passedCases) / float64(totalCases) * 100
	fmt.Printf("测试用例通过率: %.2f%%\n", passRate)
}

/*
练习4：综合应用
任务：定义一个测试环境配置
- 使用常量定义：开发环境URL、测试环境URL、生产环境URL
- 使用变量声明：当前环境、是否开启调试模式、最大重试次数
- 使用格式化输出打印所有配置信息
*/
func exercise4() {
	const (
		DevURL     = "http://dev.example.com"
		TestURL    = "http://test.example.com"
		ProductURL = "https://www.example.com"
	)
	currentEnv := "test"
	debugMode := true
	maxRetries := 3

	fmt.Printf("开发环境URL: %s\n", DevURL)
	fmt.Printf("测试环境URL: %s\n", TestURL)
	fmt.Printf("生产环境URL: %s\n", ProductURL)
	fmt.Printf("当前环境: %s,\ndebugMode: %v,\nmaxRetries: %d\n", currentEnv, debugMode, maxRetries)
}

// ==========================================
// 💡 学习提示
// ==========================================
/*
1. 运行程序：在终端执行 `go run day1.go`
2. 编译程序：`go build day1.go` 会生成可执行文件
3. 格式化代码：`go fmt day1.go` 自动格式化
4. 完成练习题后，取消注释并在main函数中调用
5. 遇到问题随时问我！

下一步：
- 完成4个练习题
- 实验不同的变量声明方式
- 尝试修改代码，观察错误信息
- 准备好了就开始 Day 2: 流程控制！
*/
