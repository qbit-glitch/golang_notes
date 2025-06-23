# Ports

A port in networking is a virtual point where network connections start and end. It helps yout computer distinguish between different types of network traffic. Ports act like separate doors for different types of data, ensuring that information gets to the right place. A computer has 65,535 ports.

**A computer has 65,535 ports**
- Well Known Ports (0-1023) : used by system or well known services like web servers.
- Registered Ports (1024-49151): assigned to specific services and applications.
- Dynamic or Private Ports(49152-65535): used for dynamic, private or ephemeral purposes used by client applications for short term.

**Common Ports:**
- Port 80 : http
- Port 443: https
- Port 25: smtp
- Port 21: ftp
- Ports 3000, 8080, 8000: commonly used for web development and running local servers.




**Analogy:**

The first principle is that port is only used for communicating with the outside world. Ports are reserved entry points to internal processes for outside world. Outside world can use these entry points which are ports to access our internal processes which hosted on our computer when executing on our computer.

The databases are meant to communicate with the outside world to store a lot of data from many remote resources. They are going to receive data from anywhere in the world and they are going to store that data into one location.

Similarly API by design is meant to be connected with the outside world and that communication with the outside world is done by using ports.

Ports are crucial for the functioning of a computer network. They allow multiple services to run simultaneously without interfering with each other.