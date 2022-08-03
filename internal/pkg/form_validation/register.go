package formvalidation

type Register struct {
	Email    string `form:"email" binding:"required,email"`
	Name     string `form:"name" binding:"required"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
