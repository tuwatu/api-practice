package service

import (
	"encoding/json"
	"os"

	"github.com/tse-estrella/go-api/project/model"
)

// ユーザー登録
func CreateUser(u model.User) bool {

	// JSON に変換
	// オブジェクトを指定したインデント付きでmarshalできる
	jsonUser, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return false
	}

	// ファイルに書き込む（存在しなければ作成）
	// os.OpenFile(ファイル名, フラグ, ファイルモード)
	// フラグ:オプションを指定してる
	// os.O_CREATE:存在しない場合は新規作成
	// os.O_WRONLY:書き込み専用で開く
	// os.O_APPEND:書き込むときに 既存の内容の末尾に追加
	// ファイルモード:8進数で指定、所有者のみOK、他は読みとりのみ
	f, err := os.OpenFile("users.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return false
	}
	defer f.Close()

	// 改行付きで追記
	if _, err := f.Write(append(jsonUser, '\n')); err != nil {
		return false
	}

	// 成功可否を返す
	return true
}

func Login(u model.User) bool {
	// ログイン可否を返す
	// JSON読み取る
	// その中にあるか判別
	// 結果を返す
}
