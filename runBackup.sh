#!/bin/sh
set -e

./automatic-website-backup

git add backup
DATE=$(date)
git commit -m"stored at $DATE"
