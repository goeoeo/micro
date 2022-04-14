#!/bin/bash
# 17 12 14 * * /bin/bash /home/yu/code/micro/scripts/autoupdate.sh
cd /home/yu/code/micro
git pull
cd /home/yu/code/micro/pkg && ./generate
cd /home/yu/code/micro
git add .
git commit -m "code update"
git push