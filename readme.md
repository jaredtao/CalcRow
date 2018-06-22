# calcRow
[![Build Status](https://travis-ci.org/wentaojia2014/CalcRow.svg?branch=master)](https://travis-ci.org/wentaojia2014/CalcRow)

用来统计代码行数

使用goroutine并行执行，快速统计。

# build

go build

# run

./CalcRow path

path为要统计的路径，程序会递归遍历所有子文件夹，并统计预定扩展名的文件

输出示例:
````
.h  file count  2  line  22
.c  file count  0  line  0
.hpp  file count  1  line  16
.cpp  file count  130  line  5841
.go  file count  0  line  0
.qml  file count  0  line  0
sum line  5879
sum file count  133
````