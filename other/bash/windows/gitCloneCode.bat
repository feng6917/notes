echo "清理存储文件夹"
cd C:/Users/Administrator/Desktop/test/
rd /s /q "ar_device_core"
echo "清理存储文件夹完成"

echo "克隆代码"
"E:/app install/Git/bin/git.exe" clone https://git.hiscene.net/hifoundry/micro-server/ar-device-core C:/Users/Administrator/Desktop/test/ar_device_core
echo "克隆代码完成"
