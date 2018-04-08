<div class="row">
  <div class="col-md-6 col-md-offset-3">
    <div class="panel panel-default">
      <div class="panel-heading">登录</div>
      <div class="panel-body">
        {{template "components/flash_error.tpl" .}}
        {{template "components/flash_success.tpl" .}}
        <form action="/login" method="post">
          <div class="form-group">
            <label for="username">用户名</label>
            <input type="text" id="username" name="username" class="form-control" placeholder="用户名">
          </div>
          <div class="form-group">
            <label for="password">密码</label>
            <input type="password" id="password" name="password" class="form-control" placeholder="密码">
          </div>
          <div class="form-group">
            <label class="radio-inline">
              <input type="radio" name="role" value="普通用户" checked> 普通用户
            </label>
            <label class="radio-inline ">
                <input type="radio" name="role"  value="管理员"> 管理员
            </label>
          </div>
          <input type="submit" class="btn btn-default" value="登录"> <a href="/register">注册</a>
        </form>
      </div>
    </div>
  </div>
</div>