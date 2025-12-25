// ==========================================
// Day 3 练习题参考答案
// 说明：先自己完成练习，遇到困难再参考这个文件
// 运行方式：go run Day3_answers.go
// ==========================================

package main

import (
	"fmt"
	"sort"
	"strings"
)

// 练习1：切片操作 - 测试结果管理器
func exercise1() {
	fmt.Println("\n========== 练习1：切片操作 ==========")

	// 测试结果管理器
	type ResultManager struct {
		results []string
	}

	// 添加测试结果
	addResult := func(rm *ResultManager, result string) {
		rm.results = append(rm.results, result)
	}

	// 统计各状态数量
	countStatus := func(rm ResultManager) map[string]int {
		counter := make(map[string]int)
		for _, result := range rm.results {
			counter[result]++
		}
		return counter
	}

	// 计算通过率
	calculatePassRate := func(rm ResultManager) float64 {
		if len(rm.results) == 0 {
			return 0
		}
		passed := 0
		for _, result := range rm.results {
			if result == "passed" {
				passed++
			}
		}
		return float64(passed) / float64(len(rm.results)) * 100
	}

	// 找出所有失败的索引
	findFailedIndexes := func(rm ResultManager) []int {
		indexes := make([]int, 0)
		for i, result := range rm.results {
			if result == "failed" {
				indexes = append(indexes, i)
			}
		}
		return indexes
	}

	// 使用示例
	manager := ResultManager{
		results: []string{},
	}

	// 添加测试结果
	testResults := []string{
		"passed", "passed", "failed", "passed", "skipped",
		"passed", "failed", "passed", "passed", "failed",
	}

	fmt.Println("添加测试结果:")
	for i, result := range testResults {
		addResult(&manager, result)
		fmt.Printf("  [%d] %s\n", i, result)
	}

	// 统计
	fmt.Println("\n统计各状态数量:")
	statusCount := countStatus(manager)
	for status, count := range statusCount {
		fmt.Printf("  %s: %d\n", status, count)
	}

	// 通过率
	passRate := calculatePassRate(manager)
	fmt.Printf("\n通过率: %.2f%%\n", passRate)

	// 失败索引
	failedIndexes := findFailedIndexes(manager)
	fmt.Printf("\n失败的测试索引: %v\n", failedIndexes)

	// 额外功能：删除失败的测试记录
	fmt.Println("\n删除失败记录:")
	filtered := make([]string, 0)
	for _, result := range manager.results {
		if result != "failed" {
			filtered = append(filtered, result)
		}
	}
	fmt.Printf("过滤后: %v\n", filtered)
	fmt.Printf("剩余: %d个\n", len(filtered))
}

// 练习2：Map应用 - 测试环境配置管理器
func exercise2() {
	fmt.Println("\n========== 练习2：Map应用 ==========")

	// 环境配置管理器
	type ConfigManager struct {
		configs map[string]map[string]interface{}
	}

	// 创建管理器
	createManager := func() ConfigManager {
		return ConfigManager{
			configs: make(map[string]map[string]interface{}),
		}
	}

	// 添加环境配置
	addConfig := func(cm *ConfigManager, env string, config map[string]interface{}) {
		cm.configs[env] = config
	}

	// 查询配置
	getConfig := func(cm ConfigManager, env string) (map[string]interface{}, bool) {
		config, exists := cm.configs[env]
		return config, exists
	}

	// 更新配置
	updateConfig := func(cm *ConfigManager, env, key string, value interface{}) bool {
		if config, exists := cm.configs[env]; exists {
			config[key] = value
			return true
		}
		return false
	}

	// 列出所有环境
	listEnvironments := func(cm ConfigManager) []string {
		envs := make([]string, 0, len(cm.configs))
		for env := range cm.configs {
			envs = append(envs, env)
		}
		sort.Strings(envs) // 排序以便有序显示
		return envs
	}

	// 显示配置
	displayConfig := func(env string, config map[string]interface{}) {
		fmt.Printf("\n环境: %s\n", env)
		fmt.Println(strings.Repeat("-", 40))
		for key, value := range config {
			fmt.Printf("  %s: %v\n", key, value)
		}
	}

	// 使用示例
	manager := createManager()

	// 添加配置
	fmt.Println("添加环境配置:")
	addConfig(&manager, "dev", map[string]interface{}{
		"url":     "http://dev.example.com",
		"timeout": 30,
		"retry":   3,
		"debug":   true,
	})

	addConfig(&manager, "test", map[string]interface{}{
		"url":     "http://test.example.com",
		"timeout": 60,
		"retry":   5,
		"debug":   false,
	})

	addConfig(&manager, "prod", map[string]interface{}{
		"url":     "https://www.example.com",
		"timeout": 120,
		"retry":   10,
		"debug":   false,
	})

	// 列出所有环境
	fmt.Println("\n所有环境:")
	envs := listEnvironments(manager)
	for i, env := range envs {
		fmt.Printf("  %d. %s\n", i+1, env)
	}

	// 显示配置
	for _, env := range envs {
		if config, exists := getConfig(manager, env); exists {
			displayConfig(env, config)
		}
	}

	// 更新配置
	fmt.Println("\n更新test环境的timeout:")
	if updateConfig(&manager, "test", "timeout", 90) {
		fmt.Println("  更新成功")
		if config, exists := getConfig(manager, "test"); exists {
			displayConfig("test", config)
		}
	}

	// 查询不存在的环境
	fmt.Println("\n查询不存在的环境:")
	if _, exists := getConfig(manager, "staging"); !exists {
		fmt.Println("  staging环境不存在")
	}
}

