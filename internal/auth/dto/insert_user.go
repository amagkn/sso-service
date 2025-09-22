package dto

type InsertUserInput struct {
	Email    string
	PassHash []byte
}
