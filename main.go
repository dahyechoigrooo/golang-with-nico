package main

import (
	"fmt"

	"github.com/saichoi/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}

	// "hello"를 "First"로 정의
	baseWord := "hello"
	dictionary.Add(baseWord, "First")

	// Search()로 단어를 검색
	dictionary.Search(baseWord)

	// Delete()로 단어를 삭제
	dictionary.Delete(baseWord)

	// 삭제된 단어를 다시 검색
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(word)
	}

}
