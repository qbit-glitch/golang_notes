# Part 5

## Download MailHog

Link: https://github.com/mailhog/MailHog

MailHog is a lightweight package that we can add to our API to include mail sending functionality in our server. MailHog is a simple and effective email testing tool that allows developers to simulate sending and receiving emails in a safe environment without actually sending them to real email addresses. It's particularly useful for testing email functionalities in applications during development. 

Some of the key features of MailHog include an SMTP server, so MailHog runs an SMTP server that captures emails sent from your application. You can configure your application to send emails to MailHog instead of real email servers or real email addresses. This way you can review the content of the emails without sending them out. So those emails will be received to a fake SMTP server that is running on your computer and you can send as many emails as possible because no other server is receiving those emails and you don't have to login to any email account to check those emails.

Mailhog provides a web interface where we can view the emails that have been captured. You can see details like the sender, recipient, subject and body of each email, making it absolutely easy to verify that your application is sending the correct information.

Setting up mailHog is straight forward. it can be run as a standalone binary or as a Docker container. This ease of use allows developers to quickly integrate MailHog into their development workflow. With MailHog you can simulate various email scenarios without any side effects. MailHog can also be integrated into automated tests. You can check that emails are send correctly and contain the right information, making it a valuable tool for maintaining code quality.

In our API we use MailHog to send password reset emails. When a user submits their email address to the forgot password route, the application generates a password reset email containing a secure link, and this email is captured by MailHog, allowing us to verify its content through the web interface without actually sending it to the user's email address.

## Handling the Forgot Password Route


The main code for the ForgotPasswordHandler looks like the one written below. It is further refactored for clean readability and good code quality.

execs.go 
```go
func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request){
	var req struct{
		Email string `json:"email"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	r.Body.Close()

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		utils.ErrorHandler(err, "Internal Error")
		return
	}
	defer db.Close()

	// Since we want a single email address from the database, we need a single row that's we'll use a QueryRow to get the Row
	var exec models.Exec
	err = db.QueryRow("SELECT id FROM execs WHERE email=?", req.Email).Scan(&exec.ID)
	if err != nil {
		utils.ErrorHandler(err, "User not found")
		return 
	}

	duration, err := strconv.Atoi(os.Getenv("RESET_TOKEN_EXP_DURATION"))
	if err!= nil {
		utils.ErrorHandler(err, "Failed to send password reset email")
		return
	}
	mins := time.Duration(duration)

	expiry := time.Now().Add(mins * time.Minute).Format(time.RFC3339)

	tokenBytes := make([]byte, 32)
	_, err = rand.Read(tokenBytes)
	if err!= nil {
		utils.ErrorHandler(err, "Failed to password reset email")
		return
	}

	log.Println("tokenBytes:", tokenBytes)
	token := hex.EncodeToString(tokenBytes)
	log.Println("token:", token)

	hashedToken := sha256.Sum256(tokenBytes)
	log.Println("hashedToken:", hashedToken)
	
	hashedTokenString := hex.EncodeToString(hashedToken[:])

	_, err = db.Exec("UPDATE execs SET password_reset_token=?, password_token_expires=? WHERE id=?", hashedTokenString, expiry, exec.ID)
	if err != nil {
		utils.ErrorHandler(err, "Failed to send password reset email")
		return
	}

	// Send to reset email
	resetURL := fmt.Sprintf("https://localhost:3000/execs/resetpassword/reset/%s", token)
	message := fmt.Sprintf("Forgot your password ? Reset your password using the following link: \n%s\nIf you didn't request a password reset, please ignore this email. This link is only valid for %d minutes", resetURL, int(mins))

	m := mail.NewMessage()    // Creates a new instance of mail message
	m.SetHeader("From", "schooladmin@school.com")
	m.SetHeader("To", req.Email)
	m.SetHeader("Subject", "Your Password reset link")
	m.SetBody("text/plain", message)

	d := mail.NewDialer("localhost", 1025, "", "")
	err = d.DialAndSend(m)
	if err != nil {
		utils.ErrorHandler(err, "Failed to send password reset email")
		return
	}

	// respond with success confirmation
	fmt.Fprintf(w, "Password reset link sent to %s", req.Email)

}
```