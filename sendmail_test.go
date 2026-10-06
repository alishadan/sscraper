package sscraper

import (
	"testing"
)

func Test_Sendmail(t *testing.T) {
	sender:="hello@demomailtrap.co"
	reciver:="alishadan84@gmail.com"
	subject:="list of prices"
	text_body:="body of email"
	token:="############"

	err := SendMail(text_body,subject,sender,reciver,token)
	if err == nil {
		println("SendMail passed")
	} else {
		println("errors exist in SendMail function")
	}

}
