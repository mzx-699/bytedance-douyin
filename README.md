[TOC]

# 写在前面

## 基本信息

- 组名 发际线与我坐队
- 组号 9033
- 姓名 麻志翔
- 手机号 15922710505
- 邮箱 978653881@qq.com

## 分工

- 独立完成，其他成员联系不上

## 进度

- 完成基础接口
- 完成扩展接口-I
- 完成扩展接口-II

## 使用到的工具

- 数据库

> 1. MySQL，SQL文件为`douyin-sql.sql`

- Go主要包

> gin/gorm/ffmpeg

- 其他

> 1. 需要安装系统工具ffmpeg，工具在项目目录下，链接到环境变量即可

# 项目结构

## 整体结构

```
douyin
├── README.md
├── controller
│   ├── comment_controller.go
│   ├── common.go
│   ├── favorite_controller.go
│   ├── feed_controller.go
│   ├── publish_controller.go
│   ├── relation_controller.go
│   └── user_controller.go
├── douyin
├── douyin-sql.sql
├── ffmpeg
├── go.mod
├── go.sum
├── main.go
├── public
│   ├── 7_bear.mp4_1655020443
│   └── 7_bear.png
├── repository
│   ├── comment_repository.go
│   ├── db_init.go
│   ├── db_test.go
│   ├── favorite_repository.go
│   ├── feed_repository.go
│   ├── relation_repository.go
│   └── user_repository.go
├── router.go
├── service
│   ├── comment_service.go
│   ├── common.go
│   ├── favorite_service.go
│   ├── feed_service.go
│   ├── relation_service.go
│   ├── service_test.go
│   └── user_service.go
└── util
    ├── feed_util.go
    ├── file_util.go
    ├── logger.go
    └── user_util.go

5 directories, 35 files
```

## 主文件

- `main.go`主文件
- `router.go`URL配置文件

## controller

- `controller`层，处理请求
- `common.go`

> 1. 响应结构体
> 2. 公共方法
> 3. 设置默认端口

## service

- `service`层，业务处理层，负责处理数据库报错以及对象转换等
- `common.go`

> 1. 公共接口
> 2. 响应对象结构体

## repository

- `repository`层，数据层，负责连接数据库以及操作数据库等

## util

- `util`工具层，包括日志、文件处理、用户信息处理和视频处理工具

## public

- 静态资源存储

## 其他

- `douyin`可执行文件
- `ffmpeg`ffmpeg工具

# 数据库结构

## users

- 用户表

| 序号 |       名称       |    描述    |           类型            |  键  | 为空 |                     额外                      |      默认值       |
| :--: | :--------------: | :--------: | :-----------------------: | :--: | :--: | :-------------------------------------------: | :---------------: |
|  1   |       `id`       |   唯一id   |            int            | PRI  |  NO  |                auto_increment                 |                   |
|  2   |      `name`      |    账户    |       varchar(255)        |      |  NO  |                                               |                   |
|  3   |  `follow_count`  |  关注数量  | int(32) unsigned zerofill |      |  NO  |                                               |                   |
|  4   | `follower_count` |  粉丝数量  | int(32) unsigned zerofill |      |  NO  |                                               |                   |
|  5   |     `token`      |   token    |       varchar(255)        | MUL  |  NO  |                                               |                   |
|  6   |   `created_at`   |  创建时间  |         datetime          |      |  NO  |               DEFAULT_GENERATED               | CURRENT_TIMESTAMP |
|  7   |   `updated_at`   |  更新时间  |         datetime          |      |  NO  | DEFAULT_GENERATED on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
|  8   |   `deleted_at`   | 伪删除标识 |         datetime          |      | YES  |                                               |                   |

## feeds

- 视频表

| 序号 |       名称       |    描述    |     类型     |  键  | 为空 |                     额外                      |      默认值       |
| :--: | :--------------: | :--------: | :----------: | :--: | :--: | :-------------------------------------------: | :---------------: |
|  1   |       `id`       |   唯一id   |     int      | PRI  |  NO  |                auto_increment                 |                   |
|  2   |    `user_id`     |   用户id   |     int      | MUL  | YES  |                                               |                   |
|  3   |     `title`      |    标题    | varchar(255) |      |  NO  |                                               |                   |
|  4   |    `play_url`    |  视频路径  | varchar(255) |      |  NO  |                                               |                   |
|  5   |   `cover_url`    |  封面路径  | varchar(255) |      |  NO  |                                               |                   |
|  6   | `favorite_count` |  点赞数量  |     int      |      |  NO  |                                               |                   |
|  7   | `comment_count`  |  评论数量  |     int      |      |  NO  |                                               |                   |
|  8   |   `created_at`   |  创建时间  |   datetime   | MUL  |  NO  |               DEFAULT_GENERATED               | CURRENT_TIMESTAMP |
|  9   |   `updated_at`   |  更新时间  |   datetime   |      |  NO  | DEFAULT_GENERATED on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
|  10  |   `deleted_at`   | 伪删除标识 |   datetime   |      | YES  |                                               |                   |

