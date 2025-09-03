# Mock HTTP Server for Third-Party Integration Testing

This project is a Go-based mock HTTP server designed to help developers test and validate integrations 
with third-party services. It simulates a real external service by providing data through HTTP endpoints, 
allowing you to verify integration logic and data handling in a controlled environment.

## Features
- __Internal data storage__: Maintains a dataset that is periodically updated to mimic real-world data changes.
- __Update control__: Start or stop automatic data updates as needed for your tests.
- __REST routes__: Retrieve the current state of the data and check the update status.


## Quick Start

### 1. Run the server
Start the application with the default configuration:
```
go run ./cmd
```
Set custom port and environment using flags:
```
go run ./cmd -port=4000 -env=development
```

### 2. Health Check
Verify the service is running:
```
GET host/v1/healthcheck
```

### 3. Data Model
The service uses an __in-memory store__ with a `Company` model as the default schema:
```
type Company struct {
	ID      int       `json:"id"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Active  bool      `json:"active"`
	Company string    `json:"companies"`
	Status  string    `json:"status"`
	Phone   string    `json:"phone"`
	Email   string    `json:"email"`
	Staff   int       `json:"staff"`
}
```
- Immutable attributes: `id`, `created`, `companies`, `status`
- Mutable attributes: `active`, `phone`, `email`, `staff`
- `updated` tracks the last update time

You can configure the __number of entities__ and the __update frequency__.

### 4. Get Update Info
Check how many entities exist and the update frequency:
```
GET host/v1/companies/updates/info
```
Response:
```json
{
  "company_info": {
    "total": 100,
    "updating": true,
    "period": 10
  }
}
```

### 5. Get Company by ID
Retrieve companies details:
```
GET host/v1/companies/{id}
```
Response:
```json
{
  "companies": {
    "id": 58,
    "created": "1984-08-29T06:56:24.611616918Z",
    "updated": "2025-09-01T12:13:04.046639Z",
    "active": true,
    "company": "BillGuard",
    "status": "private",
    "phone": "7229708716",
    "email": "hansmann@effertz.org",
    "staff": 15
  }
}
```
### 6. Stop Data Updates
Pause random updates of entities:
```
PATCH host/v1/companies/updates/stop
```
Useful for testing __data synchronization stability__.


### 7. Start Data Updates
Resume updates and set the mutation period (in seconds):
```
PATCH host/v1/companies/updates/start
```
Request body:
```json
{
  "period": 20
}
```

### 8. Synchronization Endpoint
Retrieve updated data with filters for time, status, and pagination:
```
GET host/v1/companies/updates?from={from}&to={to}&status={status}&page={page}&size={size}
```
Example:
```
GET localhost:4000/v1/companies/updates?from=1995-10-18T05:07:00Z&to=2025-03-02T19:57:22Z&status=public&page=1
```

### 9. Web Interface
View logs and update events in your browser:
```
host/static/
```

### 10 Extensibility
Currently, the service operates with a single `Company` schema, but you can __extend the codebase__ 
to implement custom data structures that reflect the schema of the third-party services you want to integrate with.
