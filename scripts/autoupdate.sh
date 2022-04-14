#!/bin/bash
# /bin/bash /home/yu/code/micro/scripts/autoupdate.sh
cd /home/yu/code/micro
git pull
cd pkg/active && go generate
git add .
git commit -m "code update"
git push