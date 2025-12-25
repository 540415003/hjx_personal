// ==========================================
// Day 3: Go语言复杂数据类型
// 主题：数组、切片、Map、结构体、指针
// ==========================================

package main

import (
	"fmt"
	"sort"
	"strings"
)

// ==========================================
// 1. 数组 Array - 固定长度的序列
// ==========================================

func arrayBasics() {
	fmt.Println("\n========== 数组基础 ==========")

	// 声明数组的多种方式
	var arr1 [5]int // 声明长度为5的int数组，默认零值
	fmt.Printf("零值数组: %v\n", arr1)

	// 声明并初始化
	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("初始化数组: %v\n", arr2)

	// 短声明
	arr3 := [3]string{"Go", "Java", "Python"}
	fmt.Printf("字符串数组: %v\n", arr3)

	// 让编译器计算长度
	arr4 := [...]int{10, 20, 30, 40}
	fmt.Printf("自动长度: %v (长度: %d)\n", arr4, len(arr4))

	// 指定索引初始化
	arr5 := [5]int{0: 100, 2: 200, 4: 300}
	fmt.Printf("指定索引: %v\n", arr5)

	// 访问和修改元素
	fmt.Println("\n访问和修改:")
	scores := [3]int{85, 90, 78}
	fmt.Printf("原数组: %v\n", scores)
	scores[1] = 95 // 修改索引1的元素
	fmt.Printf("修改后: %v\n", scores)
	fmt.Printf("第一个元素: %d\n", scores[0])
	fmt.Printf("数组长度: %d\n", len(scores))

	// 遍历数组
	fmt.Println("\n遍历数组:")
	for i := 0; i < len(scores); i++ {
		fmt.Printf("  索引%d: %d\n", i, scores[i])
	}

	// 使用range遍历
	for index, value := range scores {
		fmt.Printf("  range索引%d: %d\n", index, value)
	}

	// 数组是值类型（重要！）
	fmt.Println("\n数组是值类型:")
	a := [3]int{1, 2, 3}
	b := a // 复制整个数组
	b[0] = 100
	fmt.Printf("a: %v\n", a) // a不变
	fmt.Printf("b: %v\n", b) // b改变
	// Java对比: Java的数组是引用类型
}

func arrayLimitations() {
	fmt.Println("\n========== 数组的局限性 ==========")

	// 数组长度是类型的一部分
	var arr1 [3]int
	var arr2 [5]int
	// arr1 = arr2  // 编译错误！不同长度的数组是不同类型

	fmt.Printf("arr1类型: %T\n", arr1)
	fmt.Printf("arr2类型: %T\n", arr2)

	// 数组长度固定，不能动态增长
	// 实际开发中很少直接使用数组，而是使用切片（Slice）
	fmt.Println("由于这些局限，Go推荐使用切片而不是数组")
}

// ==========================================
// 2. 切片 Slice - 动态数组（重点！）
// ==========================================

func sliceBasics() {
	fmt.Println("\n========== 切片基础 ==========")

	// 切片声明（不指定长度）
	var s1 []int // nil切片
	fmt.Printf("nil切片: %v, len=%d, cap=%d, is nil: %v\n",
		s1, len(s1), cap(s1), s1 == nil)

	// 使用字面量创建
	s2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("字面量切片: %v, len=%d, cap=%d\n",
		s2, len(s2), cap(s2))

	// 使用make创建（推荐）
	s3 := make([]int, 5)     // 长度5，容量5
	s4 := make([]int, 3, 10) // 长度3，容量10
	fmt.Printf("make创建s3: %v, len=%d, cap=%d\n",
		s3, len(s3), cap(s3))
	fmt.Printf("make创建s4: %v, len=%d, cap=%d\n",
		s4, len(s4), cap(s4))

	// 从数组创建切片
	arr := [5]int{10, 20, 30, 40, 50}
	s5 := arr[1:4] // 索引1到3（不包括4）
	fmt.Printf("从数组切片: %v\n", s5)

	// 切片的切片
	s6 := s2[1:4]
	fmt.Printf("切片的切片: %v\n", s6)

	// 省略索引
	fmt.Println("\n切片操作:")
	nums := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("原切片: %v\n", nums)
	fmt.Printf("nums[:3] = %v\n", nums[:3])   // 从头到索引2
	fmt.Printf("nums[2:] = %v\n", nums[2:])   // 从索引2到尾
	fmt.Printf("nums[:] = %v\n", nums[:])     // 全部
	fmt.Printf("nums[1:4] = %v\n", nums[1:4]) // 索引1到3
}

