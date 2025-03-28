```markdown
# API Documentation for REST API Go Project

## Project Overview
This is a RESTful API built with Go (Golang) that implements a clean architecture pattern with modular design. The project provides CRUD operations for multiple entities including Products, Categories, and Users.

## Project Structure
The project follows a well-organized structure that separates concerns and promotes maintainability:

```plaintext
rest-api-go/
├── cmd/                  # Command-line applications
│   ├── main/             # Main application
│   │   └── main.go       # Entry point
│   └── seed/             # Database seeder
│       └── main.go       # Seeder entry point
├── data/                 # Sample data for seeding
├── internal/             # Private application code
│   ├── module/           # Business modules
│   │   └── [module]/     # Specific module
│   │       ├── entity/   # Domain models
│   │       ├── handler/  # HTTP handlers
│   │       └── service/  # Business logic
│   │       
│   └── seed/             # Seed implementations
├── pkg/                  # Public libraries
│   ├── config/           # Configuration
│   ├── database/         # Database connection
│   ├── middleware/       # HTTP middleware
│   └── utils/            # Utility functions
├── .env                  # Environment variables
├── .gitignore            # Git ignore file
├── go.mod                # Go modules
└── go.sum                # Go dependencies checksum
```

### Key Components
- cmd/ : Contains the application entry points
  - main/ : The main API server
  - seed/ : Database seeding utility
- internal/ : Private application code
  - module/ : Business modules organized by domain
    - Each module contains:
      - entity/ : Data models and validation
      - handler/ : HTTP request handlers
      - service/ : Business logic
  - seed/ : Database seeding implementations
- pkg/ : Shared libraries
  - config/ : Application configuration
  - database/ : Database connection management
  - middleware/ : HTTP middleware (CORS, logging)
  - utils/ : Utility functions (response formatting)
## API Endpoints
The API follows RESTful conventions and provides the following endpoints for each resource:

### Categories Method Endpoint Description GET

/api/categories

Get all categories GET

/api/categories/:id

Get a category by ID POST

/api/categories

Create a new category PUT

/api/categories/:id

Update a category DELETE

/api/categories/:id

Delete a category
### Products Method Endpoint Description GET

/api/products

Get all products GET

/api/products/:id

Get a product by ID GET

/api/products/category/:categoryId

Get products by category ID POST

/api/products

Create a new product PUT

/api/products/:id

Update a product DELETE

/api/products/:id

Delete a product
### Users Method Endpoint Description GET

/api/users

Get all users GET

/api/users/:id

Get a user by ID POST

/api/users

Create a new user PUT

/api/users/:id

Update a user DELETE

/api/users/:id

Delete a user
## Detailed API Documentation
### Categories API 1. Get All Categories
Endpoint: GET /api/categories

Description: Retrieves all categories with their associated products.

Request: No request body required

Response:

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Electronics",
      "products": [
        {
          "id": 1,
          "title": "Smartphone",
          "price": 599.99,
          "description": "Latest model smartphone",
          "category_id": 1,
          "created_at": "2023-07-15T10:30:00Z",
          "updated_at": "2023-07-15T10:30:00Z"
        }
      ],
      "created_at": "2023-07-15T10:00:00Z",
      "updated_at": "2023-07-15T10:00:00Z"
    },
    {
      "id": 2,
      "name": "Clothing",
      "products": [],
      "created_at": "2023-07-15T10:05:00Z",
      "updated_at": "2023-07-15T10:05:00Z"
    }
  ]
}
```
 2. Get Category by ID
Endpoint: GET /api/categories/:id

Description: Retrieves a specific category by its ID with associated products.

Parameters:

- :id - The ID of the category to retrieve
Response:

```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Electronics",
    "products": [
      {
        "id": 1,
        "title": "Smartphone",
        "price": 599.99,
        "description": "Latest model smartphone",
        "category_id": 1,
        "created_at": "2023-07-15T10:30:00Z",
        "updated_at": "2023-07-15T10:30:00Z"
      }
    ],
    "created_at": "2023-07-15T10:00:00Z",
    "updated_at": "2023-07-15T10:00:00Z"
  }
}
```

Error Response (Not Found):

```json
{
  "success": false,
  "error": "Category not found"
}
 ```
 3. Create Category
Endpoint: POST /api/categories

Description: Creates a new category.

Request Body:

```json
{
  "name": "Electronics"
}
 ```

Response:

```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Electronics",
    "products": [],
    "created_at": "2023-07-15T10:00:00Z",
    "updated_at": "2023-07-15T10:00:00Z"
  }
}
```

Error Response (Validation Error):

```json
{
  "success": false,
  "error": "Key: 'Category.Name' Error:Field validation for 'Name' failed on the 'max' tag"
}
```
 4. Update Category
Endpoint: PUT /api/categories/:id

