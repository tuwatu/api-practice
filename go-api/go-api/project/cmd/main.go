package main

import (
	"net/http"

	"github.com/tse-estrella/go-api/project/handler"
)

// 参考にしたWebサイト
// https://tech.yappli.io/entry/2022/05/16/Go%E3%81%AEnet/http%E3%83%91%E3%83%83%E3%82%B1%E3%83%BC%E3%82%B8%E3%82%92%E7%90%86%E8%A7%A3%E3%81%99%E3%82%8B
func main() {
	// ハンドラー層に渡す
	handler := func(w http.ResponseWriter, p *http.Request) {
		// ここにハンドラー層にリクエストを渡す処理
		// パッケージ名 + 関数名で呼び出す
		w = handler.RequestClassification(p)
	}

	http.HandleFunc("/", handler)
	// nilにするとDefaultServeMuxが使われる
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
