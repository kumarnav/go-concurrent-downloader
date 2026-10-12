# ⚡ Go Concurrent URL Checker & Downloader

A high-performance command-line utility built in Go that demonstrates concurrent network requests using **Goroutines**, **Channels**, and **WaitGroups**.

## 🚀 Features
- **True Concurrency:** Processes multiple HTTP requests simultaneously rather than sequentially.
- **Performance Tracking:** Measures individual request latency and total execution time.
- **Thread-Safe Results:** Uses buffered channels to safely collect data across asynchronous workers.

## 🛠️ Tech Stack
- **Language:** Go (Golang)
- **Standard Libraries:** `net/http`, `sync`, `time`

## 📦 Usage

1. **Run the tool:**
   ```bash
   go run main.go