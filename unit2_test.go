package main

import (
	"fmt"
	"src/marriage/person"
	"testing"
)

// フォルダ内のテストすべて実行(go test -v ./...)
// Test テスト(go test -v unit2_test.go)
func Test2(t *testing.T) {
	// 佐藤太郎(男性,1996/10/11生まれ) を生成
	taro := person.NewPerson("佐藤", "太郎", "male", "1996/10/11")
	fmt.Println(taro) // => {佐藤 太郎 male}

	// 山田花子(女性,1994/7/30生まれ) を生成
	hanako := person.NewPerson("山田", "花子", "female", "1994/07/30")
	fmt.Println(hanako) // => {山田 花子 female}

	// 佐藤太郎(男性,1996/10/11生まれ) は 山田花子(女性,1994/7/30生まれ) と 2014/10/11 に結婚できる
	fmt.Println(taro.JudgeMarriage(hanako, "2014/10/11")) // => true

	// 佐藤太郎(男性,1996/10/11生まれ) は 山田花子(女性,1994/7/30生まれ) と 2014/10/10 には結婚できない
	fmt.Println(taro.JudgeMarriage(hanako, "2014/10/10")) // => false
}
