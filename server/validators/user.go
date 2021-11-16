package validators

import (
	"github.com/J-Obog/pomodoro/data"
	"github.com/J-Obog/pomodoro/models"
	"github.com/gookit/validate"
)

func ValidateUserReg(m map[string]interface{}) (error) {
	v := validate.New(m)
	
	v.AddValidator("uniqueEmail", func(v interface{}) bool {
		e := data.DB.Where("email = ?", v).First(&models.User{}).Error
		return e != nil
	})

	v.ConfigRules(map[string]string {
		"email": "required|email|uniqueEmail",
		"password": "required|minLen:8", 
	})
	
	v.WithTranslates(map[string]string{
		"email": "Email", 
		"password": "Password",
	})
	v.AddMessages(map[string]string{
		"required": "{field} is required",
		"min": "{field} must be a minimum of %d characters",
		"uniqueEmail": "Account with email already exists",
	})

	if !v.Validate() {
		return v.Errors
	}

	return nil
}