#!/bin/bash
# /bin/bash /home/yu/code/micro/scripts/autoupdate.sh
cd /home/yu/code/micro
git pull
cd /home/yu/code/micro/pkg/active
whereis go >> /home/yu/scripts/active_micro.log
cd /home/yu/code/micro
git add .
git commit -m "code update"
#git push