## comments

- 评论表

| 序号 |     名称     |     描述     |   类型   |  键  | 为空 |                     额外                      |      默认值       |
| :--: | :----------: | :----------: | :------: | :--: | :--: | :-------------------------------------------: | :---------------: |
|  1   |     `id`     |    唯一id    |   int    | PRI  |  NO  |                auto_increment                 |                   |
|  2   |  `user_id`   |  评论用户id  |   int    |      |  NO  |                                               |                   |
|  3   |  `feed_id`   | 被评论视频id |   int    |      |  NO  |                                               |                   |
|  4   |  `content`   |     内容     |   text   |      |  NO  |                                               |                   |
|  5   | `created_at` |   创建时间   | datetime |  0   |  NO  |               DEFAULT_GENERATED               | CURRENT_TIMESTAMP |
|  6   | `updated_at` |   更新时间   | datetime |      |  NO  | DEFAULT_GENERATED on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
|  7   | `deleted_at` |  伪删除标识  | datetime |      | YES  |                                               |                   |

## favorites

- 点赞表

| 序号 |     名称     |     描述     |   类型   |  键  | 为空 |                     额外                      |      默认值       |
| :--: | :----------: | :----------: | :------: | :--: | :--: | :-------------------------------------------: | :---------------: |
|  1   |     `id`     |    唯一id    |   int    | PRI  |  NO  |                auto_increment                 |                   |
|  2   |    `user`    |    用于id    |   int    |      |  NO  |                                               |                   |
|  3   |    `feed`    |    视频id    |   int    |      |  NO  |                                               |                   |
|  4   |   `cancel`   | 是否取消标识 |   int    |  0   |  NO  |                                               |         0         |
|  5   | `created_at` |   创建时间   | datetime |      |  NO  |               DEFAULT_GENERATED               | CURRENT_TIMESTAMP |
|  6   | `updated_at` |   更新时间   | datetime |      |  NO  | DEFAULT_GENERATED on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
|  7   | `deleted_at` |  伪删除标识  | datetime |      | YES  |                                               |                   |

## relations

- 关注表

| 序号 |     名称     |    描述    |   类型   |  键  | 为空 |                     额外                      |      默认值       |
| :--: | :----------: | :--------: | :------: | :--: | :--: | :-------------------------------------------: | :---------------: |
|  1   |     `id`     |   唯一id   |   int    | PRI  |  NO  |                auto_increment                 |                   |
|  2   |   `follow`   |  关注者id  |   int    |      |  NO  |                                               |                   |
|  3   |  `follower`  |   粉丝id   |   int    |      |  NO  |                                               |                   |
|  4   |   `cancel`   |  取消标识  |   int    |  0   |  NO  |                                               |         0         |
|  5   | `created_at` |  创建时间  | datetime |      |  NO  |               DEFAULT_GENERATED               | CURRENT_TIMESTAMP |
|  6   | `updated_at` |  更新时间  | datetime |      |  NO  | DEFAULT_GENERATED on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
|  7   | `deleted_at` | 伪删除标识 | datetime |      | YES  |                                               |                   |

# 运行效果

- 用户基本信息

<img src="/Users/mzx/Desktop/GoGoGoGoGo/bytedance/douyin/readme-source/image-20220612170241586.png" alt="image-20220612170241586" style="zoom:25%;" />

- 喜欢

<img src="/Users/mzx/Library/Application Support/typora-user-images/image-20220612172128895.png" alt="image-20220612172128895" style="zoom:25%;" />

- 播放

<img src="/Users/mzx/Library/Application Support/typora-user-images/image-20220612165919240.png" alt="image-20220612165919240" style="zoom: 25%;" />

- 评论

<img src="/Users/mzx/Library/Application Support/typora-user-images/image-20220612170046199.png" alt="image-20220612170046199" style="zoom:25%;" />

- 关注

<img src="/Users/mzx/Library/Application Support/typora-user-images/image-20220612170830478.png" alt="image-20220612170830478" style="zoom:25%;" />

- 粉丝

<img src="/Users/mzx/Desktop/GoGoGoGoGo/bytedance/douyin/readme-source/image-20220612171529030.png" alt="image-20220612171529030" style="zoom:25%;" />
