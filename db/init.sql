-- 创建数据库
CREATE DATABASE IF NOT EXISTS user CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS video CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS chat CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS comment CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS follow CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS like_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS `notification` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS gorse CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE user;
-- 用户表
CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY COMMENT '用户ID',
    username VARCHAR(255) NOT NULL COMMENT '用户名',
    avatar VARCHAR(255) COMMENT '头像URL',
    password VARCHAR(255) NOT NULL COMMENT '密码',
    email VARCHAR(255) NOT NULL UNIQUE COMMENT '邮箱',
    signature VARCHAR(255) COMMENT '个性签名',
    follow_count BIGINT NOT NULL DEFAULT 0 COMMENT '关注数',
    follower_count BIGINT NOT NULL DEFAULT 0 COMMENT '粉丝数',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    INDEX idx_email (email) COMMENT '邮箱索引'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '用户信息表';

USE chat;
-- 消息表
CREATE TABLE IF NOT EXISTS message (
    id BIGINT PRIMARY KEY COMMENT '消息ID',
    to_user_id BIGINT NOT NULL COMMENT '接收者用户ID',
    from_user_id BIGINT NOT NULL COMMENT '发送者用户ID',
    content VARCHAR(255) NOT NULL COMMENT '消息内容',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    INDEX idx_to_user (to_user_id) COMMENT '接收者索引',
    INDEX idx_from_user (from_user_id) COMMENT '发送者索引'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '聊天消息表';

USE follow;
-- 关注表
CREATE TABLE IF NOT EXISTS follows (
    id BIGINT PRIMARY KEY COMMENT '关注记录ID',
    follower_id BIGINT NOT NULL COMMENT '关注者ID',
    followee_id BIGINT NOT NULL COMMENT '被关注者ID',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    INDEX idx_follower (follower_id) COMMENT '关注者索引',
    INDEX idx_followee (followee_id) COMMENT '被关注者索引'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '用户关注关系表';

-- 关注计数表
CREATE TABLE IF NOT EXISTS follows_count (
    id BIGINT PRIMARY KEY COMMENT '计数记录ID',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    follow_count BIGINT NOT NULL DEFAULT 0 COMMENT '关注数量',
    fans_count BIGINT NOT NULL DEFAULT 0 COMMENT '粉丝数量',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '用户关注统计表';

USE like_db;
-- 点赞记录表
CREATE TABLE IF NOT EXISTS like_records (
    id BIGINT PRIMARY KEY COMMENT '点赞记录ID',
    biz_id VARCHAR(50) NOT NULL COMMENT '业务类型ID',
    obj_id BIGINT NOT NULL COMMENT '被点赞对象ID',
    user_id BIGINT NOT NULL COMMENT '点赞用户ID',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    INDEX idx_biz_obj (biz_id, obj_id) COMMENT '业务对象联合索引',
    INDEX idx_user (user_id) COMMENT '用户索引'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '点赞记录表';

-- 点赞计数表
CREATE TABLE IF NOT EXISTS like_counts (
    id BIGINT PRIMARY KEY COMMENT '计数记录ID',
    biz_id VARCHAR(50) NOT NULL COMMENT '业务类型ID',
    obj_id BIGINT NOT NULL COMMENT '被点赞对象ID',
    like_num BIGINT NOT NULL DEFAULT 0 COMMENT '点赞数量',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    INDEX idx_biz_obj (biz_id, obj_id) COMMENT '业务对象联合索引'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '点赞计数表';

USE gorse;
-- 反馈表
CREATE TABLE IF NOT EXISTS `feedback` (
  `feedback_type` varchar(256) NOT NULL,
  `user_id` varchar(256) NOT NULL,
  `item_id` varchar(256) NOT NULL,
  `time_stamp` timestamp NOT NULL,
  `comment` text NOT NULL,
  PRIMARY KEY (`feedback_type`,`user_id`,`item_id`),
  KEY `user_id` (`user_id`),
  KEY `item_id` (`item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 用户表
CREATE TABLE IF NOT EXISTS `users` (
  `user_id` varchar(256) NOT NULL,
  `labels` json NOT NULL,
  `comment` text NOT NULL,
  `subscribe` json NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 物品表
CREATE TABLE IF NOT EXISTS `items` (
  `item_id` varchar(256) NOT NULL,
  `time_stamp` timestamp NOT NULL,
  `labels` json NOT NULL,
  `comment` text NOT NULL,
  `is_hidden` tinyint(1) NOT NULL,
  `categories` json NOT NULL,
  PRIMARY KEY (`item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

USE video;
-- 视频表
CREATE TABLE IF NOT EXISTS videos (
    id BIGINT PRIMARY KEY COMMENT '视频ID',
    user_id BIGINT NOT NULL COMMENT '作者ID',
    title VARCHAR(255) NOT NULL COMMENT '视频标题',
    description TEXT COMMENT '视频描述',
    url VARCHAR(255) NOT NULL COMMENT '视频URL',
    cover_url VARCHAR(255) COMMENT '封面URL',
    like_count BIGINT NOT NULL DEFAULT 0 COMMENT '点赞数',
    comment_count BIGINT NOT NULL DEFAULT 0 COMMENT '评论数',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    INDEX idx_user (user_id) COMMENT '作者索引',
    INDEX idx_created (created_at) COMMENT '创建时间索引'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '视频信息表';

USE comment;
-- 评论表
CREATE TABLE IF NOT EXISTS comments (
    id BIGINT PRIMARY KEY COMMENT '评论ID',
    user_id BIGINT NOT NULL COMMENT '评论用户ID',
    content TEXT NOT NULL COMMENT '评论内容',
    biz_id VARCHAR(50) NOT NULL COMMENT '业务类型ID',
    obj_id BIGINT NOT NULL COMMENT '评论对象ID',
    parent_id BIGINT DEFAULT NULL COMMENT '父评论ID',
    reply_count INT NOT NULL DEFAULT 0 COMMENT '回复数量',
    like_count INT NOT NULL DEFAULT 0 COMMENT '点赞数量',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    INDEX idx_user (user_id) COMMENT '评论用户索引',
    INDEX idx_biz_obj (biz_id, obj_id) COMMENT '评论对象索引',
    INDEX idx_parent (parent_id) COMMENT '父评论索引',
    INDEX idx_created (created_at) COMMENT '创建时间索引'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '评论表';

USE notification;
-- 通知表
CREATE TABLE IF NOT EXISTS notifications (
    id BIGINT PRIMARY KEY COMMENT '通知ID',
    user_id BIGINT NOT NULL COMMENT '接收用户ID',
    title VARCHAR(255) NOT NULL COMMENT '通知标题',
    content TEXT NOT NULL COMMENT '通知内容',
    type TINYINT NOT NULL COMMENT '通知类型: 1-系统通知 2-点赞通知 3-评论通知 4-关注通知',
    status TINYINT NOT NULL DEFAULT 0 COMMENT '状态: 0-未读 1-已读',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    INDEX idx_user (user_id) COMMENT '接收用户索引',
    INDEX idx_type (type) COMMENT '通知类型索引',
    INDEX idx_status (status) COMMENT '通知状态索引',
    INDEX idx_created (created_at) COMMENT '创建时间索引',
    INDEX idx_user_status (user_id, status) COMMENT '用户未读通知索引'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '通知表';

USE follow;
-- 关注表
CREATE TABLE IF NOT EXISTS follows (
    id BIGINT PRIMARY KEY COMMENT '关注记录ID',
    follower_id BIGINT NOT NULL COMMENT '关注者ID',
    followee_id BIGINT NOT NULL COMMENT '被关注者ID',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    INDEX idx_follower (follower_id) COMMENT '关注者索引',
    INDEX idx_followee (followee_id) COMMENT '被关注者索引',
    INDEX idx_created (created_at) COMMENT '创建时间索引',
    UNIQUE INDEX uniq_follower_followee (follower_id, followee_id) COMMENT '关注关系唯一索引'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '用户关注关系表';

-- 关注计数表
CREATE TABLE IF NOT EXISTS follows_count (
    id BIGINT PRIMARY KEY COMMENT '计数记录ID',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    follow_count BIGINT NOT NULL DEFAULT 0 COMMENT '关注数量',
    fans_count BIGINT NOT NULL DEFAULT 0 COMMENT '粉丝数量',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    UNIQUE INDEX uniq_user (user_id) COMMENT '用户ID唯一索引'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '用户关注统计表';