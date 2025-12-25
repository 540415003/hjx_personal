// ==========================================
// Day 2 练习题参考答案
// 说明：先自己完成练习，遇到困难再参考这个文件
// 运行方式：go run Day2_answers.go
// ==========================================

package main

import (
	"fmt"
	"strings"
	"time"
)

// 练习1：条件判断 - 测试用例优先级判断
func exercise1() {
	fmt.Println("\n========== 练习1：条件判断 ==========")

	// 测试函数
	judgePriority := func(testType string) {
		var priority string
		var executeTime string

		// 方法1: if-else if
		if testType == "smoke" {
			priority = "P0 (最高优先级)"
			executeTime = "每次构建后立即执行"
		} else if testType == "regression" {
			priority = "P1 (高优先级)"
			executeTime = "每日执行"
		} else if testType == "integration" {
			priority = "P2 (中优先级)"
			executeTime = "每周执行"
		} else {
			priority = "未知"
			executeTime = "待定"
		}

		fmt.Printf("测试类型: %s\n", testType)
		fmt.Printf("  优先级: %s\n", priority)
		fmt.Printf("  执行时间: %s\n\n", executeTime)
	}

	// 方法2: switch（更优雅）
	judgePrioritySwitch := func(testType string) {
		fmt.Printf("测试类型: %s\n", testType)

		switch testType {
		case "smoke":
			fmt.Println("  优先级: P0 (最高优先级)")
			fmt.Println("  执行时间: 每次构建后立即执行")
			fmt.Println("  说明: 冒烟测试，验证核心功能")
		case "regression":
			fmt.Println("  优先级: P1 (高优先级)")
			fmt.Println("  执行时间: 每日执行")
			fmt.Println("  说明: 回归测试，确保无功能退化")
		case "integration":
			fmt.Println("  优先级: P2 (中优先级)")
			fmt.Println("  执行时间: 每周执行")
			fmt.Println("  说明: 集成测试，验证模块间协作")
		default:
			fmt.Println("  优先级: 未知")
			fmt.Println("  执行时间: 待定")
		}
		fmt.Println()
	}

	// 测试不同类型
	fmt.Println("方法1: if-else")
	judgePriority("smoke")
	judgePriority("regression")
	judgePriority("integration")

	fmt.Println("方法2: switch (推荐)")
	judgePrioritySwitch("smoke")
	judgePrioritySwitch("regression")
}

// 练习2：循环遍历 - 统计测试结果
func exercise2() {
	fmt.Println("\n========== 练习2：循环遍历 ==========")

	// 测试结果数据
	results := []int{1, 0, 1, 1, 0, 1, 1, 1, 0, 1}

	// 统计变量
	total := len(results)
	passed := 0
	failed := 0
	firstFailIndex := -1

	// 遍历统计
	fmt.Println("开始分析测试结果...")
	for index, result := range results {
		if result == 1 {
			passed++
			fmt.Printf("  [%d] ✓ 通过\n", index+1)
		} else {
			failed++
			fmt.Printf("  [%d] ✗ 失败\n", index+1)

			// 记录第一个失败的索引
			if firstFailIndex == -1 {
				firstFailIndex = index
			}
		}
	}

	// 计算通过率
	passRate := float64(passed) / float64(total) * 100

	// 输出统计结果
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("统计结果:")
	fmt.Printf("  总用例数: %d\n", total)
	fmt.Printf("  通过数: %d\n", passed)
	fmt.Printf("  失败数: %d\n", failed)
	fmt.Printf("  通过率: %.2f%%\n", passRate)

	if firstFailIndex != -1 {
		fmt.Printf("  第一个失败的用例: 索引 %d (第 %d 个用例)\n",
			firstFailIndex, firstFailIndex+1)
	}

	// 额外分析：连续通过的最长序列
	maxStreak := 0
	currentStreak := 0

	for _, result := range results {
		if result == 1 {
			currentStreak++
			if currentStreak > maxStreak {
				maxStreak = currentStreak
			}
		} else {
			currentStreak = 0
		}
	}
	fmt.Printf("  最长连续通过: %d 个用例\n", maxStreak)
}

