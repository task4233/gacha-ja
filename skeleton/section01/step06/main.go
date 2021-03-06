// STEP06: レア度ごとに出る確率を変えてみよう

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())

	// 0から99までの間で乱数を生成する
	num := rand.Intn(100)

	// TODO: 変数numが0〜79のときは"ノーマル"、
	// 80〜94のときは"R"、95〜98のときは"SR"、
	// それ以外のときは"XR"と表示する
	var msg string
	switch {
	case 0 <= num && num < 80:
		msg = "ノーマル"
		break
	case 80 <= num && num < 95:
		msg = "R"
		break
	case 95 <= num && num < 98:
		msg = "SR"
		break
	default:
		msg = "XR"
		break
	}
	fmt.Println(msg)
}