func sliceOperations() {
	fmt.Println("\n========== 切片操作 ==========")

	// append: 追加元素（重要！）
	s := []int{1, 2, 3}
	fmt.Printf("原切片: %v, len=%d, cap=%d\n", s, len(s), cap(s))

	s = append(s, 4) // 追加一个元素
	fmt.Printf("追加4: %v, len=%d, cap=%d\n", s, len(s), cap(s))

	s = append(s, 5, 6, 7) // 追加多个元素
	fmt.Printf("追加5,6,7: %v, len=%d, cap=%d\n", s, len(s), cap(s))

	// 追加切片
	s2 := []int{8, 9, 10}
	s = append(s, s2...) // 注意...语法
	fmt.Printf("追加切片: %v, len=%d, cap=%d\n", s, len(s), cap(s))

	// copy: 复制切片
	fmt.Println("\ncopy操作:")
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 3)
	n := copy(dst, src) // 复制src到dst
	fmt.Printf("源切片: %v\n", src)
	fmt.Printf("目标切片: %v\n", dst)
	fmt.Printf("复制了%d个元素\n", n)

	// 删除元素（Go没有内置删除函数）
	fmt.Println("\n删除元素:")
	nums := []int{10, 20, 30, 40, 50}
	fmt.Printf("原切片: %v\n", nums)

	// 删除索引2的元素
	index := 2
	nums = append(nums[:index], nums[index+1:]...)
	fmt.Printf("删除索引2: %v\n", nums)

	// 插入元素
	fmt.Println("\n插入元素:")
	nums = []int{10, 20, 40, 50}
	fmt.Printf("原切片: %v\n", nums)

	// 在索引2插入30
	index = 2
	value := 30
	nums = append(nums[:index], append([]int{value}, nums[index:]...)...)
	fmt.Printf("插入30: %v\n", nums)
}

func sliceMemory() {
	fmt.Println("\n========== 切片的底层原理 ==========")

	// 切片是引用类型（指向底层数组）
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1 // s2和s1共享底层数组
	s2[0] = 100
	fmt.Printf("s1: %v\n", s1) // s1也被修改
	fmt.Printf("s2: %v\n", s2)
	fmt.Println("切片是引用类型，修改会影响共享底层数组的其他切片")

	// 容量和扩容
	fmt.Println("\n容量和扩容:")
	s := make([]int, 0, 3)
	fmt.Printf("初始: len=%d, cap=%d\n", len(s), cap(s))

	for i := 1; i <= 6; i++ {
		s = append(s, i)
		fmt.Printf("追加%d后: len=%d, cap=%d, %v\n",
			i, len(s), cap(s), s)
	}
	fmt.Println("当容量不足时，Go会自动扩容（通常翻倍）")
}

// ==========================================
// 3. Map - 键值对集合（哈希表）
// ==========================================

func mapBasics() {
	fmt.Println("\n========== Map基础 ==========")

	// 声明map
	var m1 map[string]int // nil map，不能直接使用
	fmt.Printf("nil map: %v, is nil: %v\n", m1, m1 == nil)

	// 使用make创建（推荐）
	m2 := make(map[string]int)
	fmt.Printf("make创建: %v\n", m2)

	// 使用字面量创建
	m3 := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 28,
	}
	fmt.Printf("字面量创建: %v\n", m3)

	// 添加和修改元素
	fmt.Println("\n添加和修改:")
	scores := make(map[string]int)
	scores["数学"] = 95
	scores["英语"] = 88
	scores["语文"] = 92
	fmt.Printf("成绩: %v\n", scores)

	scores["数学"] = 98 // 修改已存在的key
	fmt.Printf("修改后: %v\n", scores)

	// 访问元素
	fmt.Println("\n访问元素:")
	fmt.Printf("数学成绩: %d\n", scores["数学"])
	fmt.Printf("物理成绩: %d (不存在返回零值)\n", scores["物理"])

	// 检查key是否存在（重要！）
	if value, exists := scores["英语"]; exists {
		fmt.Printf("英语成绩存在: %d\n", value)
	} else {
		fmt.Println("英语成绩不存在")
	}

	// 获取长度
	fmt.Printf("map长度: %d\n", len(scores))
}

