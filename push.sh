#!/bin/bash

echo -e "\033[0;32mPushing Code to GitHub...\033[0m"


# Add changes to git.
git add .

# Commit changes.
msg="Hacker rank code update on `date`"
if [ $# -eq 1 ]
then msg="$1"
fi
git commit -m "$msg"

# Push source and build repos.
git push origin master