// 练习3：switch应用 - HTTP状态码处理器
func exercise3() {
	fmt.Println("\n========== 练习3：HTTP状态码处理 ==========")

	// 状态码处理函数
	handleStatusCode := func(code int) {
		fmt.Printf("\nHTTP状态码: %d\n", code)

		// 首先检查特殊的具体状态码
		switch code {
		case 200:
			fmt.Println("  ✓ 200 OK - 请求成功")
			fmt.Println("  含义: 请求已成功处理")
		case 201:
			fmt.Println("  ✓ 201 Created - 资源已创建")
			fmt.Println("  含义: 请求成功并创建了新资源")
		case 404:
			fmt.Println("  ✗ 404 Not Found - 资源未找到")
			fmt.Println("  含义: 请求的资源不存在")
		case 500:
			fmt.Println("  ✗ 500 Internal Server Error - 服务器内部错误")
			fmt.Println("  含义: 服务器遇到错误，无法完成请求")
		default:
			// 按范围判断
			switch {
			case code >= 200 && code < 300:
				fmt.Println("  ✓ 类别: 请求成功 (2xx)")
				fmt.Println("  说明: 请求已被成功接收、理解和处理")
			case code >= 300 && code < 400:
				fmt.Println("  ↻ 类别: 重定向 (3xx)")
				fmt.Println("  说明: 需要客户端采取进一步操作完成请求")
			case code >= 400 && code < 500:
				fmt.Println("  ✗ 类别: 客户端错误 (4xx)")
				fmt.Println("  说明: 请求包含语法错误或无法完成")
			case code >= 500 && code < 600:
				fmt.Println("  ✗ 类别: 服务器错误 (5xx)")
				fmt.Println("  说明: 服务器在处理请求时发生错误")
			default:
				fmt.Println("  ? 未知状态码")
			}
		}
	}

	// 测试各种状态码
	statusCodes := []int{200, 201, 301, 404, 500, 503}

	fmt.Println("测试常见HTTP状态码:")
	for _, code := range statusCodes {
		handleStatusCode(code)
	}
}

// 练习4：defer应用 - 测试执行包装器
func exercise4() {
	fmt.Println("\n========== 练习4：defer应用 ==========")

	// 测试执行函数
	executeTest := func(testName string, duration time.Duration) {
		// 记录开始时间
		startTime := time.Now()

		// 使用defer确保结束时打印统计
		defer func() {
			endTime := time.Now()
			elapsed := endTime.Sub(startTime)

			fmt.Println(strings.Repeat("-", 40))
			fmt.Printf("测试完成: %s\n", testName)
			fmt.Printf("  开始时间: %s\n", startTime.Format("15:04:05.000"))
			fmt.Printf("  结束时间: %s\n", endTime.Format("15:04:05.000"))
			fmt.Printf("  执行耗时: %v\n", elapsed)
			fmt.Println(strings.Repeat("-", 40))
		}()

		// 使用defer捕获可能的错误
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("  ✗ 测试执行出错: %v\n", err)
			}
		}()

		// 执行测试
		fmt.Printf("\n开始执行测试: %s\n", testName)
		fmt.Println("  步骤1: 初始化测试环境...")
		time.Sleep(duration / 3)

		fmt.Println("  步骤2: 执行测试用例...")
		time.Sleep(duration / 3)

		fmt.Println("  步骤3: 清理测试数据...")
		time.Sleep(duration / 3)

		fmt.Println("  ✓ 测试执行成功")
	}

	// 执行多个测试
	fmt.Println("测试套件开始执行:")
	executeTest("用户登录功能测试", 150*time.Millisecond)
	executeTest("数据查询性能测试", 200*time.Millisecond)
	executeTest("API接口测试", 100*time.Millisecond)
}

