

## 依赖

- [github.com/astaxie/beego](https://github.com/astaxie/beego)
- [github.com/astaxie/beego/context](https://github.com/astaxie/beego/context)
- [github.com/astaxie/beego/orm](https://github.com/astaxie/beego/orm)
- [github.com/xeonx/timeago](https://github.com/xeonx/timeago)
- [github.com/russross/blackfriday](https://github.com/russross/blackfriday)
- [github.com/sluu99/uuid](https://github.com/sluu99/uuid)
- [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)


## 注意
- 如果访问地址不是localhost,需要修改conf/app.conf文件里的cookie.domain,否则登录后不会记录登录状态  



# 开发进度
## 管理员
### 用户信息（增删改查）
- [X] 查找
- [X] 删除用户
- [ ] 拉黑邮箱

### 角色管理（增删改查）
- [X] 普通管理员
- [X] 用户
- [X] 游客

### 普通管理员——用户管理
- [X] 查看用户信息
- [ ] 冻结用户
- [X] 删除用户

### 普通管理员——帖子管理
- [X] 设置帖子分类标签
- [X] 修改删除帖子
- [ ] 帖子加精顶置
- [X] 公告公示（增删改查）

### 普通管理员——统计管理
- [ ] 用户活跃度统计
- [ ] 标签活跃度统计

## 用户

### 注册
- [X] 密码储存加密

### 登录
- [ ] 设置登录状态
- [ ] 多次登录失败输入验证码

### 找回密码
- [ ] 邮箱找回
- [ ] 冻结账号

### 好友删除
- [ ] 搜索删除
- [ ] 批量删除

### 好友分组
- [ ] 特别关注
- [ ] 黑名单

### 信息管理
- [X] 设置个人信息
- [ ] 设置信息可见性
- [ ] 登录状态管理

### 帖子管理
- [X] 发帖（支持markdown/定时删除/多标签分类）
- [X] 删贴
- [X] 查看（点赞数/评论信息）
- [X] 设置帖子（分类）
- [ ] 顶置帖子(搜索标签之后)

### 聊天室(基于主题帖)
- [X] 帖主（创建/删除/更改）
- [X] 聊天
- [X] 聊天历史记录查看



## 记录
* 发帖支持图片
* 无权限时，不显示按钮
* 伪数据完善
* 各部分操作的提示信息
* 聊天室布局
* 列表显示布局
* 整体布局不同分辨率
* div 布局

** 多对多查询时，下标为__(2个) **


 


