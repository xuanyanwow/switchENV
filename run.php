<?php

$config = require_once "config.php";
require_once "functions.php";

init($config);


$values = getopt('n:v:g::');



if (isset($values['g'])){
    // rmdir_pro(BASE_PATH);
    init_config($config);
}

if (empty($values['n'])){
    die("请使用-n 输入软件名");
}

if (empty($values['v'])){
    die("请使用-v 输入版本号");
}


$filename = BASE_PATH."{$values['n']}{$values['v']}.bat";

if (!file_exists($filename)){
    die("{$filename} 不存在 先新建才能切换");
}

copy($filename, BASE_PATH."{$values['n']}.bat");

echo "switchEnv --> 切换成功\n";