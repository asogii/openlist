#!/bin/bash

git fetch alist --tags
git checkout main
#git rebase tags/v3.45.0
git describe --tags `git rev-list --tags --max-count=1` > upstream_version
git rebase tags/$(cat upstream_version)