// 练习5：综合练习 - 测试用例过滤和执行器
func exercise5() {
	fmt.Println("\n========== 练习5：综合练习 ==========")

	// 定义测试用例结构体
	type TestCase struct {
		name    string
		tags    []string
		enabled bool
	}

	// 创建测试用例集
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

	// 测试执行器
	runTestSuite := func(filterTag string, stopOnFailure bool) {
		fmt.Printf("\n执行测试套件 (过滤标签: '%s')\n", filterTag)
		fmt.Println(strings.Repeat("=", 50))

		executed := 0
		passed := 0
		failed := 0
		skipped := 0

		for index, tc := range testCases {
			testNum := index + 1

			// 检查是否启用
			if !tc.enabled {
				fmt.Printf("[%d] ⊘ 跳过: %s (未启用)\n", testNum, tc.name)
				skipped++
				continue
			}

			// 检查标签过滤
			hasTag := false
			for _, tag := range tc.tags {
				if tag == filterTag {
					hasTag = true
					break
				}
			}

			if !hasTag {
				fmt.Printf("[%d] ⊘ 跳过: %s (标签不匹配: %v)\n",
					testNum, tc.name, tc.tags)
				skipped++
				continue
			}

			// 执行测试
			fmt.Printf("[%d] ▶ 执行: %s (标签: %v)\n",
				testNum, tc.name, tc.tags)
			executed++

			// 模拟测试执行（简单随机结果）
			isPass := (index % 3) != 1 // 模拟：大部分通过，部分失败

			if isPass {
				fmt.Printf("    ✓ 通过\n")
				passed++
			} else {
				fmt.Printf("    ✗ 失败\n")
				failed++

				// 如果失败且需要停止
				if stopOnFailure {
					fmt.Println("\n⚠️ 检测到测试失败，停止执行")
					break
				}
			}
		}

		// 输出统计
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println("测试执行摘要:")
		fmt.Printf("  总用例数: %d\n", len(testCases))
		fmt.Printf("  已执行: %d\n", executed)
		fmt.Printf("  通过: %d\n", passed)
		fmt.Printf("  失败: %d\n", failed)
		fmt.Printf("  跳过: %d\n", skipped)

		if executed > 0 {
			passRate := float64(passed) / float64(executed) * 100
			fmt.Printf("  通过率: %.2f%%\n", passRate)
		}
	}

	// 测试不同场景
	runTestSuite("smoke", false)   // 执行所有smoke测试，不停止
	runTestSuite("auth", true)     // 执行auth测试，失败时停止
	runTestSuite("payment", false) // 执行payment测试
}

// 主函数
/*
func main() {
	fmt.Println("╔════════════════════════════════════╗")
	fmt.Println("║   Day 2 练习题参考答案             ║")
	fmt.Println("╚════════════════════════════════════╝")

	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()

	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("🎓 所有练习完成！")
	fmt.Println("💡 建议：对比你的答案，理解不同实现方式")
	fmt.Println(strings.Repeat("=", 40))
}
*/

// ==========================================
// 📚 知识点总结
// ==========================================
/*
通过这5个练习，你应该掌握了：

1. ✅ if条件判断和多分支选择
2. ✅ for循环的多种形式（传统、range、无限）
3. ✅ switch语句的灵活用法（值匹配、范围判断）
4. ✅ defer的执行顺序和实际应用
5. ✅ break和continue控制循环流程

🎯 实际工作应用场景：
- 测试用例优先级管理
- 测试结果统计分析
- HTTP响应状态处理
- 测试执行计时和监控
- 测试套件过滤执行

💡 关键要点：
1. if语句可以包含初始化语句，作用域仅在if块内
2. for...range遍历时，使用_忽略不需要的索引或值
3. switch默认break，需要fall through时显式使用
4. defer按LIFO顺序执行（后进先出）
5. defer常用于资源清理、计时、错误恢复

❓ 思考题：
1. 为什么Go的switch默认不fall through？
   答：更安全，避免忘记break导致的bug

2. defer什么时候执行？
   答：函数返回前，按后进先出顺序执行

3. range遍历时修改元素会影响原数据吗？
   答：value是副本，不影响；要修改原数据需要用索引

准备好了就开始 Day 3 的学习吧！🚀
*/
