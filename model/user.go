// フォルダ名と同じにする必要あり
package model

// packageにする時は大文字
type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
