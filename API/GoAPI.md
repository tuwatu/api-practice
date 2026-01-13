# 1 リクエストを受け取るとこまで
WebサーバーはFWが色々
ginとかechoとか

curl http://localhost:8080 -UseBasicParsing
ってやると単純に確認できるよ

Goでの関数定義　
func test(string testSt) 戻り値{
    処理
    return hoge
}

h1 := func(w ほげ)

Handler（ハンドラ） = 「処理を受け取って動かすもの」
HandleFunc　→何か受け取って処理する関数とかって思いそうだね

# 2 リクエストを基に特定のハンドラー層の特定の処理を呼び出す
どうやって他の層の処理を呼び出すのかわからん、、、

# 雑記
task

ginでRestAPI～

待ち受ける場所、IPとポート
受け取る仕組み

インターフェースを満たしたら使える
→インターフェースは窓口
つまりそれを満たしてたら　それ　＝　これってことになるのよ！

DefaultServeMuxはデフォルトのServeMuxであり、ServeMuxとはHTTP requestのマルチプレクサーのことです。
マルチプレクサー＝振り分け係
HandleFuncは与えられたパターンに対応するhandler関数をDefaultServeMuxに登録する　＝振り分け係の人にこれはここね～と登録するのがHandleFunc！

はいはいそんで　http.Serverでなんかポートが何とか設定できるってことね

サーバー側ハンドラーでは ResponseWriter に書く

クライアント側では Response を読む

両方同じ「レスポンス」だけど、見る方向が逆だから型が違う