Description: Updates an existing category.

Parameters:

- :id - The ID of the category to update
Request Body:

```json
{
  "name": "Updated Electronics"
}
 ```

Response:

```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Updated Electronics",
    "products": [],
    "created_at": "2023-07-15T10:00:00Z",
    "updated_at": "2023-07-15T11:00:00Z"
  }
}
```

Error Response (Not Found):

```json
{
  "success": false,
  "error": "Category not found"
}
 ```
 5. Delete Category
Endpoint: DELETE /api/categories/:id

Description: Deletes a category by its ID.

Parameters:

- :id - The ID of the category to delete
Response:

```json
{
  "success": true,
  "data": null
}
 ```

Error Response (Not Found):

```json
{
  "success": false,
  "error": "Category not found"
}
 ```

### Products API 1. Get All Products
Endpoint: GET /api/products

Description: Retrieves all products.

Response:

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "title": "Smartphone",
      "price": 599.99,
      "description": "Latest model smartphone",
      "category_id": 1,
      "created_at": "2023-07-15T10:30:00Z",
      "updated_at": "2023-07-15T10:30:00Z"
    },
    {
      "id": 2,
      "title": "Laptop",
      "price": 1299.99,
      "description": "High-performance laptop",
      "category_id": 1,
      "created_at": "2023-07-15T10:35:00Z",
      "updated_at": "2023-07-15T10:35:00Z"
    }
  ]
}
```
 2. Get Product by ID
Endpoint: GET /api/products/:id

Description: Retrieves a specific product by its ID.

Parameters:

- :id - The ID of the product to retrieve
Response:

```json
{
  "success": true,
  "data": {
    "id": 1,
    "title": "Smartphone",
    "price": 599.99,
    "description": "Latest model smartphone",
    "category_id": 1,
    "created_at": "2023-07-15T10:30:00Z",
    "updated_at": "2023-07-15T10:30:00Z"
  }
}
```

Error Response (Not Found):

```json
{
  "success": false,
  "error": "Product not found"
}
 ```
 3. Get Products by Category ID
Endpoint: GET /api/products/category/:categoryId

Description: Retrieves all products belonging to a specific category.

Parameters:

- :categoryId - The ID of the category to filter products by
Response:

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "title": "Smartphone",
      "price": 599.99,
      "description": "Latest model smartphone",
      "category_id": 1,
      "created_at": "2023-07-15T10:30:00Z",
      "updated_at": "2023-07-15T10:30:00Z"
    },
    {
      "id": 2,
      "title": "Laptop",
      "price": 1299.99,
      "description": "High-performance laptop",
      "category_id": 1,
      "created_at": "2023-07-15T10:35:00Z",
      "updated_at": "2023-07-15T10:35:00Z"
    }
  ]
}
```
 4. Create Product
Endpoint: POST /api/products

Description: Creates a new product.

Request Body:

```json
{
  "title": "Smartphone",
  "price": 599.99,
  "description": "Latest model smartphone",
  "category_id": 1
}
```

Response:

```json
{
  "success": true,
  "data": {
    "id": 1,
    "title": "Smartphone",
    "price": 599.99,
    "description": "Latest model smartphone",
    "category_id": 1,
    "created_at": "2023-07-15T10:30:00Z",
    "updated_at": "2023-07-15T10:30:00Z"
  }
}
```

Error Response (Validation Error):

```json
{
  "success": false,
  "error": "Key: 'Product.Title' Error:Field validation for 'Title' failed on the 'max' tag"
}
```
 5. Update Product
Endpoint: PUT /api/products/:id

Description: Updates an existing product.

Parameters:

- :id - The ID of the product to update
Request Body:

```json
{
  "title": "Updated Smartphone",
  "price": 649.99,
  "description": "Latest model smartphone with improved features",
  "category_id": 1
}
```

Response:

```json
{
  "success": true,
  "data": {
    "id": 1,
    "title": "Updated Smartphone",
    "price": 649.99,
    "description": "Latest model smartphone with improved features",
    "category_id": 1,
    "created_at": "2023-07-15T10:30:00Z",
    "updated_at": "2023-07-15T11:30:00Z"
  }
}
```

Error Response (Not Found):

```json
{
  "success": false,
  "error": "Product not found"
}
 ```
 6. Delete Product
Endpoint: DELETE /api/products/:id

Description: Deletes a product by its ID.

Parameters:

- :id - The ID of the product to delete
Response:

```json
{
  "success": true,
  "data": null
}
 ```

Error Response (Not Found):

```json
{
  "success": false,
  "error": "Product not found"
}
 ```

### Users API 1. Get All Users
Endpoint: GET /api/users

Description: Retrieves all users.

