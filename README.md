# JobBuddy API

### Overview

The JobBuddy API is a comprehensive backend service designed to assist software engineers in Bangkok with tracking job applications, managing job search progress, and accessing resources for job hunting. Built using Golang, Gin, and PostgreSQL, this API provides secure and efficient endpoints for job application management, company research, and job interview preparation.


### Features

1. User Authentication: Secure user registration and login using JWT.
2. Job Application Tracking: Add, edit, delete, and retrieve job applications with details such as company, position, application status, and dates.
3. Company Research: Store and retrieve information about companies including notes, ratings, and interview experiences.
4. Job Interview Preparation: Manage interview schedules, store interview questions, and track interview feedback.
5. Networking and Events: Track job fairs, networking events, and meetups related to software engineering in Bangkok.
6. Resource Management: Access resources such as resume templates, cover letter examples, and coding challenge websites.
7. Reminders and Notifications: Set reminders for application deadlines, interview dates, and follow-ups, with email/SMS notifications.
8. Secure Data Storage: Secure storage of user data with encryption and compliance with data protection regulations.

### Technology Stack

- Backend: Golang
- Framework: Gin
- ORM: GORM
- Database: PostgreSQL
- Authentication: JWT
- Notifications: Twilio, SendGrid (for SMS and email notifications)
- Deployment: AWS

### Installation and Setup

#### Prerequisites
- Go (version 1.16 or higher)
- PostgreSQL
- Git

### Steps
1. Clone the Repository

```bash
git clone https://github.com/yourusername/jobbuddy-api.git

cd jobbuddy-api
```

2. Set Up Environment Variables
Create a .env file in the project root with the following variables:

```bash

DB_HOST=localhost
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=jobbuddy
DB_PORT=5432
JWT_SECRET=your_jwt_secret
TWILIO_ACCOUNT_SID=your_twilio_account_sid
TWILIO_AUTH_TOKEN=your_twilio_auth_token
SENDGRID_API_KEY=your_sendgrid_api_key

```

3. Install Dependencies

```bash
go mod download

```

4. Run Database Migrations
Use a tool like golang-migrate to run the database migrations.

```bash
migrate -path ./migrations -database "postgres://your_db_user:your_db_password@localhost:5432/jobbuddy?sslmode=disable" up

```

5. Start the Server

```bash
go run main.go

```


### API Endpoints

#### Authentication
* POST /api/auth/register: Register a new user
* POST /api/auth/login: User login

#### Job Applications
* POST /api/applications: Add a new job application
* GET /api/applications: Retrieve all job applications
* GET /api/applications/
: Retrieve a specific job application
* PUT /api/applications/
: Update a job application
* DELETE /api/applications/
: Delete a job application

#### Companies
* POST /api/companies: Add a new company
* GET /api/companies: Retrieve all companies
* GET /api/companies/
: Retrieve a specific company
* PUT /api/companies/
: Update a company
* DELETE /api/companies/
: Delete a company

#### Interviews
* POST /api/interviews: Schedule a new interview
* GET /api/interviews: Retrieve all interviews
* GET /api/interviews/
: Retrieve a specific interview
* PUT /api/interviews/
: Update an interview
* DELETE /api/interviews/
: Delete an interview

#### Networking Events
* POST /api/events: Add a new event
* GET /api/events: Retrieve all events
* GET /api/events/
: Retrieve a specific event
* PUT /api/events/
: Update an event
* DELETE /api/events/
: Delete an event

#### Resources
* GET /api/resources: Retrieve all resources
* POST /api/resources: Add a new resource
* GET /api/resources/
: Retrieve a specific resource
* PUT /api/resources/
: Update a resource
* DELETE /api/resources/
: Delete a resource
* Reminders and Notifications
* POST /api/reminders: Set a new reminder
* GET /api/reminders: Retrieve all reminders
* DELETE /api/reminders/
: Delete a reminder

### Testing
Run the tests using:

```bash
go test ./...
```

### Deployment
Deploy the application using Heroku, DigitalOcean, or any cloud provider of your choice. Ensure to set up environment variables and database configurations as per the provider's requirements.

### Contributing
Contributions are welcome! Please submit a pull request or open an issue to discuss changes.

### License
This project is licensed under the MIT License.

### Contact
For any inquiries, please contact me at kaungzawhtet.mm@gmail.com.




