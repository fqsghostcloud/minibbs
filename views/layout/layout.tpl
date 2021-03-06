<!doctype html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport"
        content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>{{.PageTitle}}</title>
  <link rel="stylesheet" href="/static/bootstrap/css/bootstrap.min.css"/>
  <link rel="stylesheet" href="/static/css/minibbs.css">
  <link rel="stylesheet"href="/static/editor/css/editormd.css" />
  <script src="/static/js/jquery.min.js"></script>
  <script src="/static/js/jquery-1.10.1.min.js"></script>
  <script src="/static/bootstrap/js/bootstrap.min.js"></script>
  <script src="/static/editor/editormd.js"></script>

</head>
<body>
<div class="wrapper">
  <nav class="navbar navbar-inverse">
    <div class="container">
      <div class="navbar-header">
        <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar"
                aria-expanded="false" aria-controls="navbar">
          <span c1lass="sr-only">Toggle navigation</span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
        </button>
        <a class="navbar-brand" style="color:#fff;" href="/">Mini社区</a>
      </div>
      <div Id="navbar" class="navbar-collapse collapse header-navbar">
        <ul class="nav navbar-nav navbar-right">
          {{if .IsLogin}}
          <li>
            <a href="/user/{{.UserInfo.Username}}">
              {{.UserInfo.Username}}
            </a>
          </li>
          <li>
            <a href="javascript:;" class="dropdown-toggle" data-toggle="dropdown"
               data-hover="dropdown">
              设置
              <span class="caret"></span>
            </a>
            <span class="dropdown-arrow"></span>
            <ul class="dropdown-menu">
              <li><a href="/user/setting">个人资料</a></li>
              <li><a href="/topic/manage">发贴管理</a></li>
              <li><a href="/user/{{.UserInfo.Username}}">最近动态</a></li>
              {{if haspermission .UserInfo.Id "user:list"}}<li><a href="/user/list">用户管理</a></li>{{end}}
              {{if haspermission .UserInfo.Id "role:list"}}<li><a href="/role/list">角色管理</a></li>{{end}}
              {{if haspermission .UserInfo.Id "permission:list"}}<li><a href="/permission/list">权限管理</a></li>{{end}}
               {{if haspermission .UserInfo.Id "tag:manage"}}<li><a href="/tag/manage">标签管理</a></li>{{end}}
              <li><a href="/logout">退出</a></li>
            </ul>
          </li>
          {{else}}
          <li><a href="/login">登录</a></li>
          <li><a href="/register">注册</a></li>
          {{end}}
        </ul>
      </div>
    </div>
  </nav>
  <div class="container">
    {{.LayoutContent}}
  </div>
</div>

{{.ChatRoomScript}}
{{.TopicScript}}
</body>
</html>