package dto

type CreateSubjectRequest struct {
	Name string `json:"name" binding:"required,min=2,max=20"`
}

type SubjectResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
}