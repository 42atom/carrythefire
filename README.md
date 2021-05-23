# carrythefire
files process manager

# 使用说明

## 基本用法

复制源盘到目标盘

```
plot-carrier start --src src_disk --dst target_disk -t 120
```

`--src` 源盘目录

`--dst` 目标盘目录

`-t or --interval` 扫描间隔时间

## 后台运行

```
nohup ./plot-carrier start --src test/src --dst test/dst -t 5 > plotcarrier.log &
```

`plotcarrier.log` 指定输出日志文件


## 停止运行

```
pkill plot-carrier
```