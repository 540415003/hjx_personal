// ==========================================
// Day 1 练习题参考答案
// 说明：先自己完成练习，遇到困难再参考这个文件
// 运行方式：go run day1_answers.go
// ==========================================

package main

import (
	"fmt"
)

// 练习1：变量声明和基本运算
func exercise1_answers() {
	fmt.Println("\n========== 练习1：变量声明和基本运算 ==========")

	// 声明个人信息变量
	myName := "李晓明"         // 使用短声明
	var myAge int = 26      // 使用var声明
	var mySalary = 12000.00 // 类型推断

	// 计算年收入（月薪 * 12）
	yearlyIncome := mySalary * 12

	// 格式化输出
	fmt.Printf("姓名：%s\n", myName)
	fmt.Printf("年龄：%d岁\n", myAge)
	fmt.Printf("月薪：￥%.2f\n", mySalary)
	fmt.Printf("年收入：￥%.2f\n", yearlyIncome)

	// 额外计算：税后收入（假设税率20%）
	taxRate := 0.2
	afterTaxIncome := yearlyIncome * (1 - taxRate)
	fmt.Printf("税后年收入：￥%.2f (税率%.0f%%)\n", afterTaxIncome, taxRate*100)
}

// 练习2：常量和iota
func exercise2_answers() {
	fmt.Println("\n========== 练习2：常量和iota ==========")

	// 使用iota定义测试状态
	const (
		Pending = iota + 1 // 从1开始：1
		Running            // 2
		Passed             // 3
		Failed             // 4
	)

	// 定义测试用例总数
	const TotalCases = 100

	// 输出所有常量
	fmt.Println("测试用例状态定义：")
	fmt.Printf("  Pending (待执行) = %d\n", Pending)
	fmt.Printf("  Running (执行中) = %d\n", Running)
	fmt.Printf("  Passed (通过)   = %d\n", Passed)
	fmt.Printf("  Failed (失败)   = %d\n", Failed)
	fmt.Printf("\n测试用例总数：%d\n", TotalCases)

	// 实际应用示例
	currentStatus := Running
	fmt.Printf("\n当前测试状态：%d (执行中)\n", currentStatus)
}

// 练习3：类型转换和计算
func exercise3_answers() {
	fmt.Println("\n========== 练习3：类型转换和计算 ==========")

	// 测试数据
	totalCases := 150
	passedCases := 127

	// 关键点：整数除法会丢失小数，必须转换为float64
	passRate := float64(passedCases) / float64(totalCases) * 100

	// 输出结果
	fmt.Printf("测试用例统计：\n")
	fmt.Printf("  总用例数：%d\n", totalCases)
	fmt.Printf("  通过数：%d\n", passedCases)
	fmt.Printf("  失败数：%d\n", totalCases-passedCases)
	fmt.Printf("  通过率：%.2f%%\n", passRate)

	// 错误示范（如果取消注释会得到错误结果）
	// wrongRate := passedCases / totalCases * 100  // 结果是0，因为整数除法！
	// fmt.Printf("错误的通过率：%d%%\n", wrongRate)

	// 判断测试质量
	fmt.Printf("\n测试评价：")
	if passRate >= 95 {
		fmt.Println("优秀 ⭐⭐⭐")
	} else if passRate >= 80 {
		fmt.Println("良好 ⭐⭐")
	} else {
		fmt.Println("需要改进 ⭐")
	}
}

// 练习4：综合应用
func exercise4_answers() {
	fmt.Println("\n========== 练习4：综合应用 ==========")

	// 使用常量定义环境URL
	const (
		DevURL     = "http://dev.example.com"
		TestURL    = "http://test.example.com"
		ProductURL = "https://www.example.com"
	)

	// 使用变量定义配置
	currentEnv := "test" // 当前环境
	debugMode := true    // 调试模式
	maxRetries := 3      // 最大重试次数
	timeout := 30.0      // 超时时间（秒）
	apiVersion := "v2"   // API版本

	// 格式化输出配置信息
	fmt.Println("╔═══════════════════════════════════════╗")
	fmt.Println("║      测试环境配置信息                   ║")
	fmt.Println("╚═══════════════════════════════════════╝")

	fmt.Println("\n【环境URL配置】")
	fmt.Printf("  开发环境：%s\n", DevURL)
	fmt.Printf("  测试环境：%s\n", TestURL)
	fmt.Printf("  生产环境：%s\n", ProductURL)

	fmt.Println("\n【当前运行配置】")
	fmt.Printf("  当前环境：%s\n", currentEnv)
	fmt.Printf("  调试模式：%v\n", debugMode)
	fmt.Printf("  最大重试：%d 次\n", maxRetries)
	fmt.Printf("  超时时间：%.1f 秒\n", timeout)
	fmt.Printf("  API版本：%s\n", apiVersion)

	// 根据环境选择URL
	var targetURL string
	switch currentEnv {
	case "dev":
		targetURL = DevURL
	case "test":
		targetURL = TestURL
	case "prod":
		targetURL = ProductURL
	default:
		targetURL = TestURL // 默认测试环境
	}

	fmt.Printf("\n✅ 将连接到：%s\n", targetURL)
}

// 主函数 - 运行所有练习
/*
func main() {
	fmt.Println("╔════════════════════════════════════╗")
	fmt.Println("║   Day 1 练习题参考答案             ║")
	fmt.Println("╚════════════════════════════════════╝")

	exercise1()
	exercise2()
	exercise3()
	exercise4()

	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("🎓 所有练习完成！")
	fmt.Println("💡 建议：对比你的答案和参考答案，理解不同的实现方式")
	fmt.Println(strings.Repeat("=", 40))
}
*/

// ==========================================
// 📚 知识点总结
// ==========================================
/*
通过这4个练习，你应该掌握了：

1. ✅ 变量的三种声明方式：var、var with type、:=
2. ✅ 基本数据类型：int、float64、string、bool
3. ✅ 常量声明和iota的使用
4. ✅ 类型转换：必须显式转换
5. ✅ 格式化输出：Printf的各种占位符
6. ✅ 基本数学运算

🎯 实际工作应用场景：
- 测试配置管理
- 测试结果统计
- 测试状态定义
- 环境切换
- 数据计算和报告

❓ 思考题：
1. 为什么Go要求显式类型转换？
   答：为了类型安全，避免隐式转换带来的潜在bug

2. := 和 var 的区别是什么？什么时候用哪个？
   答：:= 只能在函数内使用，更简洁；var 可以用在包级别，更明确

3. const 和 var 的主要区别？
   答：const 编译时确定，不可修改；var 运行时可以修改

准备好了就开始 Day 2 的学习吧！🚀
*/