func mapOperations() {
	fmt.Println("\n========== Map操作 ==========")

	testResults := map[string]string{
		"test_login":    "passed",
		"test_register": "passed",
		"test_payment":  "failed",
		"test_search":   "passed",
	}

	// 遍历map
	fmt.Println("遍历map:")
	for key, value := range testResults {
		symbol := "✓"
		if value == "failed" {
			symbol = "✗"
		}
		fmt.Printf("  %s %s: %s\n", symbol, key, value)
	}

	// 只遍历key
	fmt.Println("\n只遍历key:")
	for key := range testResults {
		fmt.Printf("  - %s\n", key)
	}

	// 删除元素
	fmt.Println("\n删除元素:")
	fmt.Printf("删除前: %v\n", testResults)
	delete(testResults, "test_payment")
	fmt.Printf("删除后: %v\n", testResults)

	// 删除不存在的key不会报错
	delete(testResults, "not_exist")
	fmt.Println("删除不存在的key是安全的")

	// map是无序的
	fmt.Println("\n注意: map是无序的，遍历顺序随机")
}

func mapAdvanced() {
	fmt.Println("\n========== Map高级用法 ==========")

	// 嵌套map
	fmt.Println("嵌套map:")
	testSuites := map[string]map[string]int{
		"UI测试": {
			"total":  10,
			"passed": 9,
			"failed": 1,
		},
		"API测试": {
			"total":  20,
			"passed": 18,
			"failed": 2,
		},
	}

	for suite, results := range testSuites {
		fmt.Printf("%s:\n", suite)
		fmt.Printf("  总数: %d, 通过: %d, 失败: %d\n",
			results["total"], results["passed"], results["failed"])
	}

	// map的key必须是可比较类型
	// 可以: string, int, float, bool, pointer, struct(字段都可比较)
	// 不可以: slice, map, function
	fmt.Println("\nmap的key必须是可比较类型（不能是slice、map、function）")

	// 统计单词出现次数（实际应用）
	fmt.Println("\n实际应用 - 统计测试状态:")
	statuses := []string{"passed", "failed", "passed", "passed", "skipped", "failed", "passed"}
	counter := make(map[string]int)

	for _, status := range statuses {
		// 统计每个状态出现的次数(status是map的key；counter[status]是map的value，初始值为0)
		counter[status]++
	}

	fmt.Println("统计结果:")
	for status, count := range counter {
		fmt.Printf("  %s: %d次\n", status, count)
	}
}

// ==========================================
// 4. 结构体 Struct - Go的"类"
// ==========================================

func structBasics() {
	fmt.Println("\n========== 结构体基础 ==========")

	// 定义结构体
	type TestCase struct {
		ID       int
		Name     string
		Priority string
		Status   string
		Duration float64
	}

	// 创建结构体实例 - 方式1
	var tc1 TestCase
	fmt.Printf("零值结构体: %+v\n", tc1)

	// 方式2: 字面量初始化（按顺序）
	tc2 := TestCase{1, "登录测试", "P0", "passed", 1.5}
	fmt.Printf("按顺序初始化: %+v\n", tc2)

	// 方式3: 字面量初始化（指定字段名，推荐）
	tc3 := TestCase{
		ID:       2,
		Name:     "注册测试",
		Priority: "P1",
		Status:   "passed",
		Duration: 2.3,
	}
	fmt.Printf("指定字段名: %+v\n", tc3)

	// 方式4: 部分初始化
	tc4 := TestCase{
		ID:   3,
		Name: "支付测试",
	}
	fmt.Printf("部分初始化: %+v\n", tc4)

	// 访问和修改字段
	fmt.Println("\n访问和修改字段:")
	tc := TestCase{
		ID:       4,
		Name:     "搜索测试",
		Priority: "P2",
		Status:   "pending",
	}
	fmt.Printf("原始: %+v\n", tc)

	tc.Status = "running"
	tc.Duration = 3.5
	fmt.Printf("修改后: %+v\n", tc)
	fmt.Printf("用例名称: %s\n", tc.Name)
}

func structMethods() {
	fmt.Println("\n========== 结构体方法 ==========")

	// 定义结构体
	type TestCase struct {
		ID       int
		Name     string
		Priority string
		Status   string
		Duration float64
	}

	// 定义方法（值接收者）
	type TestCaseWithMethods struct {
		ID       int
		Name     string
		Priority string
		Status   string
		Duration float64
	}

	// 注意：方法定义在这个函数外部（见下面的示例）
	// 这里演示如何使用
	fmt.Println("方法示例将在外部定义的结构体中展示")
}

