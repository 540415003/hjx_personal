// ==========================================
// Day 2: Go语言流程控制
// 主题：if/else、for循环、switch、defer
// ==========================================

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// ==========================================
// 1. if 语句 - 条件判断
// ==========================================

func ifStatement() {
	fmt.Println("\n========== if 语句 ==========")

	// 基本if语句（注意：Go的if不需要括号！）
	age := 28
	if age >= 18 {
		fmt.Println("你是成年人")
	}
	// Java对比: if (age >= 18) { ... }

	// if-else
	score := 85
	if score >= 60 {
		fmt.Println("考试通过")
	} else {
		fmt.Println("考试不通过")
	}

	// if-else if-else
	testResult := 75
	if testResult >= 90 {
		fmt.Println("优秀 ⭐⭐⭐")
	} else if testResult >= 80 {
		fmt.Println("良好 ⭐⭐")
	} else if testResult >= 60 {
		fmt.Println("及格 ⭐")
	} else {
		fmt.Println("不及格")
	}

	// 重要特性：if语句可以包含初始化语句（作用域仅在if块内）
	if num := 42; num > 0 {
		fmt.Printf("num=%d 是正数\n", num)
	}
	// 注意：这里num已经不可用了（超出作用域）
	// fmt.Println(num)  // 编译错误！

	// 实际应用：测试结果判断
	passed := 127
	failed := 23
	total := passed + failed
	if passRate := float64(passed) / float64(total) * 100; passRate >= 95 {
		fmt.Printf("✅ 测试通过率: %.2f%% (优秀)\n", passRate)
	} else if passRate >= 80 {
		fmt.Printf("⚠️  测试通过率: %.2f%% (良好)\n", passRate)
	} else {
		fmt.Printf("❌ 测试通过率: %.2f%% (需改进)\n", passRate)
	}
}

// ==========================================
// 2. for 循环 - Go唯一的循环结构
// ==========================================

func forLoop() {
	fmt.Println("\n========== for 循环 ==========")

	// 形式1: 传统for循环（类似Java）
	fmt.Println("形式1: 传统for")
	for i := 0; i < 5; i++ {
		fmt.Printf("  循环 %d\n", i)
	}
	// Java对比: for (int i = 0; i < 5; i++) { ... }

	// 形式2: 类似while循环（Go没有while关键字）
	fmt.Println("\n形式2: 类似while")
	count := 0
	for count < 3 {
		fmt.Printf("  count = %d\n", count)
		count++
	}
	// Java对比: while (count < 3) { ... }

	// 形式3: 无限循环
	fmt.Println("\n形式3: 无限循环（带break）")
	i := 0
	for {
		if i >= 3 {
			break // 退出循环
		}
		fmt.Printf("  无限循环 %d\n", i)
		i++
	}
	// Java对比: while (true) { ... }

	// 形式4: range遍历（重要！）
	fmt.Println("\n形式4: range遍历")

	// 遍历数组/切片
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("  索引: %d, 值: %d\n", index, value)
	}

	// 只要索引
	for index := range numbers {
		fmt.Printf("  索引: %d\n", index)
	}

	// 只要值（使用_忽略索引）
	for _, value := range numbers {
		fmt.Printf("  值: %d\n", value)
	}

	// 遍历字符串
	str := "Go语言"
	for index, char := range str {
		fmt.Printf("  位置: %d, 字符: %c (Unicode: %U)\n", index, char, char)
	}

	// 遍历Map
	testResults := map[string]int{
		"用例1": 1,
		"用例2": 0,
		"用例3": 1,
	}
	for name, status := range testResults {
		result := "失败"
		if status == 1 {
			result = "通过"
		}
		fmt.Printf("  %s: %s\n", name, result)
	}
}

// ==========================================
// 3. break 和 continue
// ==========================================

func breakContinue() {
	fmt.Println("\n========== break 和 continue ==========")

	// break: 终止循环
	fmt.Println("break示例:")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			fmt.Println("  遇到5，退出循环")
			break
		}
		fmt.Printf("  i = %d\n", i)
	}

	// continue: 跳过本次循环
	fmt.Println("\ncontinue示例:")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			fmt.Println("  跳过3")
			continue
		}
		fmt.Printf("  i = %d\n", i)
	}

	// 实际应用：过滤测试用例
	fmt.Println("\n实际应用：只执行标记为active的测试:")
	testCases := []struct {
		name   string
		active bool
	}{
		{"登录测试", true},
		{"注册测试", false},
		{"支付测试", true},
		{"注销测试", false},
	}

	for _, tc := range testCases {
		if !tc.active {
			fmt.Printf("  ⊘ 跳过: %s (未激活)\n", tc.name)
			continue
		}
		fmt.Printf("  ✓ 执行: %s\n", tc.name)
	}
}

