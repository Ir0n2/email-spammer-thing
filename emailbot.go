
var email string

var num int

fmt.Println("To whom is this to?")
fmt.Scanln(&email)

fmt.Println("How many emails will you send? (note: more emails take longer)")
fmt.Scan(&num)

for i := 1; i <= num; i++ {

// Configuration	
	from := "email"
	password := "password"
	to := []string{email}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("Subject: So no header?\nuh oh spahgetti hoes")



        // Create authentication
        auth := smtp.PlainAuth("", from, password, smtpHost)

        // Send actual message
        err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
        if err != nil {
                log.Fatal(err)
}

if (i == num) {
        fmt.Println("emails sent!")
}
        }
}
                                                                                                                                                                                                                                             
