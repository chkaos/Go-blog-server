CREATE TABLE `tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `desc` varchar(100) DEFAULT '' COMMENT '标签描述',
  `icon` varchar(100) DEFAULT '' COMMENT '标签对应图标名称',
  `article_num` int(10) DEFAULT 0 COMMENT '标签对应文章数目',
  `created_on` int(10) unsigned DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_on` int(10) unsigned DEFAULT 0 COMMENT '修改时间',
  `state` tinyint(3) unsigned DEFAULT 1 COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';


INSERT INTO `tag` (`name`, `icon`, `description`) VALUES ('vue', 'icon-vue', 'Vue (读音 /vjuː/，类似于 view) 是一套用于构建用户界面的渐进式框架');

CREATE TABLE `category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '文章分类名称',
  `desc` varchar(100) DEFAULT '' COMMENT '文章分类描述',
  `created_on` int(10) unsigned DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_on` int(10) unsigned DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章分类管理';

INSERT INTO `category` (`name`, `desc`) VALUES ('code', '代码');

CREATE TABLE `article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `category_id` int(10) unsigned DEFAULT '0' COMMENT '分类ID',
  `title` varchar(100) DEFAULT '' COMMENT'文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `keywords` varchar(100) DEFAULT '' COMMENT '关键词',
  `content` text,
  `rendered_content` text,
  `created_on` int(10) DEFAULT NULL,
  `modified_on` int(10) unsigned DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  `published_on` int(10) unsigned DEFAULT 0 COMMENT '发布时间',
  `thumb` varchar(100) DEFAULT '' COMMENT '缩略图链接',
  `source` tinyint(3) unsigned DEFAULT 0 COMMENT '状态 0原创 | 1转载 | 2混撰 | 3翻译',
  `reproduce_url` varchar(100) DEFAULT '' COMMENT '转载URL source 为1时需要',
  `pvs_num` int(10) unsigned DEFAULT 0 COMMENT '浏览数',
  `likes_num` int(10) unsigned DEFAULT 0 COMMENT '点赞数',
  `comments_num` int(10) unsigned DEFAULT 0 COMMENT '评论数',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为草稿 1为发布',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';


CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `blog`.`blog_auth` (`id`, `username`, `password`) VALUES (null, 'test', 'test123456');