package mydict

import "errors"

// struct가 아닌 type으로 map 생성
// Dictionary : map의 alias(별칭)
// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Can't update non-exsiting word")
	errWordExists = errors.New("That word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if exist {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	// 이미 사전에 등록되어 있는 단어라면... Search()에서 error를 받는 경우이다.
	_, err := d.Search(word) // 단어의 뜻이 필요한 것이 아니기 때문에 _를 사용하여 값을 무시한다.

	// switch 문의 경우
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil

}

// Update a word
func (d Dictionary) Update(word, definition string) error {
	// 수정하기 위한 단어가 있는지 검색한다.
	_, err := d.Search(word)
	switch err {
	// 에러가 없다. 즉 단어가 있다면 단어를 수정할 수 있다.
	case nil:
		d[word] = definition
	// 에러가 있다. 즉 단어가 없다면 단어를 수정할 수 없다.
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
