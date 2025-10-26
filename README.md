# ZAI-GO-SDK
<div align="center">
    <h1>ZAI-GO-SDK</h1>
    <img alt="GitHub License" src="https://img.shields.io/github/license/dingdinglz/ReelFlow">
    <img alt="Model" src="https://img.shields.io/badge/Model-4.1-brightgreen?logo=Model">
    <img alt="Backend" src="https://img.shields.io/badge/Go-1.24.5-blue?logo=go">
    <br/>
</div>

ZAI-GO-SDK(ZGS) 是一个基于 `Golang` 和 `智谱API` 开发的为智谱平台的用户提供 Golang SDK 工具包，旨在简化与智谱 AI 的交互，让开发者可以更为便利的在 Golang 下只用智谱 AI 的各种能力。

## 功能特性

轻量化，易用性强，目标是完全实现所有智谱官方提供的 API 功能。

## 项目结构

```
├─config    // 处理配置文件包
├─example   // 示例代码
│   └─chat  // chat 示例
├─chat.go   // chat 主文件
├─tool.go   // 工具函数
├─type.go   // 类型定义
├─go.mod    // go mod 文件
├─go.sum    // go sum 文件
└─README.md // 项目说明文件
```

## 快速使用

### 安装

```bash
go get -u github.com/pykelysia/zaigosdk
```

### 使用

在项目根目录下创建响应配置文件

```
root
└─config
    ├─.env
    └─config.yaml
```

`root/config/.env` 内容如下：

```
AI_API_KEY="your_zhipu_api_key"
```

`root/config/config.yaml` 内容如下：

```yaml
AI_URL: "https://open.bigmodel.cn/api"
```

在代码中使用相应函数即可。

```go
// https://github.com/pykelysia/zaigosdk/example/chat/chat.go

package main

import (
	"fmt"

	"github.com/pykelysia/zaigosdk"
)

func main() {
	chatClient := zaigosdk.MustDefaultChatModel()
	response := chatClient.Chat("任意回答事或否")
	fmt.Println(response)
}
```

## API 文档

[ZAI-GO-SDK API 文档]() 仍在构建中，敬请期待。

该 SDK 全部通过[智谱官方文档](https://docs.bigmodel.cn/cn/api/introduction)实现。

## 贡献

欢迎任何形式的贡献！无论是报告问题、提交代码还是改进文档，都会非常感谢您的贡献。