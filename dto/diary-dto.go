package dto

type DiaryUpdateDTO struct {
	ID     uint64 `json:"id" form:"id" binding:"required"`
	Title  string `json:"title" form:"title" binding:"required"`
	Body   string `json:"body" form:"body" binding:"required"`
	UserID uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type DiaryCreateDTO struct {
	Title  string `json:"title" form:"title" binding:"required"`
	Body   string `json:"body" form:"body" binding:"required"`
	UserID uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
