#!/bin/bash
# This sets up some local git repos that can be used with the scm plugins
cd /home
git init --bare rundeck-export.git --shared=group
git init --bare rundeck-import.git --shared=group
cd rundeck-export.git
git config receive.denyCurrentBranch ignore
cd -
cd rundeck-import.git
git config receive.denyCurrentBranch ignore
cd -
chown -R rundeck:rundeck rundeck-*.git
