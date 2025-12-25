# 📝 Day 2 快速参考卡片

> Go语言流程控制速查表

## if 条件语句

```go
// 基本形式（无需括号）
if age >= 18 {
    fmt.Println("成年人")
}

// if-else
if score >= 60 {
    fmt.Println("及格")
} else {
    fmt.Println("不及格")
}

// if-else if-else
if score >= 90 {
    fmt.Println("优秀")
} else if score >= 80 {
    fmt.Println("良好")
} else if score >= 60 {
    fmt.Println("及格")
} else {
    fmt.Println("不及格")
}

// if带初始化语句（作用域仅在if块内）
if num := computeValue(); num > 0 {
    fmt.Println(num)  // num可用
}
// num不可用

// 实际应用
if passRate := float64(passed)/float64(total)*100; passRate >= 95 {
    fmt.Printf("通过率: %.2f%%\n", passRate)
}
```

### if vs Java

| 特性 | Go | Java |
|------|----|----|
| 条件括号 | 不需要 | 必须有 |
| 花括号 | 必须有 | 单语句可省略 |
| 初始化语句 | 支持 | 不支持 |

## for 循环

```go
// 形式1: 传统for循环
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// 形式2: 类似while（Go没有while关键字）
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}

// 形式3: 无限循环
for {
    // 需要break退出
    if condition {
        break
    }
}

// 形式4: range遍历（最常用）
// 遍历切片/数组
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("索引:%d, 值:%d\n", index, value)
}

// 只要索引
for i := range numbers {
    fmt.Println(i)
}

// 只要值（忽略索引）
for _, value := range numbers {
    fmt.Println(value)
}

// 遍历map
m := map[string]int{"a": 1, "b": 2}
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}

// 遍历字符串（按rune）
for index, char := range "Go语言" {
    fmt.Printf("位置:%d, 字符:%c\n", index, char)
}
```

### Go只有for循环

| Java循环 | Go等价写法 |
|---------|-----------|
| `for(;;)` | `for` |
| `while(condition)` | `for condition` |
| `do-while` | 用for+条件判断模拟 |
| `for-each` | `for range` |

## break 和 continue

```go
// break: 退出循环
for i := 1; i <= 10; i++ {
    if i == 5 {
        break  // 退出循环
    }
    fmt.Println(i)
}

// continue: 跳过本次循环
for i := 1; i <= 10; i++ {
    if i%2 == 0 {
        continue  // 跳过偶数
    }
    fmt.Println(i)
}

// 标签：跳出多层循环
OuterLoop:
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        if i == 2 && j == 2 {
            break OuterLoop  // 跳出外层循环
        }
        fmt.Printf("%d-%d\n", i, j)
    }
}

// 标签continue
Outer:
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        if j == 2 {
            continue Outer  // 跳到外层下一次迭代
        }
        fmt.Printf("%d-%d\n", i, j)
    }
}
```

## switch 语句

```go
// 基本switch（自动break，无需手动添加）
switch day {
case 1:
    fmt.Println("星期一")
case 2:
    fmt.Println("星期二")
case 6, 7:  // 多个值
    fmt.Println("周末")
default:
    fmt.Println("其他")
}

// switch带初始化
switch status := getStatus(); status {
case 1:
    fmt.Println("运行中")
case 2:
    fmt.Println("已完成")
}

// 无表达式switch（相当于if-else链）
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
case score >= 60:
    fmt.Println("C")
default:
    fmt.Println("F")
}

// fallthrough: 强制执行下一个case
switch num {
case 1:
    fmt.Println("1")
    fallthrough
case 2:
    fmt.Println("2")  // 会执行
case 3:
    fmt.Println("3")
}

// 类型switch
var x interface{} = "hello"
switch v := x.(type) {
case int:
    fmt.Printf("整数: %d\n", v)
case string:
    fmt.Printf("字符串: %s\n", v)
case bool:
    fmt.Printf("布尔: %v\n", v)
default:
    fmt.Printf("未知: %T\n", v)
}
```

### switch vs Java

| 特性 | Go | Java |
|------|----|----|
| break | 自动break | 需要手动break |
| 多值 | `case 1, 2, 3:` | 不支持 |
| 表达式 | 可选 | 必须有 |
| 类型判断 | 支持 | 不支持 |

## defer 延迟执行

```go
// 基本用法：函数返回前执行
func example() {
    defer fmt.Println("最后执行")
    fmt.Println("先执行")
}
// 输出：先执行 -> 最后执行

// 多个defer：后进先出(LIFO)
func multiple() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
    fmt.Println("主体")
}
// 输出：主体 -> 3 -> 2 -> 1

// defer参数立即求值
func valueCapture() {
    x := 10
    defer fmt.Println(x)  // 捕获10
    x = 20
}
// 输出：10

// defer闭包：访问最新值
func closureValue() {
    x := 10
    defer func() {
        fmt.Println(x)  // 访问最新值
    }()
    x = 20
}
// 输出：20

// 实际应用1: 资源清理
func fileOperation() {
    file := openFile("test.txt")
    defer file.Close()  // 确保文件被关闭
    
    // 处理文件...
}

// 实际应用2: 执行计时
func testTiming() {
    start := time.Now()
    defer func() {
        fmt.Printf("耗时: %v\n", time.Since(start))
    }()
    
    // 执行测试...
}

// 实际应用3: 错误恢复
func errorRecover() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Printf("捕获错误: %v\n", err)
        }
    }()
    
    // 可能panic的代码...
}

// 实际应用4: 解锁
func mutexUnlock() {
    mu.Lock()
    defer mu.Unlock()  // 确保解锁
    
    // 临界区代码...
}
```

