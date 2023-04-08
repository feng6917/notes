@echo off
 

if "%1" == "h" goto begin
 

mshta vbscript:createobject("wscript.shell").run("""%~nx0"" h",0)(window.close)&&exit
 

:begin
 

REM
echo touch tody markdown file
cd C:/workspace/golang/src/feng6917/local/hiar/log/

echo set today date
set dt=%date:~0,4%%date:~5,2%%date:~8,2%-%date:~11,13%

echo mkdir today date dir
md %dt%

echo copy template markdown file 
xcopy template.md ./%dt%
cd ./%dt%

echo rename template.md today markdown file
ren template.md %dt%.md
echo create success
