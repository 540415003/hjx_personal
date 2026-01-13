# Code Guardian MCP 工具使用指南

## 目录

1. [工具概述](#1-工具概述)
2. [功能特性](#2-功能特性)
3. [环境要求](#3-环境要求)
4. [工具使用方法](#4-工具使用方法)
5. [使用示例](#5-使用示例)
6. [最佳实践](#6-最佳实践)
7. [常见问题与解决方案](#7-常见问题与解决方案)
8. [优化建议](#8-优化建议)
9. [附录](#9-附录)

---

## 1. 工具概述

### 1.1 什么是 Code Guardian

**Code Guardian** 是一个基于 MCP (Model Context Protocol) 的智能代码覆盖率提升工具。它通过与 AI 协作，自动化完成以下核心任务：

- 🔍 **覆盖率分析**: 从 DMS 平台获取代码覆盖率数据，精准识别未覆盖的代码块
- 🧪 **测试生成**: 基于未覆盖代码块智能生成单元测试
- ✅ **测试验证**: 本地执行测试，确保测试用例正确通过
- 📤 **代码推送**: 自动提交和推送测试代码到远程仓库
- 🚀 **DMS 触发**: 触发 DMS 任务执行覆盖率验证

### 1.2 核心价值

| 价值维度 | 描述 |
|---------|------|
| **效率提升** | 自动化识别未覆盖代码，减少人工分析时间 |
| **质量保障** | 系统性补充测试用例，提升代码覆盖率 |
| **标准化** | 遵循统一的测试规范，保证代码质量一致性 |
| **可追溯** | 生成详细的执行报告，便于追踪和审计 |

### 1.3 工作流程图

```
┌─────────────────────────────────────────────────────────────────────┐
│                    Code Guardian 流程图                              │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌──────────┐    ┌──────────────┐    ┌──────────────┐               │
│  │ 参数验证  │───▶│ 获取覆盖率数据  │───▶│ 解析未覆盖块   │               │
│  └──────────┘    └──────────────┘    └──────────────┘               │
│                                              │                      │
│                                              ▼                      │
│  ┌──────────┐    ┌──────────────┐    ┌──────────────┐               │
│  │ 推送代码  │◀───│  执行测试      │◀───│  生成测试文件  │               │
│  └──────────┘    └──────────────┘    └──────────────┘               │
│        │                                                            │
│        ▼                                                            │
│  ┌──────────┐    ┌──────────────┐                                   │
│  │ 触发 DMS  │───▶│  生成报告     │                                   │
│  └──────────┘    └──────────────┘                                   │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

## 2. 功能特性

### 2.1 核心工具列表

| 工具名称 | 功能描述 | 必需参数 |
|---------|---------|---------|
| `validate_and_confirm` | 验证和确认所有必填参数 | git_repo, git_branch, target_branch, env, ssc_org_id |
| `get_and_parse_uncovered_blocks` | 获取覆盖率数据并解析未覆盖代码块 | git_repo, git_branch, env, ssc_org_id |
| `generate_test_file` | 生成测试文件框架 | uncovered_blocks |
| `write_test_code` | 写入测试代码 | test_file_path, test_code |
| `execute_test` | 执行测试 | test_file_path |
| `fix_test_failure` | 修复测试失败 | test_file_path, fixed_test_code |
| `push_to_repo` | 推送代码到仓库 | (可选参数) |
| `trigger_dms_task` | 触发 DMS 任务 | git_repo, branch_name, repo_id, jira_key |
| `generate_summary` | 生成执行总结 | execution_data |
| `report_usage_to_aura` | 上报使用情况到 Aura | execution_data |

### 2.2 支持的编程语言

- ✅ Go (主要支持)

### 2.3 支持的测试框架

- Go testing 标准库
- testify/assert
- gomonkey (Mock 框架)
- gomock

---

## 3. 环境要求

### 3.1 基础环境

| 组件 | 版本要求 | 说明 |
|-----|---------|------|
| Cursor IDE | 最新版本 | 支持 MCP 的 AI 编辑器 |
| Go | 1.17+ | 项目依赖的 Go 版本 |
| Git | 2.0+ | 版本控制工具 |

### 3.2 MCP 服务配置

确保 Cursor 的 MCP 配置中已添加 Code Guardian 服务：

```json
{
  "mcpServers": {
    "code-guardian": {
      "command": "node",
      "args": ["/path/to/code-guardian/server.js"],
      "env": {
        "API_TOKEN": "your-api-token"
      }
    }
  }
}
```

### 3.3 项目依赖

确保项目中安装了必要的测试依赖：

```bash
# gomonkey - Mock 框架
go get github.com/agiledragon/gomonkey/v2

# testify - 断言库
go get github.com/stretchr/testify
```

---

## 4. 工具使用方法

### 4.1 标准使用流程

#### 步骤 1: 启动对话并提供参数

向 AI 发送以下格式的请求：

```
使用 code_guardian MCP 工具生成单元测试，执行完整流程：

## 参数
git_repo: <完整仓库路径>
git_branch: <源分支名>
work_path: <本地工作目录>
target_branch: <目标分支名>
env: <环境: test/staging/prod>
ssc_org_id: <SSC 组织 ID>
pfb: <PFB 参数>

## 要求
1. 使用 MCP Tool 完成全流程：获取覆盖率 → 生成测试 → 本地执行 → push代码 → 触发DMS任务 → 上报Aura
2. 严禁自证预言测试：测试文件必须与被测代码同目录同包，必须实际调用被测函数并验证返回值
3. 遵循 2-2-1 原则：至少 2 成功 + 2 失败 + 1 边界场景
```

#### 步骤 2: AI 自动执行完整流程

AI 将依次调用以下 MCP 工具：

| 步骤 | 工具 | 说明 |
|-----|------|------|
| 1 | `validate_and_confirm` | 验证参数 |
| 2 | `get_and_parse_uncovered_blocks` | 获取未覆盖代码 |
| 3 | `generate_test_file` + `write_test_code` | 生成测试代码 |
| 4 | `execute_test` | 本地执行验证 |
| 5 | `push_to_repo` | 推送代码（需确认） |
| 6 | `trigger_dms_task` | 触发 DMS 任务 |
| 7 | `report_usage_to_aura` | 上报数据到 Aura |
| 8 | `generate_summary` | 生成执行报告 |

### 4.2 参数说明

| 参数名 | 必填 | 说明 | 示例 |
|-------|-----|------|------|
| `git_repo` | ✅ | Git 仓库完整路径 | `shopee/bg-logistics/spx/fulfillment/fulfillment-service` |
| `git_branch` | ✅ | 源分支名 | `feature/SPXFM-186041-location-update` |
| `work_path` | ✅ | 本地工作目录绝对路径 | `/Users/xxx/code/fulfillment-service` |
| `target_branch` | ✅ | 目标分支（通常与源分支相同） | `feature/SPXFM-186041-location-update` |
| `env` | ✅ | 环境标识 | `test`, `staging`, `prod` |
| `ssc_org_id` | ✅ | SSC 组织 ID | `27` |
| `pfb` | 推荐 | PFB 参数，用于 DMS 任务筛选 | `pfb-dms-qa-vn-location-update` |

---

## 5. 使用示例

### 5.1 完整示例：fulfillment-service 仓库

#### 5.1.1 初始请求

```
使用 code_guardian MCP 工具生成单元测试，执行完整流程：

## 参数
git_repo: shopee/bg-logistics/spx/fulfillment/fulfillment-service
git_branch: feature/SPXFM-186041-location-update
work_path: /Users/jiaxuan.han/code/dev_project/fulfillment-service
target_branch: feature/SPXFM-186041-location-update
env: test
ssc_org_id: 27
pfb: pfb-dms-qa-vn-location-update

## 要求
1. 使用 MCP Tool 完成全流程：获取覆盖率 → 生成测试 → 本地执行 → push代码 → 触发DMS任务 → 上报Aura
2. 严禁自证预言测试：测试文件必须与被测代码同目录同包，必须实际调用被测函数并验证返回值
3. 遵循 2-2-1 原则：至少 2 成功 + 2 失败 + 1 边界场景
```

#### 5.1.2 执行过程

**第一阶段：参数验证**

AI 调用 `validate_and_confirm` 工具验证所有参数：

```json
{
  "git_repo": "shopee/bg-logistics/spx/fulfillment/fulfillment-service",
  "git_branch": "feature/SPXFM-186041-location-update",
  "target_branch": "feature/SPXFM-186041-location-update",
  "env": "test",
  "ssc_org_id": "27",
  "pfb": "pfb-dms-qa-vn-location-update"
}
```

**第二阶段：获取未覆盖代码块**

AI 调用 `get_and_parse_uncovered_blocks` 获取覆盖率数据：

```
识别到 6 个未覆盖的代码块：
1. apps/order/business/update_destination.go:360-386
2. apps/order/business/update_destination.go:441-447
3. thirdparty/site/site_network_router/address_service.go:163-174
4. thirdparty/site/site_network_router/base_service.go:516-531
5. apps/order/schema/order_grpc.go:7440-7465
6. apps/order/service/solution/network_solution.go:2114-2188
```

**第三阶段：生成测试文件**

AI 调用 `generate_test_file` 生成测试框架，然后使用 `write_test_code` 写入测试代码。

**重要提示**：测试文件必须与被测代码在同一目录下（同包），而不是独立的 `test/` 目录。

✅ **正确位置**:
```
thirdparty/site/site_network_router/address_service_test.go
```

❌ **错误位置**:
```
test/thirdparty/site/site_network_router/address_service_test.go
```

**第四阶段：执行测试**

AI 调用 `execute_test` 执行本地测试：

```bash
go test -gcflags=all=-l -v -coverprofile=coverage.out ./thirdparty/site/site_network_router/...
```

测试结果：
```
=== RUN   TestGetLowestAddressNameByAddressId
--- PASS: TestGetLowestAddressNameByAddressId (0.00s)
    --- PASS: success_get_address_name
    --- PASS: success_address_not_in_mapping
    --- PASS: success_nil_mapping
    --- PASS: error_get_mapping_failed
    --- PASS: boundary_zero_address_id
PASS
coverage: 1.4% of statements
```

**第五阶段：推送代码**

AI 调用 `push_to_repo` 推送代码：

```json
{
  "status": "success",
  "branch": "feature/SPXFM-186041-location-update",
  "message": "Code pushed successfully"
}
```

#### 5.1.3 最终结果

| 指标 | 结果 |
|-----|------|
| 生成测试文件 | 1 个 |
| 测试用例数 | 7 个 |
| 测试通过率 | 100% |
| 覆盖率提升 | 1.4% |
| 覆盖函数 | `GetLowestAddressNameByAddressId`, `GetLowestAddressIdNameMapping` |

### 5.1.4 生成的测试代码示例

```go
package site_network_router // <ai-gen>

import ( // <ai-gen>
    "context" // <ai-gen>
    "errors" // <ai-gen>
    "reflect" // <ai-gen>
    "testing" // <ai-gen>

    "git.garena.com/shopee/bg-logistics/spx/network-route/network-route-protobuf/pkg/address" // <ai-gen>

    "github.com/agiledragon/gomonkey/v2" // <ai-gen>
    "github.com/stretchr/testify/assert" // <ai-gen>
) // <ai-gen>

func TestGetLowestAddressNameByAddressId(t *testing.T) { // <ai-gen>
    tests := []struct { // <ai-gen>
        name            string            // <ai-gen>
        addressId       uint64            // <ai-gen>
        mockMapping     map[uint64]string // <ai-gen>
        mockErr         error             // <ai-gen>
        wantAddressName string            // <ai-gen>
        wantErr         bool              // <ai-gen>
    }{ // <ai-gen>
        { // <ai-gen>
            name:            "success_get_address_name", // <ai-gen>
            addressId:       12345, // <ai-gen>
            mockMapping:     map[uint64]string{12345: "District Name"}, // <ai-gen>
            mockErr:         nil, // <ai-gen>
            wantAddressName: "District Name", // <ai-gen>
            wantErr:         false, // <ai-gen>
        }, // <ai-gen>
        // ... 更多测试用例
    } // <ai-gen>

    for _, tc := range tests { // <ai-gen>
        t.Run(tc.name, func(t *testing.T) { // <ai-gen>
            patches := gomonkey.NewPatches() // <ai-gen>
            defer patches.Reset() // <ai-gen>

            facade := &SiteNetworkRouteFacade{serviceName: "test-service"} // <ai-gen>

            patches.ApplyMethod(reflect.TypeOf(facade), "GetLowestAddressIdNameMapping", // <ai-gen>
                func(_ *SiteNetworkRouteFacade, _ context.Context, _ *address.GetLowestAddressIdNameMappingReq) (map[uint64]string, error) { // <ai-gen>
                    return tc.mockMapping, tc.mockErr // <ai-gen>
                }) // <ai-gen>

            ctx := context.Background() // <ai-gen>
            result, err := facade.GetLowestAddressNameByAddressId(ctx, tc.addressId) // <ai-gen>

            if tc.wantErr { // <ai-gen>
                assert.Error(t, err) // <ai-gen>
            } else { // <ai-gen>
                assert.NoError(t, err) // <ai-gen>
                assert.Equal(t, tc.wantAddressName, result) // <ai-gen>
            } // <ai-gen>
        }) // <ai-gen>
    } // <ai-gen>
} // <ai-gen>
```
### 5.2 示例：fleetorder-service 仓库（问题案例）

#### 5.2.1 初始请求

```
使用 code_guardian 生成单元测试：

git_repo: shopee/bg-logistics/spx/order-center/fleetorder-service
git_branch: feature/SPXFM-186041-location-update
work_path: /Users/jiaxuan.han/code/dev_project/fleetorder-service
target_branch: feature/SPXFM-186041-location-update
env: test
ssc_org_id: 27
pfb: pfb-dms-qa-vn-location-update
```

#### 5.2.2 执行过程

**第一阶段：获取未覆盖代码块**

识别到 5 个需要覆盖的代码块：

| 序号 | 文件路径 | 行号 | 描述 |
|-----|---------|------|------|
| 1 | `apps/order/service/order_tracking_list.go` | 100-145 | `calcBandSLowestAddressName` 分支逻辑 |
| 2 | `apps/order/schema/order_grpc.go` | 4068-4080 | `GetFleetOrderInfoList` 地址解析成功路径 |
| 3 | `apps/order/schema/order_grpc.go` | 4274-4286 | `BatchGetFleetOrderInfoList` 地址解析成功路径 |
| 4 | `apps/order/service/order.go` | 1234-1260 | `ParseAddress` 方法 |
| 5 | `apps/track/service/service_point_track_handler.go` | 130-132 | JSON 解析错误分支 |

**第二阶段：生成测试用例**

成功生成 3 个测试文件，共计 15+ 个测试用例：

| 测试文件 | 测试函数数 | 覆盖的代码块 |
|---------|-----------|-------------|
| `apps/order/service/order_tracking_list_test.go` | 8 | calcBandSLowestAddressName 多个分支 |
| `apps/order/schema/order_grpc_test.go` | 4 | GetFleetOrderInfoList, BatchGetFleetOrderInfoList |
| `apps/track/service/service_point_track_handler_test.go` | 2 | AppendSPDopReceiveUpdateFieldsNew |

**第三阶段：本地测试执行**

```bash
go test -gcflags=all=-l -v ./apps/order/schema/...
go test -gcflags=all=-l -v ./apps/order/service/...
go test -gcflags=all=-l -v ./apps/track/service/...
```

结果：✅ **所有测试本地通过**

```
ok      fleetorder/apps/order/schema    10.070s
ok      fleetorder/apps/order/service   7.206s
ok      fleetorder/apps/track/service   6.914s
```

**第四阶段：DMS 执行失败** ❌

推送代码到远程仓库后，DMS 执行测试时出现 panic：

```
--- FAIL: TestServicePointGrpcSchema_GetFleetOrderInfoList_ParseAddressSuccess (0.08s)
panic: InitConfig error,CHASSIS_CONF_DIR:/workspace/gitlab/fleetorder-service/conf/http_server/SPX-FleetOrderService
```

#### 5.2.3 问题分析

| 对比项 | 本地环境 | DMS 环境 |
|-------|---------|---------|
| 操作系统 | macOS (darwin) | Linux |
| 工作目录 | `/Users/xxx/code/fleetorder-service` | `/workspace/gitlab/fleetorder-service-xxx-for-run` |
| 配置文件 | ✅ 存在 | ❌ 不存在或路径不匹配 |
| `lib.InitTestChassis()` | ✅ 初始化成功 | ❌ panic |

**根本原因**：

`lib.InitTestChassis()` 函数通过以下逻辑查找配置文件：

```go
home, _ := os.Getwd()
index := strings.Index(home, "/fleetorder-service")
if index > 0 {
    home = SubString(home, 0, index) + "/fleetorder-service"
}
chassisPath := path.Join(home, "/conf/http_server/SPX-FleetOrderService")
```

在 DMS 环境中，工作目录是 `/workspace/gitlab/fleetorder-service-46c33d4ec0df063ed1c48e02464dde2002a140c9-for-run`，配置文件路径查找逻辑失败。

#### 5.2.4 解决方案（待实施）

**方案一：修改测试初始化逻辑**

避免在测试中直接调用 `lib.InitTestChassis()`，改为只初始化必要的配置：

```go
// ❌ 避免：直接调用 InitTestChassis
lib.InitTestChassis()

// ✅ 推荐：只初始化必要配置
lib.CommonConfig = &lib.ConfigOptions{IConfig: lib.ConfigCenter{}}
```

**方案二：修改 lib.InitTestChassis() 容错逻辑**

修改 `lib/test_env_config.go`，使其在初始化失败时不 panic：

```go
err := chassis.Init()
if err != nil {
    // 不再 panic，改为降级处理
    fmt.Printf("WARNING: InitConfig error, using fallback config\n")
    CommonConfig = &ConfigOptions{IConfig: ConfigCenter{}}
    return
}
```

#### 5.2.5 经验总结

| 要点 | 说明 |
|-----|------|
| 🔴 **环境差异** | 本地测试通过不代表 DMS 一定通过，需注意环境差异 |
| 🔴 **配置依赖** | 测试应尽量减少对外部配置文件的依赖 |
| 🟡 **初始化函数** | 使用 `lib.InitTestChassis()` 前需确认 DMS 兼容性 |
| 🟢 **Mock 策略** | 通过完善的 Mock 可以避免对真实配置的依赖 |

---

## 6. 最佳实践

### 6.1 测试文件位置规范

| 规则 | 说明 |
|-----|------|
| **同目录原则** | 测试文件必须与被测代码在同一目录 |
| **同包原则** | 测试文件的 package 必须与被测代码相同 |
| **命名规范** | 测试文件命名: `{source_file}_test.go` |

### 6.2 Mock 策略

#### 推荐的 Mock 层级

```
┌─────────────────────────────────────────┐
│           被测函数                       │
├─────────────────────────────────────────┤
│  ✅ Mock 这一层: 直接依赖的方法             │
├─────────────────────────────────────────┤
│  ❌ 不要 Mock: 底层实现细节                │
└─────────────────────────────────────────┘
```

#### Mock 示例

```go
// ✅ 推荐：Mock Facade 的公开方法
patches.ApplyMethod(reflect.TypeOf(facade), "GetLowestAddressIdNameMapping",
    func(_ *SiteNetworkRouteFacade, ...) (map[uint64]string, error) {
        return mockData, nil
    })

// ❌ 避免：Mock 私有类型的底层方法
patches.ApplyMethod(reflect.TypeOf(client.Grpc), "Invoke", ...)
```

### 6.3 测试用例覆盖原则

遵循 **2-2-1** 原则：

| 场景类型 | 最少数量 | 说明 |
|---------|---------|------|
| 成功场景 | ≥2 | 正常数据 + 边界数据 |
| 失败场景 | ≥2 | 参数错误 + 外部依赖失败 |
| 边界场景 | ≥1 | 空值/零值/极值 |

---

## 7. 常见问题与解决方案

### 7.1 自证预言测试（Self-fulfilling Prophecy Tests）

**问题现象**：
- 测试通过但覆盖率为 0%
- DMS 日志显示 `coverage: [no statements]`

**问题原因**：
- 测试文件与被测代码不在同一个包
- 测试只验证了 Mock 设置，没有实际调用被测函数

**解决方案**：
1. 将测试文件移动到与被测代码相同的目录
2. 确保测试中实际调用了被测函数
3. 验证返回值而不仅仅是 Mock 是否被调用

### 7.2 Mock 私有类型失败

**问题现象**：
```
panic: retrieve method by name failed
```

**问题原因**：
- 尝试 Mock 的类型是私有的（小写开头）
- `ApplyMethodReturn` 不支持私有类型

**解决方案**：
1. 改为 Mock 公开的 Facade 方法
2. 或在被测包内创建 Mock 对象

### 7.3 Chassis 初始化 Panic

**问题现象**：
```
panic: InitConfig error, CHASSIS_CONF_DIR: ...
```

**问题原因**：
- 测试环境缺少 Chassis 配置文件

**解决方案**：
1. Mock 依赖 Chassis 的方法
2. 或使用 `lib.InitTestChassis()` 初始化测试环境

### 7.4 本地测试通过但 DMS 执行失败

**问题现象**：
- 本地执行 `go test` 全部通过
- 推送到远程后，DMS 执行时测试 panic 或失败

**典型错误信息**：
```
panic: InitConfig error,CHASSIS_CONF_DIR:/workspace/gitlab/xxx-service/conf/http_server/SPX-XXXService

[chassis_test_vn.yaml] not exist
[chassis_test.yaml] not exist
[chassis.yaml] not exist
```

**问题原因**：

| 原因类型 | 详细说明 |
|---------|---------|
| **目录结构差异** | DMS 环境的工作目录结构与本地不同，通常格式为 `/workspace/gitlab/{repo}-{hash}-for-run/` |
| **配置文件缺失** | DMS 环境中可能没有复制项目的配置文件目录 |
| **路径查找逻辑** | `lib.InitTestChassis()` 的路径查找逻辑假设目录名包含 `/fleetorder-service`，但 DMS 目录名带有 hash 后缀 |
| **操作系统差异** | 本地是 macOS/Windows，DMS 是 Linux |

**解决方案**：

**方案一（推荐）：避免使用 InitTestChassis**

```go
// ❌ 不要这样做
func TestXxx(t *testing.T) {
    lib.InitTestChassis()  // 在 DMS 中会 panic
    // ...
}

// ✅ 推荐做法：直接初始化必要的配置
func TestXxx(t *testing.T) {
    lib.CommonConfig = &lib.ConfigOptions{IConfig: lib.ConfigCenter{}}
    // 然后 Mock 所有需要的方法
    patches := gomonkey.NewPatches()
    defer patches.Reset()
    // ...
}
```

**方案二：添加 DMS 环境检测**

```go
import "runtime"

// inDms 检测是否在 DMS 环境中运行
func inDms() bool {
    return runtime.GOOS != "darwin" && runtime.GOOS != "windows"
}

func TestXxx(t *testing.T) {
    if inDms() {
        t.Skip("Skip in DMS environment: requires chassis config")
    }
    lib.InitTestChassis()
    // ...
}
```

**方案三：完善 Mock 覆盖**

如果测试必须调用 `lib.InitTestChassis()`，确保所有依赖 Chassis 配置的方法都被 Mock：

```go
func TestXxx(t *testing.T) {
    patches := gomonkey.NewPatches()
    defer patches.Reset()
    
    // Mock chassis.Init 避免初始化失败
    patches.ApplyFunc(chassis.Init, func(...chassis.Option) error {
        return nil
    })
    
    // 手动初始化必要配置
    lib.CommonConfig = &lib.ConfigOptions{IConfig: lib.ConfigCenter{}}
    
    // ... 其他测试逻辑
}
```

**最佳实践**：

| 实践 | 说明 |
|-----|------|
| ✅ 完善 Mock | 尽量 Mock 所有外部依赖，减少对真实配置的依赖 |
| ✅ 环境无关 | 测试代码应尽量做到环境无关 |
| ✅ 最小依赖 | 只初始化测试必需的最小配置集 |
| ⚠️ 谨慎使用 | 谨慎使用 `lib.InitTestChassis()` 等全局初始化函数 |

### 7.5 包级初始化导致测试 panic

**问题现象**：
- 在同一个包中添加新测试后，整个包的测试都失败
- panic 发生在 `init()` 函数或包级变量初始化中

**典型场景**：

```go
// track_service_test.go
var trackService = initService()  // 包级变量在导入时初始化

func initService() *TrackService {
    lib.InitTestChassis()  // 如果这里 panic，整个包的测试都会失败
    // ...
}
```

**问题原因**：
- Go 的包级变量在包被导入时立即初始化
- 如果包中任一测试文件有包级初始化逻辑，会影响整个包的所有测试
- 即使新添加的测试不依赖该初始化逻辑，也会被影响

**解决方案**：

1. **避免在包级变量中调用可能失败的初始化函数**

```go
// ❌ 避免
var trackService = initService()

// ✅ 推荐：延迟初始化
var trackService *TrackService
var initOnce sync.Once

func getTrackService() *TrackService {
    initOnce.Do(func() {
        if !inDms() {
            trackService = initService()
        }
    })
    return trackService
}
```

2. **将测试文件拆分到不同子包**

避免新测试受现有包级初始化的影响。

---

## 8. 优化建议

### 8.1 提示词优化

#### 8.1.1 避免自证预言测试

**建议添加到 AI 提示词**：

```markdown
## 关键测试规范

1. **测试文件位置**：测试文件必须与被测代码在同一目录下，使用相同的 package 名

2. **禁止自证预言测试**：
   - ❌ 禁止：测试只验证 Mock 是否被设置
   - ✅ 必须：实际调用被测函数并验证返回值
   
3. **Mock 层级选择**：
   - ✅ 优先 Mock 被测函数的直接依赖
   - ❌ 避免 Mock 私有类型的底层实现
```

#### 8.1.2 增强测试生成质量

```markdown
## 测试生成检查清单

在生成测试前，AI 必须确认：
- [ ] 测试文件与被测代码在同一目录
- [ ] 使用 `t.Run()` 执行子测试
- [ ] 每个测试用例都实际调用被测函数
- [ ] 断言验证返回值而非 Mock 调用
- [ ] 包含至少 2 个成功、2 个失败、1 个边界场景
```

#### 8.1.3 DMS 环境兼容性检查

```markdown
## DMS 环境兼容性检查清单

在生成测试前，AI 必须检查并确认：
- [ ] 测试是否依赖 `lib.InitTestChassis()` 或类似的全局初始化函数
- [ ] 如果依赖，是否可以通过 Mock 替代
- [ ] 如果无法替代，是否添加了 DMS 环境跳过逻辑
- [ ] 测试是否依赖本地文件系统路径
- [ ] 测试是否依赖环境变量（如 CHASSIS_CONF_DIR）

## 推荐的初始化模式

❌ 避免使用全局初始化函数：
```go
lib.InitTestChassis()  // 可能在 DMS 中失败
```

✅ 推荐使用轻量级初始化：
```go
lib.CommonConfig = &lib.ConfigOptions{IConfig: lib.ConfigCenter{}}
// 配合完善的 Mock
```

### 8.2 工具流程优化

#### 8.2.1 当前流程问题

```
当前流程:
验证参数 → 获取覆盖率 → 生成测试 → 执行测试 → 推送代码 → [结束]
                                                    ↑
                                               缺少 DMS 触发和验证
```

#### 8.2.2 建议优化后的流程

```
优化流程:
验证参数 → 获取覆盖率 → 生成测试 → 执行测试 → 推送代码 
                                            ↓
生成报告 ← 获取覆盖率结果 ← 等待 DMS 完成 ← 触发 DMS
```

#### 8.2.3 具体优化建议

| 优化项 | 当前状态 | 建议改进 |
|-------|---------|---------|
| DMS 触发 | 需手动触发 | 自动触发 DMS 任务 |
| 覆盖率验证 | 仅本地验证 | 等待 DMS 结果并验证 |
| 失败重试 | 无 | 测试失败时自动修复重试 |
| 进度显示 | 简单输出 | 添加进度条和状态提示 |
| 报告生成 | 基础信息 | 增加详细的代码块覆盖对比 |

### 8.3 技术改进建议

#### 8.3.1 提升 Mock 能力

**问题**：无法 Mock 私有类型（如 `*grpcClient`）

**建议**：
1. 在 `lib/client` 包内添加测试辅助函数
2. 为 `grpcClient` 添加接口抽象

```go
// 建议添加到 lib/client/grpc_client.go
type GrpcInvoker interface {
    Invoke(ctx context.Context, serviceName, schemaID, operationID string, in, out interface{}) error
}

var Grpc GrpcInvoker = newGrpcClient()
```

#### 8.3.2 改进测试模板

**当前模板问题**：
- 通用性不够，需要针对不同代码结构调整

**建议改进**：
1. 添加更多测试模板类型（CRUD、RPC、事件处理等）
2. 根据函数签名自动选择合适的模板
3. 支持自定义模板配置

#### 8.3.3 增强 DMS 环境兼容性

**问题**：本地测试通过但 DMS 执行失败

**现状分析**：

```
┌────────────────────────────────────────────────────────────────────┐
│                       测试执行环境对比                                │
├──────────────────┬──────────────────────┬──────────────────────────┤
│       项目        │       本地环境        │        DMS 环境           │
├──────────────────┼──────────────────────┼──────────────────────────┤
│ 操作系统          │ macOS/Windows        │ Linux                     │
│ 工作目录          │ /Users/xxx/repo      │ /workspace/gitlab/repo-   │
│                  │                      │ {hash}-for-run/          │
│ 配置文件          │ ✅ 完整存在            │ ❌ 可能缺失                │
│ 环境变量          │ 用户可控               │ DMS 预设                 │
│ 网络访问          │ 完整                  │ 受限                      │
└──────────────────┴──────────────────────┴──────────────────────────┘
```

**建议改进**：

**1. 改进 `lib.InitTestChassis()` 容错能力**

```go
// 建议修改 lib/test_env_config.go
func initTestChassisOnce() {
    // ... 现有路径查找逻辑 ...
    
    err := chassis.Init()
    if err != nil {
        // 不再 panic，改为降级处理
        fmt.Printf("WARNING: InitConfig error (CHASSIS_CONF_DIR:%s)\n", chassisPath)
        fmt.Println("Using fallback configuration for testing...")
        
        // 初始化基本配置
        CommonConfig = &ConfigOptions{IConfig: ConfigCenter{}}
        return  // 优雅退出而非 panic
    }
}
```

**2. 添加 DMS 环境检测辅助函数**

```go
// 建议添加到 lib/test_helper.go
package lib

import "runtime"

// InDmsEnvironment 检测是否在 DMS 环境中运行
func InDmsEnvironment() bool {
    return runtime.GOOS != "darwin" && runtime.GOOS != "windows"
}

// SafeInitTestChassis 安全初始化，在 DMS 环境中不 panic
func SafeInitTestChassis() error {
    if InDmsEnvironment() {
        CommonConfig = &ConfigOptions{IConfig: ConfigCenter{}}
        return nil
    }
    return InitTestChassisWithError()
}
```

**3. 在 Code Guardian 工具中添加环境检测**

建议在 `generate_test_file` 工具中：
- 自动检测被测代码是否依赖 Chassis 初始化
- 如果依赖，生成的测试代码自动添加环境检测逻辑
- 提供警告提示用户潜在的 DMS 兼容性问题

**4. 添加 DMS 预检功能**

建议在 `execute_test` 工具中：
- 模拟 DMS 环境（如设置 GOOS=linux）运行测试
- 提前发现可能在 DMS 中失败的测试
- 在本地阶段就给出修复建议

### 8.4 用户体验优化

| 优化项 | 描述 |
|-------|------|
| **交互式确认** | 在关键步骤前询问用户确认 |
| **进度可视化** | 显示当前执行步骤和整体进度 |
| **错误提示** | 提供更清晰的错误信息和修复建议 |
| **历史记录** | 保存执行历史，支持查看和对比 |
| **批量处理** | 支持同时处理多个代码块 |

---

## 9. 附录

### 9.1 相关文档链接

- [项目单元测试规范](/.cursor/rules/unit-test/)
- [Go 常见陷阱指南](/.cursor/rules/knowledge/go-common-pitfalls.mdc)
- [gomonkey 使用指南](/.cursor/rules/unit-test/knowledge/gomonkey-runtime-errors.mdc)

### 9.2 版本历史

| 版本 | 日期 | 更新内容 |
|-----|------|---------|
| v1.0 | 2026-01-13 | 初始版本发布 |
| v1.1 | 2026-01-13 | 添加 fleetorder-service 使用案例（问题案例）；新增 DMS 环境兼容性问题说明；补充本地测试通过但 DMS 失败的解决方案；添加 DMS 环境兼容性优化建议 |

### 9.3 反馈与支持

如有问题或建议，请联系：
- **Slack**: #spx-fulfillment-dev
- **邮箱**: spx-fulfillment-team@shopee.com

---

> 📌 **注意**: 本工具持续迭代优化中，如遇到问题请及时反馈，帮助我们改进工具质量。
