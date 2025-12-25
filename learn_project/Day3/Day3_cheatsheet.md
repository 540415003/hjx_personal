# 📝 Day 3 快速参考卡片

> Go语言复杂数据类型速查表

## 数组 Array

### 声明和初始化

```go
// 声明（固定长度）
var arr1 [5]int                    // 零值数组
var arr2 = [5]int{1, 2, 3, 4, 5}   // 完整初始化
arr3 := [3]string{"Go", "Java", "Python"}  // 短声明
arr4 := [...]int{10, 20, 30}       // 自动计算长度
arr5 := [5]int{0: 100, 4: 500}     // 指定索引初始化

// 访问和修改
scores := [3]int{85, 90, 78}
scores[1] = 95                      // 修改元素
fmt.Println(scores[0])              // 访问元素
fmt.Println(len(scores))            // 长度
```

### 遍历

```go
arr := [3]int{1, 2, 3}

// 方式1: 传统for
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}

// 方式2: range
for index, value := range arr {
    fmt.Printf("%d: %d\n", index, value)
}
```

### 特性

- ✅ 长度固定，编译时确定
- ✅ 值类型（赋值会复制整个数组）
- ⚠️ 长度是类型的一部分：`[3]int` ≠ `[5]int`
- ⚠️ 实际开发中很少用，推荐用切片

## 切片 Slice（重点！）

### 声明和创建

```go
// 方式1: 声明（nil切片）
var s1 []int
fmt.Println(s1 == nil)  // true

// 方式2: 字面量
s2 := []int{1, 2, 3, 4, 5}

// 方式3: make（推荐）
s3 := make([]int, 5)       // 长度5，容量5
s4 := make([]int, 3, 10)   // 长度3，容量10

// 方式4: 从数组切片
arr := [5]int{1, 2, 3, 4, 5}
s5 := arr[1:4]             // [2, 3, 4]

// 方式5: 从切片切片
s6 := s2[1:3]              // [2, 3]
```

### 切片操作

```go
s := []int{0, 1, 2, 3, 4, 5}

// 切片语法 [start:end] (不包括end)
s[:3]    // [0, 1, 2] - 从头到索引2
s[2:]    // [2, 3, 4, 5] - 从索引2到尾
s[:]     // [0, 1, 2, 3, 4, 5] - 全部
s[1:4]   // [1, 2, 3] - 索引1到3
```

### append 追加元素

```go
s := []int{1, 2, 3}

s = append(s, 4)           // 追加一个: [1, 2, 3, 4]
s = append(s, 5, 6, 7)     // 追加多个: [1, 2, 3, 4, 5, 6, 7]

s2 := []int{8, 9}
s = append(s, s2...)       // 追加切片: [1, 2, 3, 4, 5, 6, 7, 8, 9]
```

### copy 复制切片

```go
src := []int{1, 2, 3, 4, 5}
dst := make([]int, 3)

n := copy(dst, src)        // 复制src到dst
fmt.Println(dst)           // [1, 2, 3]
fmt.Println(n)             // 3（复制的元素数）
```

### 删除元素

```go
nums := []int{10, 20, 30, 40, 50}

// 删除索引2的元素
index := 2
nums = append(nums[:index], nums[index+1:]...)
// 结果: [10, 20, 40, 50]
```

### 插入元素

```go
nums := []int{10, 20, 40, 50}

// 在索引2插入30
index := 2
value := 30
nums = append(nums[:index], append([]int{value}, nums[index:]...)...)
// 结果: [10, 20, 30, 40, 50]
```

### 长度和容量

```go
s := make([]int, 3, 5)

len(s)   // 3 - 当前元素个数
cap(s)   // 5 - 底层数组容量

// 扩容: 容量不足时自动扩容（通常翻倍）
s = append(s, 1, 2, 3, 4)
fmt.Println(len(s), cap(s))  // 长度7，容量10（自动扩容）
```

### 切片特性

- ✅ 动态长度
- ✅ 引用类型（指向底层数组）
- ⚠️ 多个切片可能共享底层数组
- ⚠️ 修改切片可能影响其他切片

## Map 映射

### 声明和创建

```go
// 方式1: 声明（nil map，不能直接使用）
var m1 map[string]int

// 方式2: make（推荐）
m2 := make(map[string]int)

// 方式3: 字面量
m3 := map[string]int{
    "Alice": 25,
    "Bob":   30,
}
```

### 基本操作

```go
// 创建
m := make(map[string]int)

// 添加/修改
m["key1"] = 100
m["key2"] = 200

// 访问
value := m["key1"]         // 100
value := m["notexist"]     // 0（零值）

// 检查key是否存在（重要！）
if value, exists := m["key1"]; exists {
    fmt.Println(value)     // key存在
} else {
    fmt.Println("不存在")
}

// 删除
delete(m, "key1")

// 长度
fmt.Println(len(m))
```

### 遍历

```go
m := map[string]int{
    "a": 1,
    "b": 2,
    "c": 3,
}

// 遍历key和value
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}

// 只遍历key
for key := range m {
    fmt.Println(key)
}

// 只遍历value
for _, value := range m {
    fmt.Println(value)
}
```

