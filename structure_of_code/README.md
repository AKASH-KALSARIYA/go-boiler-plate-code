# Go Boilerplate Project Structure & Guide

This document explains how this project is structured, how to install its dependencies, and how to run it.

## 📁 Project Structure

```text
/
├── cmd/
│   └── main.go             # The entry point of the application
├── internal/
│   ├── api/
│   │   ├── controller/      # Handles HTTP requests/responses (I/O)
│   │   └── router/          # Defines API routes and path parameters
│   ├── config/              # Configuration loading (.env, etc.)
│   ├── constant/            # App-wide constants
│   ├── database/            # Database connection logic
│   ├── dto/                 # Data Transfer Objects (request/response structures)
│   ├── model/               # Database entities/schemas
│   ├── repository/          # Data access layer (Direct database interaction)
│   └── service/             # Business logic layer (The "Brain")
├── pkg/
│   └── logger/              # Utility packages (e.g., Logging)
├── Makefile                 # Shortcuts for common commands
├── go.mod                   # Go module definition and dependencies
└── .env                     # Environment variables (Port, DB credentials)
```

## 🛠️ Installation & Setup

To get this project running on your machine, follow these steps:

### 1. Install Go
Ensure شما have Go installed (Version 1.22 or higher is required for the path parameters used in this project).
Check your version with:
```bash
go version
```

### 2. Download Dependencies
Run the following command to download all necessary libraries:
```bash
go mod tidy
```

## 🚀 How to Run the Code

You can use the built-in `Makefile` commands for speed, or use standard `go` commands.

### Using Makefile (Easiest)
*   **Run the App:** `make run`
*   **Build the App:** `make build`
*   **Run Tests:** `make test`
*   **Clean Up:** `make clean`

### Using Standard Go Commands
*   **Run:** `go run ./cmd/main.go`
*   **Install Deps:** `go install`
*   **Tidy Modules:** `go mod tidy`

## 💻 Go Commands Cheat Sheet (Essential for every project)

Here is a structured list of commands you will use in your Go development journey.

---

### `go mod tidy`
- **Method Name:** Dependency Synchronization
- **Why it is used:** This is the most used command. It automatically scans your code for imported packages, downloads the missing ones, and removes any that are no longer used. It keeps your `go.mod` and `go.sum` files clean.

### `go mod init <module_name>`
- **Method Name:** Project Initialization
- **Why it is used:** Used once at the start of every project. It creates the `go.mod` file which defines the "Root" of your project and its name.

### `go run ./cmd/main.go`
- **Method Name:** Instant Execution
- **Why it is used:** Perfect for development. It compiles your code into a temporary memory location and runs it immediately so you can see your changes without creating an `.exe` file.

### `go build -o app.exe ./cmd/main.go`
- **Method Name:** Binary Compilation
- **Why it is used:** Used when you are finished coding. It creates a permanent executable file. The `-o` flag allows you to name the output file.

### `go get <package_url>`
- **Method Name:** Package Retrieval
- **Why it is used:** Downloads a specific external library from GitHub or other sources and adds it to your project's dependencies.

### `go fmt ./...`
- **Method Name:** Code Formatting
- **Why it is used:** Automatically fixes the spacing, tabs, and indentation in all your files to match the official Go standard. Use this before sharing code!

### `go test ./...`
- **Method Name:** Automated Testing
- **Why it is used:** Searches for all files ending in `_test.go` and runs the tests inside them to ensure your logic is correct.

### `go vet ./...`
- **Method Name:** Static Analysis
- **Why it is used:** Scans your code for "suspicious" things that aren't necessarily errors but could cause bugs, like unreachable code or formatting mistakes in `Printf`.

### `go env`
- **Method Name:** Environment Inspection
- **Why it is used:** Displays all the internal Go settings on your computer, like where Go is installed (`GOROOT`) and where it downloads libraries (`GOPATH`).

### `go version`
- **Method Name:** Version Check
- **Why it is used:** Simply tells you which version of Go is running on your system. Important for ensuring compatibility with features like path parameters (v1.22+).

### `go install`
- **Method Name:** Binary Installation
- **Why it is used:** Compiles and moves the executable into your `GOBIN` folder so you can run your program as a command from anywhere in your terminal.

### `go mod download`
- **Method Name:** Bulk Download
- **Why it is used:** Downloads all dependencies listed in your `go.mod` file at once. Usually used in CI/CD pipelines or when setting up a project for the first time.

### `go list -m all`
- **Method Name:** Dependency Listing
- **Why it is used:** Prints a list of every single library your project currently depends on, including their versions.

### `go doc <package_name>`
- **Method Name:** Documentation Viewer
- **Why it is used:** Displays the official documentation for any package or function directly in your terminal so you don't have to go to a website.

---

## �📝 Key Features in this Boilerplate
1.  **RESTful Routing:** Uses the latest Go 1.22+ `ServeMux` features for clean path parameters like `/users/{id}`.
2.  **Layered Architecture:** Separates database logic from API logic so you can swap databases easily.
3.  **Graceful Shutdown:** The server han dles stop signals (Ctrl+C) properly without losing data.
4.  **Logger Utility:** A pre-configured logger for tracking application events.
