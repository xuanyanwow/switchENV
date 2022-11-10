# switch ENV

开发过程中，可能需要电脑共存n个版本的软件  如`php`,`composer`,`node`,`java`等

使用本软件，可以灵活切换当前环境的软件变量（无需重启cmd）


🕣 v1使用的是php实现，第一次使用比较麻烦（也不适用于非phper 毕竟他们电脑上没有php环境）
🕣 过俩天用go重写一下！

## 使用教程

1. 你需要灵活切换版本的软件`不要`加入环境变量，纯用本软件管理（使用本软件前先删掉）
2. 本软件安装目录需要添加俩个环境变量，假设本软件安装在`D:\switchENV\`，则需要添加以下环境变量
- D:\switchENV\
- D:\switchENV\bat
3. 在`config.php`中定义`软件名+版本号`以及安装地址，以及初始化PHP版本`PHP_CLI_PATH`
4. 重启cmd，运行命令 `switchEnv -g` 根据config.php初始化（以后更新了config.php就运行一次这个命令）
5. 切换软件版本`switchEnv -n 软件名 -v 版本号`


注意事项，第一次用比较麻烦 不能直接使用`switchEnv`运行

- 先cd到本软件安装目录
- 显式使用php运行 如`D:\Soft\phpstudy_pro\Extensions\php\php7.3.4nts\php run.php`
- 第一次才用显式使用  后面就会把初始化PHP版本注入到环境中
- `switchEnv -n 软件名 -v 版本号`