---------
- u盘名操作
1.新建或者编辑/etc/fstab。
```
sudo vi /etc/fstab
```

2. 往/etc/fstab 写入数据。
```
LABEL=u盘名(大白菜U盘) none ntfs rw,auto,nobrowse
```

3. 拔出u盘重新插入。

4. 点击finder中前往文件夹｜打开页面 command+shift+G, 输入路径 /Volumes,之后就可以拷贝文件。

5. 退出时需要在/Volumes界面点击退出。

---------
- uuid操作

1. 查询硬盘的UUID
查询UUID
diskutil list
可以使用名字（NAME)或者分区ID（IDENTIFIER)查询UUID（是Volume UUID)
diskutil info /Volumes/TOSHIBA

2. 编辑fstab
UUID=B83735C0-40A9-478B-9689-FD98941041C3 none ntfs rw,auto,nobrowse

3, 推出硬盘，重新插入，发现桌面上没有硬盘图标即成功。
