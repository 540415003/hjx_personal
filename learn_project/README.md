# 🚀 Go语言15天快速入门学习计划

> 为有Java基础的QA工程师量身定制

## 📋 学习进度

- [x] **Day 1**: 环境搭建、基础语法、数据类型、变量、常量 ✅
- [x] **Day 2**: 流程控制（if、for、switch、defer） ✅
- [ ] **Day 3**: 复杂数据类型（数组、切片、Map、结构体）
- [ ] **Day 4**: 函数（多返回值、变参、闭包）
- [ ] **Day 5**: 方法和接口
- [ ] **Day 6**: 错误处理
- [ ] **Day 7**: 并发基础（Goroutine、Channel）
- [ ] **Day 8**: 包和模块
- [ ] **Day 9**: 文件和I/O
- [ ] **Day 10**: 面向对象思想
- [ ] **Day 11**: 单元测试 ⭐
- [ ] **Day 12**: 测试进阶 ⭐
- [ ] **Day 13**: 自动化测试框架 ⭐
- [ ] **Day 14**: 实战项目1（API测试）
- [ ] **Day 15**: 实战项目2（E2E测试）

## 📁 文件说明

```
learn_project/
├── README.md           # 学习指南（本文件）
├── Day1/              # Day 1 学习资料
│   ├── Day1.go           # 学习内容和示例代码
│   ├── Day1_answers.go   # 练习题参考答案
│   └── Day1_cheatsheet.md # 快速参考卡片
├── Day2/              # Day 2 学习资料
│   ├── Day2.go           # 学习内容和示例代码
│   ├── Day2_answers.go   # 练习题参考答案
│   └── Day2_cheatsheet.md # 快速参考卡片
└── Day3/              # Day 3 学习资料
    ├── Day3.go           # 学习内容和示例代码
    ├── Day3_answers.go   # 练习题参考答案
    └── Day3_cheatsheet.md # 快速参考卡片
```

## 🎯 Day 1 学习指南

### 快速开始

1. **运行学习程序**
   ```bash
   go run day1.go
   ```
   这会运行所有示例代码，展示各种Go语法特性。

2. **完成练习题**
   - 打开 `day1.go` 文件
   - 找到底部的4个练习题
   - 取消注释并编写你的代码
   - 运行程序检查结果

3. **查看参考答案**（完成练习后）
   ```bash
   go run day1_answers.go
   ```

### Day 1 学习内容

#### 1️⃣ Hello World
- Go程序的基本结构
- package 和 import
- main 函数

#### 2️⃣ 变量声明（重点）
Go有三种变量声明方式：
```go
var name string = "张三"     // 完整声明
var age = 25               // 类型推断
city := "北京"              // 短声明（最常用）
```

**与Java对比**：
- Java: `String name = "张三";`
- Go可以省略类型，更简洁

#### 3️⃣ 基本数据类型
| Go类型 | Java对应 | 说明 |
|--------|---------|------|
| bool | boolean | 布尔值 |
| string | String | 字符串（不可变）|
| int, int8, int16, int32, int64 | int, long | 整数 |
| uint, uint8, uint16, uint32, uint64 | - | 无符号整数 |
| float32, float64 | float, double | 浮点数 |
| byte | byte | uint8别名 |
| rune | char | int32别名，Unicode |

#### 4️⃣ 零值（默认值）
Go会自动初始化变量：
- 数字类型 → `0`
- 布尔类型 → `false`
- 字符串 → `""`（空字符串）

#### 5️⃣ 类型转换（重点）
**Go要求显式类型转换**：
```go
var i int = 42
var f float64 = float64(i)  // 必须显式转换
```

❌ **错误示例**：
```go
var x int = 10
var y float64 = x  // 编译错误！
```

#### 6️⃣ 常量和 iota
```go
const Pi = 3.14159

const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
)
```

**iota** 是Go特有的枚举计数器，在const块中自动递增。

### 练习题

