<script type="text/javascript">
  function up(Id) {
    var isLogin = {{.IsLogin}}; // 不可以分行，否则js代码出错
    if (isLogin) {
      $.ajax({
        url: "/reply/up",
        async: true,
        cache: false,
        type: "get",
        dataType: "json",
        data: {
          rid: Id
        },
        success: function (data) {
          if (data.Code == 200) {
            var upele = $("#up_" + Id);
            upele.text(parseInt(upele.text()) + 1);
          } else if (data.Code == 201) {
            var upele = $("#up_" + Id);
            upele.text(parseInt(upele.text()) - 1);
          } else {
            alert(data.Description)
          }
        }
      });
    } else {
      alert("请先登录");
    }
  }
</script>