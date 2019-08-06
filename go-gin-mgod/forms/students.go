package forms

type CreateStudentCommand struct {
	Name     string `json:"name" `
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Gender   string `json:"gender" `
}

type UpdateStudentCommand struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
}
