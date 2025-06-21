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
	if _, err := fmt.Fprintf(buf, "Subject: [ALERT-%s] Data Pumper Failure — Error Details Inside\n",
		time.Now().Format("02/01/2006")); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(buf, "To: %s\n", os.Getenv("MAIL_RECIPIENTS")); err != nil {
		return err
	}

	if _, err := buf.WriteString("Content-Type: text/html; charset=\"UTF-8\";\n\n"); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(buf, `
		<div style="font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; font-size: 16px; color: #333;">
  		  <p>Hi team,</p>

  		  <p>
    		The <strong>Data Pumper</strong> service has encountered an unexpected error and has stopped running.
  		  </p>

  		  <p><strong>Error Details:</strong></p>
  		  <blockquote style="background-color:#f9f9f9; border-left: 4px solid #d9534f; padding: 10px; margin: 10px 0;">
    		<code style="font-family: Consolas, monospace; font-size: 14px;">%s</code>
  		  </blockquote>

  		  <p>
    		Please investigate and resolve the issue at your earliest convenience to restore normal operation.
  		  </p>

  		  <hr style="border: none; border-top: 1px solid #ccc;" />

  		  <p style="font-size: 13px; color: #888;">
    		This is an automated notification sent by the monitoring system. Please do not reply to this email.
  		  </p>
		</div>
	`, err.Error()); err != nil {
		return err
	}

	recipients := strings.Split(os.Getenv("MAIL_RECIPIENTS"), ",")
	return smtp.SendMail(fmt.Sprintf("%s:%s", os.Getenv("MAIL_HOST"), os.Getenv("MAIL_PORT")), auth, os.Getenv("MAIL_USER"), recipients, buf.Bytes())
}
