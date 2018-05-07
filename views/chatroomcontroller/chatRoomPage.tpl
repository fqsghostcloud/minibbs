<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">
        <h3>关于--{{.TopicName}}--编号
          <span id="tid">{{.TopicId}}</span>
        </h3>
      </div>
      <div class="panel-body">
        <div class="container" style="background-color:#f5f5f5;width:600px;height:460px;overflow: auto;overflow-x:hidden">
              <h3>聊天历史</h3>
              <ul id="chatbox">
                <li>欢迎你，
                  <span id="uname">{{.UserName}}</span>
                </li>
                {{range .History}} {{range $name, $message := .}}
                <li>{{$name}}:{{$message}}</li>
                {{end}} {{end}}

              </ul>
            </div>
      </div>
      <div class="panel-footer">
        <form class="form-horizontal" role="form">
          <div class="form-group">
            <label for="firstname" class="col-sm-2 control-label">发送消息</label>
            <div class="col-sm-8">
              <textarea id="sendbox" type="text" class="form-control" placeholder="请输入消息" onkeydown="if(event.keyCode==13)return false;"
                required></textarea>
            </div>
          </div>

          <div class="form-group">
            <div class="col-sm-offset-2 col-sm-10">
              <button id="sendbtn" type="button" class="btn btn-default">发送</button>
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
  <div class="col-md-3 hidden-sm hidden-xs">
    {{template "components/user_info.tpl" .}} {{template "components/topic_create.tpl" .}} {{template "components/otherbbs.tpl".}}
  </div>
</div>