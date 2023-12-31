# 项目说明
二进制文件是使用 go 进行编译的，运行的时候端口号为 8000，这个二进制服务只是提供 http 暴露的服务，用于给 prometheus 抓取数据的作用
主要抓取了 syn_recv 和 time_wait 这两个数据，用于监控 ddos 攻击的告警

# 项目文件说明

该项目文件夹包含以下文件：

- `network-metrics-server`：主要的可执行文件，用于启动网络指标服务器，暴露 http服务，端口 8000
- `script.sh`：network-metrics-server 的依赖文件，用于筛选指标的数据
- `network-metrics.service` 用于注册系统服务
- `README.md`：当前文件，用于对项目文件进行说明和文档。
- `main.go`: go的源码文件，未编译过的

## 编译 .go 文件
```bash
GOOS=linux GOARCH=amd64 go build -o network-metrics-server main.go
```
net-metrics-server: 编译后得到的可执行的二进制文件
main.go: go的源码文件
GOARCH：选择处理器的架构，这里为amd64，即 x86 64位的
GOOS：选择操作系统，Linux，macos，Windows

## 文件说明


- `network-metrics-server`：
  - 类型：可执行文件
  - 作用：启动网络指标收集
  - 使用方法：在终端中运行 `./network-metrics-server` 启动服务器

- `script.sh`：
  - 类型：脚本文件（`.sh`）
  - 作用：用于查询当前服务器的 time_wait, syn_recv 的数据
  - 使用方法：被动的被 network-metrics-server 调用

- `network-metrics.service`
  - 类型：unit（`.service`）, 系统服务文件
  - 作用：让服务使用系统服务启动托管
  - 使用方法：使用systemctl start/stop/restart/status network-metrics 
  - 注册系统服务，将该文件复制到 /etc/systemd/system 目录下
    - 执行 systemctl daemon-reload
    - 开机启动 systemctl enable network-script
    - 启动服务 systemctl start network

- `README.md`：
  - 类型：文本文件（Markdown 格式）
  - 作用：当前文件，用于对项目文件进行说明和文档

## github action 自动编译
* 有提交推送的时候，就会自动构建打包项目
* 当前构建的可执行文件，是没把 shell 脚本集成进去的，所以运行的时候，还是需要将shell脚本跟可执行文件放在同一目录下