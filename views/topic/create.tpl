<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">
        <a href="/">主页</a> / 发布话题
      </div>
      <div class="panel-body">
        {{template "../components/flash_error.tpl" .}}
        <form method="post" action="/topic/create" enctype="multipart/form-data">
          <div class="form-group">
            <label for="title">标题</label>
            <input type="text" class="form-control" Id="title" name="title" placeholder="标题">
          </div>
          <div class="form-group">
            <label for="title">内容</label>
            <textarea name="content" Id="content" rows="15" class="form-control" placeholder="支持Markdown语法哦~"></textarea>
          </div>
          <div class="form-group">
            <label for="title">标签</label>
             <div>
              {{range .Tags}}
                <input type="checkbox" name="tids" value="{{.Id}}" id="tag_{{.Id}}">
                <label for="{{.Id}}">{{.Name}}</label>&nbsp;
              {{end}}
            </div>
          </div>
          <div class="form-group">
            <label >选择文件</label>
            <input type="file" class="form-control" name="file">
          </div>
          <button type="submit" class="btn btn-default">发布</button>
        </form>
      </div>
    </div>
  </div>
</div>