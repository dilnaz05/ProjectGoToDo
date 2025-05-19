### TODO List Web Application (Golang + Gin + PostgreSQL)

This project is a simple TODO List web application built using Golang(Gin) and PostgreSQL.

Users can register, log in, and manage their personal todo tasks.  
The system uses JWT-based authentication and supports role-based access control(admin and user roles).

---
### Authentication & Authorization

Authentication is implemented using JWT tokens.  
Each user is assigned a role: either `"user"` or `"admin"`.

### Role: `user`
Regular users can:
- Register and log in
- View their own profile (`/api/me`)
- Manage their own todos:
  - Create (`POST /api/todos`)
  - Read (`GET /api/todos`)
  - Update (`PUT /api/todos/:id`)
  - Delete (`DELETE /api/todos/:id`)

They cannot:
- Access other users' data
- Access any `/admin/...` routes

---

### Role: `admin`
Admins have all `user` privileges plus:
- View all users (`GET /admin/users`)
- View all todos (`GET /admin/todos`)
- Delete any user (`DELETE /admin/user/:id`)
- Access admin dashboard (`GET /admin/dashboard`)

---

### Backend:
- **Gin** – HTTP web framework for routing and middleware
- **Gorm** – ORM for PostgreSQL database interactions
- **PostgreSQL** – Relational database for storing user and todo data
- **JWT** – Secure token-based authentication
- **Role-based Middleware** – for authorization by user roles

