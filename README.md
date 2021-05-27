# carrythefire
files process manager

# 使用说明

## 生成配置文件

```
plot-carrier init
```
将在当前文件夹中生成config.yaml文件。文件带有非常详细的注释，配置时，请直接查看对应的配置文件。

## 基本用法

复制源盘到目标盘

```
plot-carrier start --src src_disk --dst target_disk -t 120 --config plot-carrier.yaml
```

`--src` 源盘目录

`--dst` 目标盘目录

`-t or --interval` 扫描间隔时间/s

`--config` 配置文件路径

## 后台运行

```
nohup ./plot-carrier start --src test/src --dst test/dst -t 120 --config plot-carrier.yaml > plotcarrier.log &
```

`plotcarrier.log` 指定输出日志文件


## 停止运行

```
pkill plot-carrier
```