func structComparison() {
	fmt.Println("\n========== 结构体与Java类对比 ==========")

	fmt.Println("Go结构体 vs Java类:")
	fmt.Println("1. Go没有类，使用struct + 方法")
	fmt.Println("2. Go没有构造函数，使用工厂函数")
	fmt.Println("3. Go没有继承，使用组合（嵌入）")
	fmt.Println("4. Go的方法在类型外部定义")
	fmt.Println("5. Go通过大小写控制可见性（大写公开，小写私有）")
}

// ==========================================
// 5. 结构体实际应用示例
// ==========================================

// 定义测试用例结构体
type TestCase struct {
	ID       int
	Name     string
	Priority string
	Status   string
	Duration float64
	Tags     []string
}

// 定义方法 - 值接收者
func (tc TestCase) Display() {
	symbol := "○"
	switch tc.Status {
	case "passed":
		symbol = "✓"
	case "failed":
		symbol = "✗"
	case "running":
		symbol = "▶"
	}
	fmt.Printf("[%d] %s %s (%s) - %.2fs\n",
		tc.ID, symbol, tc.Name, tc.Priority, tc.Duration)
}

// 定义方法 - 指针接收者
func (tc *TestCase) Run() {
	tc.Status = "running"
	fmt.Printf("执行测试: %s\n", tc.Name)
	// 模拟测试执行...
	tc.Status = "passed"
	tc.Duration = 1.5
}

func (tc *TestCase) AddTag(tag string) {
	tc.Tags = append(tc.Tags, tag)
}

func structExample() {
	fmt.Println("\n========== 结构体实际应用 ==========")

	// 创建测试用例
	tc := TestCase{
		ID:       1,
		Name:     "用户登录测试",
		Priority: "P0",
		Status:   "pending",
		Tags:     []string{"smoke", "auth"},
	}

	// 调用方法
	tc.Display()
	tc.Run()
	tc.Display()

	tc.AddTag("critical")
	fmt.Printf("标签: %v\n", tc.Tags)
}

// ==========================================
// 6. 结构体嵌入（组合）
// ==========================================

// 基础结构体
type BaseTest struct {
	ID       int
	Name     string
	Status   string
	Duration float64
}

// 嵌入结构体
type APITest struct {
	BaseTest // 嵌入（匿名字段）
	URL      string
	Method   string
	Expected int
}

func structEmbedding() {
	fmt.Println("\n========== 结构体嵌入（组合） ==========")

	api := APITest{
		BaseTest: BaseTest{
			ID:     1,
			Name:   "API登录测试",
			Status: "pending",
		},
		URL:      "https://api.example.com/login",
		Method:   "POST",
		Expected: 200,
	}

	// 可以直接访问嵌入结构体的字段
	fmt.Printf("ID: %d\n", api.ID)         // 直接访问
	fmt.Printf("Name: %s\n", api.Name)     // 直接访问
	fmt.Printf("URL: %s\n", api.URL)       // 自己的字段
	fmt.Printf("Method: %s\n", api.Method) // 自己的字段

	// 也可以显式访问
	fmt.Printf("Status: %s\n", api.BaseTest.Status)

	fmt.Println("\n这就是Go的'继承'方式 - 组合而非继承")
}

// ==========================================
// 7. 指针基础
// ==========================================

