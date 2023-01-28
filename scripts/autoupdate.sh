#!/bin/bash
# 17 12 14 * * /bin/bash /home/yu/code/micro/scripts/autoupdate.sh
cd /Users/yu/code/yu/micro
git pull
cd /Users/yu/code/yu/micro/pkg && ./generate
cd /Users/yu/code/yu/micro
git add .
git commit -m "code update"
git push