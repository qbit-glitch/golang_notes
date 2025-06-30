# .env File

Think of environment variables as secret settings or configurations, your API needs to work correctly. They include things like database credentials which are username, password, database name, etc. and yes, database name is something that needs to be confidential apart from just username and password.

Environment variables also include secret keys which are used for encryption and JWT tokens. Port number on which your API will run, needs to be a secret. So instead of hard-coding these values directly in our code, we keep them in the dot env file.


**Environment Variables**
- Database credentials
- API Keys
- Secret Keys
- Port Numbers

**Why use a .env File**
- Security
- Configuration Flexibility
- Maintainability

When we use `os.Getenv()`, what happens is that it will try to access the environment variables of the operating system of our computer. It will not access the .env file in our directory. So to make the .env file in our repo be treated like they are environment variables in the os, there's an external package for that: 
