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

## Reset Password using the Reset Link Generated

Put the reset link generated in the above route in the postman and add the body fields of `new_password` and `confirm_password`. The code for implementing the above functionality looks like below, before refactoring.

execs.go
```go

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request){
	token := r.PathValue("resetcode")

	type request struct {
		NewPassword string `json:"new_password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid values in request", http.StatusBadRequest)
		return
	}

	// TODO: Data validation for blank values
	if req.NewPassword == "" && req.ConfirmPassword == "" {
		http.Error(w, "Password cannot have empty values", http.StatusBadRequest)
		return
	}

	if req.NewPassword != req.ConfirmPassword {
		http.Error(w, "Passwords should match", http.StatusBadRequest)
		return
	}

	bytes, err := hex.DecodeString(token)
	if err != nil {
		utils.ErrorHandler(err, "Internal Error")
		return
	}

	hashedToken := sha256.Sum256(bytes)
	hashedTokenString := hex.EncodeToString(hashedToken[:])


	db, err := sqlconnect.ConnectDb()
	if err != nil {
		utils.ErrorHandler(err, "Internal Error")
		return
	}
	defer db.Close()

	var user models.Exec

	query := "SELECT id, email FROM execs WHERE password_reset_token=? AND password_token_expires>?"
	err = db.QueryRow(query, hashedTokenString, time.Now().Format(time.RFC3339)).Scan(&user.ID, &user.Email)
	if err != nil {
		utils.ErrorHandler(err, "Invalid or expired resetcode")
		return
	}

	// Hash the new password
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err!= nil {
		utils.ErrorHandler(err, "internal error")
		return
	}

	updateQuery := "UPDATE execs SET password=?, password_reset_token=NULL, password_token_expires=NULL, password_changed_at=? WHERE id=?"
	_, err = db.Exec(updateQuery, hashedPassword, time.Now().Format(time.RFC3339), user.ID)
	if err != nil {
		utils.ErrorHandler(err, "Internal Error")
		return
	}

	fmt.Fprintln(w, "Password reset successfully")
}
```


## CSRF (Cross Site Request Forgery)

- Cross Site Request Forgery
- Stateless Nature
- Token-based Authentication

- Best Practices for CSRF Protection in APIs
	- Use Same-Site Cookies
	- Double Submit Cookies
	- Custom Headers
	- CSRF Tokens

- Common Pitfalls in CSRF Protection
	- Ignoring Stateless APIs
	- Weak Token Generation
	- Exposing Tokens

Cross-Site Request Forgery is a type of attack where a malicious actor tricks a user into performing actions on a web-application, where they are authenticated without their knowledge. This can lead to unauthorized actions such as data theft, account manipulation, and other harmful operations. CSRF attacks exploits the trust of web-application has in a user's web browser. Without proper protection, any authenticated action like changing a password or making a transaction, can be performed without the user's consent. This compromises the integrity and security of the application and the user's data.

While traditional web applications render HTML and manage user sessions, API often operate statelessly primarily using tokens for authentication. This makes CSRF slightly different in APIs. APIs do not maintain session states, reducing the direct risk of CSRF compared to stateful applications. APIs use tokens like JWT for authentication, which helps mitigate CSRF since token need to be included in each request explicitly. 

CSRF protection is primarity needed for applications where the server and the client usually a web-browser, have a trust relationship and where the client needs to perform state changing operations like form submissions, which are authenticated by cookies or other mechanisms that the browser automatically includes with requests. If you are building a purely API based backed that does not directly interact with a web-browser, CSRF protection is generally not as necessary and there are some scenarios where CSRD is not typically needed.


## Add Pagination to the Students Route

students.go

```go

