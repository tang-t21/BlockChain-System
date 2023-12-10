#!/bin/bash

# 执行ps命令并查找server进程
output=`ps aux | grep "$1" | grep -v grep | grep -v kill`

# 获取进程ID
pid=`echo "\$output" | awk '{print \$2}'`

# 如果找到PID,则杀死进程
if [ ! -z "$pid" ]
then
    kill $pid
    echo "Killed process with PID: $pid"
else
    echo "No matching process found"
fi