### 嵌套Map

```go
// map的value也可以是map
nested := map[string]map[string]int{
    "group1": {
        "item1": 10,
        "item2": 20,
    },
    "group2": {
        "item1": 30,
        "item2": 40,
    },
}

// 访问
value := nested["group1"]["item1"]  // 10
```

### Map特性

- ✅ 动态增长
- ✅ 引用类型
- ⚠️ 无序（遍历顺序随机）
- ⚠️ key必须是可比较类型
- ⚠️ 不是线程安全的

### 可用作Map的key

| 类型 | 可以作为key |
|------|------------|
| int, float, string, bool | ✅ 可以 |
| pointer | ✅ 可以 |
| struct（字段都可比较）| ✅ 可以 |
| slice | ❌ 不可以 |
| map | ❌ 不可以 |
| function | ❌ 不可以 |

## 结构体 Struct

### 定义和创建

```go
// 定义结构体
type Person struct {
    Name string
    Age  int
    City string
}

// 创建实例 - 方式1: 零值
var p1 Person

// 方式2: 字面量（按顺序）
p2 := Person{"Alice", 25, "Beijing"}

// 方式3: 字面量（指定字段，推荐）
p3 := Person{
    Name: "Bob",
    Age:  30,
    City: "Shanghai",
}

// 方式4: 部分初始化
p4 := Person{
    Name: "Charlie",
    // Age和City使用零值
}

// 方式5: 指针（使用new）
p5 := new(Person)  // 返回*Person
```

### 访问和修改

```go
type TestCase struct {
    ID     int
    Name   string
    Status string
}

tc := TestCase{
    ID:     1,
    Name:   "登录测试",
    Status: "pending",
}

// 访问字段
fmt.Println(tc.Name)

// 修改字段
tc.Status = "passed"
tc.ID = 2
```

### 方法

```go
type TestCase struct {
    ID     int
    Name   string
    Status string
}

// 值接收者方法
func (tc TestCase) Display() {
    fmt.Printf("[%d] %s: %s\n", tc.ID, tc.Name, tc.Status)
}

// 指针接收者方法
func (tc *TestCase) Run() {
    tc.Status = "running"  // 修改接收者
    // 执行测试...
    tc.Status = "passed"
}

// 使用
tc := TestCase{ID: 1, Name: "测试"}
tc.Display()   // 值接收者
tc.Run()       // 指针接收者（会修改tc）
```

### 值接收者 vs 指针接收者

| 接收者类型 | 何时使用 | 特点 |
|-----------|---------|------|
| 值接收者 `(t Type)` | 只读方法 | 不会修改原值 |
| 指针接收者 `(t *Type)` | 需要修改 | 会修改原值 |
| 指针接收者 `(t *Type)` | 结构体很大 | 避免复制 |

### 结构体嵌入（组合）

```go
// 基础结构体
type BaseTest struct {
    ID     int
    Name   string
    Status string
}

// 嵌入结构体
type APITest struct {
    BaseTest       // 匿名字段（嵌入）
    URL    string
    Method string
}

// 使用
api := APITest{
    BaseTest: BaseTest{
        ID:     1,
        Name:   "API测试",
        Status: "pending",
    },
    URL:    "https://api.example.com",
    Method: "POST",
}

// 可以直接访问嵌入结构体的字段
fmt.Println(api.ID)      // 直接访问
fmt.Println(api.Name)    // 直接访问
fmt.Println(api.URL)     // 自己的字段

// 也可以显式访问
fmt.Println(api.BaseTest.Status)
```

### 结构体标签

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// 标签用于JSON、XML等序列化
```

## 指针 Pointer

### 基础

```go
// 声明指针
var p *int

// 获取地址
x := 42
p = &x         // &取地址符

// 解引用
fmt.Println(*p)  // *解引用，获取指针指向的值

// 修改值
*p = 100
fmt.Println(x)   // 100
```

### 指针操作符

| 操作符 | 名称 | 说明 |
|-------|------|------|
| `&` | 取地址 | `&x` 获取x的地址 |
| `*` | 解引用 | `*p` 获取p指向的值 |

### 结构体指针

```go
type Person struct {
    Name string
    Age  int
}

p := &Person{Name: "Alice", Age: 25}

// Go自动解引用，可以直接访问字段
p.Name = "Bob"     // 等价于 (*p).Name = "Bob"
fmt.Println(p.Age) // 等价于 (*p).Age
```

### 函数参数

```go
// 值传递（不会修改原值）
func modifyValue(x int) {
    x = 100
}

// 指针传递（会修改原值）
func modifyPointer(x *int) {
    *x = 100
}

// 使用
num := 10
modifyValue(num)     // num还是10
modifyPointer(&num)  // num变成100
```

### new函数

```go
// new分配内存并返回指针
p := new(int)        // 分配int零值，返回*int
*p = 100