func GetStudentsHandler(w http.ResponseWriter, r *http.Request) {

	var students []models.Student

	// Implementing the pagination
	// url?limit=x&page=y
	// database will-leave/ will-not show calculated entries from the begining. (page-1) * limit ((1-1)*50 = 0*50 = 0)
	// page y => (y-1) * x, next x entries
	page, limit := getPaginationParams(r)

	students, totalStudents, err := sqlconnect.GetStudentsDbHandler(students, r, limit, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := struct {
		Status   string           `json:"status"`
		Count    int              `json:"count"`
		Page     int              `json:"page"`
		PageSize int              `json:"page_size"`
		Data     []models.Student `json:"data"`
	}{
		Status:   "success",
		Count:    totalStudents,
		Page:     page,
		PageSize: limit,
		Data:     students,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func getPaginationParams(r *http.Request) (int, int) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}
	return page, limit
}

```

students_crud.go

```go

func GetStudentsDbHandler(students []models.Student, r *http.Request, limit, page int) ([]models.Student, int, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, 0, utils.ErrorHandler(err, "error retrieving data")
	}
	defer db.Close()

	query := "SELECT id, first_name, last_name, email, class FROM students WHERE 1=1"
	var args []interface{}

	query, args = utils.AddFilters(r, query, args)

	// Add Pagination
	offset := (page - 1) * limit
	query += " LIMIT ? OFFSET ? "
	args = append(args, limit, offset)


	query = utils.AddSorting(r, query)

	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println("err")
		return nil, 0, utils.ErrorHandler(err, "error retrieving data")
	}
	defer rows.Close()
	for rows.Next() {
		student := models.Student{}
		err = rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Class)
		if err != nil {

			return nil, 0, utils.ErrorHandler(err, "error retrieving data")
		}
		students = append(students, student)
	}

	// Get the total count of students
	var totalStudents int
	err = db.QueryRow("SELECT COUNT(*) FROM students").Scan(&totalStudents)
	if err != nil {
		utils.ErrorHandler(err, "")
		totalStudents = 0
	}

	return students, totalStudents, nil
}

```

## Data Sanitization - XSS Middleware

Sanitization is the process of cleaning and filtering user input to prevent the introduction of malicious data into a system.  This practice is essential in safeguarding applications from various security threats such as SQL injection, cross-site scripting and other forms of injection attacks. 

Data sanization plays a significant role in securing our API. 
- It protects against injection attacks by removing or escaping harmful characters. 
- It ensures that data confirms to expected formats and content maintaining system integrity.
- It prevents malicious data from degrading system performance.

Data Sanitization is crucial on the server side to ensure that all data Entering the system is clean and safe. And while it's important to sanitize data on the client side for user-feedback and immediate security, it should not be solely relied upon as client side sanitization can be bypassed.

*Importance*
- Security
- Integrity
- Performance

*Areas Of Application*
- API / Server-Side
- Frontend Development

*Data Sanitization in APIs / Server-Side Development*
- Input Sanitization
- Output Sanitization
- Database Interaction

*How Data Sanitization is Implemented*
- Escaping : `>` to `&gt`, `<` to `&lt`
- Validation :  checking if an email address has a valid format before sending it to the database.
- Encoding : transform data into a safe format. Encoding data to be safely included in htmls or urls.
- Whitelist Filtering : allowing only known safe data to passthrough. eg: restricting input to only alphabetic characters for a name field.

*Best Practices*
- Sanitize all user inputs
- Use established libraries
- Sanitize at Multiple Layers
- Contextual Escaping
- Regularly Update

*Common Pitfalls*
- Relying Solely on client-side sanitization
- Incomplete Sanitization
- Improper Context Handling
- Neglecting Output Sanitization
- Over-Sanitization

*Examples of Data Sanitization*
- Preventing SQL Injection
- Preventing XSS
- Preventing URL Injection


io.ReadCloser vs io.Reader : An instance of io.ReadCloser needs to be read and it needs to be closed as well once it is read. So we have read method associated with io.ReadCloser as well as Close method assosciated with io.ReadCloser. 

```go
package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"school_management_api/pkg/utils"
	"strings"
	"github.com/microcosm-cc/bluemonday"
)

