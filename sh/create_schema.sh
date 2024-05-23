#!/bin/zsh

# 切换到指定目录
cd /Users/xialiqun/Desktop/learn-go/learn_go_test/rss_feed_aggregator/sql/schema

# 提示用户输入xxxx的内容
echo "请输入创建的file_name: "
read input

# 执行命令
goose create "$input" sql