// 练习3：结构体设计 - 测试用例管理系统
func exercise3() {
	fmt.Println("\n========== 练习3：结构体设计 ==========")

	// 定义TestCase结构体
	type TestCase struct {
		// 基本信息
		ID          int
		Name        string
		Description string

		// 测试属性
		Priority string
		Type     string
		Tags     []string

		// 执行信息
		Status       string
		Duration     float64
		ErrorMessage string
	}

	// Display方法：显示用例信息
	display := func(tc TestCase) {
		symbol := "○"
		switch tc.Status {
		case "passed":
			symbol = "✓"
		case "failed":
			symbol = "✗"
		case "running":
			symbol = "▶"
		case "skipped":
			symbol = "⊘"
		}

		fmt.Printf("\n%s [%d] %s\n", symbol, tc.ID, tc.Name)
		fmt.Printf("  描述: %s\n", tc.Description)
		fmt.Printf("  优先级: %s | 类型: %s | 标签: %v\n",
			tc.Priority, tc.Type, tc.Tags)
		fmt.Printf("  状态: %s | 耗时: %.2fs\n", tc.Status, tc.Duration)
		if tc.ErrorMessage != "" {
			fmt.Printf("  错误: %s\n", tc.ErrorMessage)
		}
	}

	// Execute方法：模拟执行测试
	execute := func(tc *TestCase) {
		fmt.Printf("\n执行测试: %s\n", tc.Name)
		tc.Status = "running"

		// 模拟测试执行（简单随机）
		tc.Duration = 1.5 + float64(tc.ID)*0.3

		// 模拟测试结果（P0优先级更容易通过）
		if tc.Priority == "P0" || tc.ID%3 != 0 {
			tc.Status = "passed"
			tc.ErrorMessage = ""
			fmt.Println("  ✓ 测试通过")
		} else {
			tc.Status = "failed"
			tc.ErrorMessage = "断言失败: 期望值与实际值不匹配"
			fmt.Println("  ✗ 测试失败")
		}
	}

	// IsPass方法：判断是否通过
	isPass := func(tc TestCase) bool {
		return tc.Status == "passed"
	}

	// 创建测试用例
	testCases := []TestCase{
		{
			ID:          1,
			Name:        "用户登录功能测试",
			Description: "验证用户可以使用正确的用户名和密码登录系统",
			Priority:    "P0",
			Type:        "smoke",
			Tags:        []string{"auth", "smoke", "critical"},
			Status:      "pending",
		},
		{
			ID:          2,
			Name:        "用户注册功能测试",
			Description: "验证新用户可以成功注册账号",
			Priority:    "P1",
			Type:        "functional",
			Tags:        []string{"auth", "registration"},
			Status:      "pending",
		},
		{
			ID:          3,
			Name:        "密码强度验证测试",
			Description: "验证系统对密码强度的校验规则",
			Priority:    "P2",
			Type:        "functional",
			Tags:        []string{"auth", "security"},
			Status:      "pending",
		},
	}

	// 执行所有测试
	fmt.Println("开始执行测试套件:")
	fmt.Println(strings.Repeat("=", 60))

	passed := 0
	failed := 0

	for i := range testCases {
		execute(&testCases[i])
		display(testCases[i])

		if isPass(testCases[i]) {
			passed++
		} else {
			failed++
		}
	}

	// 显示统计
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("执行完成 | 总计: %d | 通过: %d | 失败: %d\n",
		len(testCases), passed, failed)

	if len(testCases) > 0 {
		passRate := float64(passed) / float64(len(testCases)) * 100
		fmt.Printf("通过率: %.2f%%\n", passRate)
	}
}

