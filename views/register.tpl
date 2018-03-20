<div class="row">
  <div class="col-md-6">
    <div class="panel panel-default">
      <div class="panel-heading">注册</div>
      <div class="panel-body">
        {{template "components/flash_error.tpl" .}}
        {{template "components/flash_success.tpl" .}}
        <form action="/register" method="post">
          <div class="form-group">
            <label for="username">用户名</label>
            <input type="text" id="username" name="username" class="form-control" placeholder="用户名">
          </div>
          <div class="form-group">
            <label for="password">密码</label>
            <input type="password" id="password" name="password" class="form-control" placeholder="密码">
          </div>
           <div class="form-group">
            <label for="email">邮箱</label>
            <input type="email" id="email" name="email" class="form-control" placeholder="邮箱">
          </div>
          <input type="submit" class="btn btn-sm btn-default" value="注册"> <a href="/login">登录</a>
        </form>
      </div>
    </div>
  </div>
</div>