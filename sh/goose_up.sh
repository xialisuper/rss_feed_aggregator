#!/bin/zsh

# 切换到指定目录
cd /Users/xialiqun/Desktop/learn-go/learn_go_test/rss_feed_aggregator/sql/schema
# 执行 goose 命令
goose postgres "user=xialiqun dbname=blogator sslmode=disable" up
