package models

type Student struct {
	Uid     int64     `json:"uid" validate:"required,min=4,max=32"`
	Name    string    `json:"name" validate:"required,min=4,max=32"`
	Gender  bool      `json:"gender" validate:"required,man=true,woman=false"`
	Age     int64     `json:"age" validate:"required,min=1,max=99"`
	Class   Class     `json:"class" validate:"dive"`
	Subject []Subject `json:"subject" validate:"dive"`
}

type Subject struct {
	Name    string  `json:"name" validate:"required,min=4,max=32"`
	Record  int64   `json:"record" validate:"required,min=4,max=32"`
	Teacher Teacher `json:"teacher"`
}