// ==========================================
// 4. 标签和多层循环控制
// ==========================================

func labeledLoop() {
	fmt.Println("\n========== 标签和多层循环 ==========")

	// 使用标签跳出多层循环
	fmt.Println("查找第一个匹配的测试用例:")

OuterLoop: // 定义标签
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("  检查 用例组%d-用例%d\n", i, j)
			if i == 2 && j == 2 {
				fmt.Println("  ✓ 找到匹配项，退出所有循环")
				break OuterLoop // 跳出到标签位置
			}
		}
	}
	fmt.Println("循环结束")

	// continue也可以使用标签
	fmt.Println("\n跳过特定组合:")
Outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				fmt.Printf("  跳过组%d的第2个\n", i)
				continue Outer // 跳到外层循环的下一次迭代
			}
			fmt.Printf("  处理 %d-%d\n", i, j)
		}
	}
}

// ==========================================
// 5. switch 语句 - Go的switch更强大
// ==========================================

func switchStatement() {
	fmt.Println("\n========== switch 语句 ==========")

	// 基本switch（不需要break！）
	day := 3
	fmt.Printf("今天是星期%d，", day)
	switch day {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6, 7: // 多个值
		fmt.Println("周末")
	default:
		fmt.Println("无效的日期")
	}
	// Java对比: 需要break，否则会fall through

	// switch可以有初始化语句
	fmt.Println("\n测试状态:")
	switch status := 2; status {
	case 1:
		fmt.Println("待执行")
	case 2:
		fmt.Println("执行中")
	case 3:
		fmt.Println("已通过")
	case 4:
		fmt.Println("已失败")
	default:
		fmt.Println("未知状态")
	}

	// switch条件表达式（无需变量）
	score := 85
	fmt.Printf("\n分数 %d 的等级: ", score)
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}

	// fallthrough: 强制执行下一个case
	fmt.Println("\nfallthrough示例:")
	num := 1
	switch num {
	case 1:
		fmt.Println("  输出1")
		fallthrough
	case 2:
		fmt.Println("  输出2")
		fallthrough
	case 3:
		fmt.Println("  输出3")
	default:
		fmt.Println("  其他")
	}

	// 类型switch（判断接口类型）
	fmt.Println("\n类型switch:")
	var x interface{} = "hello"
	switch v := x.(type) {
	case int:
		fmt.Printf("  整数: %d\n", v)
	case string:
		fmt.Printf("  字符串: %s\n", v)
	case bool:
		fmt.Printf("  布尔: %v\n", v)
	default:
		fmt.Printf("  未知类型: %T\n", v)
	}
}

// ==========================================
// 6. defer 延迟执行 - Go的特色功能
// ==========================================

func deferBasics() {
	fmt.Println("\n========== defer 基础 ==========")

	// defer会在函数返回前执行
	defer fmt.Println("这是第一个defer（最后执行）")
	defer fmt.Println("这是第二个defer（倒数第二）")
	defer fmt.Println("这是第三个defer（倒数第三）")

	fmt.Println("正常执行的代码")
	fmt.Println("继续执行")

	// 输出顺序：
	// 正常执行的代码
	// 继续执行
	// 这是第三个defer（倒数第三）
	// 这是第二个defer（倒数第二）
	// 这是第一个defer（最后执行）
}

func deferStack() {
	fmt.Println("\n========== defer 栈特性 ==========")

	fmt.Println("defer遵循后进先出(LIFO)原则:")
	for i := 1; i <= 5; i++ {
		defer fmt.Printf("  defer %d\n", i)
	}
	fmt.Println("循环结束")
	// 输出: 循环结束 -> defer 5 -> defer 4 -> defer 3 -> defer 2 -> defer 1
}

func deferWithVariable() {
	fmt.Println("\n========== defer 变量捕获 ==========")

	// defer会立即捕获参数的值
	x := 10
	defer fmt.Printf("defer中的x = %d (捕获时的值)\n", x)

	x = 20
	fmt.Printf("修改后的x = %d\n", x)

	// 使用闭包可以访问最新值
	defer func() {
		fmt.Printf("闭包中的x = %d (最新值)\n", x)
	}()

	x = 30
}

