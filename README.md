# rest-api-go

## 🚀 Project Structure

```
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

## 🛠️ Commands

### Generate CRUD

Generate a new CRUD module:

```bash
henotic generate Product name:string price:number description:text
```

### Run the application

```bash
go run cmd/main/main.go
```

### Seed the database

```bash
go run cmd/seed/main.go
```

## 📚 API Endpoints

After generating a module, the following endpoints will be available:

- **GET** `/api/[resources]`: Get all resources
- **GET** `/api/[resources]/:id`: Get a resource by ID
- **POST** `/api/[resources]`: Create a new resource
- **PUT** `/api/[resources]/:id`: Update a resource
- **DELETE** `/api/[resources]/:id`: Delete a resource

## 🔧 Configuration

Edit the `.env` file to configure your database connection and other settings.
