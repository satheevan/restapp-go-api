package RestaurantController

import (
	"net/http"

	"gopkg.in/gomail.v2"
)

// type RestaurantController struct {
// 	controllers.BaseController
// }

func (rc RestaurantController) Index() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rc.Json(rw, &struct{ Data string }{
			Data: "restaurants",
		})
	}
}

// email services
func SendSignupEmail(toEmail, username string) error {
	// Compose the email content
	email := gomail.NewMessage()
	email.SetHeader("From")
	email.SetHeader("To", toEmail)
	email.SetHeader("Subject", "Welcome to Your App!")
	email.SetBody("text/html", "Hello "+username+",<br><br>Welcome to Your App!")

	// Configure the SMTP sender
	d := gomail.NewDialer("smtp.example.com", 587, "your-email@example.com", "your-password")

	// Send the email
	if err := d.DialAndSend(email); err != nil {
		return err
	}

	return nil
}
