package models

type Teacher struct {
	Uid      int64   `json:"uid" validate:"required,min=4,max=32"`
	Name     string  `json:"name" validate:"required,min=4,max=32"`
	Gender   bool    `json:"gender" validate:"required,man=true,woman=false"`
	Age      int64   `json:"age" validate:"required,min=1,max=99"`
	Class    []Class `json:"class" validate:"dive"`
	Password string  `json:"password" validate:"required,number"`
	IsAdmin  bool    `json:"is_admin" validate:"required"`
	Email    string  `json:"email" validate:"required,email,min=6,max=32"`
	Phone    int64   `json:"phone" validate:"required,min=6,max=32"`
	Job      Job     `json:"job" validate:"dive"`
}

type Job struct {
	Type   string `json:"type" validate:"required,min=4,max=32"`
	Salary string `json:"salary" validate:"required,min=4,max=32"`
}

type Class struct {
	Number        int64 `json:"number" validate:"required,min=0,max=255"`
	StudentNumber int64 `json:"student_number" validate:"required,min=0,max=255"`
}
