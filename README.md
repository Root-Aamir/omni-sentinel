# 🚀 Omni-Sentinel

**Omni-Sentinel** is a high-performance, multi-threaded framework built in **Go (Golang)**. It is designed to handle concurrent tasks across different domains, specifically focusing on **Cybersecurity Scanning** and **Trading Intelligence**.

---

## 🛠️ Key Features

-   **Interface-Based Architecture:** Modular design allowing easy integration of new security or trading modules.
-   **High Concurrency:** Utilizes Go's `goroutines` and `sync.WaitGroups` for parallel task execution.
-   **Scout Module:** Fast TCP port scanner for network reconnaissance.
-   **Trading Module:** Real-time market monitoring (XAU/USD) with simulated intelligence.

## 📂 Project Structure

-   `cmd/sentinel/`: Main entry point and task orchestrator.
-   `pkg/scanner/`: Security modules (Port scanners, reconnaissance tools).
-   `pkg/trading/`: Financial intelligence modules (Price fetchers, indicators).
-   `pkg/utils/`: Shared utilities (Logging, network helpers).

## 🚀 Getting Started

### Prerequisites
-   Go 1.20+
-   Git

### Installation & Run
```bash
# Clone the repository
git clone [https://github.com/Root-Aamir/omni-sentinel.git](https://github.com/Root-Aamir/omni-sentinel.git)

# Initialize modules
go mod tidy

# Run the engine
go run cmd/sentinel/main.go