Response:

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "username": "johndoe",
      "email": "john@example.com",
      "password": "********",
      "created_at": "2023-07-15T09:00:00Z",
      "updated_at": "2023-07-15T09:00:00Z"
    },
    {
      "id": 2,
      "username": "janedoe",
      "email": "jane@example.com",
      "password": "********",
      "created_at": "2023-07-15T09:05:00Z",
      "updated_at": "2023-07-15T09:05:00Z"
    }
  ]
}
```
 2. Get User by ID
Endpoint: GET /api/users/:id

Description: Retrieves a specific user by their ID.

Parameters:

- :id - The ID of the user to retrieve
Response:

```json
{
  "success": true,
  "data": {
    "id": 1,
    "username": "johndoe",
    "email": "john@example.com",
    "password": "********",
    "created_at": "2023-07-15T09:00:00Z",
    "updated_at": "2023-07-15T09:00:00Z"
  }
}
```

Error Response (Not Found):

```json
{
  "success": false,
  "error": "User not found"
}
 ```
 3. Create User
Endpoint: POST /api/users

Description: Creates a new user.

Request Body:

```json
{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "securepassword"
}
 ```

Response:

```json
{
  "success": true,
  "data": {
    "id": 1,
    "username": "johndoe",
    "email": "john@example.com",
    "password": "********",
    "created_at": "2023-07-15T09:00:00Z",
    "updated_at": "2023-07-15T09:00:00Z"
  }
}
```

Error Response (Validation Error):

```json
{
  "success": false,
  "error": "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'max' tag"
}
```
 4. Update User
Endpoint: PUT /api/users/:id

Description: Updates an existing user.

Parameters:

- :id - The ID of the user to update
Request Body:

```json
{
  "username": "johndoe_updated",
  "email": "john_updated@example.com",
  "password": "newsecurepassword"
}
```

Response:

```json
{
  "success": true,
  "data": {
    "id": 1,
    "username": "johndoe_updated",
    "email": "john_updated@example.com",
    "password": "********",
    "created_at": "2023-07-15T09:00:00Z",
    "updated_at": "2023-07-15T10:00:00Z"
  }
}
```

Error Response (Not Found):

```json
{
  "success": false,
  "error": "User not found"
}
 ```
 5. Delete User
Endpoint: DELETE /api/users/:id

Description: Deletes a user by their ID.

Parameters:

- :id - The ID of the user to delete
Response:

```json
{
  "success": true,
  "data": null
}
 ```

Error Response (Not Found):

```json
{
  "success": false,
  "error": "User not found"
}
 ```

## Testing with Postman
### Setting Up Postman
1. Download and Install Postman : If you haven't already, download and install Postman from https://www.postman.com/downloads/ .
2. Create a New Collection :
   
   - Click on "Collections" in the sidebar
   - Click the "+" button to create a new collection
   - Name it "REST API Go"
3. Set Base URL :
   
   - Click on the collection you just created
   - Go to the "Variables" tab
   - Add a variable named "base_url" with an initial value of " http://localhost:8080/api "
### Testing Endpoints Categories
1. Get All Categories :
   
   - Method: GET
   - URL: {{base_url}}/categories
   - Click "Send" to execute the request
2. Get Category by ID :
   
   - Method: GET
   - URL: {{base_url}}/categories/1
   - Click "Send" to execute the request
3. Create Category :
   
   - Method: POST
   - URL: {{base_url}}/categories
   - Headers: Add "Content-Type: application/json"
   - Body: Select "raw" and "JSON", then enter:
     ```json
     {
       "name": "Electronics"
     }
      ```
   - Click "Send" to execute the request
4. Update Category :
   
   - Method: PUT
   - URL: {{base_url}}/categories/1
   - Headers: Add "Content-Type: application/json"
   - Body: Select "raw" and "JSON", then enter:
     ```json
     {
       "name": "Updated Electronics"
     }
     ```
   - Click "Send" to execute the request
5. Delete Category :
   
   - Method: DELETE
   - URL: {{base_url}}/categories/1
   - Click "Send" to execute the request Products
1. Get All Products :
   
   - Method: GET
   - URL: {{base_url}}/products
   - Click "Send" to execute the request
2. Get Product by ID :
   
   - Method: GET
   - URL: {{base_url}}/products/1
   - Click "Send" to execute the request
3. Get Products by Category ID :
   
   - Method: GET
   - URL: {{base_url}}/products/category/1
   - Click "Send" to execute the request
4. Create Product :
   
   - Method: POST
   - URL: {{base_url}}/products
   - Headers: Add "Content-Type: application/json"
   - Body: Select "raw" and "JSON", then enter:
     ```json
     {
       "title": "Smartphone",
       "price": 599.99,
       "description": "Latest model smartphone",
       "category_id": 1
     }
     ```
   - Click "Send" to execute the request
5. Update Product :
   
   - Method: PUT
   - URL: {{base_url}}/products/1
   - Headers: Add "Content-Type: application/json"
   - Body: Select "raw" and "JSON", then enter:
     ```json
     {
       "title": "Updated Smartphone",
       "price": 649.99,
       "description": "Latest model smartphone with improved features",
       "category_id": 1
     }
     ```
   - Click "Send" to execute the request
6. Delete Product :
   
   - Method: DELETE
   - URL: {{base_url}}/products/1
   - Click "Send" to execute the request Users
1. Get All Users :
   
   - Method: GET
   - URL: {{base_url}}/users
   - Click "Send" to execute the request
2. Get User by ID :
   
   - Method: GET
   - URL: {{base_url}}/users/1
   - Click "Send" to execute the request
3. Create User :
   
   - Method: POST
   - URL: {{base_url}}/users
   - Headers: Add "Content-Type: application/json"
   - Body: Select "raw" and "JSON", then enter:
     ```json
     {
       "username": "johndoe",
       "email": "john@example.com",
       "password": "securepassword"
     }
     ```
   - Click "Send" to execute the request
4. Update User :
   
   - Method: PUT
   - URL: {{base_url}}/users/1
   - Headers: Add "Content-Type: application/json"
   - Body: Select "raw" and "JSON", then enter:
     ```json
     {
       "username": "johndoe_updated",
       "email": "john_updated@example.com",
       "password": "newsecurepassword"
     }
     ```
   - Click "Send" to execute the request
5. Delete User :
   
   - Method: DELETE
   - URL: {{base_url}}/users/1
   - Click "Send" to execute the request
### Automating Tests with Postman
You can also create test scripts in Postman to automate testing:

1. Test for Successful Response :
   
   ```javascript
   pm.test("Status code is 200", function () {
       pm.response.to.have.status(200);
   });
   
   pm.test("Response has success:true", function () {
       var jsonData = pm.response.json();
       pm.expect(jsonData.success).to.eql(true);
   });
   ```
2. Test for Created Resource :
   
   ```javascript
   pm.test("Status code is 201", function () {
       pm.response.to.have.status(201);
   });
   
   pm.test("Resource created successfully", function () {
       var jsonData = pm.response.json();
       pm.expect(jsonData.success).to.eql(true);
       pm.expect(jsonData.data).to.have.property("id");
   });
   ```
3. Test for Not Found Error :
   
   ```javascript
   pm.test("Status code is 404", function () {
       pm.response.to.have.status(404);
   });
   
   pm.test("Error message is correct", function () {
       var jsonData = pm.response.json();
       pm.expect(jsonData.success).to.eql(false);
       pm.expect(jsonData.error).to.include("not found");
   });
    ```
   ```
