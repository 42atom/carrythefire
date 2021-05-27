package service

import (
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SMTPMailTo(subject, msg string) error {
	host := viper.GetString("email.host")
	port := viper.GetInt("email.port")
	userName := viper.GetString("email.username")
	password := viper.GetString("email.password")
	to := viper.GetStringSlice("email.to")
	if len(to) <= 0 {
		return fmt.Errorf("config email.to is empty")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", userName)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", msg)

	d := gomail.NewDialer(host, port, userName, password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
