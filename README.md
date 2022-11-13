
# switch ENV

开发过程中，可能需要电脑共存n个版本的软件  如`php`,`composer`,`node`,`java`等

使用本软件，可以灵活切换当前环境的软件变量（无需重启cmd）

[![zFtPyD.md.png](https://s1.ax1x.com/2022/11/13/zFtPyD.md.png)](https://imgse.com/i/zFtPyD)
[![zFYAP0.png](https://s1.ax1x.com/2022/11/13/zFYAP0.md.png)](https://imgse.com/i/zFYAP0)
[![zFY0II.md.png](https://s1.ax1x.com/2022/11/13/zFY0II.md.png)](https://imgse.com/i/zFY0II)

## 使用教程

1. 你需要灵活切换版本的软件`不要`加入环境变量，纯用本软件管理（使用本软件前先删掉）
2. 本软件安装目录需要添加俩个环境变量，假设本软件安装在`D:\switchENV\`，则需要添加以下环境变量
- D:\switchENV\
- D:\switchENV\bat
3. 启动本软件: cmd 输入`./switchENV.exe` 或者运行 `swtichENV.bat`
4. 页面可视化配置和切换


## 配置

`默认`不需要什么配置，直接cmd运行软件即可，但是如果你需要自定义端口和目录等场景，可以按以下步骤新建配置文件

- 在软件目录同级新建`set.ini`
- 放置以下文件内容
```ini
app_mode = dev

db_path = D:\CodeSoft\switchEnv
db_name = switch.db

bat_path = D:\CodeSoft\switchEnv
bat_identify = bat

[gin]
address = 127.0.0.1
port = 8899
```

## 开发

- 下载代码
- `go run main.go`
- 前端代码说明看`kernel/public/README.md`
- 打包命令`go build -o switchENV.exe`

### 规划

- [x] github CICD 自动打包exe
- [ ] github CI中 添加set.ini的初始化到压缩包
- [ ] 前端 - 添加重名重复版本  API有做拦截  前端没做显示
- [ ] 前端 - 添加软件成功后 需要马上调用API刷新列表
- [ ] 前端 - 所在路径必须/结尾  做提示和检测
- [ ] 前端 - 新增弹框的时候，取消点击外面阴影关闭弹窗
- [ ] 前端 - 软件列表  路径布局优化
- [ ] 前端 - 终止程序对接