#### 练习1：变量声明和基本运算
声明你的个人信息变量（姓名、年龄、工资），计算年收入。

#### 练习2：常量和iota
定义测试用例状态常量：Pending、Running、Passed、Failed。

#### 练习3：类型转换和计算
计算测试用例通过率（注意类型转换）。

#### 练习4：综合应用
定义测试环境配置（开发、测试、生产环境）。

### 关键知识点

#### 🔴 Go与Java的主要区别

1. **类型声明位置不同**
   - Java: `String name`
   - Go: `name string`

2. **无隐式类型转换**
   - Go必须显式转换所有类型

3. **未使用的变量**
   - Go: 编译错误 ❌
   - Java: 编译警告 ⚠️

4. **自增自减**
   - Go: 只支持 `i++`（后置）
   - Java: 支持 `++i` 和 `i++`

5. **字符串不可变**
   - 两者都不可变
   - 但处理方式略有不同

## 🛠️ 常用命令

```bash
# 运行程序
go run day1.go

# 编译程序（生成可执行文件）
go build day1.go

# 格式化代码（自动整理代码格式）
go fmt day1.go

# 查看Go版本
go version

# 查看环境信息
go env
```

## 💡 学习建议

1. **每天坚持学习**：建议每天投入1-2小时
2. **动手实践**：光看不练假把式，一定要写代码
3. **对比学习**：利用你的Java知识，对比理解
4. **及时总结**：完成每天学习后做个小总结
5. **应用导向**：思考如何应用到工作中的测试场景

## 🎓 Day 1 学习检查清单

完成以下内容后，你就可以开始Day 2了：

- [ ] 成功运行 `day1.go`
- [ ] 理解三种变量声明方式
- [ ] 掌握基本数据类型
- [ ] 理解零值概念
- [ ] 掌握显式类型转换
- [ ] 理解常量和iota
- [ ] 完成4个练习题
- [ ] 能用Go写简单的计算程序

## 📚 参考资源

