package person

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

// Person 人
type Person struct {
	familyName string
	firstName  string
	gender     string
	birthday   string
}

// NewPerson 人を生成
func NewPerson(familyName string, firstName string, gender string, birthday string) Person {

	if familyName == "" || firstName == "" {
		err := errors.New("Error: 名字か名前がありません。")
		fmt.Println(err)
		os.Exit(1)
	}

	if gender != "male" && gender != "female" {
		err := errors.New("Error: 性別はmaleかfemaleにしてください。")
		fmt.Println(err)
		os.Exit(1)
	}

	if !check(birthday) {
		err := errors.New("Error: 正しい日付を入力してください。")
		fmt.Println(err)
		os.Exit(1)
	}

	return Person{familyName: familyName, firstName: firstName, gender: gender, birthday: birthday}
}

// GetFamilyName 名字を取得
func (p Person) GetFamilyName() string {
	return p.familyName
}

// GetFirstName 名前を取得
func (p Person) GetFirstName() string {
	return p.firstName
}

// GetFullName 名前を取得
func (p Person) GetFullName() string {
	return fmt.Sprintf("%s%s", p.familyName, p.firstName)
}

// GetGender 性別を取得
func (p Person) GetGender() string {
	return p.gender
}

// GetBirthday 生年月日を取得
func (p Person) GetBirthday() string {
	return p.birthday
}

// JudgeMale 男かどうか判定
func (p Person) JudgeMale() bool {
	if p.gender == "male" {
		return true
	}
	return false
}

// JudgeFemale 女かどうか判定
func (p Person) JudgeFemale() bool {
	if p.gender == "female" {
		return true
	}
	return false
}

// JudgeMarriage 結婚できるかどうか判定
func (p Person) JudgeMarriage(partner Person, dateStr string) bool {
	if p.gender == partner.gender {
		return false
	}
	switch p.gender {
	case "male":
		if p.HowOld(dateStr) < 18 {
			return false
		}
	case "female":
		if p.HowOld(dateStr) < 16 {
			return false
		}
	}
	switch partner.gender {
	case "male":
		if p.HowOld(dateStr) < 18 {
			return false
		}
	case "female":
		if p.HowOld(dateStr) < 16 {
			return false
		}
	}
	return true
}

// HowOld 何歳か
func (p Person) HowOld(dateStr string) int {
	if !check(dateStr) {
		err := errors.New("Error: 正しい日付を入力してください。")
		fmt.Println(err)
		os.Exit(1)
	}

	birthday, _ := time.Parse("20060102", dateFormat(p.birthday))
	date, _ := time.Parse("20060102", dateFormat(dateStr))

	age, err := calcAge(birthday, date)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return age

}

func dateFormat(dateStr string) string {
	// 削除する文字列を定義
	reg := regexp.MustCompile(`[-|/|:| |　]`)

	// 指定文字を削除
	str := reg.ReplaceAllString(dateStr, "")

	return str
}

// 存在する日付かチェック
func check(dateStr string) bool {
	str := dateFormat(dateStr)

	// 数値の値に対してフォーマットを定義
	format := string([]rune("20060102")[:len(str)])

	// パース処理 → 日付ではない場合はエラー
	_, err := time.Parse(format, str)
	return err == nil
}

func calcAge(t time.Time, t1 time.Time) (int, error) {
	// 現在日時を数値のみでフォーマット (YYYYMMDD)
	dateFormatOnlyNumber := "20060102" // YYYYMMDD

	now := t1.Format(dateFormatOnlyNumber)
	birthday := t.Format(dateFormatOnlyNumber)

	// 日付文字列をそのまま数値化
	nowInt, err := strconv.Atoi(now)
	if err != nil {
		return 0, err
	}
	birthdayInt, err := strconv.Atoi(birthday)
	if err != nil {
		return 0, err
	}

	// (今日の日付 - 誕生日) / 10000 = 年齢
	age := (nowInt - birthdayInt) / 10000
	return age, nil
}
