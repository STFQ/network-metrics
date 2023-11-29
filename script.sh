#!/bin/bash

# 获取SYN_RECV状态的连接数
syn_recv=$(ss -tun state syn-recv | wc -l)

# 获取TIME_WAIT状态的连接数
time_wait=$(ss -tun state time-wait | wc -l)

# 获取外部连接的TIME_WAIT的连接数
time_wait_foreign=$(ss -tun state time-wait | awk 'NR>1 && $5 !~/172.21/ {print}'| wc -l)

# 打印出Prometheus指标格式
echo "# HELP node_syn_recv_total The total number of TCP connections in SYN_RECV state."
echo "# TYPE node_syn_recv_total counter"
echo "node_syn_recv_total $syn_recv"

echo "# HELP node_time_wait_total The total number of TCP connections in TIME_WAIT state."
echo "# TYPE node_time_wait_total counter"
echo "node_time_wait_total $time_wait"

echo "# HELP node_time_wait_foreign_total The total number of TCP connections in TIME_WAIT state."
echo "# TYPE node_time_wait_foreign_total counter"
echo "node_time_wait_foreign_total $time_wait_foreign"