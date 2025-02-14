# DataKit

<p align="center">
  <img alt="datakit logo" src="datakit-logo.png" height="150" />
</p>

[![Slack Status](https://img.shields.io/badge/slack-join_chat-orange?logo=slack&style=plastic)](https://app.slack.com/client/T032YB4B6TA/)
[![MIT License](https://img.shields.io/badge/license-MIT-green?style=plastic)](LICENSE)

<h2>
  <a href="https://datakit.tools">官网</a>
  <span> • </span>
  <a href="https://www.yuque.com/dataflux/datakit">文档</a>
</h2>

DataKit 是一款开源、一体式的数据采集 Agent，它提供全平台操作系统（Linux/Windows/macOS）支持，拥有全面数据采集能力，涵盖主机、容器、中间件、Tracing、日志以及安全巡检等各种场景。

## 主要功能点

- 支持主机、中间件、日志、APM 等领域的指标、日志以及 Tracing 几大类数据采集
- 完整支持 Kubernetes 云原生生态
- [Pipeline](https://www.yuque.com/dataflux/datakit/pipeline)：简便的结构化数据提取
- 支持接入其它第三方数据采集
	- [Telegraf](https://www.yuque.com/dataflux/datakit/telegraf)
	- [Prometheus](https://www.yuque.com/dataflux/datakit/prom)
	- [Statsd](https://www.yuque.com/dataflux/datakit/statsd)
	- [Fluentd](https://www.yuque.com/dataflux/datakit/logstreaming#a653042e)
	- [Function](https://www.yuque.com/dataflux/func/write-data-via-datakit)
	- Tracing 相关（[OpenTelemetry]()/[DDTrace]()/[Zipkin]()/[Jaeger]()/[Skywalking]()）

## 发布历史

DataKit 发布历史参见[这里](https://www.yuque.com/dataflux/datakit/changelog).

## 操作系统最低要求

| 操作系统 | 架构 | 安装路径 |
| --- | --- | --- |
| Linux 内核 2.6.23 或更高版本 | amd64/386/arm/arm64 | `/usr/local/datakit` |
| macOS 10.12 或更高版本([原因](https://github.com/golang/go/issues/25633)) | amd64 | `/usr/local/datakit` |
| Windows 7, Server 2008R2 或更高版本 | amd64/386 | 64位：`C:\Program Files\datakit`<br />32位：`C:\Program Files(32)\datakit` |

## DataKit 安装

我们可以直接在观测云平台获取 DataKit 安装命令，主流平台的安装命令大概如下：

- Linux & Mac
```shell
DK_DATAWAY="https://openway.guance.com?token=<YOUR-TOKEN>" bash -c "$(curl -L https://static.guance.com/datakit/install.sh)"
```

- Windows

```powershell
$env:DK_DATAWAY="https://openway.guance.com?token=<YOUR-TOKEN>";Set-ExecutionPolicy Bypass -scope Process -Force; Import-Module bitstransfer; start-bitstransfer -source https://static.guance.com/datakit/install.ps1 -destination .install.ps1; powershell .install.ps1;
```

- [Kubernetes DaemonSet](https://www.yuque.com/dataflux/datakit/datakit-daemonset-deploy)

更多关于安装的文档，参见[这里](https://www.yuque.com/dataflux/datakit/datakit-install)。

### 安装社区版

同时我们也发布了 DataKit 的社区版，相比上面的版本，社区版会激进的加入一些新的功能，但不一定具有较好的兼容性以及稳定。可通过如下方式安装社区版：

- Linux & Mac

```bash
DK_DATAWAY="https://openway.guance.com?token=<YOUR-TOKEN>" bash -c "$(curl -L https://static.guance.com/datakit/community/install.sh)"
```

- Windows

```powershell
$env:DK_DATAWAY="https://openway.guance.com?token=<YOUR-TOKEN>";Set-ExecutionPolicy Bypass -scope Process -Force; Import-Module bitstransfer; start-bitstransfer -source https://static.guance.com/datakit/community/install.ps1 -destination .install.ps1; powershell .install.ps1;
```

- [Kubernetes DaemonSet](https://www.yuque.com/dataflux/datakit/datakit-daemonset-deploy)

```bash
# 我们须替换上文中的 yaml 地址
wget https://static.guance.com/datakit/community/datakit.yaml
```

## 源码编译

DataKit 开发过程中依赖了一些外部工具，我们必须先将这些工具准备好才能比较顺利的编译 DataKit。

以下依赖（库/工具）主要用于 DataKit 自身的编译、打包以及发布流程。其中，**不建议在 Windows 上编译 DataKit**。

- Go-1.16.4 及以上版本
- `apt-get install gcc-multilib`: 用于编译 Oracle 采集器
- `apt-get install tree`: Makefile 中用于显示编译结果
- `packr2`: 用于打包一些资源文件
- `go get -u golang.org/x/tools/cmd/goyacc`: 用于生成 Pipeline 语法代码
- Docker 用于生成 DataKit 镜像
- lint 相关
	- `go install mvdan.cc/gofumpt@latest` 用于规范化 Golang 代码格式
	- [golangci-lint 1.42.1](https://github.com/golangci/golangci-lint/releases/tag/v1.42.1)
- eBPF 相关
	- clang 10.0+
	- llvm 10.0+
	- `apt install go-bindata`
- 文档相关
	- [waque 1.13.1+](https://github.com/yesmeck/waque)

### 编译

1. 拉取代码：

```shell
git clone https://github.com/DataFlux-cn/datakit.git
```

2. 编译：

```shell
cd datakit
make
```

如果编译通过，将在当前目录的 *dist* 目录下生成如下文件：

```
dist/
├── datakit-linux-amd64
│   ├── datakit             # DataKit 主程序
│   └── externals      
│       ├── datakit-ebpf    # eBPF 相关采集器
│       ├── logfwd          # logfwd 采集器
│       └── oracle          # Oracle 采集器
└── local
    ├── installer-linux-amd64 # Linux 平台安装程序
    └── version               # version 信息描述文件
```

如果要编译全平台版本，执行：

```shell
make testing
```

## DataKit 基本使用

可通过如下命令查看更多使用方法：

```shell
datakit help
```

## 如何贡献代码

在为我们贡献代码之前：

- 可尝试阅读 DataKit [基本架构介绍](https://www.yuque.com/dataflux/datakit/datakit-arch)
- 请先查看我们的[开发指南](https://www.yuque.com/dataflux/datakit/development)

## 文档

- [DataKit 文档库](https://www.yuque.com/dataflux/datakit)
- [DataKit 社区版文档库](https://www.yuque.com/dataflux/datakit-community)
