
##### 基础命令
```
# 查看帮助
helm help

# 查看当前安装的charts
helm list --namespace namespace-name

# 搜索 charts
helm search chart-name

# 查看chart 状态
helm status chart-name

# 删除charts 不同版本命令不一致
helm delete --purge chart-name

# 安装chart
helm install -name chart-name --namespace namespace-name chart-path/

# 卸载chart 
helm uninstall chart-name --namespace namespace-name
```