func pointerBasics() {
	fmt.Println("\n========== 指针基础 ==========")

	// 声明指针
	var p *int // int类型的指针
	fmt.Printf("nil指针: %v\n", p)

	// 获取变量地址
	x := 42
	p = &x // &取地址符
	fmt.Printf("x的值: %d\n", x)
	fmt.Printf("x的地址: %p\n", &x)
	fmt.Printf("p的值（地址）: %p\n", p)
	fmt.Printf("p指向的值: %d\n", *p) // *解引用

	// 通过指针修改值
	*p = 100
	fmt.Printf("修改后x的值: %d\n", x)

	// 结构体指针
	fmt.Println("\n结构体指针:")
	tc := TestCase{
		ID:   1,
		Name: "测试用例",
	}

	ptr := &tc
	// Go可以直接通过指针访问字段，无需->
	ptr.Name = "修改后的名称"
	// 等价于: (*ptr).Name = "修改后的名称"

	fmt.Printf("通过指针修改: %+v\n", tc)

	// new函数：分配内存并返回指针
	fmt.Println("\nnew函数:")
	p2 := new(int) // 分配int的零值，返回指针
	fmt.Printf("new(int): %v, 值: %d\n", p2, *p2)

	// 函数参数：值传递 vs 指针传递
	fmt.Println("\n值传递 vs 指针传递:")
	num := 10
	fmt.Printf("原始值: %d\n", num)

	// 值传递（不会修改原值）
	modifyByValue := func(n int) {
		n = 20
	}
	modifyByValue(num)
	fmt.Printf("值传递后: %d (未改变)\n", num)

	// 指针传递（会修改原值）
	modifyByPointer := func(n *int) {
		*n = 30
	}
	modifyByPointer(&num)
	fmt.Printf("指针传递后: %d (已改变)\n", num)
}

// ==========================================
// 8. 综合示例：测试报告系统
// ==========================================

// 测试套件结构体
type TestSuite struct {
	Name      string
	TestCases []TestCase
	Summary   TestSummary
}

// 测试摘要结构体
type TestSummary struct {
	Total    int
	Passed   int
	Failed   int
	Skipped  int
	Duration float64
}

// 计算摘要
func (ts *TestSuite) CalculateSummary() {
	ts.Summary = TestSummary{}
	for _, tc := range ts.TestCases {
		ts.Summary.Total++
		ts.Summary.Duration += tc.Duration

		switch tc.Status {
		case "passed":
			ts.Summary.Passed++
		case "failed":
			ts.Summary.Failed++
		case "skipped":
			ts.Summary.Skipped++
		}
	}
}

// 显示报告
func (ts TestSuite) DisplayReport() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("测试套件: %s\n", ts.Name)
	fmt.Println(strings.Repeat("=", 60))

	// 显示所有测试用例
	for _, tc := range ts.TestCases {
		tc.Display()
	}

	// 显示摘要
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("总计: %d | 通过: %d | 失败: %d | 跳过: %d\n",
		ts.Summary.Total, ts.Summary.Passed,
		ts.Summary.Failed, ts.Summary.Skipped)
	fmt.Printf("总耗时: %.2fs\n", ts.Summary.Duration)

	if ts.Summary.Total > 0 {
		passRate := float64(ts.Summary.Passed) / float64(ts.Summary.Total) * 100
		fmt.Printf("通过率: %.2f%%\n", passRate)
	}
	fmt.Println(strings.Repeat("=", 60))
}

func comprehensiveExample() {
	fmt.Println("\n========== 综合示例：测试报告系统 ==========")

	// 创建测试套件
	suite := TestSuite{
		Name: "用户模块测试",
		TestCases: []TestCase{
			{
				ID:       1,
				Name:     "用户登录",
				Priority: "P0",
				Status:   "passed",
				Duration: 1.2,
				Tags:     []string{"smoke", "auth"},
			},
			{
				ID:       2,
				Name:     "用户注册",
				Priority: "P1",
				Status:   "passed",
				Duration: 1.8,
				Tags:     []string{"smoke", "auth"},
			},
			{
				ID:       3,
				Name:     "密码修改",
				Priority: "P2",
				Status:   "failed",
				Duration: 2.1,
				Tags:     []string{"auth", "security"},
			},
			{
				ID:       4,
				Name:     "个人信息更新",
				Priority: "P1",
				Status:   "passed",
				Duration: 1.5,
				Tags:     []string{"profile"},
			},
			{
				ID:       5,
				Name:     "账户注销",
				Priority: "P2",
				Status:   "skipped",
				Duration: 0,
				Tags:     []string{"cleanup"},
			},
		},
	}

	// 计算摘要
	suite.CalculateSummary()

	// 显示报告
	suite.DisplayReport()

	// 按标签过滤
	fmt.Println("\n过滤标签'smoke'的测试用例:")
	smokeCases := filterByTag(suite.TestCases, "smoke")
	for _, tc := range smokeCases {
		tc.Display()
	}
}

// 辅助函数：按标签过滤
func filterByTag(cases []TestCase, tag string) []TestCase {
	result := make([]TestCase, 0)
	for _, tc := range cases {
		for _, t := range tc.Tags {
			if t == tag {
				result = append(result, tc)
				break
			}
		}
	}
	return result
}

