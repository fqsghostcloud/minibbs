<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
     {{template "components/flash_error.tpl" .}}
     {{template "components/flash_success.tpl" .}}
     <div class="panel-heading">
        <form class="form-inline" role="form" action="/topic/manage" method="get">
          <div class="form-group">
            <input type="text" class="form-control" name="topicName" placeholder="请输入帖子名称">
          </div>
          <button type="submit" class="btn btn-default">搜索</button>
        </form>
      </div>
      <div class="panel-heading">
        帖子管理
        <span class="pull-right">{{.Page.TotalCount}}篇</span>
      </div>
      <div class="table-responsive">
        <table class="table table-striped">
          <tbody>
          {{range .Page.List}}
          <tr>
            <td><a href="/topic/{{.Id}}">{{.Title}}</a></td>
            <td>
              {{if $.UserInfo | isAdmin}} 
                {{if .IsApproval}}
                  <a href="/topic/manage/{{.Id}}/notapproval" class="btn btn-xs btn-danger">打回帖子</a>
                {{else}}
                  <a href="/topic/manage/{{.Id}}/approval" class="btn btn-xs btn-success">允许发布</a>
                {{end}}
              {{else}}
                {{if .IsApproval}}
                  <span class="label label-success">审批通过</span>
                {{else}}
                  <span class="label label-warning">审批未通过</span>
                {{end}}
              {{end}}   
             
              <a href="javascript:if(confirm('确认删除吗?')) location.href='/topic/delete/{{.Id}}'" class="btn btn-xs btn-danger">删除</a>
            </td>
          </tr>
          {{end}}
          </tbody>
        </table>
      </div>
      <div class="panel-body" style="padding: 0 15px;">
        <ul id="page"></ul>
      </div>
    </div>
  </div>
  <div class="col-md-3 hidden-sm hidden-xs">

  </div>
</div>
<script type="text/javascript" src="/static/js/bootstrap-paginator.min.js"></script>
<script type="text/javascript">
  $(function () {
    $("#page").bootstrapPaginator({
      currentPage: '{{.Page.PageNo}}',
      totalPages: '{{.Page.TotalPage}}',
      bootstrapMajorVersion: 3,
      size: "small",
      onPageClicked: function(e,originalEvent,type,page){
        window.location.href = "/?p=" + page
      }
    });
  });
</script>