## Data Models
### Category
```go
type Category struct {
    ID          uint                `json:"id"`
    Name        string              `json:"name"`
    Products    []Product           `json:"products,omitempty"`
    CreatedAt   time.Time           `json:"created_at"`
    UpdatedAt   time.Time           `json:"updated_at"`
}
```

### Product
```go
type Product struct {
    ID          uint      `json:"id"`
    Title       string    `json:"title"`
    Price       float64   `json:"price"`
    Description string    `json:"description"`
    CategoryID  uint      `json:"category_id"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

### User
```go
type User struct {
    ID          uint      `json:"id"`
    Username    string    `json:"username"`
    Email       string    `json:"email"`
    Password    string    `json:"password"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

## Getting Started
### Prerequisites
- Go 1.16+
- MariaDB/MySQL
### Configuration
Edit the .env file or modify the configuration in pkg/config/config.go to set your database connection parameters:

```go
// Default configuration
DBHost:     "localhost"
DBPort:     "3306"
DBUser:     "root"
DBPassword: "mariadb"
DBName:     "learning-go-DB"
ServerPort: "8080"
 ```

### Running the Application
1. Start the API server:
```bash
go run cmd/main/main.go
 ```

2. Seed the database with initial data:
```bash
go run cmd/seed/main.go
 ```

## Architecture
The project follows a clean architecture pattern with:

1. Entity Layer : Domain models with validation logic
2. Service Layer : Business logic and data manipulation
3. Handler Layer : HTTP request handling and response formatting
4. Bootstrap : Module initialization and dependency injection
Each module is self-contained with its own entity, service, and handler components, making the codebase modular and maintainable.

## Middleware
The API includes middleware for:

- CORS : Cross-Origin Resource Sharing support
- Logging : Request logging
## Database
The application uses GORM as an ORM with MariaDB/MySQL. Database operations include:

- Auto-migration for schema creation
- Seeding for initial data population
- Relationship management (one-to-many between Category and Product)
## Error Handling
The API implements consistent error handling with appropriate HTTP status codes and formatted error messages.

## Validation
Data validation is performed at the entity level using the validator package, ensuring data integrity before database operations.