// ==========================================
// 9. 排序和比较
// ==========================================

func sortingExample() {
	fmt.Println("\n========== 排序示例 ==========")

	// 排序基本类型
	nums := []int{5, 2, 8, 1, 9, 3}
	fmt.Printf("原切片: %v\n", nums)
	sort.Ints(nums)
	fmt.Printf("排序后: %v\n", nums)

	strs := []string{"banana", "apple", "cherry", "date"}
	fmt.Printf("原字符串: %v\n", strs)
	sort.Strings(strs)
	fmt.Printf("排序后: %v\n", strs)

	// 自定义排序
	testCases := []TestCase{
		{ID: 3, Name: "测试C", Priority: "P0", Duration: 2.5},
		{ID: 1, Name: "测试A", Priority: "P1", Duration: 1.2},
		{ID: 2, Name: "测试B", Priority: "P2", Duration: 3.1},
	}

	// 按ID排序
	sort.Slice(testCases, func(i, j int) bool {
		return testCases[i].ID < testCases[j].ID
	})
	fmt.Println("\n按ID排序:")
	for _, tc := range testCases {
		fmt.Printf("  ID:%d, Name:%s\n", tc.ID, tc.Name)
	}

	// 按Duration排序
	sort.Slice(testCases, func(i, j int) bool {
		return testCases[i].Duration < testCases[j].Duration
	})
	fmt.Println("\n按Duration排序:")
	for _, tc := range testCases {
		fmt.Printf("  %s: %.2fs\n", tc.Name, tc.Duration)
	}
}

// ==========================================
// 主函数 - 程序入口
// ==========================================

func main() {
	fmt.Println("╔════════════════════════════════════╗")
	fmt.Println("║   Go语言 Day 3: 复杂数据类型       ║")
	fmt.Println("╚════════════════════════════════════╝")

	// 依次运行各个示例
	arrayBasics()
	arrayLimitations()
	sliceBasics()
	sliceOperations()
	sliceMemory()
	mapBasics()
	mapOperations()
	mapAdvanced()
	structBasics()
	structMethods()
	structComparison()
	structExample()
	structEmbedding()
	pointerBasics()
	comprehensiveExample()
	sortingExample()

	fmt.Println("========== 练习题 ==========")
	exercise_1()
	exercise_2()
	exercise_3()
	exercise_4()
	exercise_5()

	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("🎉 恭喜！Day 3 学习完成！")
	fmt.Println(strings.Repeat("=", 40))
}

// ==========================================
// 📝 Day 3 练习题（在下面编写答案）
// ==========================================

/*
练习1：切片操作
任务：实现测试结果管理器
- 创建一个测试结果切片（passed, failed, skipped等状态）
- 实现功能：
  * 添加测试结果
  * 统计各状态数量
  * 计算通过率
  * 找出所有失败的索引
*/

