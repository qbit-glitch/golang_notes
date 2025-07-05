# Part - 4

## Student Database Creation

```sql
USE school;
CREATE INDEX idx_class ON teachers(class);

CREATE TABLE IF NOT EXISTS students (
	id INT AUTO_INCREMENT PRIMARY KEY,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR(255) NOT NULL,
	email VARCHAR(255) UNIQUE NOT NULL,
	class VARCHAR(255) NOT NULL,
	INDEX(email),
	FOREIGN KEY (class) REFERENCES teachers(class)
) AUTO_INCREMENT=1000
```

## CRUD for Students Route

