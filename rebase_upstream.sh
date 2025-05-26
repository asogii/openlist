#!/bin/bash

git fetch alist --tags
git checkout main
#change tag to newest
git rebase tags/v3.45.0