func deferPractical() {
	fmt.Println("\n========== defer 实际应用 ==========")

	// 实际应用1: 资源清理（类似Java的finally）
	fmt.Println("模拟打开和关闭文件:")
	file := "test.txt"
	fmt.Printf("  打开文件: %s\n", file)
	defer fmt.Printf("  关闭文件: %s\n", file)
	fmt.Println("  处理文件内容...")
	// Java对比: try-finally 或 try-with-resources

	// 实际应用2: 测试计时
	fmt.Println("\n测试执行计时:")
	testName := "登录功能测试"
	startTime := time.Now()
	defer func() {
		duration := time.Since(startTime)
		fmt.Printf("  ⏱ %s 耗时: %v\n", testName, duration)
	}()

	fmt.Printf("  开始执行: %s\n", testName)
	time.Sleep(100 * time.Millisecond) // 模拟测试执行
	fmt.Println("  测试执行中...")

	// 实际应用3: 错误恢复（后面会详细学习）
	fmt.Println("\n错误恢复示例:")
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("  捕获到错误: %v\n", err)
		}
	}()
}

// ==========================================
// 7. 综合示例：测试用例执行器
// ==========================================

func testExecutor() {
	fmt.Println("\n========== 综合示例：测试执行器 ==========")

	// 定义测试用例
	type TestCase struct {
		name     string
		enabled  bool
		priority int
	}

	testSuite := []TestCase{
		{"用户登录测试", true, 1},
		{"用户注册测试", true, 2},
		{"密码重置测试", false, 3},
		{"个人信息修改", true, 1},
		{"账户注销测试", false, 2},
		{"权限验证测试", true, 1},
	}

	// 执行测试统计
	executed := 0
	passed := 0
	failed := 0
	skipped := 0

	// 开始执行
	fmt.Println("开始执行测试套件...")
	fmt.Println(strings.Repeat("-", 50))

	for index, tc := range testSuite {
		testNum := index + 1

		// 跳过未启用的测试
		if !tc.enabled {
			fmt.Printf("[%d] ⊘ 跳过: %s (未启用)\n", testNum, tc.name)
			skipped++
			continue
		}

		// 执行测试
		fmt.Printf("[%d] ▶ 执行: %s (优先级: %d)\n", testNum, tc.name, tc.priority)
		executed++

		// 模拟测试执行和结果
		rand.Seed(time.Now().UnixNano() + int64(index))
		result := rand.Intn(10) // 0-9随机数

		// 根据优先级调整通过率（优先级越高，通过率越高）
		threshold := 3
		if tc.priority == 1 {
			threshold = 2 // 优先级1，80%通过率
		}

		if result > threshold {
			fmt.Printf("    ✓ 通过\n")
			passed++
		} else {
			fmt.Printf("    ✗ 失败\n")
			failed++
		}

		// 优先级1的失败用例立即停止
		if tc.priority == 1 && result <= threshold {
			fmt.Println("\n⚠️ 发现高优先级测试失败，停止执行")
			break
		}
	}

	// 输出统计结果
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("测试执行完成！")
	fmt.Printf("总计: %d | 执行: %d | 通过: %d | 失败: %d | 跳过: %d\n",
		len(testSuite), executed, passed, failed, skipped)

	// 计算通过率
	if executed > 0 {
		passRate := float64(passed) / float64(executed) * 100
		fmt.Printf("通过率: %.2f%%\n", passRate)

		// 评估结果
		switch {
		case passRate >= 95:
			fmt.Println("评价: 优秀 ⭐⭐⭐")
		case passRate >= 80:
			fmt.Println("评价: 良好 ⭐⭐")
		case passRate >= 60:
			fmt.Println("评价: 及格 ⭐")
		default:
			fmt.Println("评价: 需要改进")
		}
	}
}

// ==========================================
// 8. Go vs Java 流程控制对比
// ==========================================

func goVsJavaControl() {
	fmt.Println("\n========== Go vs Java 流程控制对比 ==========")
	fmt.Println("1. if语句: Go不需要括号，但必须有花括号")
	fmt.Println("2. for循环: Go只有for，没有while/do-while")
	fmt.Println("3. switch: Go不需要break，默认不会fall through")
	fmt.Println("4. range: Go特有，类似Java的增强for循环")
	fmt.Println("5. defer: Go特有，类似Java的finally但更灵活")
	fmt.Println("6. 标签: 两者都支持，但Go更常用于多层循环")
}

// ==========================================
// 主函数 - 程序入口
// ==========================================

