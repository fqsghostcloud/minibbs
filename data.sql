/*
 Navicat Premium Data Transfer

 Source Server         : mydata
 Source Server Type    : MySQL
 Source Server Version : 50615
 Source Host           : 192.168.34.145:3306
 Source Schema         : minibbs

 Target Server Type    : MySQL
 Target Server Version : 50615
 File Encoding         : 65001

 Date: 04/05/2018 11:24:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL DEFAULT 0,
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 29 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of permission
-- ----------------------------
INSERT INTO `permission` VALUES (1, 0, '', '', '话题节点');
INSERT INTO `permission` VALUES (2, 1, '/topic/create', 'topic:add', '创建话题');
INSERT INTO `permission` VALUES (3, 1, '/topic/edit/[0-9]+', 'topic:edit', '编辑话题');
INSERT INTO `permission` VALUES (4, 1, '/topic/delete/[0-9]+', 'topic:delete', '删除话题');
INSERT INTO `permission` VALUES (5, 0, '', '', '回复节点');
INSERT INTO `permission` VALUES (6, 5, '/reply/delete/[0-9]+', 'reply:delete', '删除回复');
INSERT INTO `permission` VALUES (7, 5, '/reply/save', 'reply:save', '创建回复');
INSERT INTO `permission` VALUES (8, 5, '/reply/up', 'reply:up', '点赞回复');
INSERT INTO `permission` VALUES (12, 0, '', '', '权限节点');
INSERT INTO `permission` VALUES (13, 12, '/user/list', 'user:list', '用户列表');
INSERT INTO `permission` VALUES (14, 12, '/user/edit/[0-9]+', 'user:edit', '角色配置');
INSERT INTO `permission` VALUES (15, 12, '/user/delete/[0-9]+', 'user:delete', '用户删除');
INSERT INTO `permission` VALUES (16, 12, '/role/list', 'role:list', '角色列表');
INSERT INTO `permission` VALUES (17, 12, '/role/add', 'role:add', '添加角色');
INSERT INTO `permission` VALUES (18, 12, '/role/delete/[0-9]+', 'role:delete', '删除角色');
INSERT INTO `permission` VALUES (20, 12, '/role/edit/[0-9]+', 'role:edit', '编辑角色');
INSERT INTO `permission` VALUES (21, 12, '/permission/list', 'permission:list', '权限列表');
INSERT INTO `permission` VALUES (22, 12, '/permission/add', 'permission:add', '添加权限');
INSERT INTO `permission` VALUES (23, 12, '/permission/edit/[0-9]+', 'permission:edit', '编辑权限');
INSERT INTO `permission` VALUES (24, 12, '/permission/delete/[0-9]+', 'permission:delete', '删除权限');
INSERT INTO `permission` VALUES (25, 12, '/topic/manage', 'topic:manage', '帖子管理');
INSERT INTO `permission` VALUES (26, 12, '/about/edit', 'topic:edit', '公告编辑');
INSERT INTO `permission` VALUES (27, 1, '/topic/join/ws', 'topic:chat', '聊天室');
INSERT INTO `permission` VALUES (28, 1, '/tag/manage', 'tag:manage', '标签管理');

-- ----------------------------
-- Table structure for reply
-- ----------------------------
DROP TABLE IF EXISTS `reply`;
CREATE TABLE `reply`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `topic_id` int(11) NOT NULL,
  `content` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `user_id` int(11) NOT NULL,
  `up` int(11) NOT NULL DEFAULT 0,
  `in_time` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of reply
-- ----------------------------
INSERT INTO `reply` VALUES (5, 1, '请大家遵守。', 1, 0, '2018-04-19 15:46:02');
INSERT INTO `reply` VALUES (6, 5, '## 因为Equaler的Equal函数需要的类型不同，正确的实现方式为\r\n```\r\ntype T2 int\r\nfunc (t T2) Equal(u Equaler) bool { return t == u.(T2) }  // satisfies Equaler\r\n\r\n另一个例子\r\ntype Opener interface {\r\n   Open() Reader\r\n}\r\n\r\nfunc (t T3) Open() *os.File\r\n//T3 does not satisfy Opener, although it might in another language.\r\n```', 2, 0, '2018-04-19 15:53:57');
INSERT INTO `reply` VALUES (7, 1, '必须的！', 2, 0, '2018-04-19 16:13:38');
INSERT INTO `reply` VALUES (8, 8, '666', 2, 2, '2018-04-19 16:16:35');
INSERT INTO `reply` VALUES (9, 8, '多谢', 1, 1, '2018-04-19 16:46:40');

-- ----------------------------
-- Table structure for reply_up_log
-- ----------------------------
DROP TABLE IF EXISTS `reply_up_log`;
CREATE TABLE `reply_up_log`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `reply_id` int(11) NOT NULL,
  `in_time` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of reply_up_log
-- ----------------------------
INSERT INTO `reply_up_log` VALUES (5, 2, 8, '2018-04-19 16:46:19');
INSERT INTO `reply_up_log` VALUES (7, 1, 8, '2018-04-19 16:46:34');
INSERT INTO `reply_up_log` VALUES (8, 1, 9, '2018-04-19 16:46:42');

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (5, '普通用户');
INSERT INTO `role` VALUES (3, '管理员');

-- ----------------------------
-- Table structure for role_permissions
-- ----------------------------
DROP TABLE IF EXISTS `role_permissions`;
CREATE TABLE `role_permissions`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `permission_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 117 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of role_permissions
-- ----------------------------
INSERT INTO `role_permissions` VALUES (47, 4, 3);
INSERT INTO `role_permissions` VALUES (48, 4, 4);
INSERT INTO `role_permissions` VALUES (49, 4, 6);
INSERT INTO `role_permissions` VALUES (81, 3, 2);
INSERT INTO `role_permissions` VALUES (82, 3, 3);
INSERT INTO `role_permissions` VALUES (83, 3, 4);
INSERT INTO `role_permissions` VALUES (84, 3, 27);
INSERT INTO `role_permissions` VALUES (85, 3, 28);
INSERT INTO `role_permissions` VALUES (86, 3, 6);
INSERT INTO `role_permissions` VALUES (87, 3, 7);
INSERT INTO `role_permissions` VALUES (88, 3, 8);
INSERT INTO `role_permissions` VALUES (89, 3, 13);
INSERT INTO `role_permissions` VALUES (90, 3, 14);
INSERT INTO `role_permissions` VALUES (91, 3, 15);
INSERT INTO `role_permissions` VALUES (92, 3, 16);
INSERT INTO `role_permissions` VALUES (93, 3, 17);
INSERT INTO `role_permissions` VALUES (94, 3, 18);
INSERT INTO `role_permissions` VALUES (95, 3, 20);
INSERT INTO `role_permissions` VALUES (96, 3, 21);
INSERT INTO `role_permissions` VALUES (97, 3, 22);
INSERT INTO `role_permissions` VALUES (98, 3, 23);
INSERT INTO `role_permissions` VALUES (99, 3, 24);
INSERT INTO `role_permissions` VALUES (100, 3, 25);
INSERT INTO `role_permissions` VALUES (101, 3, 26);
INSERT INTO `role_permissions` VALUES (109, 5, 2);
INSERT INTO `role_permissions` VALUES (110, 5, 3);
INSERT INTO `role_permissions` VALUES (111, 5, 4);
INSERT INTO `role_permissions` VALUES (112, 5, 27);
INSERT INTO `role_permissions` VALUES (113, 5, 6);
INSERT INTO `role_permissions` VALUES (114, 5, 7);
INSERT INTO `role_permissions` VALUES (115, 5, 8);
INSERT INTO `role_permissions` VALUES (116, 5, 25);

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of tag
-- ----------------------------
INSERT INTO `tag` VALUES (1, 'docker');
INSERT INTO `tag` VALUES (3, 'golang');
INSERT INTO `tag` VALUES (2, 'kubernetes');
INSERT INTO `tag` VALUES (4, '公告');

-- ----------------------------
-- Table structure for topic
-- ----------------------------
DROP TABLE IF EXISTS `topic`;
CREATE TABLE `topic`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `content` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `in_time` datetime(0) NOT NULL,
  `user_id` int(11) NOT NULL,
  `view` int(11) NOT NULL DEFAULT 0,
  `reply_count` int(11) NOT NULL DEFAULT 0,
  `last_reply_user_id` int(11) NULL DEFAULT NULL,
  `last_reply_time` datetime(0) NOT NULL,
  `is_approval` tinyint(1) NOT NULL DEFAULT 0,
  `file` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of topic
-- ----------------------------
INSERT INTO `topic` VALUES (1, '新人指导', '* 帖子标题尽量描述清晰大致内容 （避免类似“新手第一次”“初次多指教”“大神快来”含糊不清标题）\r\n* 求助解决后选出一个满意的答复，这样帮助你的人也会得到一个详细的了解\r\n* 其余不明白跟帖回复', '2018-04-19 15:13:43', 1, 24, 2, NULL, '2018-04-19 15:13:43', 1, '');
INSERT INTO `topic` VALUES (2, 'Golang的特点和发起目的', 'Golang的特点以及要解决的问题概括起来就是三点: \r\n1. concurrent : 多核 解决方式-> 语言层级并发, goroutine \r\n2. garbage-collected language : c/c++的不足 \r\n3. fast compilation : c/c++等的不足, 依赖简单，类型系统简单，非传统OO。开发更简单快捷。 \r\n这种**简单设计**的特点很容易让人和C++对比，对比C++就是砍了90%特性，减少90%的麻烦。更好的对比可能是C，better c with goroutine and garbage-collection。', '2018-04-19 15:50:41', 2, 1, 0, NULL, '2018-04-19 15:50:41', 1, '');
INSERT INTO `topic` VALUES (3, 'Golang设计原则', '1. felicity of programming ： 尽可能的简化代码编写规则，这点在各种解释语言，c++11等里面都可以体现一部分了，在golang上的体现就是如包的定义，编译安装，没有头文件，no forward declarations，:= 类型推断等等\r\n2. orthogonality of concepts ： 另一个原则是概念设计尽可能正交orthogonal，这样理解使用会更简单。 Methods can be implemented for any type; structures represent data while interfaces represent abstraction; and so on. Orthogonality makes it easier to understand what happens when things combine.当然一旦设计正交，需要的概念也变得很少。\r\n3. speed of compilation', '2018-04-19 15:51:35', 2, 1, 0, NULL, '2018-04-19 15:51:35', 1, '');
INSERT INTO `topic` VALUES (4, '为什么没有泛型', '这点是golang遭受用户(尤其是c++，java用户)诟病的重要原因，实际上Golang提供了panic,recover语法类似try catch。但是个人理解为什么没有只是一个选择问题，而不是技术问题。在很多语言的编码风格里尤其是Objective-C，一般都是使用Error Object来传递错误，虽然现在try catch的性能损失可以忽略不计，但是try catch的坏处是容易滥用，导致用户忽略error和exception的区别，另外Golang提供的多返回值也方便了error传递这种风格的使用，我个人对这种设计并不反感。', '2018-04-19 15:51:52', 2, 4, 0, NULL, '2018-04-19 15:51:52', 1, 'static/upload/users/havefun/files/111.jpeg');
INSERT INTO `topic` VALUES (5, ' interface的一个有疑问的例子Why doesn’t type T satisfy the Equal interface', '```\r\ntype Equaler interface {\r\n    Equal(Equaler) bool\r\n}\r\ntype T int\r\nfunc (t T) Equal(u T) bool { return t == u } // does not satisfy Equaler\r\n```', '2018-04-19 15:52:30', 2, 5, 1, NULL, '2018-04-19 15:52:30', 1, '');
INSERT INTO `topic` VALUES (6, 'docker入门时碰到的代理设置问题', '如果是在公司内网环境，本身虚拟机就使用NAT连接外网时，环境变量中就会配置proxy。除了bashrc中为vm系统配置proxy，还要为docker本身配置代理。网上的文件不一定适应于自己的环境，我的操作系统是ubuntu-14.04.4-desktop-amd64，是在/etc/default/docker中做了修改，文件本身也有注释讲如何做配置\r\n\r\n# If you need Docker to use an HTTP proxy, it can also bespecified here.\r\n\r\n#export http_proxy=http://127.0.0.1:3128/\r\n\r\n修改完文件之后需要service docker restart，配置才会生效，然后再执行docker run hello-world，就能链接到docker的hub。', '2018-04-19 15:55:10', 2, 1, 0, NULL, '2018-04-19 15:55:10', 1, '');
INSERT INTO `topic` VALUES (7, 'Docker：设置代理proxy', '相信刚开始用docker的一定会遇到下面这种情况：\r\n```\r\ndocker@boot2docker:~$ sudo docker search ubuntu \r\nFATA[0000] Error response from daemon: Get https://index.docker.io/v1/search?q=ubuntu: \r\ndial tcp: lookup index.docker.io: no such host\r\n```', '2018-04-19 15:56:01', 2, 2, 0, NULL, '2018-04-19 15:56:01', 1, '');
INSERT INTO `topic` VALUES (8, '使用Kubernetes需要注意的一些问题（FAQ of k8s)', '本篇文章并不是介绍K8S 或者Docker的，而仅仅是使用过程中一些常见问题的汇总。\r\n1. 重启策略：http://kubernetes.io/docs/user-guide/pod-states/， 对于一个服务，默认的设置是RestartAlways，而其他的比如Job，重启策略则是Never or OnFailure，\r\n2. 如果docker重启，kubelet也要重启，否则pod的状态会变成Completed\r\n3. 如果报错是报image找不到，可能是因为与docker registry的认证关系没有建立，可以通过docker login 来解决\r\n需要pause，可以docker tag修改源。\r\n4. 需要升级内核至少3.10版本，只能使用root\r\n5. 使用Nginx，尤其是对外网服务的时候可能会很管用。', '2018-04-19 15:57:28', 1, 44, 2, NULL, '2018-04-19 15:57:28', 1, '');
INSERT INTO `topic` VALUES (9, 'ttt', '# 1', '2018-04-20 02:37:06', 2, 1, 0, NULL, '2018-04-20 02:37:06', 0, '');

-- ----------------------------
-- Table structure for topic_tags
-- ----------------------------
DROP TABLE IF EXISTS `topic_tags`;
CREATE TABLE `topic_tags`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `topic_id` int(11) NOT NULL,
  `tag_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of topic_tags
-- ----------------------------
INSERT INTO `topic_tags` VALUES (3, 1, 4);
INSERT INTO `topic_tags` VALUES (4, 2, 3);
INSERT INTO `topic_tags` VALUES (5, 3, 3);
INSERT INTO `topic_tags` VALUES (8, 5, 3);
INSERT INTO `topic_tags` VALUES (9, 6, 1);
INSERT INTO `topic_tags` VALUES (10, 7, 1);
INSERT INTO `topic_tags` VALUES (11, 8, 2);
INSERT INTO `topic_tags` VALUES (12, 4, 1);
INSERT INTO `topic_tags` VALUES (13, 4, 3);
INSERT INTO `topic_tags` VALUES (14, 4, 2);
INSERT INTO `topic_tags` VALUES (15, 9, 1);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `token` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `image` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `signature` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `in_time` datetime(0) NOT NULL,
  `active` tinyint(1) NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 0,
  `last_time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE,
  UNIQUE INDEX `token`(`token`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'root', '$2a$10$9LjYG/iOJaifWlJeAUwXOuFjDjhHrTRHdpdXk6zKYbNDMJ25Hexwq', 'fcd1cb8e-b71f-46c3-9974-7225997b40c7', '/static/imgs/default.png', '1178996513@qq.com', '欢迎大家访问我的小站', '2016-08-26 09:22:16', 1, 0, '1970-01-01 00:00:00');
INSERT INTO `user` VALUES (2, 'havefun', '$2a$10$gBA3bbxWxdO7tLBvz9LUNeFFF/FxdDbljkEcQEQOqVvIJ.BMdnZ6G', 'a181197c-dee9-4474-a47b-eab0c66f5d0f', '/static/upload/users/havefun/avatar/头像.jpeg', '23457@1qq.com', '学习使我快乐', '2018-04-19 15:46:53', 1, 0, '');

-- ----------------------------
-- Table structure for user_roles
-- ----------------------------
DROP TABLE IF EXISTS `user_roles`;
CREATE TABLE `user_roles`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `role_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of user_roles
-- ----------------------------
INSERT INTO `user_roles` VALUES (5, 1, 3);
INSERT INTO `user_roles` VALUES (7, 2, 5);

SET FOREIGN_KEY_CHECKS = 1;