func XSSMiddleware(next http.Handler) http.Handler {
	fmt.Println("++++++++++++ Initializing XSSMiddleware +++++++++++")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("+++++++++++++++ XSS Middleware Ran ")

		// Sanitize the URL Path
		sanitizePath, err := clean(r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("Original Path:", r.URL.Path)
		fmt.Println("Sanitized Path:", sanitizePath)

		// Sanitize the query Params
		params := r.URL.Query()
		sanitizedQuery := make(map[string][]string)
		for key, values := range params {
			sanitizedKey, err := clean(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			var sanitizedValues []string
			for _, value := range values {
				cleanValue, err := clean(value)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				sanitizedValues = append(sanitizedValues, cleanValue.(string))
			}
			sanitizedQuery[sanitizedKey.(string)] = sanitizedValues
			fmt.Printf("Original Query %s: %s\n", key, strings.Join(values, ", "))
			fmt.Printf("Sanitized Query %s: %s\n", sanitizedKey, strings.Join(sanitizedValues, ", "))
		}

		r.URL.Path = sanitizePath.(string)
		r.URL.RawQuery = url.Values(sanitizedQuery).Encode()
		fmt.Println("Updated URL:", r.URL.String())

		// Sanitize request body
		if r.Header.Get("Content-Type") == "appplication/json" {
			if r.Body != nil {
				bodyBytes, err := io.ReadAll(r.Body)
				if err != nil {
					http.Error(w, utils.ErrorHandler(err, "Error reading request body").Error(), http.StatusBadRequest)
					return
				}
				bodyString := strings.TrimSpace(string(bodyBytes))
				fmt.Println("Original Body:", bodyString)

				// Reset the request Body
				r.Body = io.NopCloser(bytes.NewReader([]byte(bodyString)))

				if len(bodyString) > 0 {
					var inputData interface{}	
					err := json.NewDecoder(bytes.NewReader([]byte(bodyString))).Decode(&inputData)
					if err != nil {
						http.Error(w, utils.ErrorHandler(err, "Invalid JSON body").Error(), http.StatusBadRequest)
						return
					}
					fmt.Println("Original JSON data:", inputData)

					// Sanitize the JSON body
					sanitizedData, err := clean(inputData)
					if err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					fmt.Println("Sanitized JSON data:", sanitizedData)

					// Marshall the sanitized data back to the body
					sanitizedBody, err := json.Marshal(sanitizedData)
					if err != nil {
						http.Error(w, utils.ErrorHandler(err, "Error sanitizing body").Error(), http.StatusBadRequest)
						return
					}

					r.Body = io.NopCloser(bytes.NewReader(sanitizedBody))
					fmt.Println("Sanitized body:", string(sanitizedBody))

				} else {
					log.Println("Request body is empty")
				}

			} else {
				log.Println("No body in the request")
			}
		} else if r.Header.Get("Content-Type") != "" {
			log.Printf("Received request with unsupported Content-Type: %s. Expected application/json.\n", r.Header.Get("Content-Type"))
			http.Error(w, "Unsupported Content-Type. please use application/json.", http.StatusUnsupportedMediaType)
			return
		}

		next.ServeHTTP(w, r)
		fmt.Println("Sending response from XSSMiddleware Ran")
	})
}

// Clean sanitizes input data to prevent XSS attacks
func clean(data interface{}) (interface{}, error) {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			v[key] = sanitizeValue(value)
		}
		return v, nil
	case []interface{}:
		for i, value := range v {
			v[i] = sanitizeValue(value)
		}
		return v, nil
	case string:
		return sanitizeString(v), nil
	default:
		// Error
		return nil, utils.ErrorHandler(fmt.Errorf("unsupported type: %T", data), fmt.Sprintf("unsupported type: %T", data))
	}
}

func sanitizeValue(data interface{}) interface{} {
	switch v := data.(type) {
	case string:
		return sanitizeString(v)
	case map[string]interface{}:
		for k, value := range v {
			v[k] = sanitizeValue(value)
		}
		return v
	case []interface{}:
		for i, value := range v {
			v[i] = sanitizeValue(value)
		}
		return v
	default:
		return v
	}
}

func sanitizeString(value string) string {
	return bluemonday.UGCPolicy().Sanitize(value)
}
```

## Authorization

```go
package utils

import "errors"

type ContextKey string

func AuthorizeUser(userRole string, allowedRoles ...string) (bool, error){
	for _, allowedRole := range allowedRoles {
		if userRole == allowedRole {
			return true, nil
		}
	}
	return false, errors.New(("user not authorized"))
}
```