// 练习4：综合应用 - 测试报告生成器
func exercise4() {
	fmt.Println("\n========== 练习4：综合应用 ==========")

	// 测试用例结构体
	type TestCase struct {
		ID       int
		Name     string
		Priority string
		Status   string
		Duration float64
	}

	// 测试套件结构体
	type TestSuite struct {
		Name      string
		TestCases []TestCase
	}

	// 统计结构体
	type Statistics struct {
		Total    int
		Passed   int
		Failed   int
		Skipped  int
		Duration float64
		PassRate float64
	}

	// 添加测试用例
	addTestCase := func(ts *TestSuite, tc TestCase) {
		ts.TestCases = append(ts.TestCases, tc)
	}

	// 执行所有测试
	executeAll := func(ts *TestSuite) {
		fmt.Printf("\n执行测试套件: %s\n", ts.Name)
		fmt.Println(strings.Repeat("-", 60))

		for i := range ts.TestCases {
			tc := &ts.TestCases[i]
			fmt.Printf("[%d] 执行: %s (%s)\n", tc.ID, tc.Name, tc.Priority)

			// 模拟执行
			tc.Duration = 1.0 + float64(i)*0.5

			// 模拟结果（高优先级更容易通过）
			if tc.Priority == "P0" || i%3 != 1 {
				tc.Status = "passed"
				fmt.Println("    ✓ 通过")
			} else {
				tc.Status = "failed"
				fmt.Println("    ✗ 失败")
			}
		}
	}

	// 生成统计报告
	generateReport := func(ts TestSuite) Statistics {
		stats := Statistics{}

		for _, tc := range ts.TestCases {
			stats.Total++
			stats.Duration += tc.Duration

			switch tc.Status {
			case "passed":
				stats.Passed++
			case "failed":
				stats.Failed++
			case "skipped":
				stats.Skipped++
			}
		}

		if stats.Total > 0 {
			stats.PassRate = float64(stats.Passed) / float64(stats.Total) * 100
		}

		return stats
	}

	// 按优先级分组显示
	displayByPriority := func(ts TestSuite) {
		fmt.Println("\n按优先级分组:")
		fmt.Println(strings.Repeat("=", 60))

		// 使用map分组
		groups := make(map[string][]TestCase)
		for _, tc := range ts.TestCases {
			groups[tc.Priority] = append(groups[tc.Priority], tc)
		}

		// 按优先级顺序显示
		priorities := []string{"P0", "P1", "P2"}
		for _, priority := range priorities {
			if cases, exists := groups[priority]; exists {
				fmt.Printf("\n%s (高优先级测试):\n", priority)
				for _, tc := range cases {
					symbol := "✓"
					if tc.Status != "passed" {
						symbol = "✗"
					}
					fmt.Printf("  %s [%d] %s - %.2fs\n",
						symbol, tc.ID, tc.Name, tc.Duration)
				}
			}
		}
	}

	// 创建测试套件
	suite := TestSuite{
		Name: "完整功能测试套件",
	}

	// 添加测试用例
	testCases := []TestCase{
		{ID: 1, Name: "用户登录", Priority: "P0", Status: "pending"},
		{ID: 2, Name: "用户注册", Priority: "P1", Status: "pending"},
		{ID: 3, Name: "密码重置", Priority: "P2", Status: "pending"},
		{ID: 4, Name: "个人资料", Priority: "P1", Status: "pending"},
		{ID: 5, Name: "权限验证", Priority: "P0", Status: "pending"},
		{ID: 6, Name: "数据导出", Priority: "P2", Status: "pending"},
	}

	for _, tc := range testCases {
		addTestCase(&suite, tc)
	}

	// 执行测试
	executeAll(&suite)

	// 生成报告
	stats := generateReport(suite)

	// 显示总体统计
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("测试统计报告:")
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("总计: %d\n", stats.Total)
	fmt.Printf("通过: %d\n", stats.Passed)
	fmt.Printf("失败: %d\n", stats.Failed)
	fmt.Printf("跳过: %d\n", stats.Skipped)
	fmt.Printf("总耗时: %.2fs\n", stats.Duration)
	fmt.Printf("通过率: %.2f%%\n", stats.PassRate)

	// 按优先级分组
	displayByPriority(suite)

	fmt.Println(strings.Repeat("=", 60))
}

