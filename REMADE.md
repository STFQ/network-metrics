# 项目文件说明

该项目文件夹包含以下文件：

- `network-metrics-server`：主要的可执行文件，用于启动网络指标服务器，暴露 http服务，端口 8000
- `script.sh`：network-metrics-server 的依赖文件，用于筛选指标的数据
- `network-metrics.service` 用于注册系统服务
- `README.md`：当前文件，用于对项目文件进行说明和文档。

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

- `README.md`：
  - 类型：文本文件（Markdown 格式）
  - 作用：当前文件，用于对项目文件进行说明和文档