func main() {
	fmt.Println("╔════════════════════════════════════╗")
	fmt.Println("║   Go语言 Day 2: 流程控制学习        	║")
	fmt.Println("╚════════════════════════════════════╝")

	// 依次运行各个示例
	ifStatement()
	forLoop()
	breakContinue()
	labeledLoop()
	switchStatement()
	deferBasics()
	deferStack()
	deferWithVariable()
	deferPractical()
	testExecutor()
	goVsJavaControl()

	// 练习题
	exercise_1()
	exercise_2()
	exercise_3()
	exercise_4()
	exercise_5()

	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("🎉 恭喜！Day 2 学习完成！")
	fmt.Println(strings.Repeat("=", 40))
}

// ==========================================
// 📝 Day 2 练习题（在下面编写答案）
// ==========================================

/*
练习1：条件判断
任务：编写一个测试用例优先级判断函数
- 输入：测试类型（smoke/regression/integration）
- 根据类型判断优先级：
  * smoke -> P0（最高优先级）
  * regression -> P1
  * integration -> P2
- 输出优先级和建议执行时间
*/

func exercise_1() {
	// 在这里编写你的代码
	var priority string
	var executeTime string

	judgePriority := func(testType string) {
		switch testType {
		case "smoke":
			priority = "P0 (最高优先级)"
			executeTime = "每次构建后立即执行"
		case "regression":
			priority = "P1 (高优先级)"
			executeTime = "每日执行"
		case "integration":
			priority = "P2 (中优先级)"
			executeTime = "每周执行"
		default:
			priority = "未知"
			executeTime = "待定"
		}

		fmt.Printf("测试类型: %s\n", testType)
		fmt.Printf("  优先级: %s\n", priority)
		fmt.Printf("  执行时间: %s\n", executeTime)
	}

	judgePriority("smoke")
	judgePriority("regression")
	judgePriority("integration")
}

/*
练习2：循环遍历
任务：统计测试用例执行结果
- 给定一个测试结果切片：[]int{1, 0, 1, 1, 0, 1, 1, 1, 0, 1}
  （1表示通过，0表示失败）
- 使用for循环统计：
  * 总用例数
  * 通过数
  * 失败数
  * 通过率
- 找出第一个失败的用例索引
*/

func exercise_2() {
	// 在这里编写你的代码
	results := []int{1, 0, 1, 1, 0, 1, 1, 1, 0, 1}
	var total int
	var passed int
	var failed int
	var passRate float64
	var firstFailIndex int

	for index, result := range results {
		if result == 1 {
			passed++
		} else {
			failed++
			if firstFailIndex == 0 {
				firstFailIndex = index
			}
		}
		total++
	}

	passRate = float64(passed) / float64(total) * 100

	fmt.Printf("总用例数: %d\n", total)
	fmt.Printf("通过数: %d\n", passed)
	fmt.Printf("失败数: %d\n", failed)
	fmt.Printf("通过率: %.2f%%\n", passRate)
	fmt.Printf("第一个失败的用例: %d\n", firstFailIndex)
}

/*
练习3：switch应用
任务：HTTP状态码处理器
- 输入：HTTP状态码（200, 404, 500等）
- 使用switch判断并输出：
  * 200-299: 请求成功
  * 300-399: 重定向
  * 400-499: 客户端错误
  * 500-599: 服务器错误
- 特殊处理：200, 201, 404, 500单独输出具体信息
*/

func exercise_3() {
	// 在这里编写你的代码
	handleStatusCode := func(code int) {
		switch {
		case code > 200 && code < 300 && code != 201:
			fmt.Println("请求成功(2xx)，请求已被成功接收、理解和处理")
		case code >= 300 && code < 400:
			fmt.Println("重定向(3xx)，需要客户端采取进一步操作完成请求")
		case code >= 400 && code < 500 && code != 404:
			fmt.Println("客户端错误(4xx)，请求包含语法错误或无法完成")
		case code > 500 && code < 600:
			fmt.Println("服务器错误(5xx)，服务器在处理请求时发生错误")
		case code == 200:
			fmt.Println("请求成功(200)，请求已成功处理")
		case code == 201:
			fmt.Println("资源已创建(201)，请求成功并创建了新资源")
		case code == 404:
			fmt.Println("资源未找到(404)，请求的资源不存在")
		case code == 500:
			fmt.Println("服务器内部错误(500)，服务器遇到错误，无法完成请求")
		default:
			fmt.Println("未知状态码")
		}
	}
	handleStatusCode(200)
	handleStatusCode(201)
	handleStatusCode(301)
	handleStatusCode(404)
	handleStatusCode(500)
	handleStatusCode(503)
}

/*
练习4：defer应用
任务：测试执行包装器
- 创建一个函数，模拟测试用例执行
- 使用defer实现：
  * 记录测试开始和结束时间
  * 计算测试执行耗时
  * 确保测试结束后打印统计信息（无论是否出错）
- 模拟一个测试执行过程（可以用time.Sleep）
*/

