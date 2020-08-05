package utils

// FormAddMenu
type FormAddMenu struct {
	Name     string `form:"name" binding:"required"`
	ParentId uint   `form:"parentId" binding:"required"`
	Router   string `form:"router" binding:"required"`
	Icon     string `form:"icon" binding:"required"`
}
