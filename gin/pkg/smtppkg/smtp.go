package smtppkg

import "net/smtp"

func PlainAuth(identity, username, password, host string) smtp.Auth {
	return smtp.PlainAuth(identity, username, password, host)
}

// template example:
//   t, _ := template.ParseFiles("template.html")

//   var body bytes.Buffer

//   mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
//   body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

//	t.Execute(&body, struct {
//	  Name    string
//	  Message string
//	}{
//
//	  Name:    "Name",
//	  Message: "Name",
//	})
func SendMail(host, port string, auth smtp.Auth, from string, to []string, message []byte) error {
	return smtp.SendMail(host+":"+port, auth, from, to, message)
}