// 练习5：数据处理 - 测试数据分析工具
func exercise5() {
	fmt.Println("\n========== 练习5：数据处理 ==========")

	// 测试记录结构体
	type TestRecord struct {
		Date     string
		TestName string
		Status   string
	}

	// 每日统计
	type DailyStats struct {
		Date     string
		Total    int
		Passed   int
		Failed   int
		PassRate float64
	}

	// 测试历史数据
	history := []TestRecord{
		{"2024-01-01", "登录测试", "passed"},
		{"2024-01-01", "注册测试", "passed"},
		{"2024-01-01", "支付测试", "failed"},
		{"2024-01-02", "登录测试", "passed"},
		{"2024-01-02", "注册测试", "failed"},
		{"2024-01-02", "支付测试", "failed"},
		{"2024-01-03", "登录测试", "passed"},
		{"2024-01-03", "注册测试", "passed"},
		{"2024-01-03", "支付测试", "passed"},
		{"2024-01-03", "搜索测试", "passed"},
	}

	// 按日期分组统计
	fmt.Println("按日期分组统计:")
	fmt.Println(strings.Repeat("-", 60))

	dailyMap := make(map[string]*DailyStats)

	for _, record := range history {
		if _, exists := dailyMap[record.Date]; !exists {
			dailyMap[record.Date] = &DailyStats{
				Date: record.Date,
			}
		}

		stats := dailyMap[record.Date]
		stats.Total++

		if record.Status == "passed" {
			stats.Passed++
		} else {
			stats.Failed++
		}
	}

	// 计算通过率
	dailySlice := make([]DailyStats, 0, len(dailyMap))
	for _, stats := range dailyMap {
		if stats.Total > 0 {
			stats.PassRate = float64(stats.Passed) / float64(stats.Total) * 100
		}
		dailySlice = append(dailySlice, *stats)
	}

	// 排序日期
	sort.Slice(dailySlice, func(i, j int) bool {
		return dailySlice[i].Date < dailySlice[j].Date
	})

	// 显示每日统计
	for _, stats := range dailySlice {
		fmt.Printf("%s: 总计=%d, 通过=%d, 失败=%d, 通过率=%.2f%%\n",
			stats.Date, stats.Total, stats.Passed, stats.Failed, stats.PassRate)
	}

	// 找出通过率最高和最低的日期
	fmt.Println(strings.Repeat("-", 60))
	if len(dailySlice) > 0 {
		highest := dailySlice[0]
		lowest := dailySlice[0]

		for _, stats := range dailySlice {
			if stats.PassRate > highest.PassRate {
				highest = stats
			}
			if stats.PassRate < lowest.PassRate {
				lowest = stats
			}
		}

		fmt.Printf("通过率最高: %s (%.2f%%)\n", highest.Date, highest.PassRate)
		fmt.Printf("通过率最低: %s (%.2f%%)\n", lowest.Date, lowest.PassRate)
	}

	// 计算平均通过率
	totalPassRate := 0.0
	for _, stats := range dailySlice {
		totalPassRate += stats.PassRate
	}
	avgPassRate := totalPassRate / float64(len(dailySlice))
	fmt.Printf("平均通过率: %.2f%%\n", avgPassRate)

	// 找出最常失败的测试用例
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println("测试用例失败统计:")

	failureCount := make(map[string]int)
	for _, record := range history {
		if record.Status == "failed" {
			failureCount[record.TestName]++
		}
	}

	// 转换为切片并排序
	type FailureStats struct {
		TestName string
		Count    int
	}

	failures := make([]FailureStats, 0, len(failureCount))
	for name, count := range failureCount {
		failures = append(failures, FailureStats{name, count})
	}

	sort.Slice(failures, func(i, j int) bool {
		return failures[i].Count > failures[j].Count
	})

	// 显示失败次数最多的测试
	for i, f := range failures {
		fmt.Printf("  %d. %s: 失败%d次\n", i+1, f.TestName, f.Count)
	}

	if len(failures) > 0 {
		fmt.Printf("\n最常失败的测试: %s (失败%d次)\n",
			failures[0].TestName, failures[0].Count)
	}
}

// 主函数
/*
func main() {
	fmt.Println("╔════════════════════════════════════╗")
	fmt.Println("║   Day 3 练习题参考答案             ║")
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

1. ✅ 切片的创建、append、copy、切片操作
2. ✅ Map的创建、添加、删除、遍历、嵌套
3. ✅ 结构体的定义、实例化、方法
4. ✅ 指针接收者的使用场景
5. ✅ 综合应用：切片+Map+结构体

🎯 实际工作应用场景：
- 测试结果收集和统计
- 多环境配置管理
- 测试用例数据建模
- 测试报告生成
- 历史数据分析

💡 关键要点：
1. 切片是引用类型，传递时会共享底层数组
2. Map的key必须是可比较类型
3. 结构体方法：需要修改接收者用指针，否则用值
4. 使用make创建切片和map
5. 切片容量不足时会自动扩容

❓ 思考题：
1. 切片和数组的区别？
   答：数组长度固定且是值类型；切片长度可变且是引用类型

2. 什么时候用值接收者，什么时候用指针接收者？
   答：需要修改接收者或接收者很大时用指针；否则用值

3. Map是线程安全的吗？
   答：不是，并发访问需要加锁或使用sync.Map

准备好了就开始 Day 4 的学习吧！🚀
*/
