CREATE TABLE `tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `desc` varchar(100) DEFAULT '' COMMENT '标签描述',
  `slug` varchar(100) DEFAULT '' COMMENT '标签对应图标名称',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp DEFAULT null COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';

CREATE TABLE `category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '文章分类名称',
  `desc` varchar(100) DEFAULT '' COMMENT '文章分类描述',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp DEFAULT null COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章分类管理';

CREATE TABLE `article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `category_id` int(10) unsigned DEFAULT '0' COMMENT '分类ID',
  `title` varchar(100) DEFAULT '' COMMENT'文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `keywords` varchar(100) DEFAULT '' COMMENT '关键词',
  `content` text,
  `rendered_content` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp DEFAULT null COMMENT '删除时间',
  `published_on` timestamp DEFAULT null COMMENT '发布时间',
  `thumb` varchar(100) DEFAULT '' COMMENT '缩略图链接',
  `source` tinyint(3) unsigned DEFAULT 0 COMMENT '状态 0原创 | 1转载 | 2混撰 | 3翻译',
  `reproduce_url` varchar(100) DEFAULT '' COMMENT '转载URL source 为1时需要',
  `pvs_num` int(10) unsigned DEFAULT 0 COMMENT '浏览数',
  `likes_num` int(10) unsigned DEFAULT 0 COMMENT '点赞数',
  `comments_num` int(10) unsigned DEFAULT 0 COMMENT '评论数',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为草稿 1为发布',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';

CREATE TABLE `tag_relation` (
  `article_id` int(10) NOT NULL COMMENT '文章id',
  `tag_id` int(10) NOT NULL COMMENT '标签id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签映射表';

CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL COMMENT '昵称',
  `password` varchar(50) NOT NULL COMMENT '密码',
  `url` varchar(100) DEFAULT '' COMMENT '个人网站链接',
  `avatar` varchar(100) DEFAULT '' COMMENT '缩略图链接',
  `role` tinyint(3) unsigned DEFAULT 0 COMMENT '状态 0普通群组 | 1管理员 ',
  `source` tinyint(3) unsigned DEFAULT 0 COMMENT '状态 0主站 | 1Github | 2wechat | 3qq',
  `is_muted` tinyint(3) unsigned DEFAULT 0 COMMENT '状态 0为正常 1为禁言',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp DEFAULT null COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户管理';

CREATE TABLE `file` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL COMMENT '文件名',
  `url` varchar(100) DEFAULT '' COMMENT '文件oss链接',
  `type` varchar(100) DEFAULT '' COMMENT '文件后缀',
  `size` bigint(100) DEFAULT '' COMMENT '文件大小',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp DEFAULT null COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文件管理';

CREATE TABLE `system_config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `oss_access_key_id` varchar(100) NOT NULL COMMENT 'ossID',
  `oss_access_key_sercet` varchar(100) DEFAULT '' COMMENT 'oss密钥',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统设置';

CREATE TABLE `comment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(10) unsigned NOT NULL COMMENT '用户id',
  `content` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp DEFAULT null COMMENT '删除时间',
  `likes_num` int(10) unsigned DEFAULT 0 COMMENT '点赞数',
  `state` tinyint(3) DEFAULT 0 COMMENT '-2 垃圾评论 | -1 隐藏 | 0 待审核 | 1 通过',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='评论管理';

CREATE TABLE `comment-parent-child` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(10) NOT NULL COMMENT '父评论id',
  `child_id` int(10) NOT NULL COMMENT '子评论id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='评论父子表';

INSERT INTO `tag` (`name`, `slug`, `desc`) VALUES ('vue', 'icon-vue', 'Vue (读音 /vjuː/，类似于 view) 是一套用于构建用户界面的渐进式框架'), ('react', 'icon-react', '用于构建用户界面的JavaScript库');
INSERT INTO `category` (`name`, `desc`) VALUES ('code', '代码'), ('lefe', '生活');
INSERT INTO `article` (`category_id`, `title`, `desc`, `keywords`, `content`, `rendered_content`) VALUES (1, '测试文章', '这是一篇测试文章', '额。。', '渲染前的内容啊', '<div>选然后的内容额</div>');
INSERT INTO `article-tag-map` (`article_id`, `tag_id`) VALUES (1, 1), (1, 2);
INSERT INTO `comment` (`uid`, `content`, `likes_num`, `state`) VALUES (1, '测试评论1', 3, 1), (1, '测试评论2', 0, 1), (1, '测试评论3', 0, 1);
INSERT INTO `comment-parent-child` (`parent_id`, `child_id`) VALUES (1, 2), (1, 3);
INSERT INTO `user` (`username`, `password`, `role`) VALUES ('admin', 'admin', 1);