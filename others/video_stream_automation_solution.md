# 视频流自动化测试方案

> 文档版本：v1.0  
> 创建日期：2026-01-22  
> 作者：jiaxuan.han  

---

## 目录

- [一、背景与目标](#一背景与目标)
- [二、测试场景与优先级](#二测试场景与优先级)
- [三、技术架构](#三技术架构)
- [四、技术选型](#四技术选型)
- [五、检测算法原理](#五检测算法原理)
- [六、核心实现代码](#六核心实现代码)
- [七、项目目录结构](#七项目目录结构)
- [八、部署与运行](#八部署与运行)
- [九、测试用例设计](#九测试用例设计)
- [十、扩展功能](#十扩展功能)
- [十一、FAQ](#十一faq)

---

## 一、背景与目标

### 1.1 背景

CCTV系统提供摄像头直播和回放功能，当前自动化测试主要覆盖API接口层面的验证（HTTP状态码、业务retcode等），但缺少对**视频流画面质量**的自动化检测能力。

实际生产环境中可能出现以下问题：
- 摄像头画面黑屏
- 视频流卡顿/冻屏
- 画面出现雪花屏/花屏
- 画面模糊不清
- 亮度异常（过暗/过曝）

### 1.2 目标

建立一套视频流画面质量自动化检测系统，能够：

1. **自动抓取**：从直播/回放流中自动抓取视频帧
2. **智能检测**：使用图像算法检测画面质量问题
3. **集成测试**：与现有Go测试框架无缝集成
4. **持续监控**：支持接入CI/CD流水线进行持续质量监控

---

## 二、测试场景与优先级

### 2.1 画面质量检测场景

| 测试场景 | 检测内容 | 优先级 | 说明 |
|----------|----------|--------|------|
| **黑屏检测** | 画面全黑或接近全黑 | P0 | 最常见的播放问题 |
| **静止画面检测** | 画面长时间无变化（卡顿/冻屏） | P0 | 影响用户体验 |
| **雪花屏/花屏检测** | 画面出现大量噪点 | P1 | 信号问题导致 |
| **画面亮度异常** | 过暗或过曝 | P1 | 摄像头配置问题 |
| **画面模糊检测** | 对焦不清晰 | P2 | 摄像头硬件问题 |
| **帧率检测** | 实际帧率是否符合预期 | P2 | 编码/网络问题 |
| **色彩偏差检测** | 偏绿/偏红/偏蓝 | P3 | 白平衡问题 |

### 2.2 功能测试场景

| 测试场景 | 验证点 | 优先级 |
|----------|--------|--------|
| 直播开启 | session_id非空，流URL返回正确 | P0 |
| 直播心跳 | 保持直播不中断 | P0 |
| 直播取消 | 正确释放资源 | P0 |
| 回放开启 | 指定时间段的回放流正常 | P0 |
| 回放搜索录像 | 返回指定时间段的录像记录 | P0 |
| 流URL可访问性 | HLS/FLV URL返回200 | P1 |
| 流类型切换 | 主流/辅流切换正常 | P1 |
| 倍速播放 | 快进/慢放功能正常 | P2 |

---

## 三、技术架构

### 3.1 整体架构图

```
┌────────────────────────────────────────────────────────────────────┐
│                        视频流自动化测试架构                           │
├────────────────────────────────────────────────────────────────────┤
│                                                                    │
│   ┌──────────────┐    ┌──────────────┐    ┌──────────────────────┐│
│   │  CCTV API    │───>│ 获取流URL    │───>│   FFmpeg/OpenCV      ││
│   │  (Go测试框架) │    │ HLS/FLV/RTSP │    │   抓取视频帧          ││
│   └──────────────┘    └──────────────┘    └──────────┬───────────┘│
│                                                       │            │
│                                           ┌───────────▼──────────┐│
│                                           │   图像质量分析引擎   ││
│                                           │   (Python + OpenCV)  ││
│                                           └───────────┬──────────┘│
│                                                       │            │
│       ┌───────────────────────────────────────────────┼──────────┐│
│       │                      检测模块                  │          ││
│       │  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌───────┴─┐        ││
│       │  │ 黑屏    │ │ 卡顿    │ │ 花屏    │ │ 模糊    │        ││
│       │  │ 检测    │ │ 检测    │ │ 检测    │ │ 检测    │        ││
│       │  └────┬────┘ └────┬────┘ └────┬────┘ └────┬────┘        ││
│       └───────┼───────────┼───────────┼───────────┼─────────────┘│
│               └───────────┴───────────┴───────────┘              │
│                               │                                   │
│                    ┌──────────▼──────────┐                       │
│                    │   测试报告生成       │                       │
│                    │   (JSON/HTML)       │                       │
│                    └─────────────────────┘                       │
└────────────────────────────────────────────────────────────────────┘
```

### 3.2 数据流程

```
┌─────────────────────────────────────────────────────────────────┐
│                          测试流程                                │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────┐     ┌─────────┐     ┌─────────┐     ┌─────────┐  │
│  │ Step 1  │────>│ Step 2  │────>│ Step 3  │────>│ Step 4  │  │
│  │ 开启直播 │     │ 获取流URL│     │ 抓取帧   │     │ 图像分析 │  │
│  └─────────┘     └─────────┘     └─────────┘     └─────────┘  │
│       │               │               │               │        │
│       ▼               ▼               ▼               ▼        │
│  ┌─────────┐     ┌─────────┐     ┌─────────┐     ┌─────────┐  │
│  │SessionID│     │HLS/FLV  │     │JPEG帧   │     │质量报告  │  │
│  │         │     │URL      │     │         │     │         │  │
│  └─────────┘     └─────────┘     └─────────┘     └─────────┘  │
│                                                                 │
│                          ┌─────────┐                           │
│                          │ Step 5  │                           │
│                          │ 断言验证 │                           │
│                          └─────────┘                           │
│                               │                                 │
│                               ▼                                 │
│                          ┌─────────┐                           │
│                          │测试结果  │                           │
│                          │PASS/FAIL│                           │
│                          └─────────┘                           │
└─────────────────────────────────────────────────────────────────┘
```

---

## 四、技术选型

### 4.1 组件选型

| 组件 | 推荐方案 | 备选方案 | 说明 |
|------|----------|----------|------|
| **视频帧抓取** | FFmpeg | GStreamer | FFmpeg稳定支持多种流协议 |
| **图像分析** | OpenCV (Python) | Pillow + NumPy | OpenCV算法丰富 |
| **测试框架** | Go Test | - | 与现有框架保持一致 |
| **服务通信** | HTTP REST API | gRPC | 简单易用，调试方便 |
| **容器化** | Docker | - | 便于部署和环境隔离 |

### 4.2 流媒体协议支持

| 协议 | 说明 | FFmpeg支持 |
|------|------|------------|
| **HLS** | HTTP Live Streaming | ✅ 完全支持 |
| **FLV** | Flash Video (HTTP-FLV) | ✅ 完全支持 |
| **RTSP** | Real Time Streaming Protocol | ✅ 完全支持 |
| **RTMP** | Real Time Messaging Protocol | ✅ 完全支持 |

### 4.3 语言选择说明

选择 **Python + OpenCV** 作为图像分析引擎的原因：

1. **生态成熟**：OpenCV的Python绑定功能完善，文档丰富
2. **开发效率**：Python开发效率高，适合快速迭代
3. **算法丰富**：NumPy + OpenCV提供丰富的图像处理算法
4. **易于扩展**：可方便引入机器学习模型进行更复杂的检测

---

## 五、检测算法原理

### 5.1 黑屏检测

**原理**：统计灰度图中像素值低于阈值的像素占比

```python
def detect_black_screen(frame, threshold=15, ratio=0.98):
    """
    检测黑屏
    
    算法步骤：
    1. 将图像转换为灰度图
    2. 统计灰度值 < threshold 的像素数量
    3. 计算黑色像素占总像素的比例
    4. 如果比例 > ratio，则判定为黑屏
    
    参数说明：
    - threshold: 黑色像素的灰度阈值，默认15
    - ratio: 黑色像素占比阈值，默认98%
    """
    gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
    black_pixels = np.sum(gray < threshold)
    total_pixels = gray.size
    black_ratio = black_pixels / total_pixels
    
    return black_ratio >= ratio, black_ratio
```

### 5.2 画面冻结检测

**原理**：比较前后帧的相似度，相似度过高则判定为冻结

```python
def detect_frozen_frame(frame1, frame2, threshold=0.99):
    """
    检测画面冻结
    
    算法步骤：
    1. 将两帧转换为灰度图
    2. 计算两帧的直方图
    3. 使用直方图相关性比较相似度
    4. 同时计算像素级差异
    5. 如果相似度 > threshold 且像素差异很小，则判定为冻结
    """
    gray1 = cv2.cvtColor(frame1, cv2.COLOR_BGR2GRAY)
    gray2 = cv2.cvtColor(frame2, cv2.COLOR_BGR2GRAY)
    
    # 直方图比较
    hist1 = cv2.calcHist([gray1], [0], None, [256], [0, 256])
    hist2 = cv2.calcHist([gray2], [0], None, [256], [0, 256])
    similarity = cv2.compareHist(hist1, hist2, cv2.HISTCMP_CORREL)
    
    # 像素差异
    diff = cv2.absdiff(gray1, gray2)
    non_zero_ratio = np.count_nonzero(diff) / diff.size
    
    return similarity >= threshold and non_zero_ratio < 0.01
```

### 5.3 模糊检测

**原理**：使用Laplacian算子计算图像二阶导数的方差，方差越小说明越模糊

```python
def detect_blur(frame, threshold=100):
    """
    检测画面模糊
    
    算法原理：
    - Laplacian算子是图像的二阶导数，用于检测边缘
    - 清晰图像有丰富的边缘，Laplacian响应的方差大
    - 模糊图像边缘不明显，Laplacian响应的方差小
    
    参数说明：
    - threshold: 方差阈值，低于此值判定为模糊
    """
    gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
    laplacian = cv2.Laplacian(gray, cv2.CV_64F)
    variance = laplacian.var()
    
    return variance < threshold, variance
```

### 5.4 雪花屏检测

**原理**：检测图像中的高频噪点成分

```python
def detect_snow_noise(frame, threshold=50):
    """
    检测雪花屏/噪点
    
    算法步骤：
    1. 对图像进行高斯模糊，去除正常细节
    2. 计算原图与模糊图的差异
    3. 差异的均值越大，说明高频噪点越多
    """
    gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
    blurred = cv2.GaussianBlur(gray, (5, 5), 0)
    diff = cv2.absdiff(gray, blurred)
    noise_score = np.mean(diff)
    
    return noise_score > threshold, noise_score
```

### 5.5 亮度异常检测

**原理**：计算灰度图的平均值，判断是否过暗或过曝

```python
def detect_brightness_issue(frame, low=30, high=240):
    """
    检测亮度异常
    
    参数说明：
    - low: 低亮度阈值，低于此值判定为过暗
    - high: 高亮度阈值，高于此值判定为过曝
    """
    gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
    avg_brightness = np.mean(gray)
    
    if avg_brightness < low:
        return "low_brightness", avg_brightness
    elif avg_brightness > high:
        return "over_exposure", avg_brightness
    
    return None, avg_brightness
```

### 5.6 检测阈值配置

| 检测类型 | 参数名 | 默认值 | 说明 |
|----------|--------|--------|------|
| 黑屏检测 | black_threshold | 15 | 黑色像素灰度阈值 |
| 黑屏检测 | black_ratio | 0.98 | 黑色像素占比阈值 |
| 冻结检测 | frozen_threshold | 0.99 | 帧相似度阈值 |
| 模糊检测 | blur_threshold | 100 | Laplacian方差阈值 |
| 噪点检测 | noise_threshold | 50 | 噪点分数阈值 |
| 亮度检测 | low_brightness | 30 | 低亮度阈值 |
| 亮度检测 | high_brightness | 240 | 高亮度阈值 |

---

## 六、核心实现代码

### 6.1 Python图像分析服务

创建 `video_analyzer/analyzer.py`：

```python
#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""视频流质量分析服务"""

import cv2
import numpy as np
import subprocess
import tempfile
import os
from typing import Dict, List, Tuple, Optional
from dataclasses import dataclass
from enum import Enum
from flask import Flask, request, jsonify


class VideoQualityIssue(Enum):
    """视频质量问题类型"""
    BLACK_SCREEN = "black_screen"      # 黑屏
    FROZEN_FRAME = "frozen_frame"      # 画面冻结/卡顿
    SNOW_NOISE = "snow_noise"          # 雪花屏
    BLUR = "blur"                      # 模糊
    LOW_BRIGHTNESS = "low_brightness"  # 亮度过低
    OVER_EXPOSURE = "over_exposure"    # 过曝


@dataclass
class AnalysisResult:
    """分析结果"""
    is_normal: bool                    # 画面是否正常
    issues: List[VideoQualityIssue]    # 检测到的问题
    details: Dict                      # 详细指标
    frame_count: int                   # 分析的帧数


class VideoStreamAnalyzer:
    """视频流分析器"""
    
    def __init__(self):
        # 阈值配置
        self.black_threshold = 15
        self.black_ratio = 0.98
        self.frozen_threshold = 0.99
        self.blur_threshold = 100
        self.noise_threshold = 50
        self.low_brightness_threshold = 30
        self.high_brightness_threshold = 240
    
    def capture_frames_from_stream(
        self,
        stream_url: str,
        duration: int = 5,
        fps: int = 2
    ) -> List[np.ndarray]:
        """从视频流中抓取帧"""
        frames = []
        
        with tempfile.TemporaryDirectory() as tmpdir:
            output_pattern = os.path.join(tmpdir, "frame_%04d.jpg")
            
            cmd = [
                "ffmpeg",
                "-i", stream_url,
                "-t", str(duration),
                "-vf", f"fps={fps}",
                "-q:v", "2",
                "-y",
                output_pattern
            ]
            
            try:
                subprocess.run(
                    cmd,
                    capture_output=True,
                    timeout=duration + 30,
                    check=True
                )
            except subprocess.TimeoutExpired:
                print(f"FFmpeg抓取超时: {stream_url}")
            except subprocess.CalledProcessError as e:
                print(f"FFmpeg抓取失败: {e.stderr.decode()}")
            
            for filename in sorted(os.listdir(tmpdir)):
                if filename.endswith(".jpg"):
                    frame_path = os.path.join(tmpdir, filename)
                    frame = cv2.imread(frame_path)
                    if frame is not None:
                        frames.append(frame)
        
        return frames
    
    def detect_black_screen(self, frame: np.ndarray) -> Tuple[bool, float]:
        """检测黑屏"""
        gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
        black_pixels = np.sum(gray < self.black_threshold)
        total_pixels = gray.size
        black_ratio = black_pixels / total_pixels
        
        return black_ratio >= self.black_ratio, black_ratio
    
    def detect_frozen_frame(
        self,
        frame1: np.ndarray,
        frame2: np.ndarray
    ) -> Tuple[bool, float]:
        """检测画面冻结"""
        gray1 = cv2.cvtColor(frame1, cv2.COLOR_BGR2GRAY)
        gray2 = cv2.cvtColor(frame2, cv2.COLOR_BGR2GRAY)
        
        hist1 = cv2.calcHist([gray1], [0], None, [256], [0, 256])
        hist2 = cv2.calcHist([gray2], [0], None, [256], [0, 256])
        
        cv2.normalize(hist1, hist1)
        cv2.normalize(hist2, hist2)
        
        similarity = cv2.compareHist(hist1, hist2, cv2.HISTCMP_CORREL)
        
        diff = cv2.absdiff(gray1, gray2)
        non_zero_ratio = np.count_nonzero(diff) / diff.size
        
        is_frozen = similarity >= self.frozen_threshold and non_zero_ratio < 0.01
        
        return is_frozen, similarity
    
    def detect_blur(self, frame: np.ndarray) -> Tuple[bool, float]:
        """检测画面模糊"""
        gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
        laplacian = cv2.Laplacian(gray, cv2.CV_64F)
        variance = laplacian.var()
        
        return variance < self.blur_threshold, variance
    
    def detect_snow_noise(self, frame: np.ndarray) -> Tuple[bool, float]:
        """检测雪花屏/噪点"""
        gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
        blurred = cv2.GaussianBlur(gray, (5, 5), 0)
        diff = cv2.absdiff(gray, blurred)
        noise_score = np.mean(diff)
        
        return noise_score > self.noise_threshold, noise_score
    
    def detect_brightness_issue(
        self,
        frame: np.ndarray
    ) -> Tuple[Optional[VideoQualityIssue], float]:
        """检测亮度异常"""
        gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
        avg_brightness = np.mean(gray)
        
        if avg_brightness < self.low_brightness_threshold:
            return VideoQualityIssue.LOW_BRIGHTNESS, avg_brightness
        elif avg_brightness > self.high_brightness_threshold:
            return VideoQualityIssue.OVER_EXPOSURE, avg_brightness
        
        return None, avg_brightness
    
    def analyze_stream(
        self,
        stream_url: str,
        duration: int = 5,
        fps: int = 2
    ) -> AnalysisResult:
        """分析视频流质量"""
        issues = []
        details = {
            "black_screen_frames": 0,
            "frozen_frames": 0,
            "blur_frames": 0,
            "noise_frames": 0,
            "brightness_issues": 0,
            "avg_brightness": 0,
            "avg_sharpness": 0,
        }
        
        frames = self.capture_frames_from_stream(stream_url, duration, fps)
        
        if not frames:
            return AnalysisResult(
                is_normal=False,
                issues=[VideoQualityIssue.BLACK_SCREEN],
                details={"error": "无法抓取视频帧"},
                frame_count=0
            )
        
        brightness_values = []
        sharpness_values = []
        
        for i, frame in enumerate(frames):
            # 黑屏检测
            is_black, _ = self.detect_black_screen(frame)
            if is_black:
                details["black_screen_frames"] += 1
            
            # 模糊检测
            is_blur, sharpness = self.detect_blur(frame)
            sharpness_values.append(sharpness)
            if is_blur:
                details["blur_frames"] += 1
            
            # 雪花屏检测
            is_snow, _ = self.detect_snow_noise(frame)
            if is_snow:
                details["noise_frames"] += 1
            
            # 亮度检测
            brightness_issue, brightness = self.detect_brightness_issue(frame)
            brightness_values.append(brightness)
            if brightness_issue:
                details["brightness_issues"] += 1
            
            # 冻结检测
            if i > 0:
                is_frozen, _ = self.detect_frozen_frame(frames[i-1], frame)
                if is_frozen:
                    details["frozen_frames"] += 1
        
        details["avg_brightness"] = float(np.mean(brightness_values))
        details["avg_sharpness"] = float(np.mean(sharpness_values))
        
        total_frames = len(frames)
        threshold = 0.5
        
        if details["black_screen_frames"] / total_frames > threshold:
            issues.append(VideoQualityIssue.BLACK_SCREEN)
        
        if details["frozen_frames"] / max(total_frames - 1, 1) > threshold:
            issues.append(VideoQualityIssue.FROZEN_FRAME)
        
        if details["blur_frames"] / total_frames > threshold:
            issues.append(VideoQualityIssue.BLUR)
        
        if details["noise_frames"] / total_frames > threshold:
            issues.append(VideoQualityIssue.SNOW_NOISE)
        
        return AnalysisResult(
            is_normal=len(issues) == 0,
            issues=issues,
            details=details,
            frame_count=total_frames
        )


# Flask HTTP API服务
app = Flask(__name__)
analyzer = VideoStreamAnalyzer()


@app.route("/analyze", methods=["POST"])
def analyze_video():
    """分析视频流API"""
    data = request.json
    stream_url = data.get("stream_url")
    duration = data.get("duration", 5)
    fps = data.get("fps", 2)
    
    if not stream_url:
        return jsonify({"error": "stream_url is required"}), 400
    
    result = analyzer.analyze_stream(stream_url, duration, fps)
    
    return jsonify({
        "is_normal": result.is_normal,
        "issues": [issue.value for issue in result.issues],
        "details": result.details,
        "frame_count": result.frame_count
    })


@app.route("/health", methods=["GET"])
def health_check():
    return jsonify({"status": "ok"})


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8088)
```

### 6.2 Python依赖文件

创建 `video_analyzer/requirements.txt`：

```txt
opencv-python==4.8.1.78
numpy==1.24.3
flask==3.0.0
gunicorn==21.2.0
```

### 6.3 Go测试用例集成

创建 `automation/business/testcase/video_stream/stream_quality_test.go`：

```go
package video_stream

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "testing"
    "time"

    "git.garena.com/shopee/bg-logistics/qa/spx/cctv/automation/business/service/interface/parcel_cctv_server"
    "git.garena.com/shopee/bg-logistics/qa/tst/common-lib/pkg/constant/caseinfo"
    "git.garena.com/shopee/bg-logistics/qa/tst/common-lib/pkg/constant/cid"
    "git.garena.com/shopee/bg-logistics/qa/tst/common-lib/pkg/constant/env"
    "git.garena.com/shopee/bg-logistics/qa/tst/common-lib/pkg/framework/report"
    "github.com/stretchr/testify/require"
)

// VideoAnalyzerURL Python图像分析服务地址
const VideoAnalyzerURL = "http://localhost:8088"

// AnalyzeRequest 分析请求
type AnalyzeRequest struct {
    StreamURL string `json:"stream_url"`
    Duration  int    `json:"duration"`
    FPS       int    `json:"fps"`
}

// AnalyzeResponse 分析响应
type AnalyzeResponse struct {
    IsNormal   bool                   `json:"is_normal"`
    Issues     []string               `json:"issues"`
    Details    map[string]interface{} `json:"details"`
    FrameCount int                    `json:"frame_count"`
}

// analyzeVideoStream 调用Python服务分析视频流
func analyzeVideoStream(streamURL string, duration, fps int) (*AnalyzeResponse, error) {
    reqBody := AnalyzeRequest{
        StreamURL: streamURL,
        Duration:  duration,
        FPS:       fps,
    }
    
    jsonData, _ := json.Marshal(reqBody)
    
    client := &http.Client{Timeout: 60 * time.Second}
    resp, err := client.Post(
        VideoAnalyzerURL+"/analyze",
        "application/json",
        bytes.NewBuffer(jsonData),
    )
    if err != nil {
        return nil, fmt.Errorf("调用分析服务失败: %v", err)
    }
    defer resp.Body.Close()
    
    var result AnalyzeResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf("解析响应失败: %v", err)
    }
    
    return &result, nil
}

// TestLiveStreamQuality 测试直播流画面质量
func TestLiveStreamQuality(t *testing.T) {
    new(report.Configure).WithEnv(env.STAGING).
        WithCid(cid.TH).
        WithScene(caseinfo.P0).WithScene(caseinfo.ONLINE).
        WithOperator("jiaxuan.han@shopee.com").Parse()

    p := *TestParams
    t.Run(p.Info(), func(t *testing.T) {
        cfg := VideoTestData[p.Env][p.Cid]
        
        // 1. 开启直播获取流URL
        startReq := parcel_cctv_server.LiveStartReq{
            CameraID:   cfg.CameraID,
            StreamType: 2,
        }
        var startResp parcel_cctv_server.LivePlaybackStartResp
        status, errs := parcel_cctv_server.DoPostWithAdminHeaders(
            p.Env, p.Cid, p.PfbName,
            parcel_cctv_server.LiveStart,
            startReq, &startResp,
        )
        require.Equal(t, 200, status, "开启直播HTTP状态码异常")
        require.Empty(t, errs, "开启直播接口调用错误")
        require.Equal(t, 0, startResp.Retcode, "开启直播失败: %s", startResp.Message)
        
        sessionID := startResp.Data.SessionID
        streamURL := startResp.Data.FlvURL
        if streamURL == "" {
            streamURL = startResp.Data.HlsURL
        }
        require.NotEmpty(t, streamURL, "流URL不应为空")
        t.Logf("获取到流URL: %s", streamURL)
        
        // 确保测试结束后取消直播
        defer func() {
            cancelReq := parcel_cctv_server.CancelReq{
                SessionID: sessionID,
                CameraID:  cfg.CameraID,
            }
            parcel_cctv_server.DoPostWithAdminHeaders(
                p.Env, p.Cid, p.PfbName,
                parcel_cctv_server.LiveCancel,
                cancelReq, nil,
            )
            t.Log("直播已取消")
        }()
        
        // 2. 等待流稳定
        time.Sleep(3 * time.Second)
        
        // 3. 调用图像分析服务检测画面质量
        analyzeResult, err := analyzeVideoStream(streamURL, 5, 2)
        if err != nil {
            t.Skipf("图像分析服务不可用，跳过画面质量检测: %v", err)
            return
        }
        
        t.Logf("画面分析结果: 帧数=%d, 正常=%v, 问题=%v",
            analyzeResult.FrameCount,
            analyzeResult.IsNormal,
            analyzeResult.Issues,
        )
        t.Logf("详细指标: %+v", analyzeResult.Details)
        
        // 4. 断言画面质量
        require.True(t, analyzeResult.FrameCount > 0, "未能抓取到视频帧")
        require.True(t, analyzeResult.IsNormal,
            "直播画面存在问题: %v", analyzeResult.Issues)
    })
}

// TestBlackScreenDetection 专项测试：黑屏检测
func TestBlackScreenDetection(t *testing.T) {
    new(report.Configure).WithEnv(env.STAGING).
        WithCid(cid.TH).
        WithScene(caseinfo.P1).WithScene(caseinfo.ONLINE).
        WithOperator("jiaxuan.han@shopee.com").Parse()

    p := *TestParams
    t.Run(p.Info(), func(t *testing.T) {
        cfg := VideoTestData[p.Env][p.Cid]
        
        // 开启直播并获取流URL
        // ... 省略开启直播代码
        
        streamURL := "获取到的流URL"
        
        // 延长检测时间，更可靠地检测黑屏
        analyzeResult, err := analyzeVideoStream(streamURL, 10, 1)
        require.NoError(t, err)
        
        // 检查黑屏帧占比
        blackFrames := analyzeResult.Details["black_screen_frames"].(float64)
        totalFrames := float64(analyzeResult.FrameCount)
        blackRatio := blackFrames / totalFrames
        
        t.Logf("黑屏帧占比: %.2f%% (%d/%d)",
            blackRatio*100, int(blackFrames), analyzeResult.FrameCount)
        
        // 黑屏帧不应超过10%
        require.Less(t, blackRatio, 0.1,
            "黑屏帧占比过高，可能存在画面问题")
    })
}
```

---

## 七、项目目录结构

```
cctv_test/
├── automation/
│   └── business/
│       ├── service/
│       │   └── interface/
│       │       └── parcel_cctv_server/
│       │           └── model.go           # 包含流URL相关模型
│       └── testcase/
│           ├── video_view/                # 原有视频测试
│           │   ├── main_test.go
│           │   ├── video_api_test.go
│           │   └── video_view_test_data.go
│           └── video_stream/              # 新增：视频流质量测试
│               ├── main_test.go
│               ├── stream_quality_test.go
│               └── stream_test_data.go
│
├── video_analyzer/                        # Python图像分析服务
│   ├── analyzer.py                        # 核心分析代码
│   ├── requirements.txt                   # Python依赖
│   ├── Dockerfile                         # Docker镜像
│   ├── docker-compose.yml                 # Docker编排
│   └── tests/                             # Python单元测试
│       └── test_analyzer.py
│
├── docs/
│   └── video_stream_automation_solution.md  # 本文档
│
└── README.md
```

---

## 八、部署与运行

### 8.1 部署Python分析服务

#### 方式一：Docker部署（推荐）

创建 `video_analyzer/Dockerfile`：

```dockerfile
FROM python:3.10-slim

# 安装FFmpeg和OpenCV依赖
RUN apt-get update && apt-get install -y \
    ffmpeg \
    libgl1-mesa-glx \
    libglib2.0-0 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# 安装Python依赖
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# 复制代码
COPY analyzer.py .

EXPOSE 8088

# 使用gunicorn生产级部署
CMD ["gunicorn", "-w", "2", "-b", "0.0.0.0:8088", "analyzer:app"]
```

创建 `video_analyzer/docker-compose.yml`：

```yaml
version: '3.8'

services:
  video-analyzer:
    build: .
    container_name: video-analyzer
    ports:
      - "8088:8088"
    restart: unless-stopped
    environment:
      - PYTHONUNBUFFERED=1
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8088/health"]
      interval: 30s
      timeout: 10s
      retries: 3
```

启动服务：

```bash
cd video_analyzer
docker-compose up -d

# 检查服务状态
curl http://localhost:8088/health
```

#### 方式二：本地Python环境

```bash
cd video_analyzer

# 创建虚拟环境
python3 -m venv venv
source venv/bin/activate

# 安装依赖
pip install -r requirements.txt

# 确保FFmpeg已安装
ffmpeg -version

# 启动服务
python analyzer.py
```

### 8.2 运行测试

```bash
# 确保分析服务已启动
curl http://localhost:8088/health

# 运行所有视频流质量测试
cd cctv_test
go test -v ./automation/business/testcase/video_stream/...

# 运行特定测试
go test -v ./automation/business/testcase/video_stream/... \
    -run TestLiveStreamQuality

# 运行黑屏检测测试
go test -v ./automation/business/testcase/video_stream/... \
    -run TestBlackScreenDetection
```

### 8.3 CI/CD集成

在 `.gitlab-ci.yml` 中添加：

```yaml
video_stream_test:
  stage: test
  services:
    - name: video-analyzer:latest
      alias: video-analyzer
  variables:
    VIDEO_ANALYZER_URL: "http://video-analyzer:8088"
  script:
    - go test -v ./automation/business/testcase/video_stream/...
  only:
    - schedules  # 定时任务触发
```

---

## 九、测试用例设计

### 9.1 直播画面质量测试

| 用例ID | 用例名称 | 测试步骤 | 预期结果 | 优先级 |
|--------|----------|----------|----------|--------|
| VS-001 | 直播画面正常性检测 | 1.开启直播 2.抓取5秒视频帧 3.分析画面质量 | 画面无黑屏、无冻结、无花屏 | P0 |
| VS-002 | 直播黑屏检测 | 1.开启直播 2.抓取10秒视频帧 3.统计黑屏帧占比 | 黑屏帧占比<10% | P0 |
| VS-003 | 直播画面冻结检测 | 1.开启直播 2.抓取10秒视频帧 3.检测连续相同帧 | 冻结帧占比<20% | P0 |
| VS-004 | 直播画面清晰度检测 | 1.开启直播 2.抓取帧 3.计算Laplacian方差 | 清晰度分数>100 | P1 |

### 9.2 回放画面质量测试

| 用例ID | 用例名称 | 测试步骤 | 预期结果 | 优先级 |
|--------|----------|----------|----------|--------|
| VS-101 | 回放画面正常性检测 | 1.开启回放 2.抓取5秒视频帧 3.分析画面质量 | 画面无黑屏、无花屏 | P0 |
| VS-102 | 无录像时段回放 | 1.开启无录像时段回放 2.检测返回 | 返回正确错误码或空画面 | P1 |

### 9.3 测试数据配置

```go
// stream_test_data.go
package video_stream

var TestParams = &TestConfig{}

type TestConfig struct {
    Env     string
    Cid     string
    PfbName string
}

type VideoConfig struct {
    CameraID  string
    VMSID     string
    StationID int
}

type PlaybackConfig struct {
    CameraID   string
    CameraName string
    StartTime  int64
    EndTime    int64
}

// 测试数据
var VideoTestData = map[string]map[string]VideoConfig{
    "staging": {
        "th": {
            CameraID:  "v_hik_001_808",
            VMSID:     "vms_001",
            StationID: 12345,
        },
    },
}

var PlaybackTestData = map[string]map[string]PlaybackConfig{
    "staging": {
        "th": {
            CameraID:   "v_hik_001_808",
            CameraName: "Test Camera",
            StartTime:  1705909975,
            EndTime:    1705910095,
        },
    },
}
```

---

## 十、扩展功能

### 10.1 帧率检测

```python
def detect_frame_rate(self, stream_url: str, duration: int = 10) -> float:
    """检测实际帧率"""
    cmd = [
        "ffprobe",
        "-v", "error",
        "-select_streams", "v:0",
        "-count_frames",
        "-show_entries", "stream=nb_read_frames",
        "-of", "csv=p=0",
        "-t", str(duration),
        stream_url
    ]
    result = subprocess.run(cmd, capture_output=True, timeout=duration + 30)
    frame_count = int(result.stdout.decode().strip())
    return frame_count / duration
```

### 10.2 色彩偏差检测

```python
def detect_color_cast(self, frame: np.ndarray) -> Tuple[bool, str]:
    """检测色彩偏差（偏绿/偏红等）"""
    b, g, r = cv2.split(frame)
    avg_b, avg_g, avg_r = np.mean(b), np.mean(g), np.mean(r)
    
    threshold = 30
    if avg_r - avg_b > threshold and avg_r - avg_g > threshold:
        return True, "red_cast"
    elif avg_g - avg_r > threshold and avg_g - avg_b > threshold:
        return True, "green_cast"
    elif avg_b - avg_r > threshold and avg_b - avg_g > threshold:
        return True, "blue_cast"
    
    return False, "normal"
```

### 10.3 视频编码信息检测

```python
def get_stream_info(self, stream_url: str) -> dict:
    """获取视频流编码信息"""
    cmd = [
        "ffprobe",
        "-v", "quiet",
        "-print_format", "json",
        "-show_streams",
        stream_url
    ]
    result = subprocess.run(cmd, capture_output=True, timeout=30)
    info = json.loads(result.stdout.decode())
    
    video_stream = next(
        (s for s in info.get("streams", []) if s["codec_type"] == "video"),
        {}
    )
    
    return {
        "codec": video_stream.get("codec_name"),
        "width": video_stream.get("width"),
        "height": video_stream.get("height"),
        "fps": eval(video_stream.get("r_frame_rate", "0/1")),
        "bitrate": video_stream.get("bit_rate"),
    }
```

### 10.4 机器学习增强（未来规划）

可以引入预训练的CNN模型进行更复杂的检测：

- 物体遮挡检测
- 摄像头角度偏移检测
- 画面内容异常检测

---

## 十一、FAQ

### Q1: 为什么选择Python而不是Go进行图像分析？

**A**: OpenCV的Python绑定更加成熟，算法库更丰富。Go虽然也有gocv，但生态不如Python完善。通过HTTP API的方式集成，可以充分利用两种语言的优势。

### Q2: 如何处理RTSP流需要认证的情况？

**A**: 可以在流URL中包含认证信息：
```
rtsp://username:password@192.168.1.100:554/stream
```

或者使用FFmpeg的认证参数：
```bash
ffmpeg -rtsp_transport tcp -i "rtsp://192.168.1.100:554/stream" \
    -headers "Authorization: Basic xxx"
```

### Q3: 如何调整检测阈值？

**A**: 可以通过环境变量或配置文件调整：

```python
class VideoStreamAnalyzer:
    def __init__(self):
        self.black_threshold = int(os.getenv("BLACK_THRESHOLD", 15))
        self.blur_threshold = int(os.getenv("BLUR_THRESHOLD", 100))
        # ...
```

### Q4: 分析服务超时怎么办？

**A**: 
1. 增加FFmpeg超时时间
2. 减少抓取时长（duration参数）
3. 降低抓取帧率（fps参数）
4. 检查网络连接质量

### Q5: 如何获取更详细的分析报告？

**A**: 可以扩展API返回更多信息：
- 每帧的详细分析数据
- 问题帧的截图（Base64编码）
- 时间线分析图

---

## 附录

### A. 参考资料

- [OpenCV官方文档](https://docs.opencv.org/)
- [FFmpeg官方文档](https://ffmpeg.org/documentation.html)
- [HLS协议规范](https://datatracker.ietf.org/doc/html/rfc8216)

### B. 更新记录

| 版本 | 日期 | 更新内容 | 作者 |
|------|------|----------|------|
| v1.0 | 2026-01-22 | 初始版本 | jiaxuan.han |
