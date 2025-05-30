package utils

import (
	"context"
	"fmt"
	"net/smtp"
	"regexp"
	"time"

	"github.com/jordan-wright/email"
	"go-cloud-disk/conf"
)

func VerifyEmailFormat(email string) bool {
	pattern := `^[^\s@]+@[^\s@]+\.[^\s@]+$` //match email
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// sendMessage use defaultSmtpAuth to send email, If it runs for more
// than 900ms, it is automatically considered to have been sent successfully.
func sendMessage(ctx context.Context, em *email.Email) {
	c, cancel := context.WithTimeout(ctx, time.Millisecond*900)
	go func() {
		em.Send(conf.EmailSMTPServer+":25", smtp.PlainAuth("", conf.EmailAddr, conf.EmailSecretKey, conf.EmailSMTPServer))
		defer cancel()
	}()

	select {
	case <-c.Done():
		return
	case <-time.After(time.Millisecond * 900):
		return
	}
}

// SendConfirmMessage send confirm code to target mailbox,
// this func will return err when send email exceed 5 second
// or connect send email web err
func SendConfirmMessage(targetMailBox string, code string) error {
	em := email.NewEmail()
	em.From = fmt.Sprintf("Go-Cloud-Disk <%s>", conf.EmailAddr)
	em.To = []string{targetMailBox}

	// email title
	em.Subject = "Email Confirm Code " + code

	// build email content
	emailContentCode := "you confirm code is " + code + ", Your code will expire in 30 minutes"
	emailContentEmail := "you confirm email is " + targetMailBox
	emailContent := emailContentCode + "\n" + emailContentEmail
	em.Text = []byte(emailContent)

	// send message
	sendMessage(context.Background(), em)

	return nil
}
