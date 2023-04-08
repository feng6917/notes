#!/bin/bash

#参考链接：https://blog.csdn.net/SudongJang/article/details/125444498

#保存备份个数，备份7天数据
number=7
#备份保存路径  路径名可自定义
backup_dir=/root/mmysql/mysqlbackup
#日期
dd=`date +%Y-%m-%d-%H-%M-%S`
#备份工具
tool=mysqldump
#用户名
username=root
#密码  自己的数据库密码
password=root
#将要备份的数据库
database_name=db_myz
#host
host=192.168.23.68
#port
port=31506

#如果文件夹不存在则创建
if [ ! -d $backup_dir ];
then
    mkdir -p $backup_dir;
fi

#简单写法 mysqldump -u root -p123456 users > /root/mysqlbackup/users-$filename.sql
#变量写法  本实例采用变量写法，这样增强脚本可移植性、可读性，后期维护时只需修改变量名即可
$tool -h$host -P$port -u $username -p$password $database_name > $backup_dir/$database_name-$dd.sql

#写创建备份日志
echo "创建数据部备份文件 $backup_dir/$database_name-$dd.sql" >> $backup_dir/log.txt

#找出需要删除的备份
delfile=`ls -l -crt $backup_dir/*.sql | awk '{print $9 }' | head -1`

#判断现在的备份数量是否大于$number
count=`ls -l -crt $backup_dir/*.sql | awk '{print $9 }' | wc -l`

if [ $count -gt $number ]
then
  #删除最早生成的备份，只保留number数量的备份
  rm $delfile
  #写删除文件日志
  echo "删除过期本份文件 $delfile" >> $backup_dir/log.txt
fi

#service crond start //启动服务
#service crond stop //关闭服务
#service crond restart //重启服务
#service crond reload //重新载入配置
#service crond status //查看服务状态
#cron 文件内容
#0 2 * * * /root/mysql_backup_script.sh
#cron 执行
#crontab mysqlRollback.cron
#cron 状态
#crontab -l


