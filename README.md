
# Golang Playground 🚀

This repository is a personal learning and experimentation space for mastering **Go (Golang)**.  
It contains multiple standalone modules, each focusing on a specific **Go language concept**, organized for clarity and easy navigation.

---

## 📂 Project Structure

```plaintext
.
├── Basic/        # Go syntax fundamentals & basic programs
├── Command/      # Working with command-line arguments and flags
├── Error/        # Error handling patterns and best practices
├── Interface/    # Interfaces and polymorphism in Go
├── Logging/      # Logging techniques and structured logging
├── Pointer/      # Pointer usage, memory references & dereferencing
├── Struct/       # Structs, methods, and composition
└── go.work       # Workspace file linking all submodules
````

---

## 📚 Modules Overview

### 1️⃣ Basic

* Introduction to Go syntax
* Variables, constants, and types
* Control structures (`if`, `for`, `switch`)
* Functions and packages

### 2️⃣ Command

* Accessing **command-line arguments** via `os.Args`
* Using `flag` package for argument parsing
* Writing CLI-style programs

### 3️⃣ Error

* Creating and returning errors using `errors.New` and `fmt.Errorf`
* Custom error types
* Error wrapping and unwrapping
* Best practices for error handling

### 4️⃣ Interface

* Defining interfaces
* Implementing interface methods
* Type assertions and type switches
* Real-world interface use cases

### 5️⃣ Logging

* Basic logging with Go’s `log` package
* Logging levels and prefixes
* Structured logging using third-party libraries (e.g., `logrus`, `zap`)
* Best practices for logging in production

### 6️⃣ Pointer

* Understanding memory addresses
* Pointer creation and dereferencing (`&`, `*`)
* Passing pointers to functions
* Pointer vs value semantics in Go

### 7️⃣ Struct

* Declaring and initializing structs
* Value and pointer receivers in methods
* Struct embedding (composition)
* Real-world struct use cases

---

## 🛠 Go Workspace Setup

This repository uses a **Go Workspace** (`go.work`) to manage multiple modules in a single repo.

### Why Workspace?

* No need for `replace` directives in `go.mod`
* Import any module locally without versioning
* Clean separation of concepts

### How to Run a Module

```bash
# Run the Basic module
cd Basic
go run .

# Or run directly from workspace root
go run ./Logging
```

---

## 📌 Requirements

* Go 1.24+ installed
* Git installed for version control

---

## 📥 Cloning & Running

```bash
git clone https://github.com/SaptatirthaChoudhury/Golang.git
cd Golang
go run ./Basic
```

---

## ✨ Author

# *Saptatirtha Choudhury*
## Learning, building, and experimenting with Go.

