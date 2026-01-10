#!/bin/bash

git fetch openlist --tags
git checkout main
UPSTREAM_TAG=$(git tag -l "v[0-9]*" --merged openlist/main | tail -n 1)
if [ -z "$UPSTREAM_TAG" ]; then
    echo "错误：找不到有效的上游标签！"
    exit 1
fi
git rebase "$UPSTREAM_TAG" --reapply-cherry-picks || { echo "Rebase 冲突，请手动处理"; exit 1; }

echo "下一步，push的时候要加-f"

