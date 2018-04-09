<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">
        <a href="/">主页</a> / 编辑话题
      </div>
      <div class="panel-body">
        {{template "../components/flash_error.tpl" .}}
        <form method="post" action="/topic/edit/{{.Topic.Id}}">
          <div class="form-group">
            <label for="title">标题</label>
            <input type="text" class="form-control" Id="title" name="title" value="{{.Topic.Title}}" placeholder="标题">
          </div>
          <div class="form-group">
            <label for="title">内容</label>
            <textarea name="content" Id="content" rows="15" class="form-control" placeholder="支持Markdown语法哦~">{{.Topic.Content}}</textarea>
          </div>
          <div class="form-group">
            <label for="title">标签</label>
            <select name="tid" Id="tid" class="form-control">
              {{range .Tags}}
                <option value="{{.Id}}">{{.Name}}</option>
              {{end}}
            </select>
          </div>
          <button type="submit" class="btn btn-default">发布</button>
        </form>
      </div>
    </div>
  </div>
</div>
<script type="text/javascript">
  $(function () {
    {{range .TopicTags}}
    $("#tid").val('{{.Id}}')
    {{end}}
  });
</script>