#!/bin/bash

git pull
cd pkg/active && go generate
git add .
git commit -m "code update"
git push