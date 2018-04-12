<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">
        <ul class="nav nav-pills">
          <li id="tab_0" style="font-size:20px"><a href="/?tab=all">全部</a></li>
          {{range .Tags}}
          <li id="tab_{{.Id}}" style="font-size:20px"><a href="/?tagId={{.Id}}">{{.Name}}</a></li>
          {{end}}
        </ul>
      </div>
      <div class="panel-body paginate-bot">
        {{range .Page.List}}
        <div class="media">
          <div class="media-left">
            <a href="/user/{{.User.Username}}"><img src="{{.User.Image}}" class="avatar" alt="{{.User.Username}}"></a>
          </div>
          <div class="media-body">
            <div class="title">
              <a href="/topic/{{.Id}}">{{.Title}}</a>
            </div>
            <p class="gray">
             {{range  .Title | getTopicTags}}
            <span class="label label-primary">{{.Name}}</span>
            {{end}}
              <span>•</span>
              <span><a href="/user/{{.User.Username}}">{{.User.Username}}</a></span>
              <span class="hidden-sm hidden-xs">•</span>
              <span class="hidden-sm hidden-xs">{{.ReplyCount}}个回复</span>
              <span class="hidden-sm hidden-xs">•</span>
              <span class="hidden-sm hidden-xs">{{.View}}次浏览</span>
              <span>•</span>
              <span>{{.InTime | timeago}}</span>
              {{if .LastReplyUser}}
                <span>•</span>
                <span>最后回复来自 <a href="/user/{{.LastReplyUser.Username}}">{{.LastReplyUser.Username}}</a></span>
              {{end}}
               <span>•</span>
               <a href="/topic/join/ws?uname={{$.UserInfo.Username}}&tid={{.Id}}"><span class="hidden-sm hidden-xs">聊天室</span></a>
            </p>
          </div>
        </div>
        <div class="divide mar-top-5"></div>
        {{end}}
        <ul id="page"></ul>
      </div>
    </div>
  </div>
  <div class="col-md-3 hidden-sm hidden-xs">
    {{if .IsLogin}}
      {{template "components/user_info.tpl" .}}
      {{template "components/topic_create.tpl" .}}
    {{else}}
      {{template "components/welcome.tpl" .}}
    {{end}}
    {{template "components/otherbbs.tpl" .}}
  </div>
</div>
<script type="text/javascript" src="/static/js/bootstrap-paginator.min.js"></script>
<script type="text/javascript">
  $(function () {
    $("#tab_{{.TagId}}").addClass("active");
    $("#page").bootstrapPaginator({
      currentPage: '{{.Page.PageNo}}',
      totalPages: '{{.Page.TotalPage}}',
      bootstrapMajorVersion: 3,
      size: "small",
      onPageClicked: function(e,originalEvent,type,page){
        var tagId = {{.TagId}};
        if (tagId > 0) {
          window.location.href = "/?page=" + page + "&tagId={{.TagId}}"
        } else {
          window.location.href = "/?page=" + page
        }
      }
    });
  });
</script>