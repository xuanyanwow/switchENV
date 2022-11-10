<?php

function init($config)
{
    if (!is_dir(__DIR__.DIRECTORY_SEPARATOR."bat")){
        mkdir(__DIR__.DIRECTORY_SEPARATOR."bat");
        file_put_contents(BASE_PATH."php.bat", batTemplet(PHP_CLI_PATH));
        init_config($config);
    }
}

function init_config($config)
{
    foreach ($config as $key => $value) {
        foreach ($value as $version => $path) {
            $tempFilename = $key.$version;
            file_put_contents(BASE_PATH."{$tempFilename}.bat", batTemplet($path));
        }
    }
}

function rmdir_pro($path, $skip = ['php.bat'])
{
    //打开目录句柄
    $handle = opendir($path);
    //readdir() 从目录句柄中读取条目
    //返回目录中下一个文件的文件名。文件名以在文件系统中的排序返回
    while (false !== $filename = readdir($handle)) {
        if ($filename == '.' || $filename == '..') {  //跳过 . ..文件夹
            continue;
        }
        //判断是否为目录
        if (is_dir($path . '/' . $filename)) {
            $current_func = __FUNCTION__;
            //是目录，递归删除
            $current_func($path . '/' . $filename, $skip);
        } else {
            if (in_array($filename, $skip)) continue;
            //是文件，删除
            unlink($path . '/' . $filename); //unlink() 删除文件 返回bool
        }
    }
    //目录删除完毕
    closedir($handle); //关闭目录句柄
    return @rmdir($path); // 删除目录返回结果
}

function batTemplet($path)
{
    return <<<text
@echo OFF
:: in case DelayedExpansion is on and a path contains ! 
setlocal DISABLEDELAYEDEXPANSION
{$path} %*
text;
}