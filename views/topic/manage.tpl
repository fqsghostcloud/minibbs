<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
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
              <a href="/topic/edit/{{.Id}}" class="btn btn-xs btn-warning">修改帖子</a>
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