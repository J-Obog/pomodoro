package validators

import (
	"github.com/gookit/validate"
)

func ValidateNewTask(m map[string]interface{}) (error) {
	v := validate.New(m)
	
	v.ConfigRules(map[string]string {
		"title": "required", 
	})
	
	v.WithTranslates(map[string]string{
		"title": "Task title", 
	})
	
	v.AddMessages(map[string]string{
		"required": "Tasks must have a title",
	})

	if !v.Validate() {
		return v.Errors
	}

	return nil
}