### defer执行时机

```
函数开始
  ↓
defer语句1（注册）
  ↓
defer语句2（注册）
  ↓
正常代码执行
  ↓
函数返回前
  ↓
执行defer2（后进先出）
  ↓
执行defer1
  ↓
函数返回
```

## 常见模式

### 模式1: 测试结果统计

```go
results := []int{1, 0, 1, 1, 0, 1}
passed := 0
failed := 0

for _, result := range results {
    if result == 1 {
        passed++
    } else {
        failed++
    }
}

passRate := float64(passed) / float64(len(results)) * 100
fmt.Printf("通过率: %.2f%%\n", passRate)
```

### 模式2: 条件过滤

```go
testCases := []TestCase{...}

for _, tc := range testCases {
    // 跳过未启用
    if !tc.enabled {
        continue
    }
    
    // 执行测试
    runTest(tc)
}
```

### 模式3: 查找第一个匹配

```go
items := []int{10, 20, 30, 40, 50}
target := 30
found := false
index := -1

for i, v := range items {
    if v == target {
        found = true
        index = i
        break
    }
}

if found {
    fmt.Printf("找到: 索引=%d\n", index)
}
```

### 模式4: HTTP状态码处理

```go
switch {
case code >= 200 && code < 300:
    fmt.Println("成功")
case code >= 400 && code < 500:
    fmt.Println("客户端错误")
case code >= 500:
    fmt.Println("服务器错误")
}
```

### 模式5: 测试执行包装

```go
func runTest(name string) {
    start := time.Now()
    defer func() {
        duration := time.Since(start)
        fmt.Printf("%s 耗时: %v\n", name, duration)
    }()
    
    // 执行测试
    executeTestCase(name)
}
```

## 测试场景应用

### 场景1: 测试套件执行

```go
for i, test := range testSuite {
    if !test.enabled {
        fmt.Printf("跳过: %s\n", test.name)
        continue
    }
    
    fmt.Printf("执行: %s\n", test.name)
    result := runTest(test)
    
    if !result && test.priority == "P0" {
        fmt.Println("高优先级失败，停止执行")
        break
    }
}
```

### 场景2: 重试机制

```go
maxRetries := 3
for i := 0; i < maxRetries; i++ {
    if success := tryRequest(); success {
        break
    }
    fmt.Printf("重试 %d/%d\n", i+1, maxRetries)
    time.Sleep(time.Second)
}
```

### 场景3: 标签过滤

```go
filterTag := "smoke"

for _, tc := range testCases {
    hasTag := false
    for _, tag := range tc.tags {
        if tag == filterTag {
            hasTag = true
            break
        }
    }
    
    if !hasTag {
        continue
    }
    
    runTest(tc)
}
```

## 常见错误

### 错误1: if条件加括号

```go
// ❌ 错误（可以编译，但不是Go风格）
if (x > 0) {
    fmt.Println(x)
}

// ✅ 正确
if x > 0 {
    fmt.Println(x)
}
```

### 错误2: 缺少花括号

```go
// ❌ 错误
if x > 0
    fmt.Println(x)

// ✅ 正确
if x > 0 {
    fmt.Println(x)
}
```

### 错误3: switch添加break

```go
// ❌ 不必要（Go自动break）
switch x {
case 1:
    fmt.Println("1")
    break  // 不需要
}

// ✅ 正确
switch x {
case 1:
    fmt.Println("1")
}
```

### 错误4: range循环修改值

```go
numbers := []int{1, 2, 3}

// ❌ 错误：value是副本，不会修改原数组
for _, value := range numbers {
    value = value * 2
}

// ✅ 正确：使用索引修改
for i := range numbers {
    numbers[i] = numbers[i] * 2
}
```

### 错误5: defer在循环中

```go
// ❌ 问题：defer会累积到函数结束才执行
func processFiles(files []string) {
    for _, file := range files {
        f := openFile(file)
        defer f.Close()  // 所有关闭都会延迟到函数结束
    }
}

// ✅ 正确：使用匿名函数
func processFiles(files []string) {
    for _, file := range files {
        func() {
            f := openFile(file)
            defer f.Close()  // 每次循环结束都会关闭
            // 处理文件
        }()
    }
}
```

## 性能提示

1. **range复制**: range遍历大结构体时会复制，考虑用指针
2. **字符串range**: 按rune遍历，中文安全但较慢
3. **map遍历**: 顺序随机，需要顺序时先排序key
4. **defer开销**: 有轻微性能开销，不要在性能关键路径过度使用

## 记忆技巧

- **if无括号**: Go更简洁，去掉多余括号
- **for万能**: 记住for可以模拟所有循环
- **switch自动break**: 更安全，不会忘记break
- **defer后进先出**: 栈结构，最后注册最先执行
- **range双返回**: 索引和值，用_忽略不需要的

## 快速查询

| 需求 | 代码 |
|------|------|
| 条件判断 | `if condition { }` |
| 条件+初始化 | `if x := f(); x > 0 { }` |
| 计数循环 | `for i := 0; i < n; i++ { }` |
| 条件循环 | `for condition { }` |
| 无限循环 | `for { }` |
| 遍历数组 | `for i, v := range arr { }` |
| 只要索引 | `for i := range arr { }` |
| 只要值 | `for _, v := range arr { }` |
| 跳过 | `continue` |
| 退出 | `break` |
| 多分支 | `switch x { case 1: ... }` |
| 范围判断 | `switch { case x > 0: ... }` |
| 延迟执行 | `defer cleanup()` |

---

**下一步**: 完成Day 2的练习题，掌握流程控制！🚀

