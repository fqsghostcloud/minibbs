<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">
        标签管理
        <span class="pull-right">{{.Page.TotalCount}}个</span>
      </div>
      <div class="table-responsive">
        <table class="table table-striped">
          <tbody>
          {{range .Page.List}}
          <tr>
           <td >{{.Name}}</td>
            <td id={{.Name}}>
              <a href="#" class="btn btn-xs btn-warning" onClick="inputBox({{.Name}}, {{.Id}})">修改标签</a>
              <a href="javascript:if(confirm('确认删除吗?')) location.href='/tag/manage/delete/{{.Id}}'" class="btn btn-xs btn-danger">删除</a>
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
    <div class="panel panel-default">
      <div class="panel-heading">
        添加标签
      </div>
      <div class="panel-body">
        <form class="form-inline" role="form" action="/tag/manage/save" method="post">
          <div class="form-group">
            <label for="tagName">标签名称</label>
            <input type="text" class="form-control" name="tagName" id="tagName" value="">
          </div>
          <button type="submit" class="btn btn-default">提交</button>
        </form>
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

<script type="text/javascript">
  var oldEle
  function inputBox(name, id) {
    htmlEle = document.getElementById(name)
    oldEle = htmlEle.innerHTML 
    htmlEle.innerHTML = "<form class='form-inline' role='form' action='/tag/manage/update' method='post'><div class='form-group'><label for='tagName'>标签名称</label><input type='text' class='form-control' name='tagName' id='tagName'><input type='hidden' name='id' value="
    +id+"></div><button type='submit' class='btn btn-default'>提交</button><button id='cancle' type='submit' class='btn btn-default'>取消</button></form>"
  }

  document.getElementById("cancle").addEventListener("click",function(){   
        text.innerHTML = oldEle;      
    });  
</script>









