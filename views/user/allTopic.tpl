<div class="row">
  <div class="col-md-9">
    {{if .CurrentUserInfo}}
    <div class="panel panel-default">
      <div class="panel-heading">{{.CurrentUserInfo.Username}}创建的话题</div>
      <div class="panel-body">
        {{range .Page.List}}
        <div class="media">
          <div class="media-body">
            <div class="title">
              <a href="/topic/{{.Id}}">{{.Title}}</a>
            </div>
            <p>
            {{range .Tags}}
              <a href="/?tab={{.Id}}">{{.Name}}</a>
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
            </p>
          </div>
        </div>
        <div class="divide mar-top-5"></div>
        {{end}}
      </div>
      <div class="panel-footer">
        <ul id="page"></ul>
      </div>
    </div>
    {{else}}
    <div class="panel panel-default">
      <div class="panel-body">用户不存在</div>
    </div>
    {{end}}
  </div>
  <div class="col-md-3 hidden-sm hidden-xs">
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
          window.location.href = "/user/{{.CurrentUserInfo.Username}}/topics/?page=" + page + "&tagId={{.TagId}}"
        } else {
          window.location.href = "/user/{{.CurrentUserInfo.Username}}/topics/?page=" + page
        }
      }
    });
  });
</script>