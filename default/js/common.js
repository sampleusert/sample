$(document).ready(function(){
    $(".modal").modal();
});
    
$('#titleok').click(function(){
    alert("abc");
});

var titleCreate = function(){
    appendProgress();
    
    $.ajax({
    // リクエストメソッド(GET,POST,PUT,DELETEなど)
    type: 'GET',
    // リクエストURL
    url: '/titleCreate',
    // タイムアウト(ミリ秒)
    timeout: 10000,
    // キャッシュするかどうか
    cache: false,
    // サーバに送信するデータ(name: value)
    /*data: {
      'param1': 'ほげ',
      'foo': 'データ'
    },*/
    data: $("form").serialize(),
    // レスポンスを受け取る際のMIMEタイプ(html,json,jsonp,text,xml,script)
    // レスポンスが適切なContentTypeを返していれば自動判別します。
    dataType: 'json',
    // Ajax通信前処理
    beforeSend: function(jqXHR) {
      // falseを返すと処理を中断
      return true;
    }//,
    // コールバックにthisで参照させる要素(DOMなど)
    //context: domobject
  }).done(function(response, textStatus, jqXHR) {
    // 成功時処理
    alert("ok");
    removeProgress();
    //レスポンスデータはパースされた上でresponseに渡される
  }).fail(function(jqXHR, textStatus, errorThrown ) {
    // 失敗時処理
  }).always(function(data_or_jqXHR, textStatus, jqXHR_or_errorThrown) {
    // doneまたはfail実行後の共通処理
  });
};

var appendProgress = function() {
    $('#status').append('<div class="progress" id="progress"><div class="indeterminate"></div></div>');
}

var removeProgress = function() {
    $('#progress').remove();
}