# Clinic App

A simple Golang-based web application for managing patient data in a clinic. This app consists of a receptionist portal and a doctor portal with proper authentication and authorization. The application allows receptionists to register patients and perform CRUD operations, while doctors can view and update patient-related data.

## Features

- **Login & Registration**: Single API for both receptionist and doctor login.
- **Patient Management**: Receptionists can register, update, delete, and view patients.
- **Doctor Access**: Doctors can view and update patient details.
- **Authentication & Authorization**: JWT-based authentication with role-based access control (RBAC).
- **Database**: PostgreSQL used for storing patient data.
- **Middleware**: Custom middleware for role validation and database injection.
- **Unit Tests**: Unit tests for API endpoints (optional).
- **API Documentation**: API documentation using Postman or Swagger.

## Technology Stack

- **Backend**: Golang (Fiber Framework)
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **ORM**: GORM (Optional)
- **Middleware**: Custom JWT & Role-based authorization middleware
- **Testing**: Optional (for unit tests)
- **Documentation**: Postman or Swagger


## Setup and Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/<your-username>/clinic-app.git
    cd clinic-app
    ```

2. Install the necessary dependencies:

    ```bash
    go mod tidy
    ```

3. Configure PostgreSQL database:
   - Create a PostgreSQL database and update the `config` folder with the database credentials.
   - Alternatively, you can use environment variables to store database credentials.

4. Run the application:

    ```bash
    go run main.go
    ```

   This will start the app on `http://localhost:3000`.

## API Endpoints

### Public Routes

- **POST** `/login` - Login API for receptionists and doctors (returns a JWT token).
- **POST** `/register` - Register API (only for receptionists, for adding a new user).

### Protected Routes (Requires JWT)

- **POST** `/api/patients` - Register a new patient (Role: `receptionist`).
- **GET** `/api/patients` - Get a list of all patients (Role: `receptionist`).
- **GET** `/api/patients/:id` - Get a specific patient by ID (Role: `receptionist`).
- **PUT** `/api/patients/:id` - Update patient data (Role: `receptionist` or `doctor`).
- **DELETE** `/api/patients/:id` - Delete a patient (Role: `receptionist`).

### Authentication & Authorization

- **JWT Token**: When a user logs in, a JWT token is returned. This token must be sent in the `Authorization` header as `Bearer <token>` for accessing the protected routes.
- **Role-based Access**: 
  - `receptionist` can register, update, view, and delete patients.
  - `doctor` can view and update patient details.

