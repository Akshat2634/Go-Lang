# Go Programming Repository

![Golang Gopher](https://github.com/user-attachments/assets/20681416-5107-4147-84f6-75bd5c888a2b)
[![Go Version](https://img.shields.io/badge/Go-1.19+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen.svg)](https://github.com/username/go-lang)

## Overview
This repository serves as a comprehensive collection of Go programming exercises, projects, and learning materials. It demonstrates proficiency in Go language fundamentals, advanced concepts, and real-world application development.

## Repository Structure

### Core Components
- **Fundamentals**: Basic Go syntax, data types, and control structures
- **Concurrency**: Goroutines, channels, and concurrent programming patterns
- **Projects**: Production-ready applications and utilities
- **Best Practices**: Code organization, testing, and Go idioms
- **Package Development**: Custom packages and module management

## Learning Objectives

- Master Go language syntax and semantics
- Implement concurrent programming with goroutines and channels
- Develop scalable applications following Go best practices
- Explore the Go ecosystem and standard library
- Build real-world projects demonstrating Go's capabilities

## Technologies & Concepts

- **Language**: Go (Golang)
- **Concurrency**: Goroutines, Channels, Select statements
- **Testing**: Unit testing with Go's testing package
- **Modules**: Go modules for dependency management
- **Standard Library**: Extensive use of Go's built-in packages
- **Design Patterns**: Implementation of common patterns in Go

## Getting Started

### Prerequisites
- Go 1.19 or higher installed on your system
- Basic understanding of programming concepts
- Terminal/command line access

### Quick Start

1. **Clone the repository**
   ```bash
   git clone https://github.com/Akshat2634/Go-Lang.git
   cd go-lang
   ```

2. **Initialize Go modules** (if needed)
   ```bash
   go mod init go-lang
   go mod tidy
   ```

3. **Run specific projects**
   ```bash
   # Run the TodoList application
   go run Basics/main.go
   
   # Run the Quiz Game
   go run QuizGame/quiz.go
   ```

4. **Run tests**
   ```bash
   go test ./...
   ```

5. **Build executables**
   ```bash
   # Build TodoList app
   go build -o bin/todolist Basics/main.go
   
   # Build Quiz Game
   go build -o bin/quizgame QuizGame/quiz.go
   ```

### Featured Projects

#### 📝 TodoList Application (`Basics/main.go`)
- Interactive command-line todo manager
- Multiple input methods demonstration
- Web server with HTTP handlers
- Task management functionality

#### 🎯 Quiz Game (`QuizGame/quiz.go`)
- Interactive quiz application
- Game logic implementation
- User interaction patterns

### Project Structure

This repository contains:

- ✏️ **Practice exercises** - Fundamental Go concepts and syntax
- ⚙️ **Mini projects and utilities** - Real-world applications like TodoList and Quiz Game
- 📚 **Notes and references** - Go syntax, concurrency, and modules documentation
- 🔬 **Experiments** - Package development and design pattern implementations

---

Feel free to fork or star if you're learning Go too!  
Let's build something cool with Go! 🚀
