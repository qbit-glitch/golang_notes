# API Planning

In this project we are going to assume that we have been contracted to create a backend server/API for a school. The school is our client and we are going to plan the API as per our client requirements. 

So the first stage is understanding the project requirements.

### Project Goal: 
Create an API for a school management system that administrative staff can use to manage students, teachers, and other staff members.

#### Key Requirements:
- Addition of student/teaches/staff/exec entry
- Modification of student/teacher/staff/exec entry
- Delete student/teacher/staff/exec entry
- Get list of all students/teachers/staff/execs
- Authentication: login, logout
- Bulk Modifications: students/teachers/staff/execs
- Class Management:
    - Total count of a class with class teacher
    - List of all students in a class with class teacher

#### Security and Rate Limiting:
- Rate Limit the application
- Password reset mechanisms (forgot password, update password)
- Deactivate user

## Fields:

| Student       | Teacher       | Executives        | 
|---------------|---------------|-------------------|
| First Name    | First Name    | First Name        |
| Last Name     | Last Name     | Last Name         |
| Class         | Subject       | Role              |
| Email         | Class         | Email             |
|               | Email         | Username          |
|               |               | Password          |

## Endpoints

**Executives**
- GET `/execs` : Get list of executives
- POST `/execs` : Add a new executive
- PATCH `/execs` : Modify multiple executives
- GET `/execs/{id}`: Get a specific executive
- PATCH `/execs/{id}` : Modify a specific executive
- DELETE `/execs/{id}` : Delete a specific executive
- POST `/execs/login` : Login
- POST `/execs/logout` : Logout
- POST `/execs/forgotpassword` : Forgot Password
- POST `/execs/resetpassword/reset/{resetcode}` : Reset Password

**Students**
- GET `/students` : Get list of students
- POST `/students` : Add a new students
- PATCH `/students` : Modify multiple students
- DELETE `/students` : Delete multiple students
- GET `/students/{id}`: Get a specific student
- PATCH `/students/{id}` : Modify a specific student
- PUT `/students/{id}` : Update a specific student
- DELETE `/students/{id}` : Delete a specific student


**Teachers**
- GET `/teachers` : Get list of teachers
- POST `/teachers` : Add a new teachers
- PATCH `/teachers` : Modify multiple teachers
- DELETE `/teachers` : Delete multiple teachers
- GET `/teachers/{id}`: Get a specific teacher
- PATCH `/teachers/{id}` : Modify a specific teacher
- PUT `/teachers/{id}` : Update a specific teacher
- DELETE `/teachers/{id}` : Delete a specific teacher
- GET `/teachers/{id}/students`: Get students of a specific teacher
- GET `/teachers/{id}/studentcount`: Get student count for a specific teacher


## Best Practices and Common Pitfalls

- **Best Practices**
    - Modularity
    - Documentation
    - Error Handling
    - Security
    - Testing

- **Common Pitfalls**
    - Overcomplicating the API
    - Ignoring Security
    - Poor Documentation
    - Inadequate Testing

By breaking down project requirements into tasks and subsequently into endpoints, you create a clear roadmap for development. Following best practices and avoiding common pitfalls will ensure your API is robust, secure and easy to use.