- [Go官方文档](https://go.dev/doc/)
- [Go语言之旅](https://tour.go-zh.org/)
- [Go by Example](https://gobyexample.com/)

## ❓ 常见问题

**Q: 为什么Go的变量类型写在后面？**  
A: 这是Go的设计哲学，让代码更易读，特别是在函数返回值时更清晰。

**Q: 必须显式类型转换太麻烦了？**  
A: 这是为了类型安全，避免隐式转换带来的潜在bug，对测试质量有好处。

**Q: := 和 var 什么时候用？**  
A: 函数内优先用 `:=`（简洁），包级别变量必须用 `var`。

**Q: 为什么未使用的变量会报错？**  
A: Go鼓励写简洁的代码，未使用的变量可能是潜在bug，强制删除可以保持代码质量。

## 🎯 Day 2 学习指南

### 快速开始

1. **运行学习程序**
   ```bash
   cd Day2
   go run Day2.go
   ```
   这会运行所有流程控制的示例代码。

2. **完成练习题**
   - 打开 `Day2/Day2.go` 文件
   - 找到底部的5个练习题
   - 编写你的代码
   - 运行程序检查结果

3. **查看参考答案**（完成练习后）
   ```bash
   go run Day2_answers.go
   ```

### Day 2 学习内容

#### 1️⃣ if 条件语句（无需括号）
```go
// 基本形式
if age >= 18 {
    fmt.Println("成年人")
}

// if带初始化（重点！）
if passRate := float64(passed)/float64(total)*100; passRate >= 95 {
    fmt.Printf("优秀: %.2f%%\n", passRate)
}
```

**与Java对比**：
- Go的if不需要括号
- Go支持在if中初始化变量，作用域仅在if块内

#### 2️⃣ for 循环（唯一的循环）
Go只有for循环，但可以实现所有循环形式：

```go
// 传统for
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// 类似while
for condition {
    // ...
}

// 无限循环
for {
    // ...
}

// range遍历（最常用）
for index, value := range slice {
    fmt.Printf("%d: %v\n", index, value)
}
```

**重点**：range是最常用的遍历方式，支持数组、切片、map、字符串。

#### 3️⃣ switch 语句（自动break）
```go
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
// 不需要break！

// 无表达式switch（相当于if-else链）
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
}
```

**与Java对比**：
- Go的switch自动break，不会fall through
- 支持多个值在一个case中
- case可以是表达式，不仅限于常量

#### 4️⃣ defer 延迟执行（Go特色）
```go
func example() {
    defer fmt.Println("最后执行")
    defer fmt.Println("倒数第二")
    fmt.Println("先执行")
}
// 输出：先执行 -> 倒数第二 -> 最后执行
```

**重要特性**：
- defer在函数返回前执行
- 多个defer按后进先出(LIFO)顺序执行
- 常用于资源清理、计时、错误恢复

**实际应用**：
```go
// 资源清理
file := openFile("test.txt")
defer file.Close()

// 测试计时
start := time.Now()
defer func() {
    fmt.Printf("耗时: %v\n", time.Since(start))
}()
```

### 练习题

#### 练习1：条件判断
编写测试用例优先级判断函数（smoke/regression/integration）。

#### 练习2：循环遍历
统计测试结果切片，计算通过率，找出第一个失败的用例。

#### 练习3：switch应用
实现HTTP状态码处理器（2xx/3xx/4xx/5xx）。

#### 练习4：defer应用
创建测试执行包装器，记录开始/结束时间和耗时。

#### 练习5：综合练习
实现测试用例过滤和执行器（标签过滤、启用检查、结果统计）。

### 关键知识点

#### 🔴 Go与Java流程控制对比

1. **if语句**
   - Go: 无需括号，但必须有花括号
   - Go支持初始化语句

2. **循环**
   - Go: 只有for，没有while/do-while
   - Java: for/while/do-while都有

3. **switch**
   - Go: 自动break，默认不fall through
   - Java: 需要手动break

4. **range**
   - Go特有，类似Java的增强for循环
   - 更简洁，功能更强

5. **defer**
   - Go特有，类似Java的finally但更灵活
   - 可以有多个，按LIFO顺序执行

## 🎓 Day 2 学习检查清单

完成以下内容后，你就可以开始Day 3了：

- [ ] 理解if语句不需要括号
- [ ] 掌握for循环的4种形式
- [ ] 熟练使用range遍历
- [ ] 理解switch自动break的特性
- [ ] 掌握defer的执行顺序（LIFO）
- [ ] 完成5个练习题
- [ ] 能用流程控制编写测试逻辑

## 🎯 Day 3 学习指南

### 快速开始

1. **运行学习程序**
   ```bash
   cd Day3
   go run Day3.go
   ```
   这会运行所有复杂数据类型的示例代码。

2. **完成练习题**
   - 打开 `Day3/Day3.go` 文件
   - 找到底部的5个练习题
   - 编写你的代码

3. **查看参考答案**（完成练习后）
   ```bash
   go run Day3_answers.go
   ```

### Day 3 学习内容

#### 1️⃣ 数组 Array（基础）

```go
// 固定长度
var arr [5]int                  // 零值数组
arr := [3]int{1, 2, 3}          // 初始化
arr := [...]int{1, 2, 3, 4}     // 自动推断长度
```

**特点**：
- 长度固定，编译时确定
- 值类型（赋值会复制）
- 实际开发很少用

#### 2️⃣ 切片 Slice（重点！）

```go
// 动态数组
s := []int{1, 2, 3}              // 字面量
s := make([]int, 5)              // make创建，长度5
s := make([]int, 3, 10)          // 长度3，容量10

// 操作
s = append(s, 4)                 // 追加元素
copy(dst, src)                   // 复制切片
s[1:4]                           // 切片操作

// 重要属性
len(s)                           // 长度
cap(s)                           // 容量
```

**重点**：
- 动态长度，最常用
- 引用类型（指向底层数组）
- append时容量不足会自动扩容
- 多个切片可能共享底层数组

#### 3️⃣ Map 映射（哈希表）

```go
// 创建
m := make(map[string]int)
m := map[string]int{
    "key1": 100,
    "key2": 200,
}

// 操作
m["key"] = 100                   // 添加/修改
value := m["key"]                // 访问
if v, ok := m["key"]; ok {       // 检查存在
    // key存在
}
delete(m, "key")                 // 删除

// 遍历
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}
```

**重点**：
- key-value存储
- 无序（遍历随机）
- key必须是可比较类型
- 引用类型

#### 4️⃣ 结构体 Struct（Go的"类"）

```go
// 定义
type TestCase struct {
    ID       int
    Name     string
    Priority string
    Status   string
}

// 创建
tc := TestCase{
    ID:     1,
    Name:   "登录测试",
    Status: "pending",
}

// 方法 - 值接收者
func (tc TestCase) Display() {
    fmt.Println(tc.Name)
}

// 方法 - 指针接收者（可修改）
func (tc *TestCase) Run() {
    tc.Status = "running"
}

// 使用
tc.Display()
tc.Run()
```

**重点**：
- Go没有类，用struct + 方法
- 值接收者：只读，不修改
- 指针接收者：可修改，大结构体用指针
- 通过组合（嵌入）实现继承

#### 5️⃣ 指针 Pointer

```go
// 基础
var p *int                       // 指针声明
x := 42
p = &x                           // &取地址
*p = 100                         // *解引用

// 结构体指针
tc := &TestCase{ID: 1}
tc.Name = "测试"                  // 自动解引用

// 函数参数
func modify(p *int) {
    *p = 100                     // 修改原值
}
```

**重点**：
- `&` 取地址
- `*` 解引用
- 指针传递可以修改原值
- 结构体指针自动解引用

### 练习题

#### 练习1：切片操作
实现测试结果管理器（添加、统计、计算通过率、找失败索引）。

#### 练习2：Map应用
实现多环境配置管理器（添加、查询、更新配置）。

#### 练习3：结构体设计
设计测试用例管理系统（定义结构体、方法、创建实例）。

#### 练习4：综合应用
实现测试报告生成器（使用切片、map、结构体、指针）。

#### 练习5：数据处理
测试数据分析工具（按日期统计、找最高/最低、平均值）。

### 关键知识点

#### 🔴 数据类型对比

| 类型 | 长度 | 类型 | 使用场景 |
|------|------|------|----------|
| 数组 | 固定 | 值类型 | 很少用 |
| 切片 | 动态 | 引用类型 | 最常用 |
| Map | 动态 | 引用类型 | 键值存储 |
| 结构体 | 固定 | 值类型 | 数据建模 |

#### 🔴 与Java对比

1. **数组**
   - Java: 引用类型
   - Go: 值类型，很少用

2. **切片 vs ArrayList**
   - 切片更底层、更高效
   - 使用append而不是add

3. **Map vs HashMap**
   - Go的Map是内置类型
   - 使用更简洁

4. **Struct vs Class**
   - Go没有类
   - 使用struct + 方法
   - 组合而非继承

## 🎓 Day 3 学习检查清单

完成以下内容后，你就可以开始Day 4了：

- [ ] 理解数组和切片的区别
- [ ] 掌握切片的append、copy、切片操作
- [ ] 熟练使用make创建切片和map
- [ ] 掌握map的增删改查和遍历
- [ ] 理解map是无序的
- [ ] 掌握结构体的定义和实例化
- [ ] 理解值接收者和指针接收者的区别
- [ ] 掌握指针的基本使用
- [ ] 完成5个练习题
- [ ] 能用复杂数据类型建模测试数据

## 🚀 下一步

完成Day 3后，准备开始Day 4：

```bash
# Day 4 主题：函数
# 内容：多返回值、变参、闭包、匿名函数
```

---

💪 加油！继续保持学习热情！有任何问题随时问我！

