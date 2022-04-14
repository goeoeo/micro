# 自动更新
保证 git make 命令在su 模式下可执行  

```
git config --global --add safe.directory /home/yu/code/micro
```

# cron 
15 13 14 * * /bin/bash /home/yu/code/micro/scripts/autoupdate.sh
