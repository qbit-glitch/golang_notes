# Part 5

## Download MailHog

Link: https://github.com/mailhog/MailHog

MailHog is a lightweight package that we can add to our API to include mail sending functionality in our server. MailHog is a simple and effective email testing tool that allows developers to simulate sending and receiving emails in a safe environment without actually sending them to real email addresses. It's particularly useful for testing email functionalities in applications during development. 

Some of the key features of MailHog include an SMTP server, so MailHog runs an SMTP server that captures emails sent from your application. You can configure your application to send emails to MailHog instead of real email servers or real email addresses. This way you can review the content of the emails without sending them out. So those emails will be received to a fake SMTP server that is running on your computer and you can send as many emails as possible because no other server is receiving those emails and you don't have to login to any email account to check those emails.

Mailhog provides a web interface where we can view the emails that have been captured. You can see details like the sender, recipient, subject and body of each email, making it absolutely easy to verify that your application is sending the correct information.

Setting up mailHog is straight forward. it can be run as a standalone binary or as a Docker container. This ease of use allows developers to quickly integrate MailHog into their development workflow. With MailHog you can simulate various email scenarios without any side effects. MailHog can also be integrated into automated tests. You can check that emails are send correctly and contain the right information, making it a valuable tool for maintaining code quality. 

In our API we use MailHog to send password reset emails. When a user submits their email address to the forgot password route, the application generates a password reset email containing a secure link, and this email is captured by MailHog, allowing us to verify its content through the web interface without actually sending it to the user's email address.