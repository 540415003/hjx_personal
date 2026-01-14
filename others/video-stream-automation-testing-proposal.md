# 视频监控系统自动化测试方案

> **文档版本**: v1.0  
> **创建日期**: 2026-01-14  
> **作者**: Video Monitor Service Team  
> **状态**: 待评审

---

## 目录

- [1. 项目背景与目标](#1-项目背景与目标)
- [2. 现状分析](#2-现状分析)
- [3. 技术方案设计](#3-技术方案设计)
- [4. 详细测试方案](#4-详细测试方案)
- [5. 系统架构设计](#5-系统架构设计)
- [6. 实施计划](#6-实施计划)
- [7. 预期收益](#7-预期收益)
- [8. 风险评估与应对](#8-风险评估与应对)
- [9. 附录](#9-附录)

---

## 1. 项目背景与目标

### 1.1 项目背景

视频监控系统作为物流场站的核心基础设施，承担着实时监控、安全防护、业务追溯等关键职能。当前系统已接入多种品牌的视频管理系统（VMS），包括海康威视、大华、宇视、群晖、Viettel 等，管理着大量摄像头设备。

随着业务规模的扩大，人工巡检摄像头状态的方式已无法满足运维需求，存在以下痛点：

| 问题类型 | 具体表现 | 业务影响 |
|---------|---------|---------|
| **故障发现滞后** | 摄像头离线/黑屏等问题依赖人工发现 | 关键时段监控盲区 |
| **排查效率低** | 逐个检查摄像头状态耗时耗力 | 运维人力成本高 |
| **回放验证困难** | 无法自动验证录像完整性 | 事件追溯时发现录像缺失 |
| **质量无监控** | 画面模糊、卡顿等问题无感知 | 监控实效性降低 |

### 1.2 项目目标

建立一套**完整的视频流自动化测试体系**，实现：

1. **自动化健康检查**：7×24 小时自动检测摄像头连通性与流可用性
2. **智能异常检测**：自动识别黑屏、卡顿、断流等常见问题
3. **录像完整性验证**：定时验证回放录像的可用性与完整性
4. **可视化监控看板**：实时展示摄像头健康状态与关键指标
5. **自动告警通知**：异常情况及时推送至运维人员

### 1.3 关键指标（KPI）

| 指标名称 | 目标值 | 说明 |
|---------|-------|------|
| 故障发现时间 | ≤ 5 分钟 | 从故障发生到系统检测到的时间 |
| 摄像头覆盖率 | 100% | 所有在线摄像头纳入自动化测试 |
| 检测准确率 | ≥ 99% | 正确识别异常的比例 |
| 误报率 | ≤ 1% | 误报次数占总告警的比例 |
| 测试执行频率 | 每 5 分钟 | 单轮全量检测的周期 |

---

## 2. 现状分析

### 2.1 现有系统架构

当前 `video-monitor-service` 已具备以下核心能力：

```
┌─────────────────────────────────────────────────────────────────┐
│                     Video Monitor Service                        │
├─────────────────────────────────────────────────────────────────┤
│  Application Layer                                               │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐           │
│  │  Video   │ │  Camera  │ │Connection│ │   VMS    │           │
│  │ Service  │ │ Service  │ │ Service  │ │ Service  │           │
│  └──────────┘ └──────────┘ └──────────┘ └──────────┘           │
├─────────────────────────────────────────────────────────────────┤
│  Infrastructure Layer                                            │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │              CCTV Monitor (多品牌适配层)                   │   │
│  │  ┌─────┐ ┌─────┐ ┌────────┐ ┌────────┐ ┌───────┐        │   │
│  │  │ HIK │ │Dahua│ │Synology│ │ Uniview│ │Viettel│        │   │
│  │  └─────┘ └─────┘ └────────┘ └────────┘ └───────┘        │   │
│  └──────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────┘
```

### 2.2 已有能力盘点

| 模块 | 能力 | 可复用性 |
|------|------|---------|
| `MonitorService.ConnectTest` | VMS 连接测试 | ✅ 直接复用 |
| `MonitorService.GetRealStreamURL` | 获取实时流地址 | ✅ 直接复用 |
| `MonitorService.GetReplayStreamURL` | 获取回放流地址 | ✅ 直接复用 |
| `GCPHealthCheck` | GCP 设备健康检查 | ⚠️ 参考实现 |
| `WorkerHealthCheck` | FFmpeg 任务健康检查 | ⚠️ 参考实现 |

### 2.3 现有测试能力评估

| 测试类型 | 当前状态 | 需补充内容 |
|---------|---------|-----------|
| VMS 连通性测试 | ✅ 已实现 | 需增加定时执行 |
| RTSP 流可用性测试 | ❌ 未实现 | 完整实现 |
| 回放流测试 | ❌ 未实现 | 完整实现 |
| 画面质量检测 | ❌ 未实现 | 完整实现 |
| 性能指标采集 | ❌ 未实现 | 完整实现 |

---

## 3. 技术方案设计

### 3.1 技术选型

#### 3.1.1 核心技术栈

| 技术组件 | 版本要求 | 用途 | 选型理由 |
|---------|---------|------|---------|
| **Go** | 1.21+ | 主测试框架 | 与现有系统一致，便于集成 |
| **FFmpeg** | 5.0+ | 视频流探测与解码 | 业界标准，支持所有流协议 |
| **FFprobe** | 5.0+ | 流信息分析 | FFmpeg 配套工具 |
| **Python** | 3.10+ | 图像分析脚本 | OpenCV 生态成熟 |
| **OpenCV** | 4.8+ | 画面质量分析 | 图像处理标准库 |
| **Prometheus** | 2.45+ | 指标采集存储 | 云原生标准 |
| **Grafana** | 10.0+ | 可视化看板 | 与 Prometheus 深度集成 |
| **Saturn** | - | 定时任务调度 | 现有系统已集成 |
| **SeaTalk** | - | 告警通知 | 现有系统已集成 |

#### 3.1.2 技术依赖关系

```
                    ┌─────────────────────┐
                    │   Saturn Scheduler  │
                    └──────────┬──────────┘
                               │
                               ▼
┌──────────────────────────────────────────────────────────────┐
│                    Stream Test Service                        │
│  ┌─────────────────┐  ┌─────────────────┐  ┌──────────────┐ │
│  │ Connectivity    │  │ Stream Probe    │  │ Quality      │ │
│  │ Tester          │  │ (FFprobe)       │  │ Analyzer     │ │
│  │                 │  │                 │  │ (OpenCV)     │ │
│  └────────┬────────┘  └────────┬────────┘  └──────┬───────┘ │
│           │                    │                   │         │
│           ▼                    ▼                   ▼         │
│  ┌─────────────────────────────────────────────────────────┐│
│  │                    Test Result Collector                 ││
│  └─────────────────────────────────────────────────────────┘│
└──────────────────────────────────────────────────────────────┘
           │                     │                    │
           ▼                     ▼                    ▼
   ┌───────────────┐    ┌───────────────┐    ┌───────────────┐
   │  Prometheus   │    │    MySQL      │    │   SeaTalk     │
   │  (Metrics)    │    │  (Results)    │    │   (Alert)     │
   └───────────────┘    └───────────────┘    └───────────────┘
           │
           ▼
   ┌───────────────┐
   │    Grafana    │
   │  (Dashboard)  │
   └───────────────┘
```

### 3.2 测试分层设计

测试体系分为四个层级，由浅入深：

```
Layer 4: 业务层测试 ──────────────────────────────────
         │ 端到端直播/回放场景验证
         │ 多用户并发访问测试
         ▼
Layer 3: 质量层测试 ──────────────────────────────────
         │ 黑屏/卡顿/模糊检测
         │ 画面异常识别
         ▼
Layer 2: 流层测试 ────────────────────────────────────
         │ RTSP/HLS/FLV 流可用性验证
         │ 帧率/分辨率/码率检测
         ▼
Layer 1: 连接层测试 ──────────────────────────────────
         │ VMS 连接测试
         │ 摄像头在线状态检测
         │ 网络连通性验证
```

### 3.3 核心流程设计

#### 3.3.1 健康检查主流程

```
┌─────────────────────────────────────────────────────────────────┐
│                     定时任务触发 (每5分钟)                        │
└──────────────────────────┬──────────────────────────────────────┘
                           │
                           ▼
               ┌───────────────────────┐
               │ 获取待检测摄像头列表    │
               │ (分页 + 并发控制)      │
               └───────────┬───────────┘
                           │
           ┌───────────────┼───────────────┐
           ▼               ▼               ▼
   ┌───────────────┐ ┌───────────────┐ ┌───────────────┐
   │ Worker 1      │ │ Worker 2      │ │ Worker N      │
   │ (Camera Pool) │ │ (Camera Pool) │ │ (Camera Pool) │
   └───────┬───────┘ └───────┬───────┘ └───────┬───────┘
           │                 │                 │
           ▼                 ▼                 ▼
   ┌─────────────────────────────────────────────────────┐
   │                   单摄像头测试流程                    │
   │  ┌─────────┐   ┌─────────┐   ┌─────────┐           │
   │  │连接测试  │ → │ 流探测   │ → │质量分析  │           │
   │  │         │   │(FFprobe)│   │(OpenCV) │           │
   │  └─────────┘   └─────────┘   └─────────┘           │
   └──────────────────────┬──────────────────────────────┘
                          │
                          ▼
              ┌───────────────────────┐
              │     汇总测试结果        │
              └───────────┬───────────┘
                          │
          ┌───────────────┼───────────────┐
          ▼               ▼               ▼
   ┌───────────┐   ┌───────────┐   ┌───────────┐
   │ 写入数据库 │   │上报Metrics│   │ 触发告警   │
   │  (MySQL)  │   │(Prometheus)│  │(SeaTalk)  │
   └───────────┘   └───────────┘   └───────────┘
```

#### 3.3.2 单摄像头测试流程

```
                    ┌─────────────────┐
                    │ 开始测试摄像头   │
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │ 1. VMS连接测试   │
                    │ ConnectTest()   │
                    └────────┬────────┘
                             │
              ┌──────────────┴──────────────┐
              │                             │
         [成功]                          [失败]
              │                             │
              ▼                             ▼
     ┌─────────────────┐          ┌─────────────────┐
     │ 2. 获取RTSP地址  │          │ 记录连接失败     │
     │GetRealStreamURL │          │ 标记为离线       │
     └────────┬────────┘          └────────┬────────┘
              │                             │
              ▼                             │
     ┌─────────────────┐                    │
     │ 3. FFprobe探测   │                    │
     │ 超时: 10秒       │                    │
     └────────┬────────┘                    │
              │                             │
       ┌──────┴──────┐                      │
       │             │                      │
  [成功]          [失败]                    │
       │             │                      │
       ▼             ▼                      │
┌────────────┐ ┌────────────┐              │
│解析流信息   │ │记录流异常   │              │
│帧率/分辨率  │ │            │              │
└─────┬──────┘ └─────┬──────┘              │
      │              │                      │
      ▼              │                      │
┌────────────┐       │                      │
│4. 质量检测  │       │                      │
│(可选/采样)  │       │                      │
└─────┬──────┘       │                      │
      │              │                      │
      ▼              ▼                      ▼
┌─────────────────────────────────────────────┐
│              汇总测试结果                    │
│  TestResult {                               │
│    CameraID, Status, Latency,               │
│    FrameRate, Resolution, ErrorMsg...       │
│  }                                          │
└─────────────────────────────────────────────┘
```

---

## 4. 详细测试方案

### 4.1 Layer 1: 连接层测试

#### 4.1.1 VMS 连接测试

**测试目的**: 验证视频管理系统的可达性与认证有效性

**测试方法**:
- 调用各品牌 VMS 的 `ConnectTest` 接口
- 验证账号密码有效性
- 检测网络连通性

**实现代码**:

```go
// domain/stream_test/connectivity_tester.go

package stream_test

import (
    "context"
    "time"
    
    "video-monitor-service/infrastructure/thirdparty/cctv_monitor"
    "video-monitor-service/infrastructure/thirdparty/cctv_monitor/define"
)

// ConnectivityTester VMS连接测试器
type ConnectivityTester struct {
    MonitorService cctv_monitor.MonitorService
    Timeout        time.Duration
}

// ConnectivityTestResult 连接测试结果
type ConnectivityTestResult struct {
    VMSID        string        `json:"vms_id"`
    Success      bool          `json:"success"`
    Latency      time.Duration `json:"latency_ms"`
    ErrorCode    string        `json:"error_code,omitempty"`
    ErrorMessage string        `json:"error_message,omitempty"`
    TestedAt     time.Time     `json:"tested_at"`
}

// TestVMSConnectivity 测试VMS连接
func (t *ConnectivityTester) TestVMSConnectivity(
    ctx context.Context,
    vmsInfo *VMSInfo,
) ConnectivityTestResult {
    result := ConnectivityTestResult{
        VMSID:    vmsInfo.VMSID,
        TestedAt: time.Now(),
    }
    
    ctx, cancel := context.WithTimeout(ctx, t.Timeout)
    defer cancel()
    
    startTime := time.Now()
    
    isConnected, err := t.MonitorService.ConnectTest(ctx, define.ConnectTestReq{
        UserInfo: define.UserInfo{
            Account:  vmsInfo.Account,
            Password: vmsInfo.Password,
            IP:       vmsInfo.Host,
            Port:     vmsInfo.Port,
        },
    })
    
    result.Latency = time.Since(startTime)
    
    if err != nil {
        result.Success = false
        result.ErrorCode = "CONNECTION_ERROR"
        result.ErrorMessage = err.Error()
        return result
    }
    
    if !isConnected {
        result.Success = false
        result.ErrorCode = "AUTH_FAILED"
        result.ErrorMessage = "VMS authentication failed"
        return result
    }
    
    result.Success = true
    return result
}
```

**判定标准**:

| 结果 | 判定条件 | 处理方式 |
|------|---------|---------|
| 成功 | `ConnectTest` 返回 true | 进入下一步测试 |
| 认证失败 | 返回认证错误 | 告警 + 检查密码 |
| 网络超时 | 10秒内无响应 | 告警 + 检查网络 |
| 拒绝连接 | TCP连接被拒绝 | 告警 + 检查设备状态 |

---

#### 4.1.2 摄像头在线状态检测

**测试目的**: 验证具体摄像头的在线状态

**测试方法**:
- 查询 VMS 获取通道列表
- 检查目标摄像头的 `ChannelStatus`

**实现代码**:

```go
// CheckCameraOnlineStatus 检测摄像头在线状态
func (t *ConnectivityTester) CheckCameraOnlineStatus(
    ctx context.Context,
    vmsID string,
    channelID int,
) (*CameraStatusResult, error) {
    channels, err := t.MonitorService.GetChannelList(ctx, vmsID)
    if err != nil {
        return nil, fmt.Errorf("获取通道列表失败: %w", err)
    }
    
    for _, ch := range channels {
        if ch.ChannelID == channelID {
            return &CameraStatusResult{
                ChannelID:   ch.ChannelID,
                IsOnline:    ch.ChannelStatus == 1,
                CameraModel: ch.CameraModel,
            }, nil
        }
    }
    
    return nil, fmt.Errorf("未找到通道: %d", channelID)
}
```

---

### 4.2 Layer 2: 流层测试

#### 4.2.1 实时流可用性测试

**测试目的**: 验证 RTSP/RTMP 等实时视频流的可用性

**测试方法**:
- 获取 RTSP 流地址
- 使用 FFprobe 探测流信息
- 验证视频/音频轨道存在性

**实现代码**:

```go
// domain/stream_test/stream_prober.go

package stream_test

import (
    "context"
    "encoding/json"
    "fmt"
    "os/exec"
    "strconv"
    "strings"
    "time"
)

// StreamProber 流探测器
type StreamProber struct {
    FFprobePath string
    Timeout     time.Duration
}

// FFProbeResult FFprobe输出结构
type FFProbeResult struct {
    Streams []StreamInfo `json:"streams"`
    Format  FormatInfo   `json:"format"`
}

// StreamInfo 流信息
type StreamInfo struct {
    Index          int    `json:"index"`
    CodecType      string `json:"codec_type"`      // video / audio
    CodecName      string `json:"codec_name"`      // h264 / hevc / aac
    Width          int    `json:"width,omitempty"`
    Height         int    `json:"height,omitempty"`
    RFrameRate     string `json:"r_frame_rate"`    // 如 "25/1"
    AvgFrameRate   string `json:"avg_frame_rate"`
    BitRate        string `json:"bit_rate,omitempty"`
    Profile        string `json:"profile,omitempty"`
}

// FormatInfo 格式信息
type FormatInfo struct {
    FormatName string `json:"format_name"`
    Duration   string `json:"duration,omitempty"`
    BitRate    string `json:"bit_rate,omitempty"`
    ProbeScore int    `json:"probe_score"`
}

// StreamProbeResult 流探测结果
type StreamProbeResult struct {
    URL           string        `json:"url"`
    Success       bool          `json:"success"`
    HasVideo      bool          `json:"has_video"`
    HasAudio      bool          `json:"has_audio"`
    VideoCodec    string        `json:"video_codec,omitempty"`
    AudioCodec    string        `json:"audio_codec,omitempty"`
    Width         int           `json:"width,omitempty"`
    Height        int           `json:"height,omitempty"`
    FrameRate     float64       `json:"frame_rate,omitempty"`
    BitRate       int64         `json:"bit_rate,omitempty"`
    Latency       time.Duration `json:"latency_ms"`
    ErrorCode     string        `json:"error_code,omitempty"`
    ErrorMessage  string        `json:"error_message,omitempty"`
    TestedAt      time.Time     `json:"tested_at"`
}

// ProbeStream 探测RTSP流
func (p *StreamProber) ProbeStream(ctx context.Context, rtspURL string) StreamProbeResult {
    result := StreamProbeResult{
        URL:      rtspURL,
        TestedAt: time.Now(),
    }
    
    ctx, cancel := context.WithTimeout(ctx, p.Timeout)
    defer cancel()
    
    startTime := time.Now()
    
    // 构建FFprobe命令
    args := []string{
        "-v", "quiet",
        "-print_format", "json",
        "-show_streams",
        "-show_format",
        "-rtsp_transport", "tcp",        // 使用TCP传输，更可靠
        "-analyzeduration", "3000000",   // 分析时长3秒
        "-probesize", "5000000",         // 探测大小5MB
        "-i", rtspURL,
    }
    
    cmd := exec.CommandContext(ctx, p.FFprobePath, args...)
    output, err := cmd.Output()
    
    result.Latency = time.Since(startTime)
    
    if err != nil {
        result.Success = false
        if ctx.Err() == context.DeadlineExceeded {
            result.ErrorCode = "TIMEOUT"
            result.ErrorMessage = "流探测超时"
        } else {
            result.ErrorCode = "PROBE_ERROR"
            result.ErrorMessage = fmt.Sprintf("FFprobe执行失败: %v", err)
        }
        return result
    }
    
    // 解析FFprobe输出
    var probeResult FFProbeResult
    if err := json.Unmarshal(output, &probeResult); err != nil {
        result.Success = false
        result.ErrorCode = "PARSE_ERROR"
        result.ErrorMessage = fmt.Sprintf("解析FFprobe输出失败: %v", err)
        return result
    }
    
    // 分析流信息
    for _, stream := range probeResult.Streams {
        switch stream.CodecType {
        case "video":
            result.HasVideo = true
            result.VideoCodec = stream.CodecName
            result.Width = stream.Width
            result.Height = stream.Height
            result.FrameRate = parseFrameRate(stream.RFrameRate)
        case "audio":
            result.HasAudio = true
            result.AudioCodec = stream.CodecName
        }
    }
    
    // 解析码率
    if probeResult.Format.BitRate != "" {
        result.BitRate, _ = strconv.ParseInt(probeResult.Format.BitRate, 10, 64)
    }
    
    // 验证视频流是否存在
    if !result.HasVideo {
        result.Success = false
        result.ErrorCode = "NO_VIDEO_STREAM"
        result.ErrorMessage = "未检测到视频流"
        return result
    }
    
    result.Success = true
    return result
}

// parseFrameRate 解析帧率 (如 "25/1" -> 25.0)
func parseFrameRate(rateStr string) float64 {
    parts := strings.Split(rateStr, "/")
    if len(parts) != 2 {
        return 0
    }
    
    num, _ := strconv.ParseFloat(parts[0], 64)
    den, _ := strconv.ParseFloat(parts[1], 64)
    
    if den == 0 {
        return 0
    }
    
    return num / den
}
```

**判定标准**:

| 指标 | 正常范围 | 异常处理 |
|------|---------|---------|
| 探测时间 | ≤ 10秒 | 超时告警 |
| 视频流 | 必须存在 | 缺失告警 |
| 帧率 | ≥ 15 fps | 低于阈值告警 |
| 分辨率 | ≥ 640×480 | 低于阈值记录 |
| 编码格式 | H.264/H.265 | 格式记录 |

---

#### 4.2.2 回放流测试

**测试目的**: 验证历史录像回放功能正常

**测试方法**:
- 获取指定时间段的回放地址
- 验证回放流可用性
- 检查录像时间段连续性

**实现代码**:

```go
// PlaybackTestResult 回放测试结果
type PlaybackTestResult struct {
    CameraID      string        `json:"camera_id"`
    TimeRange     TimeRange     `json:"time_range"`
    Success       bool          `json:"success"`
    HasRecording  bool          `json:"has_recording"`
    RecordBlocks  []RecordBlock `json:"record_blocks,omitempty"`
    GapCount      int           `json:"gap_count"`
    TotalGapSec   int64         `json:"total_gap_seconds"`
    CoverageRate  float64       `json:"coverage_rate"`
    ErrorCode     string        `json:"error_code,omitempty"`
    ErrorMessage  string        `json:"error_message,omitempty"`
    TestedAt      time.Time     `json:"tested_at"`
}

// TimeRange 时间范围
type TimeRange struct {
    Start time.Time `json:"start"`
    End   time.Time `json:"end"`
}

// RecordBlock 录像段
type RecordBlock struct {
    Begin time.Time `json:"begin"`
    End   time.Time `json:"end"`
}

// TestPlayback 测试回放
func (p *StreamProber) TestPlayback(
    ctx context.Context,
    cameraID string,
    startTime, endTime time.Time,
    getPlaybackURL func(ctx context.Context, cameraID string, start, end time.Time) (string, []RecordBlock, error),
) PlaybackTestResult {
    result := PlaybackTestResult{
        CameraID: cameraID,
        TimeRange: TimeRange{
            Start: startTime,
            End:   endTime,
        },
        TestedAt: time.Now(),
    }
    
    // 1. 获取回放地址
    playbackURL, recordBlocks, err := getPlaybackURL(ctx, cameraID, startTime, endTime)
    if err != nil {
        result.Success = false
        result.ErrorCode = "GET_URL_FAILED"
        result.ErrorMessage = err.Error()
        return result
    }
    
    // 2. 分析录像覆盖率
    result.RecordBlocks = recordBlocks
    if len(recordBlocks) > 0 {
        result.HasRecording = true
        result.GapCount, result.TotalGapSec = analyzeRecordingGaps(recordBlocks, startTime, endTime)
        totalDuration := endTime.Sub(startTime).Seconds()
        coveredDuration := totalDuration - float64(result.TotalGapSec)
        result.CoverageRate = coveredDuration / totalDuration * 100
    }
    
    // 3. 验证回放流可用性
    probeResult := p.ProbeStream(ctx, playbackURL)
    if !probeResult.Success {
        result.Success = false
        result.ErrorCode = probeResult.ErrorCode
        result.ErrorMessage = probeResult.ErrorMessage
        return result
    }
    
    result.Success = true
    return result
}

// analyzeRecordingGaps 分析录像间隙
func analyzeRecordingGaps(blocks []RecordBlock, start, end time.Time) (gapCount int, totalGapSec int64) {
    if len(blocks) == 0 {
        return 1, int64(end.Sub(start).Seconds())
    }
    
    // 排序录像段
    sort.Slice(blocks, func(i, j int) bool {
        return blocks[i].Begin.Before(blocks[j].Begin)
    })
    
    // 计算间隙
    lastEnd := start
    for _, block := range blocks {
        if block.Begin.After(lastEnd) {
            gapCount++
            totalGapSec += int64(block.Begin.Sub(lastEnd).Seconds())
        }
        if block.End.After(lastEnd) {
            lastEnd = block.End
        }
    }
    
    // 检查结尾是否有间隙
    if lastEnd.Before(end) {
        gapCount++
        totalGapSec += int64(end.Sub(lastEnd).Seconds())
    }
    
    return gapCount, totalGapSec
}
```

**判定标准**:

| 指标 | 正常范围 | 异常处理 |
|------|---------|---------|
| 录像存在 | 必须存在 | 缺失告警 |
| 覆盖率 | ≥ 95% | 低于阈值告警 |
| 单次间隙 | ≤ 5分钟 | 大间隙告警 |

---

### 4.3 Layer 3: 质量层测试

#### 4.3.1 黑屏检测

**测试目的**: 识别摄像头画面全黑的异常情况

**测试方法**:
- 从视频流中抓取帧
- 计算画面平均亮度
- 低于阈值判定为黑屏

**实现代码** (Python):

```python
#!/usr/bin/env python3
# tools/video_quality_analyzer/black_screen_detector.py

import cv2  # <ai-gen>
import numpy as np  # <ai-gen>
import subprocess  # <ai-gen>
import json  # <ai-gen>
import sys  # <ai-gen>
from dataclasses import dataclass, asdict  # <ai-gen>
from typing import Optional, Tuple  # <ai-gen>
import tempfile  # <ai-gen>
import os  # <ai-gen>


@dataclass  # <ai-gen>
class BlackScreenResult:  # <ai-gen>
    """黑屏检测结果"""  # <ai-gen>
    is_black_screen: bool  # <ai-gen>
    mean_brightness: float  # <ai-gen>
    threshold: float  # <ai-gen>
    confidence: float  # <ai-gen>
    error: Optional[str] = None  # <ai-gen>


class BlackScreenDetector:  # <ai-gen>
    """黑屏检测器"""  # <ai-gen>
    
    def __init__(self, brightness_threshold: float = 10.0):  # <ai-gen>
        """  # <ai-gen>
        初始化黑屏检测器  # <ai-gen>
        
        Args:  # <ai-gen>
            brightness_threshold: 亮度阈值 (0-255), 低于此值判定为黑屏  # <ai-gen>
        """  # <ai-gen>
        self.brightness_threshold = brightness_threshold  # <ai-gen>
    
    def capture_frame_from_rtsp(  # <ai-gen>
        self,  # <ai-gen>
        rtsp_url: str,  # <ai-gen>
        timeout: int = 10  # <ai-gen>
    ) -> Tuple[Optional[np.ndarray], Optional[str]]:  # <ai-gen>
        """  # <ai-gen>
        从RTSP流抓取一帧  # <ai-gen>
        
        使用FFmpeg而非OpenCV的VideoCapture，更可靠  # <ai-gen>
        """  # <ai-gen>
        with tempfile.NamedTemporaryFile(suffix='.jpg', delete=False) as tmp:  # <ai-gen>
            tmp_path = tmp.name  # <ai-gen>
        
        try:  # <ai-gen>
            # 使用FFmpeg抓取单帧  # <ai-gen>
            cmd = [  # <ai-gen>
                'ffmpeg',  # <ai-gen>
                '-rtsp_transport', 'tcp',  # <ai-gen>
                '-i', rtsp_url,  # <ai-gen>
                '-frames:v', '1',  # <ai-gen>
                '-q:v', '2',  # <ai-gen>
                '-y',  # <ai-gen>
                tmp_path  # <ai-gen>
            ]  # <ai-gen>
            
            result = subprocess.run(  # <ai-gen>
                cmd,  # <ai-gen>
                capture_output=True,  # <ai-gen>
                timeout=timeout  # <ai-gen>
            )  # <ai-gen>
            
            if result.returncode != 0:  # <ai-gen>
                return None, f"FFmpeg failed: {result.stderr.decode()}"  # <ai-gen>
            
            # 读取抓取的帧  # <ai-gen>
            frame = cv2.imread(tmp_path)  # <ai-gen>
            if frame is None:  # <ai-gen>
                return None, "Failed to read captured frame"  # <ai-gen>
            
            return frame, None  # <ai-gen>
            
        except subprocess.TimeoutExpired:  # <ai-gen>
            return None, "Frame capture timeout"  # <ai-gen>
        except Exception as e:  # <ai-gen>
            return None, str(e)  # <ai-gen>
        finally:  # <ai-gen>
            if os.path.exists(tmp_path):  # <ai-gen>
                os.remove(tmp_path)  # <ai-gen>
    
    def detect(self, frame: np.ndarray) -> BlackScreenResult:  # <ai-gen>
        """  # <ai-gen>
        检测画面是否为黑屏  # <ai-gen>
        
        Args:  # <ai-gen>
            frame: BGR格式的图像帧  # <ai-gen>
            
        Returns:  # <ai-gen>
            BlackScreenResult: 检测结果  # <ai-gen>
        """  # <ai-gen>
        # 转换为灰度图  # <ai-gen>
        gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)  # <ai-gen>
        
        # 计算平均亮度  # <ai-gen>
        mean_brightness = float(np.mean(gray))  # <ai-gen>
        
        # 计算标准差，用于评估置信度  # <ai-gen>
        std_brightness = float(np.std(gray))  # <ai-gen>
        
        # 判定黑屏  # <ai-gen>
        is_black = mean_brightness < self.brightness_threshold  # <ai-gen>
        
        # 计算置信度: 标准差越小，置信度越高  # <ai-gen>
        # 真正的黑屏应该是均匀的黑色  # <ai-gen>
        if is_black:  # <ai-gen>
            confidence = max(0, 1 - (std_brightness / 50))  # <ai-gen>
        else:  # <ai-gen>
            confidence = min(1, mean_brightness / 100)  # <ai-gen>
        
        return BlackScreenResult(  # <ai-gen>
            is_black_screen=is_black,  # <ai-gen>
            mean_brightness=mean_brightness,  # <ai-gen>
            threshold=self.brightness_threshold,  # <ai-gen>
            confidence=confidence  # <ai-gen>
        )  # <ai-gen>
    
    def detect_from_rtsp(self, rtsp_url: str) -> BlackScreenResult:  # <ai-gen>
        """从RTSP流检测黑屏"""  # <ai-gen>
        frame, error = self.capture_frame_from_rtsp(rtsp_url)  # <ai-gen>
        
        if error:  # <ai-gen>
            return BlackScreenResult(  # <ai-gen>
                is_black_screen=False,  # <ai-gen>
                mean_brightness=0,  # <ai-gen>
                threshold=self.brightness_threshold,  # <ai-gen>
                confidence=0,  # <ai-gen>
                error=error  # <ai-gen>
            )  # <ai-gen>
        
        return self.detect(frame)  # <ai-gen>


def main():  # <ai-gen>
    """命令行入口"""  # <ai-gen>
    if len(sys.argv) < 2:  # <ai-gen>
        print("Usage: python black_screen_detector.py <rtsp_url>", file=sys.stderr)  # <ai-gen>
        sys.exit(1)  # <ai-gen>
    
    rtsp_url = sys.argv[1]  # <ai-gen>
    threshold = float(sys.argv[2]) if len(sys.argv) > 2 else 10.0  # <ai-gen>
    
    detector = BlackScreenDetector(brightness_threshold=threshold)  # <ai-gen>
    result = detector.detect_from_rtsp(rtsp_url)  # <ai-gen>
    
    print(json.dumps(asdict(result), indent=2))  # <ai-gen>
    
    sys.exit(0 if not result.is_black_screen else 1)  # <ai-gen>


if __name__ == '__main__':  # <ai-gen>
    main()  # <ai-gen>
```

#### 4.3.2 画面卡顿检测

**测试目的**: 识别视频流卡顿/冻结的异常情况

**测试方法**:
- 连续抓取多帧
- 计算帧间相似度
- 高相似度判定为卡顿

**实现代码** (Python):

```python
# tools/video_quality_analyzer/frozen_detector.py

import cv2  # <ai-gen>
import numpy as np  # <ai-gen>
from dataclasses import dataclass  # <ai-gen>
from typing import List, Optional  # <ai-gen>


@dataclass  # <ai-gen>
class FrozenFrameResult:  # <ai-gen>
    """画面卡顿检测结果"""  # <ai-gen>
    is_frozen: bool  # <ai-gen>
    similarity_scores: List[float]  # <ai-gen>
    avg_similarity: float  # <ai-gen>
    threshold: float  # <ai-gen>
    frozen_duration_frames: int  # <ai-gen>
    error: Optional[str] = None  # <ai-gen>


class FrozenFrameDetector:  # <ai-gen>
    """画面卡顿检测器"""  # <ai-gen>
    
    def __init__(  # <ai-gen>
        self,  # <ai-gen>
        similarity_threshold: float = 0.995,  # <ai-gen>
        min_frozen_frames: int = 3  # <ai-gen>
    ):  # <ai-gen>
        """  # <ai-gen>
        初始化卡顿检测器  # <ai-gen>
        
        Args:  # <ai-gen>
            similarity_threshold: 相似度阈值，超过此值判定为相同帧  # <ai-gen>
            min_frozen_frames: 连续相同帧数阈值，超过此值判定为卡顿  # <ai-gen>
        """  # <ai-gen>
        self.similarity_threshold = similarity_threshold  # <ai-gen>
        self.min_frozen_frames = min_frozen_frames  # <ai-gen>
    
    def calculate_similarity(  # <ai-gen>
        self,  # <ai-gen>
        frame1: np.ndarray,  # <ai-gen>
        frame2: np.ndarray  # <ai-gen>
    ) -> float:  # <ai-gen>
        """  # <ai-gen>
        计算两帧的相似度 (使用结构相似性)  # <ai-gen>
        """  # <ai-gen>
        # 转换为灰度图  # <ai-gen>
        gray1 = cv2.cvtColor(frame1, cv2.COLOR_BGR2GRAY)  # <ai-gen>
        gray2 = cv2.cvtColor(frame2, cv2.COLOR_BGR2GRAY)  # <ai-gen>
        
        # 确保尺寸一致  # <ai-gen>
        if gray1.shape != gray2.shape:  # <ai-gen>
            gray2 = cv2.resize(gray2, (gray1.shape[1], gray1.shape[0]))  # <ai-gen>
        
        # 计算归一化互相关  # <ai-gen>
        result = cv2.matchTemplate(gray1, gray2, cv2.TM_CCOEFF_NORMED)  # <ai-gen>
        
        return float(result[0][0])  # <ai-gen>
    
    def detect_from_frames(self, frames: List[np.ndarray]) -> FrozenFrameResult:  # <ai-gen>
        """  # <ai-gen>
        从帧列表检测卡顿  # <ai-gen>
        
        Args:  # <ai-gen>
            frames: 按时间顺序排列的帧列表  # <ai-gen>
            
        Returns:  # <ai-gen>
            FrozenFrameResult: 检测结果  # <ai-gen>
        """  # <ai-gen>
        if len(frames) < 2:  # <ai-gen>
            return FrozenFrameResult(  # <ai-gen>
                is_frozen=False,  # <ai-gen>
                similarity_scores=[],  # <ai-gen>
                avg_similarity=0,  # <ai-gen>
                threshold=self.similarity_threshold,  # <ai-gen>
                frozen_duration_frames=0,  # <ai-gen>
                error="需要至少2帧进行检测"  # <ai-gen>
            )  # <ai-gen>
        
        # 计算相邻帧的相似度  # <ai-gen>
        similarities = []  # <ai-gen>
        for i in range(1, len(frames)):  # <ai-gen>
            sim = self.calculate_similarity(frames[i-1], frames[i])  # <ai-gen>
            similarities.append(sim)  # <ai-gen>
        
        avg_similarity = np.mean(similarities)  # <ai-gen>
        
        # 统计连续高相似度帧数  # <ai-gen>
        consecutive_frozen = 0  # <ai-gen>
        max_frozen = 0  # <ai-gen>
        
        for sim in similarities:  # <ai-gen>
            if sim > self.similarity_threshold:  # <ai-gen>
                consecutive_frozen += 1  # <ai-gen>
                max_frozen = max(max_frozen, consecutive_frozen)  # <ai-gen>
            else:  # <ai-gen>
                consecutive_frozen = 0  # <ai-gen>
        
        is_frozen = max_frozen >= self.min_frozen_frames  # <ai-gen>
        
        return FrozenFrameResult(  # <ai-gen>
            is_frozen=is_frozen,  # <ai-gen>
            similarity_scores=similarities,  # <ai-gen>
            avg_similarity=float(avg_similarity),  # <ai-gen>
            threshold=self.similarity_threshold,  # <ai-gen>
            frozen_duration_frames=max_frozen  # <ai-gen>
        )  # <ai-gen>
    
    def detect_from_rtsp(  # <ai-gen>
        self,  # <ai-gen>
        rtsp_url: str,  # <ai-gen>
        sample_count: int = 10,  # <ai-gen>
        sample_interval_ms: int = 500  # <ai-gen>
    ) -> FrozenFrameResult:  # <ai-gen>
        """  # <ai-gen>
        从RTSP流检测卡顿  # <ai-gen>
        
        Args:  # <ai-gen>
            rtsp_url: RTSP流地址  # <ai-gen>
            sample_count: 采样帧数  # <ai-gen>
            sample_interval_ms: 采样间隔(毫秒)  # <ai-gen>
        """  # <ai-gen>
        cap = cv2.VideoCapture(rtsp_url)  # <ai-gen>
        
        if not cap.isOpened():  # <ai-gen>
            return FrozenFrameResult(  # <ai-gen>
                is_frozen=False,  # <ai-gen>
                similarity_scores=[],  # <ai-gen>
                avg_similarity=0,  # <ai-gen>
                threshold=self.similarity_threshold,  # <ai-gen>
                frozen_duration_frames=0,  # <ai-gen>
                error="无法打开视频流"  # <ai-gen>
            )  # <ai-gen>
        
        frames = []  # <ai-gen>
        try:  # <ai-gen>
            for _ in range(sample_count):  # <ai-gen>
                ret, frame = cap.read()  # <ai-gen>
                if ret:  # <ai-gen>
                    frames.append(frame)  # <ai-gen>
                cv2.waitKey(sample_interval_ms)  # <ai-gen>
        finally:  # <ai-gen>
            cap.release()  # <ai-gen>
        
        if len(frames) < 2:  # <ai-gen>
            return FrozenFrameResult(  # <ai-gen>
                is_frozen=False,  # <ai-gen>
                similarity_scores=[],  # <ai-gen>
                avg_similarity=0,  # <ai-gen>
                threshold=self.similarity_threshold,  # <ai-gen>
                frozen_duration_frames=0,  # <ai-gen>
                error="采样帧数不足"  # <ai-gen>
            )  # <ai-gen>
        
        return self.detect_from_frames(frames)  # <ai-gen>
```

#### 4.3.3 质量分析集成

**Go 调用 Python 脚本的封装**:

```go
// domain/stream_test/quality_analyzer.go

package stream_test

import (
    "context"
    "encoding/json"
    "fmt"
    "os/exec"
    "time"
)

// QualityAnalyzer 画面质量分析器
type QualityAnalyzer struct {
    PythonPath       string
    ScriptsDir       string
    Timeout          time.Duration
}

// BlackScreenResult 黑屏检测结果
type BlackScreenResult struct {
    IsBlackScreen    bool    `json:"is_black_screen"`
    MeanBrightness   float64 `json:"mean_brightness"`
    Threshold        float64 `json:"threshold"`
    Confidence       float64 `json:"confidence"`
    Error            string  `json:"error,omitempty"`
}

// FrozenFrameResult 卡顿检测结果
type FrozenFrameResult struct {
    IsFrozen             bool      `json:"is_frozen"`
    SimilarityScores     []float64 `json:"similarity_scores"`
    AvgSimilarity        float64   `json:"avg_similarity"`
    Threshold            float64   `json:"threshold"`
    FrozenDurationFrames int       `json:"frozen_duration_frames"`
    Error                string    `json:"error,omitempty"`
}

// QualityTestResult 质量测试综合结果
type QualityTestResult struct {
    CameraID      string            `json:"camera_id"`
    BlackScreen   BlackScreenResult `json:"black_screen"`
    FrozenFrame   FrozenFrameResult `json:"frozen_frame"`
    IsHealthy     bool              `json:"is_healthy"`
    Issues        []string          `json:"issues,omitempty"`
    TestedAt      time.Time         `json:"tested_at"`
}

// AnalyzeQuality 分析画面质量
func (a *QualityAnalyzer) AnalyzeQuality(
    ctx context.Context,
    cameraID string,
    rtspURL string,
) QualityTestResult {
    result := QualityTestResult{
        CameraID:  cameraID,
        IsHealthy: true,
        TestedAt:  time.Now(),
    }
    
    // 并行执行黑屏和卡顿检测
    blackScreenCh := make(chan BlackScreenResult, 1)
    frozenFrameCh := make(chan FrozenFrameResult, 1)
    
    go func() {
        blackScreenCh <- a.detectBlackScreen(ctx, rtspURL)
    }()
    
    go func() {
        frozenFrameCh <- a.detectFrozenFrame(ctx, rtspURL)
    }()
    
    // 收集结果
    result.BlackScreen = <-blackScreenCh
    result.FrozenFrame = <-frozenFrameCh
    
    // 汇总问题
    if result.BlackScreen.IsBlackScreen {
        result.IsHealthy = false
        result.Issues = append(result.Issues, "检测到黑屏")
    }
    
    if result.FrozenFrame.IsFrozen {
        result.IsHealthy = false
        result.Issues = append(result.Issues, "检测到画面卡顿")
    }
    
    if result.BlackScreen.Error != "" {
        result.Issues = append(result.Issues, 
            fmt.Sprintf("黑屏检测失败: %s", result.BlackScreen.Error))
    }
    
    if result.FrozenFrame.Error != "" {
        result.Issues = append(result.Issues, 
            fmt.Sprintf("卡顿检测失败: %s", result.FrozenFrame.Error))
    }
    
    return result
}

// detectBlackScreen 调用Python脚本检测黑屏
func (a *QualityAnalyzer) detectBlackScreen(
    ctx context.Context,
    rtspURL string,
) BlackScreenResult {
    ctx, cancel := context.WithTimeout(ctx, a.Timeout)
    defer cancel()
    
    scriptPath := fmt.Sprintf("%s/black_screen_detector.py", a.ScriptsDir)
    cmd := exec.CommandContext(ctx, a.PythonPath, scriptPath, rtspURL)
    
    output, err := cmd.Output()
    if err != nil {
        return BlackScreenResult{
            Error: fmt.Sprintf("执行检测脚本失败: %v", err),
        }
    }
    
    var result BlackScreenResult
    if err := json.Unmarshal(output, &result); err != nil {
        return BlackScreenResult{
            Error: fmt.Sprintf("解析结果失败: %v", err),
        }
    }
    
    return result
}

// detectFrozenFrame 调用Python脚本检测卡顿
func (a *QualityAnalyzer) detectFrozenFrame(
    ctx context.Context,
    rtspURL string,
) FrozenFrameResult {
    ctx, cancel := context.WithTimeout(ctx, a.Timeout)
    defer cancel()
    
    scriptPath := fmt.Sprintf("%s/frozen_detector.py", a.ScriptsDir)
    cmd := exec.CommandContext(ctx, a.PythonPath, scriptPath, rtspURL)
    
    output, err := cmd.Output()
    if err != nil {
        return FrozenFrameResult{
            Error: fmt.Sprintf("执行检测脚本失败: %v", err),
        }
    }
    
    var result FrozenFrameResult
    if err := json.Unmarshal(output, &result); err != nil {
        return FrozenFrameResult{
            Error: fmt.Sprintf("解析结果失败: %v", err),
        }
    }
    
    return result
}
```

---

### 4.4 Layer 4: 业务层测试

#### 4.4.1 端到端直播场景测试

**测试目的**: 模拟用户完整的直播观看流程

**测试流程**:
1. 调用 `StartLive` 接口开始直播
2. 获取 HLS/FLV 播放地址
3. 验证播放地址可用性
4. 调用 `Heartbeat` 保持会话
5. 调用 `StopLive` 结束直播

```go
// domain/stream_test/e2e_tester.go

package stream_test

import (
    "context"
    "time"
)

// E2ELiveTestResult 端到端直播测试结果
type E2ELiveTestResult struct {
    CameraID        string        `json:"camera_id"`
    UserID          uint64        `json:"user_id"`
    Success         bool          `json:"success"`
    SessionID       string        `json:"session_id,omitempty"`
    
    // 各阶段耗时
    StartLiveLatency    time.Duration `json:"start_live_latency_ms"`
    FirstFrameLatency   time.Duration `json:"first_frame_latency_ms"`
    TotalTestDuration   time.Duration `json:"total_test_duration_ms"`
    
    // 流质量
    StreamInfo      *StreamProbeResult `json:"stream_info,omitempty"`
    QualityResult   *QualityTestResult `json:"quality_result,omitempty"`
    
    // 错误信息
    FailedPhase     string        `json:"failed_phase,omitempty"`
    ErrorCode       string        `json:"error_code,omitempty"`
    ErrorMessage    string        `json:"error_message,omitempty"`
    TestedAt        time.Time     `json:"tested_at"`
}

// E2ETester 端到端测试器
type E2ETester struct {
    ConnectionService ConnectionService
    StreamProber      *StreamProber
    QualityAnalyzer   *QualityAnalyzer
}

// TestLiveE2E 端到端直播测试
func (t *E2ETester) TestLiveE2E(
    ctx context.Context,
    cameraID string,
    testUserID uint64,
) E2ELiveTestResult {
    result := E2ELiveTestResult{
        CameraID: cameraID,
        UserID:   testUserID,
        TestedAt: time.Now(),
    }
    
    testStart := time.Now()
    defer func() {
        result.TotalTestDuration = time.Since(testStart)
    }()
    
    // Phase 1: 开始直播
    startLiveStart := time.Now()
    liveResp, err := t.ConnectionService.StartLive(ctx, &StartLiveReq{
        CameraID: cameraID,
        UserID:   testUserID,
    })
    result.StartLiveLatency = time.Since(startLiveStart)
    
    if err != nil {
        result.Success = false
        result.FailedPhase = "start_live"
        result.ErrorCode = "START_LIVE_FAILED"
        result.ErrorMessage = err.Error()
        return result
    }
    
    result.SessionID = liveResp.SessionID
    
    // 确保测试结束后停止直播
    defer func() {
        _ = t.ConnectionService.StopLive(ctx, liveResp.SessionID)
    }()
    
    // Phase 2: 探测流
    firstFrameStart := time.Now()
    probeResult := t.StreamProber.ProbeStream(ctx, liveResp.PlayURL)
    result.FirstFrameLatency = time.Since(firstFrameStart)
    result.StreamInfo = &probeResult
    
    if !probeResult.Success {
        result.Success = false
        result.FailedPhase = "stream_probe"
        result.ErrorCode = probeResult.ErrorCode
        result.ErrorMessage = probeResult.ErrorMessage
        return result
    }
    
    // Phase 3: 心跳测试
    err = t.ConnectionService.Heartbeat(ctx, &HeartbeatReq{
        SessionID: liveResp.SessionID,
        CameraID:  cameraID,
    })
    
    if err != nil {
        result.Success = false
        result.FailedPhase = "heartbeat"
        result.ErrorCode = "HEARTBEAT_FAILED"
        result.ErrorMessage = err.Error()
        return result
    }
    
    // Phase 4: 画面质量检测 (可选，采样执行)
    if t.QualityAnalyzer != nil {
        qualityResult := t.QualityAnalyzer.AnalyzeQuality(ctx, cameraID, liveResp.PlayURL)
        result.QualityResult = &qualityResult
        
        if !qualityResult.IsHealthy {
            result.Success = false
            result.FailedPhase = "quality_check"
            result.ErrorCode = "QUALITY_ISSUE"
            result.ErrorMessage = fmt.Sprintf("质量问题: %v", qualityResult.Issues)
            return result
        }
    }
    
    result.Success = true
    return result
}
```

---

## 5. 系统架构设计

### 5.1 整体架构图

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                            Video Stream Test System                          │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐                  │
│  │  Saturn      │    │   REST API   │    │   Dashboard  │                  │
│  │  Scheduler   │    │   (手动触发)  │    │   (Grafana)  │                  │
│  └──────┬───────┘    └──────┬───────┘    └──────┬───────┘                  │
│         │                   │                   │                           │
│         └───────────────────┼───────────────────┘                           │
│                             │                                               │
│                             ▼                                               │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │                      Stream Test Orchestrator                        │   │
│  │  ┌─────────────────────────────────────────────────────────────┐   │   │
│  │  │                    Test Task Queue (Redis)                   │   │   │
│  │  └─────────────────────────────────────────────────────────────┘   │   │
│  │                                                                      │   │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐                 │   │
│  │  │ Worker Pool │  │ Worker Pool │  │ Worker Pool │                 │   │
│  │  │  (Pod 1)    │  │  (Pod 2)    │  │  (Pod N)    │                 │   │
│  │  └─────────────┘  └─────────────┘  └─────────────┘                 │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                             │                                               │
│         ┌───────────────────┼───────────────────┐                          │
│         ▼                   ▼                   ▼                          │
│  ┌─────────────┐     ┌─────────────┐     ┌─────────────┐                  │
│  │Connectivity │     │  Stream     │     │  Quality    │                  │
│  │  Tester     │     │  Prober     │     │  Analyzer   │                  │
│  │             │     │ (FFprobe)   │     │ (Python+CV) │                  │
│  └─────────────┘     └─────────────┘     └─────────────┘                  │
│         │                   │                   │                          │
│         └───────────────────┼───────────────────┘                          │
│                             ▼                                               │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │                       Result Collector                               │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│         │                   │                   │                          │
│         ▼                   ▼                   ▼                          │
│  ┌─────────────┐     ┌─────────────┐     ┌─────────────┐                  │
│  │   MySQL     │     │ Prometheus  │     │  SeaTalk    │                  │
│  │ (历史记录)   │     │  (Metrics)  │     │  (Alert)    │                  │
│  └─────────────┘     └─────────────┘     └─────────────┘                  │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                              External Systems                                │
│                                                                              │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │  Hikvision  │  │   Dahua     │  │  Uniview    │  │  Synology   │        │
│  │    VMS      │  │    VMS      │  │    VMS      │  │    NAS      │        │
│  └─────────────┘  └─────────────┘  └─────────────┘  └─────────────┘        │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 5.2 数据库设计

#### 5.2.1 测试结果表

```sql
-- 摄像头测试结果表
CREATE TABLE `camera_test_result` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `camera_id` VARCHAR(64) NOT NULL COMMENT '摄像头ID',
    `vms_id` VARCHAR(64) NOT NULL COMMENT 'VMS ID',
    `test_type` VARCHAR(32) NOT NULL COMMENT '测试类型: connectivity/stream/playback/quality',
    `test_batch_id` VARCHAR(64) NOT NULL COMMENT '测试批次ID',
    
    -- 测试结果
    `success` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否成功',
    `error_code` VARCHAR(64) DEFAULT NULL COMMENT '错误码',
    `error_message` TEXT DEFAULT NULL COMMENT '错误信息',
    
    -- 流信息
    `has_video` TINYINT(1) DEFAULT NULL COMMENT '是否有视频流',
    `has_audio` TINYINT(1) DEFAULT NULL COMMENT '是否有音频流',
    `video_codec` VARCHAR(32) DEFAULT NULL COMMENT '视频编码',
    `resolution_width` INT DEFAULT NULL COMMENT '分辨率宽度',
    `resolution_height` INT DEFAULT NULL COMMENT '分辨率高度',
    `frame_rate` DECIMAL(10,2) DEFAULT NULL COMMENT '帧率',
    `bit_rate` BIGINT DEFAULT NULL COMMENT '码率(bps)',
    
    -- 质量指标
    `is_black_screen` TINYINT(1) DEFAULT NULL COMMENT '是否黑屏',
    `is_frozen` TINYINT(1) DEFAULT NULL COMMENT '是否卡顿',
    `mean_brightness` DECIMAL(10,4) DEFAULT NULL COMMENT '平均亮度',
    
    -- 性能指标
    `latency_ms` INT DEFAULT NULL COMMENT '响应延迟(毫秒)',
    
    -- 时间戳
    `tested_at` DATETIME NOT NULL COMMENT '测试时间',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    
    PRIMARY KEY (`id`),
    INDEX `idx_camera_id` (`camera_id`),
    INDEX `idx_vms_id` (`vms_id`),
    INDEX `idx_test_batch_id` (`test_batch_id`),
    INDEX `idx_tested_at` (`tested_at`),
    INDEX `idx_success` (`success`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='摄像头测试结果表';

-- 测试批次表
CREATE TABLE `test_batch` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `batch_id` VARCHAR(64) NOT NULL COMMENT '批次ID',
    `trigger_type` VARCHAR(32) NOT NULL COMMENT '触发类型: scheduled/manual',
    `trigger_user` VARCHAR(64) DEFAULT NULL COMMENT '触发用户(手动触发时)',
    
    -- 统计信息
    `total_cameras` INT NOT NULL DEFAULT 0 COMMENT '总摄像头数',
    `success_count` INT NOT NULL DEFAULT 0 COMMENT '成功数',
    `failed_count` INT NOT NULL DEFAULT 0 COMMENT '失败数',
    `success_rate` DECIMAL(5,2) DEFAULT NULL COMMENT '成功率(%)',
    
    -- 时间信息
    `started_at` DATETIME NOT NULL COMMENT '开始时间',
    `completed_at` DATETIME DEFAULT NULL COMMENT '完成时间',
    `duration_seconds` INT DEFAULT NULL COMMENT '耗时(秒)',
    
    `status` VARCHAR(32) NOT NULL DEFAULT 'running' COMMENT '状态: running/completed/failed',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_batch_id` (`batch_id`),
    INDEX `idx_started_at` (`started_at`),
    INDEX `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='测试批次表';

-- 告警记录表
CREATE TABLE `test_alert` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `camera_id` VARCHAR(64) NOT NULL COMMENT '摄像头ID',
    `vms_id` VARCHAR(64) NOT NULL COMMENT 'VMS ID',
    `alert_type` VARCHAR(64) NOT NULL COMMENT '告警类型',
    `alert_level` VARCHAR(32) NOT NULL COMMENT '告警级别: warning/critical',
    `alert_message` TEXT NOT NULL COMMENT '告警信息',
    
    -- 状态
    `status` VARCHAR(32) NOT NULL DEFAULT 'open' COMMENT '状态: open/acknowledged/resolved',
    `acknowledged_by` VARCHAR(64) DEFAULT NULL COMMENT '确认人',
    `acknowledged_at` DATETIME DEFAULT NULL COMMENT '确认时间',
    `resolved_at` DATETIME DEFAULT NULL COMMENT '解决时间',
    
    -- 关联
    `test_result_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '关联的测试结果ID',
    
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    PRIMARY KEY (`id`),
    INDEX `idx_camera_id` (`camera_id`),
    INDEX `idx_alert_type` (`alert_type`),
    INDEX `idx_status` (`status`),
    INDEX `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='测试告警记录表';
```

### 5.3 监控指标设计

#### 5.3.1 Prometheus Metrics

```go
// infrastructure/metrics/stream_test_metrics.go

package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
    // 测试执行指标
    TestExecutionTotal = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "stream_test_execution_total",
            Help: "Total number of stream tests executed",
        },
        []string{"test_type", "vms_brand", "result"},
    )
    
    TestExecutionDuration = promauto.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "stream_test_execution_duration_seconds",
            Help:    "Duration of stream test execution",
            Buckets: []float64{0.5, 1, 2, 5, 10, 30, 60},
        },
        []string{"test_type", "vms_brand"},
    )
    
    // 摄像头状态指标
    CameraHealthStatus = promauto.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "camera_health_status",
            Help: "Camera health status (1=healthy, 0=unhealthy)",
        },
        []string{"camera_id", "vms_id", "station_id"},
    )
    
    CameraOnlineStatus = promauto.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "camera_online_status",
            Help: "Camera online status (1=online, 0=offline)",
        },
        []string{"camera_id", "vms_id", "station_id"},
    )
    
    // 流质量指标
    StreamFrameRate = promauto.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "stream_frame_rate",
            Help: "Stream frame rate in fps",
        },
        []string{"camera_id"},
    )
    
    StreamBitRate = promauto.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "stream_bit_rate_bps",
            Help: "Stream bit rate in bps",
        },
        []string{"camera_id"},
    )
    
    StreamLatencyMs = promauto.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "stream_latency_milliseconds",
            Help: "Stream probe latency in milliseconds",
        },
        []string{"camera_id"},
    )
    
    // 异常统计
    AnomalyDetectedTotal = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "stream_anomaly_detected_total",
            Help: "Total number of anomalies detected",
        },
        []string{"anomaly_type", "vms_brand"},
    )
    
    // 活跃告警数
    ActiveAlertCount = promauto.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "stream_test_active_alerts",
            Help: "Number of active alerts",
        },
        []string{"alert_type", "alert_level"},
    )
)
```

#### 5.3.2 Grafana Dashboard 设计

**Dashboard 主要面板**:

| 面板名称 | 类型 | 数据源 | 说明 |
|---------|------|--------|------|
| 摄像头健康总览 | Stat | Prometheus | 在线/离线/异常数量 |
| 测试成功率趋势 | Time Series | Prometheus | 按时间展示成功率变化 |
| 故障摄像头列表 | Table | MySQL | 当前异常摄像头详情 |
| 按 VMS 品牌统计 | Pie Chart | Prometheus | 各品牌健康状况占比 |
| 按场站统计 | Bar Gauge | Prometheus | 各场站摄像头状态 |
| 测试延迟分布 | Histogram | Prometheus | 响应时间分布 |
| 活跃告警 | Alert List | Prometheus | 未解决告警列表 |
| 历史告警趋势 | Time Series | MySQL | 告警数量趋势 |

---

## 6. 实施计划

### 6.1 项目阶段划分

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                              项目实施路线图                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  Phase 1: 基础能力建设 (第1-2周)                                              │
│  ══════════════════════                                                      │
│  • 搭建测试框架基础结构                                                        │
│  • 实现连接层测试 (Layer 1)                                                   │
│  • 实现流层测试 (Layer 2)                                                     │
│  • 数据库表设计与创建                                                          │
│                                                                              │
│  Phase 2: 质量检测能力 (第3-4周)                                              │
│  ══════════════════════                                                      │
│  • 开发 Python 质量分析脚本                                                   │
│  • 实现黑屏/卡顿检测                                                          │
│  • Go 调用 Python 封装                                                       │
│  • 单元测试覆盖                                                               │
│                                                                              │
│  Phase 3: 调度与告警 (第5-6周)                                                │
│  ══════════════════════                                                      │
│  • Saturn 定时任务集成                                                        │
│  • 并发执行控制                                                               │
│  • SeaTalk 告警集成                                                          │
│  • 告警收敛与去重                                                             │
│                                                                              │
│  Phase 4: 监控与可视化 (第7-8周)                                              │
│  ══════════════════════                                                      │
│  • Prometheus Metrics 上报                                                   │
│  • Grafana Dashboard 开发                                                    │
│  • 历史数据查询 API                                                           │
│  • 运维文档编写                                                               │
│                                                                              │
│  Phase 5: 试运行与优化 (第9-10周)                                             │
│  ══════════════════════                                                      │
│  • 小范围试运行                                                               │
│  • 性能优化                                                                   │
│  • 误报率调优                                                                 │
│  • 全量推广                                                                   │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 6.2 详细任务分解

| 阶段 | 任务 | 负责人 | 工作量(人天) | 交付物 |
|------|------|--------|-------------|--------|
| **Phase 1** | 测试框架搭建 | 后端开发 | 2 | 代码框架 |
| | 连接测试实现 | 后端开发 | 2 | ConnectivityTester |
| | 流探测实现 | 后端开发 | 3 | StreamProber |
| | 数据库设计 | DBA | 1 | DDL脚本 |
| **Phase 2** | 黑屏检测脚本 | 算法/后端 | 3 | Python脚本 |
| | 卡顿检测脚本 | 算法/后端 | 3 | Python脚本 |
| | Go-Python 集成 | 后端开发 | 2 | QualityAnalyzer |
| | 单元测试 | 测试/开发 | 2 | 测试用例 |
| **Phase 3** | Saturn 任务 | 后端开发 | 2 | 定时任务 |
| | 并发控制 | 后端开发 | 2 | Worker Pool |
| | 告警集成 | 后端开发 | 2 | AlertService |
| | 告警规则 | 运维 | 1 | 告警配置 |
| **Phase 4** | Metrics 开发 | 后端开发 | 2 | Metrics 代码 |
| | Dashboard | 运维/前端 | 3 | Grafana JSON |
| | 查询 API | 后端开发 | 2 | REST API |
| | 文档编写 | 全员 | 2 | 运维手册 |
| **Phase 5** | 试运行 | 全员 | 5 | 试运行报告 |
| | 优化迭代 | 后端开发 | 5 | 优化版本 |
| **合计** | - | - | **40人天** | - |

### 6.3 里程碑

| 里程碑 | 时间点 | 验收标准 |
|--------|--------|---------|
| M1: 基础测试能力 | 第2周末 | 连接测试+流探测功能可用 |
| M2: 完整检测能力 | 第4周末 | 黑屏/卡顿检测上线 |
| M3: 自动化运行 | 第6周末 | 定时任务+告警正常运行 |
| M4: 监控可视化 | 第8周末 | Dashboard 上线 |
| M5: 正式上线 | 第10周末 | 全量摄像头覆盖 |

---

## 7. 预期收益

### 7.1 定量收益

| 指标 | 现状 | 预期 | 提升幅度 |
|------|------|------|---------|
| 故障发现时间 | ~2小时 | ≤5分钟 | 提升 96% |
| 人工巡检工时/天 | 4小时 | 0.5小时 | 节省 87.5% |
| 摄像头可用率 | ~95% | ≥99% | 提升 4% |
| 事件追溯成功率 | ~85% | ≥98% | 提升 15% |

### 7.2 定性收益

1. **运维效率提升**: 从被动响应转为主动预警，减少人工巡检工作量
2. **故障影响降低**: 及时发现并处理故障，减少监控盲区时间
3. **服务质量提升**: 保障视频监控的高可用性，提升业务满意度
4. **数据积累**: 建立摄像头健康状态历史数据，支持分析和决策
5. **标准化运营**: 建立统一的测试标准和告警规范

---

## 8. 风险评估与应对

### 8.1 技术风险

| 风险项 | 概率 | 影响 | 应对措施 |
|--------|------|------|---------|
| FFmpeg 部署兼容性 | 中 | 高 | 使用容器化部署，统一环境 |
| 大规模并发测试性能 | 中 | 中 | Worker Pool + 限流控制 |
| Python 脚本稳定性 | 低 | 中 | 完善异常处理 + 超时控制 |
| 误报率过高 | 中 | 中 | 建立反馈机制，持续调优阈值 |
| VMS 接口兼容性 | 低 | 高 | 复用现有适配层，增量开发 |

### 8.2 运营风险

| 风险项 | 概率 | 影响 | 应对措施 |
|--------|------|------|---------|
| 测试任务影响业务 | 低 | 高 | 控制并发数，避开业务高峰 |
| 告警泛滥 | 中 | 中 | 告警收敛 + 分级推送 |
| 运维人员学习成本 | 低 | 低 | 完善文档 + 培训 |

### 8.3 资源风险

| 风险项 | 概率 | 影响 | 应对措施 |
|--------|------|------|---------|
| 开发人力不足 | 中 | 高 | 按优先级分阶段实施 |
| 服务器资源不足 | 低 | 中 | 评估资源需求，提前申请 |

---

## 9. 附录

### 9.1 术语表

| 术语 | 全称 | 说明 |
|------|------|------|
| VMS | Video Management System | 视频管理系统 |
| RTSP | Real Time Streaming Protocol | 实时流协议 |
| HLS | HTTP Live Streaming | HTTP 直播流协议 |
| FLV | Flash Video | 流媒体格式 |
| NVR | Network Video Recorder | 网络硬盘录像机 |
| FFmpeg | - | 开源音视频处理工具 |
| FFprobe | - | FFmpeg 的流分析工具 |
| OpenCV | Open Source Computer Vision | 开源计算机视觉库 |

### 9.2 参考资料

1. [FFmpeg 官方文档](https://ffmpeg.org/documentation.html)
2. [OpenCV Python 教程](https://docs.opencv.org/4.x/d6/d00/tutorial_py_root.html)
3. [Prometheus 监控最佳实践](https://prometheus.io/docs/practices/)
4. [Grafana Dashboard 设计指南](https://grafana.com/docs/grafana/latest/dashboards/)

### 9.3 相关系统接口

| 系统 | 接口 | 用途 |
|------|------|------|
| cctv_monitor | ConnectTest | VMS 连接测试 |
| cctv_monitor | GetRealStreamURL | 获取实时流地址 |
| cctv_monitor | GetReplayStreamURL | 获取回放流地址 |
| cctv_monitor | GetChannelList | 获取通道列表 |
| connection_service | StartLive | 开始直播 |
| connection_service | StopLive | 停止直播 |
| connection_service | Heartbeat | 会话心跳 |

---

**文档审批**:

| 角色 | 姓名 | 签字 | 日期 |
|------|------|------|------|
| 技术负责人 | | | |
| 产品负责人 | | | |
| 运维负责人 | | | |

---

*本文档最后更新于 2026-01-14*
