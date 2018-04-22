<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-body topic-detail-header">
        <div class="media">
          <div class="media-body">
            <h2 class="topic-detail-title">{{.Topic.Title}}</h2>
            <p class="gray">
              {{range .TopicTags}}
              <span class="label label-primary">{{.Name}}</span>
              {{end}}
              <span>•</span>
              <span>
                <a href="/user/{{.Topic.User.Username}}">{{.Topic.User.Username}}</a>
              </span>
              <span>•</span>
              <span>{{.Topic.InTime | timeago}}</span>
              <span>•</span>
              <span>{{.Topic.View}}次点击</span>
              {{if isTopicUser .UserInfo.Id .Topic.Id }}
              <span>•</span>
              <span>
                <a href="/topic/edit/{{.Topic.Id}}">编辑</a>
              </span>
              <span>•</span>
              <span>
                <a href="javascript:if(confirm('确定删除吗?')) location.href='/topic/delete/{{.Topic.Id}}'">删除</a>
              </span>
              {{end}}


              {{if .Topic.File | hasFile}}
                <span>•</span>
                <span><a href="/topic/{{.Topic.Id}}/download">下载附件</a></span>
              {{end}}

            </p>
          </div>
          <div class="media-right">
            <img src="{{.Topic.User.Image}}" alt="{{.Topic.User.Username}}" class="avatar-lg">
          </div>
        </div>
      </div>
      <div class="divide"></div>
      <div class="panel-body topic-detail-content">
        {{str2html (.Topic.Content | markdown)}}
      </div>
    </div>
    {{if eq .Topic.ReplyCount 0}}
    <div class="panel panel-default">
      <div class="panel-body text-center">目前暂无回复</div>
    </div>
    {{else}}
    <div class="panel panel-default">
      <div class="panel-heading">{{.Topic.ReplyCount}} 条回复</div>
      <div class="panel-body paginate-bot">
        {{range .Replies}}
        <div class="media">
          <div class="media-left">
            <a href="/user/{{.User.Username}}">
              <img src="{{.User.Image}}" class="avatar" alt="{{.User.Username}}">
            </a>
          </div>
          <div class="media-body reply-content">
            <div class="media-heading gray">
              <a href="/user/{{.User.Username}}">{{.User.Username}} </a>
              <span>{{.InTime | timeago}}</span>
              <span class="pull-right">
                {{if haspermission $.UserInfo.Id "reply:delete"}}
                <a href="javascript:if(confirm('确定删除吗?')) location.href='/reply/delete/{{.Id}}'">删除</a>{{end}} {{if $.IsLogin}}
                <a href="javascript:up('{{.Id}}');">
                  <span class="glyphicon glyphicon-thumbs-up"></span>
                </a>{{end}}
                <span Id="up_{{.Id}}">{{.Up}}</span>
                <span>赞</span>
              </span>
            </div>
            {{str2html (.Content | markdown)}}
          </div>
        </div>
        <div class="divide mar-top-5"></div>
        {{end}}
      </div>
    </div>
    {{end}} {{if .IsLogin}}
    <div class="panel panel-default">
      <div class="panel-heading">
        添加一条新回复
      </div>
      <div class="panel-body">
        <form action="/reply/save" method="post">
          <input type="hidden" value="{{.Topic.Id}}" name="tid">
          <div class="form-group">
            <textarea name="content" rows="5" class="form-control" placeholder="支持Markdown语法哦~"></textarea>
          </div>
          <button type="submit" class="btn btn-default">回复</button>
        </form>
      </div>
    </div>
    {{end}}
  </div>
  <div class="col-md-3 hidden-sm hidden-xs">
    {{template "components/author_info.tpl" .}}
  </div>
</div>