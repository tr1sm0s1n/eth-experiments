package helpers

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"strings"
	"time"
)

func SendAlert(err error) error {
	auth := smtp.PlainAuth("", os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASS"), os.Getenv("MAIL_HOST"))

	buf := bytes.NewBuffer(nil)
	if _, err := fmt.Fprintf(buf, "Subject: [%s] lab/data-pumper: Error Occured!\n", time.Now().Format("02/01/2006")); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(buf, "To: %s\n", os.Getenv("MAIL_RECIPIENTS")); err != nil {
		return err
	}

	if _, err := buf.WriteString("Content-Type: text/html; charset=\"UTF-8\";\n\n"); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(buf, "<div style='font-size:large;font-family:EB Garamond;'><p>Dear all,</p><p>Unfortunately, the data pumper has experienced an error and crashed.</p><p>The error is as follows:</p><p><b><q>%s</q></b></p><p>Kindly resolve this error as soon as possible.</p><p>P.S. This is a system-genrated mail. Please do not reply.</p><br/>", err.Error()); err != nil {
		return err
	}

	recipients := strings.Split(os.Getenv("MAIL_RECIPIENTS"), ",")
	return smtp.SendMail(fmt.Sprintf("%s:%s", os.Getenv("MAIL_HOST"), os.Getenv("MAIL_PORT")), auth, os.Getenv("MAIL_USER"), recipients, buf.Bytes())
}
