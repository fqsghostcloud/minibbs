<div class="row">
  <div class="col-md-9">
    {{if .CurrentUserInfo}}
    <div class="panel panel-default">
      <div class="panel-heading">{{.CurrentUserInfo.Username}}回复话题</div>
      <table class="table table-striped">
        <tbody>
        {{range .Page.List}}
        <tr>
          <td>
            {{.InTime | timeago}}
            回复了
            <a href="/user/{{.Topic.Id | getTopicUser}}">{{.Topic.Id | getTopicUser}}</a>
            创建的话题 › <a href="/topic/{{.Topic.Id}}">{{.Topic.Title}}</a>
          </td>
        </tr>
        <tr>
          <td><p>{{str2html (.Content | markdown)}}</p></td>
        </tr>
        {{end}}
        </tbody>
      </table>
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
          window.location.href = "/user/{{.CurrentUserInfo.Username}}/replies/?page=" + page + "&tagId={{.TagId}}"
        } else {
          window.location.href = "/user/{{.CurrentUserInfo.Username}}/replies/?page=" + page
        }
      }
    });
  });
</script>