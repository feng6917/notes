@echo off
  
echo jump log path
cd /d H:\workspace\src\zhst_self\log\2024

echo set today date
set dt=%date:~0,4%.%date:~5,2%.%date:~8,2%-%date:~11,13%

echo copy template markdown file 
xcopy H:\workspace\src\zhst_self\log\template.md .
cd ./%dt%

echo rename template.md today markdown file
ren template.md %dt%.md
del template.md

echo create success