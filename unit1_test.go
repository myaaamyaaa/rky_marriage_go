package main

import (
	"fmt"
	"src/marriage/person"
	"testing"
)

// フォルダ内のテストすべて実行(go test -v ./...)
// Test テスト(go test -v unit1_test.go)
func Test1(t *testing.T) {
	// 佐藤太郎(男性,1996/10/11生まれ) を生成
	taro := person.NewPerson("佐藤", "太郎", "male", "1996/10/11")
	fmt.Println(taro) // => {佐藤 太郎 male 19961011}

	// 佐藤太郎(男性,1996/10/11生まれ) の生年月日は 1996/10/11 である
	fmt.Println(taro.GetBirthday()) // => 1996/10/11
}