func exercise_4() {
	// 在这里编写你的代码
	executeTest := func(testName string, duration time.Duration) {
		startTime := time.Now()

		defer func() {
			endTime := time.Now()
			elapsed := endTime.Sub(startTime)

			fmt.Printf("开始执行测试: %s\n", testName)
			time.Sleep(2 * time.Second)
			fmt.Printf("测试完成: %s\n", testName)
			fmt.Printf("  开始时间: %s\n", startTime.Format("15:04:05.000"))
			fmt.Printf("  结束时间: %s\n", endTime.Format("15:04:05.000"))
			fmt.Printf("  执行耗时: %v\n", elapsed)
			time.Sleep(time.Second)
		}()
	}

	executeTest("用户登录功能测试", 150*time.Millisecond)
	executeTest("数据查询性能测试", 200*time.Millisecond)
	executeTest("API接口测试", 100*time.Millisecond)
}

/*
练习5：综合练习
任务：测试用例过滤和执行器
- 定义测试用例结构体：名称、标签(tags)、是否启用
- 创建至少5个测试用例
- 实现功能：
  * 只执行包含特定标签的用例（如"smoke"）
  * 跳过未启用的用例
  * 统计执行结果
  * 如果遇到失败，询问是否继续（用布尔变量模拟）
*/

func exercise_5() {
	// 在这里编写你的代码
	type TestCase struct {
		name    string
		tags    []string
		enabled bool
	}

	testCases := []TestCase{
		{"用户登录测试", []string{"smoke", "auth"}, true},
		{"用户注册测试", []string{"smoke", "auth"}, true},
		{"密码修改测试", []string{"auth", "security"}, false},
		{"商品搜索测试", []string{"smoke", "search"}, true},
		{"订单创建测试", []string{"order", "payment"}, true},
		{"支付流程测试", []string{"smoke", "payment"}, true},
		{"数据导出测试", []string{"data", "report"}, false},
		{"权限验证测试", []string{"smoke", "security"}, true},
	}

	runTestCase := func(tag string, stopOnFailure bool) {
		var executed int
		var passed int
		var failed int
		var skipped int

		for index, tc := range testCases {
			testNum := index + 1

			if tag == "smoke" && tc.enabled {
				fmt.Printf("[%d]  ✓ 执行: %s (标签: %v)\n", testNum, tc.name, tc.tags)
				executed++
			} else if tag == "smoke" && !tc.enabled {
				fmt.Printf("[%d]  ⊘ 跳过: %s (未启用)\n", testNum, tc.name)
				skipped++
				continue
			} else if tag != "smoke" && tc.enabled {
				fmt.Printf("[%d]  ✓ 执行: %s (标签: %v)\n", testNum, tc.name, tc.tags)
				executed++
			} else if tag != "smoke" && !tc.enabled {
				fmt.Printf("[%d]  ⊘ 跳过: %s (未启用)\n", testNum, tc.name)
				skipped++
				continue
			}

			isPass := (index % 3) != 1
			if isPass {
				fmt.Printf("[%d]  ✓ 通过\n", testNum)
				passed++
			} else {
				fmt.Printf("[%d]  ✗ 失败\n", testNum)
				failed++
				if stopOnFailure {
					fmt.Println("\n⚠️ 检测到测试失败，停止执行")
					break
				}
			}
		}

		fmt.Printf("总用例数: %d\n", len(testCases))
		fmt.Printf("已执行: %d\n", executed)
		fmt.Printf("通过: %d\n", passed)
		fmt.Printf("失败: %d\n", failed)
		fmt.Printf("跳过: %d\n", skipped)

		if executed > 0 {
			passRate := float64(passed) / float64(executed) * 100
			fmt.Printf("通过率: %.2f%%\n", passRate)
		}
	}

	runTestCase("smoke", false)
	fmt.Println(strings.Repeat("-", 30))
	runTestCase("auth", true)
	fmt.Println(strings.Repeat("-", 30))
	runTestCase("payment", false)
}

// ==========================================
// 💡 学习提示
// ==========================================
/*
1. 运行程序：在终端执行 `go run Day2.go`
2. Go的if不需要括号，但花括号必须写
3. for是Go唯一的循环，可以模拟while和do-while
4. switch默认自动break，无需手动添加
5. defer按后进先出(LIFO)顺序执行
6. range非常实用，记得用_忽略不需要的变量

下一步：
- 完成5个练习题
- 特别注意defer的执行顺序
- 多实践range的各种用法
- 准备好了就开始 Day 3: 复杂数据类型！
*/
