# Go勉強
## トレーニング案
これは配列！スライス！みたいにぱっとみで読めるように、単語帳みたいな流れやりたい

## その他
fmt.Sprintf("project_id = '%s'", projectId)
文字列作成、'%s'はカンマ以降の値がそこに入るよという意味

strings.Join(idConditions, " or ")
ここで " or " を間に入れて全部つなぎます
例えば idConditions = ["project_id = '123'", "project_id = '456'"] なら
→ "project_id = '123' or project_id = '456'"

strings.Join(projectIds, "','")
リストを,で繋ぐ

## map 
. mapとは？
mapは キーと値のペア を覚えておける箱です
「キーを指定したら値が取り出せる辞書」のようなイメージ

m := map[string]int{}
string → キーの型
int → 値の型
{} → 空のmapを作る

例
m["りんご"] = 100
m["みかん"] = 50
「りんご」の値は100
「みかん」の値は50
取り出すときは
fmt.Println(m["りんご"]) // 100

2. mapの特徴
キーはユニーク
同じキーを追加すると上書きされる
m["りんご"] = 200
fmt.Println(m["りんご"]) // 200

v, ok := m["ぶどう"]
if ok {
    fmt.Println("あるよ:", v)
} else {
    fmt.Println("ないよ")
}

3. struct{} を使う理由
値は不要で「キーだけを覚えたい」場合があります
そのときは空の構造体 struct{} を値にする
userIds := map[string]struct{}{}
userIds["user1"] = struct{}{}
userIds["user2"] = struct{}{}

v, ok := m["ぶどう"]
v → 値（存在すれば本当の値、なければゼロ値）
ok → bool型で 「キーが存在するかどうか」 を返す

## make
1. makeとは？

Goでは スライス（配列みたいなもの）やmapやチャネル を作るときに使う関数です

これらの型は 単純に var だけでは中身が作られない ので、初期化が必要です

書き方
make(型, 長さ, 容量)


型 → []int, map[string]int, chan int など

長さ → スライスの場合は初期の要素数

容量 → スライスの場合は余裕を持たせたい容量（省略可）

2. スライスでの例
s := make([]int, 3) // 長さ3のスライス
fmt.Println(s)      // [0 0 0]


長さ3なので、3個のゼロ値（intなら0）が入ったスライスが作られる

s := make([]int, 0, 10) // 長さ0、容量10
fmt.Println(s)          // []


長さ0 → 要素はまだない

容量10 → 後で最大10個まで追加してもメモリ再確保が少ない

3. mapでの例
m := make(map[string]int)
m["りんご"] = 100


var m map[string]int だけだと m は nil で使えない

make(map[string]int) で使える状態になる

4. あなたのコードの場合
ids := make([]string, 0, len(userIds))


ids は文字列スライス

長さ0 → まだ要素はない

容量 = len(userIds) → 追加する要素数が分かっているのでメモリを先に確保

その後の append(ids, k) が効率的に動く

make は「スライスやmapを使える状態に作る魔法の関数」
必要な長さや容量を指定して効率よくメモリを確保できる
## 例
cond := builder.NewCond()
for k := range userMap {
    cond = cond.Or(builder.Eq{"id": k})
}
err = session.Where(cond).Find(&userRecords)
ポイント
builder.NewCond() で 条件の箱 を作る

builder.Eq{"id": k} → 「id = k」という条件

cond.Or(...) → 条件を ORでつなぐ

例: id = 1 OR id = 2 OR id = 3

session.Where(cond).Find(&userRecords) → 条件に合うレコードを取得

Where Find この条件で探してきて