func exercise_1() {
	// 在这里编写你的代码
	type ResultManager struct {
		results []string
	}

	addResult := func(rm *ResultManager, result string) {
		rm.results = append(rm.results, result)
	}

	countStatus := func(rm *ResultManager) map[string]int {
		counter := make(map[string]int)
		for _, result := range rm.results {
			counter[result]++
		}
		return counter
	}

	calculatePassRate := func(rm *ResultManager) float64 {
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

	findFailedIndexes := func(rm *ResultManager) []int {
		indexes := make([]int, 0)
		for index, result := range rm.results {
			if result == "failed" {
				indexes = append(indexes, index)
			}
		}
		return indexes
	}

	resultManager := ResultManager{
		results: []string{},
	}

	testResaults := []string{
		"passed", "passed", "failed", "passed", "skipped", "passed", "failed", "passed", "passed", "failed",
	}

	fmt.Println("添加测试结果:")
	for index, result := range testResaults {
		addResult(&resultManager, result)
		fmt.Printf("[%d] %s\n", index, result)
	}

	fmt.Println("统计各状态数量:")
	statusCount := countStatus(&resultManager)
	for status, count := range statusCount {
		fmt.Printf("%s: %d\n", status, count)
	}

	fmt.Println("计算通过率:")
	passRate := calculatePassRate(&resultManager)
	fmt.Printf("通过率: %.2f%%\n", passRate)

	fmt.Println("找出所有失败的索引:")
	failedIndexes := findFailedIndexes(&resultManager)
	fmt.Printf("失败的测试索引: %v\n", failedIndexes)
}

/*
练习2：Map应用
任务：测试环境配置管理器
- 创建一个map存储不同环境的配置：
  * key: 环境名称（dev, test, prod）
  * value: 另一个map，包含（url, timeout, retry等配置）
- 实现功能：
  * 添加环境配置
  * 查询配置
  * 更新配置
  * 列出所有环境
*/

func exercise_2() {
	// 在这里编写你的代码
	type ConfigManager struct {
		configs map[string]map[string]interface{}
	}

	createManager := func() ConfigManager {
		return ConfigManager{
			configs: make(map[string]map[string]interface{}),
		}
	}

	addConfigs := func(cm *ConfigManager, env string, config map[string]interface{}) {
		cm.configs[env] = config
	}

	getConfigs := func(cm ConfigManager, env string) (map[string]interface{}, bool) {
		config, exists := cm.configs[env]
		return config, exists
	}

	updateConfigs := func(cm *ConfigManager, env, key string, value interface{}) bool {
		if config, exists := cm.configs[env]; exists {
			config[key] = value
			return true
		}
		return false
	}

	listConfigs := func(cm ConfigManager) []string {
		envs := make([]string, 0, len(cm.configs))
		for env := range cm.configs {
			envs = append(envs, env)
		}
		return envs
	}

	displayConfigs := func(env string, config map[string]interface{}) {
		fmt.Printf("\n环境: %s\n", env)
		fmt.Println(strings.Repeat("-", 40))
		for key, value := range config {
			fmt.Printf("  %s: %v\n", key, value)
		}
	}

	manager := createManager()

	// 添加配置
	addConfigs(&manager, "dev", map[string]interface{}{
		"url":     "http://dev.example.com",
		"timeout": 30,
	})

	addConfigs(&manager, "test", map[string]interface{}{
		"url":     "http://test.example.com",
		"timeout": 60,
	})

	addConfigs(&manager, "prod", map[string]interface{}{
		"url":     "https://www.example.com",
		"timeout": 120,
	})

	// 更新配置
	updateConfigs(&manager, "test", "timeout", 90)
	if config, exists := getConfigs(manager, "test"); exists {
		displayConfigs("test", config)
	}

	// 列出所有环境
	fmt.Printf("所有环境: \n")
	for _, env := range listConfigs(manager) {
		fmt.Printf("  %s\n", env)
	}
}

/*
练习3：结构体设计
任务：设计一个测试用例管理系统
- 定义TestCase结构体，包含：
  * 基本信息（ID、名称、描述）
  * 测试属性（优先级、类型、标签）
  * 执行信息（状态、耗时、错误信息）
- 定义方法：
  * Display(): 显示用例信息
  * Execute(): 模拟执行测试
  * IsPass(): 判断是否通过
- 创建至少3个测试用例实例并调用方法
*/

func exercise_3() {
	// 在这里编写你的代码
}

/*
练习4：综合应用
任务：实现一个简单的测试报告生成器
- 创建TestSuite结构体（包含多个TestCase）
- 实现功能：
  * 添加测试用例
  * 执行所有测试（修改状态和耗时）
  * 生成统计报告（总数、通过、失败、跳过、通过率）
  * 按优先级分组显示
- 要求使用切片、map、结构体、指针
*/

func exercise_4() {
	// 在这里编写你的代码
}

/*
练习5：数据处理
任务：测试数据分析工具
- 给定一个测试结果切片（包含多次测试的历史数据）
- 实现功能：
  * 按日期分组统计
  * 找出通过率最高和最低的日期
  * 计算平均通过率
  * 找出最常失败的测试用例
- 使用map存储统计结果，使用结构体表示测试记录
*/

func exercise_5() {
	// 在这里编写你的代码
}

// ==========================================
// 💡 学习提示
// ==========================================
/*
1. 运行程序：go run Day3.go
2. 切片是引用类型，修改会影响底层数组
3. map的遍历顺序是随机的
4. 结构体方法：值接收者 vs 指针接收者
5. Go通过组合而非继承实现代码复用

重点掌握：
- 切片的append、copy、切片操作
- map的创建、添加、删除、检查key
- 结构体的定义、实例化、方法
- 指针的使用场景

下一步：
- 完成5个练习题
- 特别注意切片的底层原理
- 理解值接收者和指针接收者的区别
- 准备好了就开始 Day 4: 函数！
*/
