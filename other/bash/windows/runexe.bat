@echo off
 

if "%1" == "h" goto begin
 

mshta vbscript:createobject("wscript.shell").run("""%~nx0"" h",0)(window.close)&&exit
 

:begin
 
echo 打开打卡定时图片任务...
start /B C:/workspace/golang/src/feng6917/local/bash/gohome_bat.exe