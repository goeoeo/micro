# 自动更新
保证 go git make 命令在su 模式下可执行  

```
git config --global --add safe.directory /home/yu/code/micro
```

# cron 
35 10 14 * * cd /home/yu/code/micro && make