tc := new(TestCase)  // 分配TestCase零值，返回*TestCase
tc.ID = 1
```

## 常见模式

### 模式1: 切片作为栈

```go
// 创建栈
stack := []int{}

// 压栈
stack = append(stack, 1)
stack = append(stack, 2)

// 弹栈
if len(stack) > 0 {
    top := stack[len(stack)-1]
    stack = stack[:len(stack)-1]
}
```

### 模式2: 切片作为队列

```go
// 创建队列
queue := []int{}

// 入队
queue = append(queue, 1)
queue = append(queue, 2)

// 出队
if len(queue) > 0 {
    first := queue[0]
    queue = queue[1:]
}
```

### 模式3: Map计数

```go
items := []string{"a", "b", "a", "c", "b", "a"}
counter := make(map[string]int)

for _, item := range items {
    counter[item]++
}
// 结果: {"a": 3, "b": 2, "c": 1}
```

### 模式4: Map分组

```go
type Item struct {
    Category string
    Name     string
}

items := []Item{
    {"fruit", "apple"},
    {"fruit", "banana"},
    {"vegetable", "carrot"},
}

groups := make(map[string][]Item)
for _, item := range items {
    groups[item.Category] = append(groups[item.Category], item)
}
```

### 模式5: 结构体工厂函数

```go
// 类似构造函数
func NewTestCase(id int, name string) *TestCase {
    return &TestCase{
        ID:     id,
        Name:   name,
        Status: "pending",
    }
}

// 使用
tc := NewTestCase(1, "登录测试")
```

## 测试场景应用

### 场景1: 测试结果收集

```go
type TestResult struct {
    Name     string
    Status   string
    Duration float64
}

// 使用切片收集
results := []TestResult{}
results = append(results, TestResult{
    Name:     "测试1",
    Status:   "passed",
    Duration: 1.5,
})
```

### 场景2: 环境配置

```go
// 使用map存储配置
configs := map[string]map[string]string{
    "dev": {
        "url":     "http://dev.example.com",
        "timeout": "30",
    },
    "prod": {
        "url":     "https://www.example.com",
        "timeout": "120",
    },
}
```

### 场景3: 测试用例建模

```go
type TestCase struct {
    ID       int
    Name     string
    Priority string
    Tags     []string
    Status   string
}

func (tc *TestCase) Run() {
    tc.Status = "running"
    // 执行测试...
    tc.Status = "passed"
}
```

### 场景4: 测试套件

```go
type TestSuite struct {
    Name      string
    TestCases []*TestCase
}

func (ts *TestSuite) AddTest(tc *TestCase) {
    ts.TestCases = append(ts.TestCases, tc)
}

func (ts *TestSuite) RunAll() {
    for _, tc := range ts.TestCases {
        tc.Run()
    }
}
```

## 常见错误

### 错误1: nil map

```go
// ❌ 错误
var m map[string]int
m["key"] = 1  // panic: nil map

// ✅ 正确
m := make(map[string]int)
m["key"] = 1
```

### 错误2: 切片越界

```go
// ❌ 错误
s := []int{1, 2, 3}
fmt.Println(s[5])  // panic: index out of range

// ✅ 正确
if len(s) > 5 {
    fmt.Println(s[5])
}
```

### 错误3: 修改range的value

```go
// ❌ 错误：value是副本
nums := []int{1, 2, 3}
for _, value := range nums {
    value = value * 2  // 不会修改原切片
}

// ✅ 正确：使用索引
for i := range nums {
    nums[i] = nums[i] * 2
}
```

### 错误4: 忘记接收append返回值

```go
// ❌ 错误
s := []int{1, 2, 3}
append(s, 4)  // 没有接收返回值

// ✅ 正确
s = append(s, 4)
```

### 错误5: 结构体方法接收者类型错误

```go
// ❌ 如果需要修改，不能用值接收者
func (tc TestCase) Run() {
    tc.Status = "running"  // 不会修改原结构体
}

// ✅ 正确：使用指针接收者
func (tc *TestCase) Run() {
    tc.Status = "running"  // 会修改原结构体
}
```

## 性能提示

1. **预分配切片容量**: `make([]int, 0, 100)` 比动态扩容快
2. **大结构体使用指针**: 避免值传递的复制开销
3. **Map预分配**: `make(map[string]int, 100)` 避免频繁扩容
4. **避免在循环中append到共享切片**: 可能导致频繁内存分配

## 快速查询

| 需求 | 代码 |
|------|------|
| 创建切片 | `s := make([]int, 0)` |
| 追加元素 | `s = append(s, 1)` |
| 复制切片 | `copy(dst, src)` |
| 创建map | `m := make(map[string]int)` |
| 检查key | `if v, ok := m[k]; ok {}` |
| 删除key | `delete(m, key)` |
| 定义结构体 | `type T struct { ... }` |
| 值接收者 | `func (t T) Method() {}` |
| 指针接收者 | `func (t *T) Method() {}` |
| 获取地址 | `&x` |
| 解引用 | `*p` |

---

**下一步**: 完成Day 3的练习题，掌握复杂数据类型！🚀

