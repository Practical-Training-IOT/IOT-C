# ITO-C
物联网项目C端


# IOT-C 项目目录结构

## 目录概述

本项目 `IOT-C` 的目录结构如下所示，每个文件夹和文件都有其特定的用途和功能。

## 根目录

- **common**: 存放通用代码或资源。
- **idl**: 包含接口定义语言（IDL）文件，用于描述服务接口。
    - `base.thrift`: 通用的Thrift定义文件。
    - `user.thrift`: 用户相关服务的Thrift定义文件。
- **kitex_gen**: Kitex生成的代码目录。
    - **iot**: IoT相关的生成代码。
        - **base**: 基础服务的生成代码。
        - **user**: 用户服务的生成代码。
    - **pkg**: 可能包含一些打包或辅助工具的代码。
- **rpc**: 远程过程调用（RPC）相关的代码和配置。
    - **user**: 用户服务的RPC实现。
        - **script**: 脚本文件。
            - `build.sh`: 构建脚本。
        - `handler.go`: 处理器逻辑实现。
        - `kitex_info.yaml`: Kitex配置文件。
        - `main.go`: 主程序入口。
- `go.mod`: Go模块管理文件，记录项目的依赖信息。
- `README.md`: 项目说明文件，通常包含项目介绍、使用方法等信息。

## 其他

- **外部库**: 非项目源码的外部库文件。
- **临时文件和控制台**: 临时文件和控制台输出，可能不包含在版本控制系统中。

---

以上是 `IOT-C` 项目的目录结构概览，每个部分都服务于项目的不同方面，确保了项目的模块化和可维护性。

---

## 使用方法

写完thrift后在根目录执行 kitex -module github.com/Practical-Training-IOT/IOT-C idl/你的thrift


然后在rpc文件内创建一个你服务的文件然后执行kitex -module github.com/Practical-Training-IOT/IOT-C -service iot.user -use github.com/Practical-Training-IOT/IOT-C/kitex_gen ../../idl/你的thrift