package dto

type SaveUserInput struct {
	Email    string
	PassHash []byte
}
