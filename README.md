# 🛡️ OMNI-SENTINEL v6.5
### Advanced Cyber-Financial Reconnaissance & Monitoring Suite

**Omni-Sentinel** is a high-performance, concurrent engine built in Go. It integrates network security reconnaissance (Subnet Discovery & Port Scanning) with real-time financial intelligence (Gold/XAU-USD Scalping Alerts) and a remote Command & Control (C2) interface via Telegram.

---

## 🚀 Key Features

* **⚡ Concurrent Scanning Engine**: Leverages Go-Routines for non-blocking, parallel scanning of multiple targets and ports.
* **📡 Subnet Discovery**: Automatically identifies active devices on the local network (ARP/ICMP Discovery).
* **📈 Quant Intelligence**: Monitors XAU/USD (Gold) market price targets and sends real-time trading signals.
* **🕹️ Remote C2 (Command & Control)**: Fully interactive Telegram bot interface to monitor status, check prices, and trigger scans remotely.
* **🔐 Hardened Security**: Zero-leak architecture using Environment Variables for API secrets. No sensitive data is ever stored in the repository.
* **📜 Automated Logging**: Detailed JSON-based logging for audit trails and session analysis.

---

## 🛠️ Tech Stack

* **Language**: Go (Golang) 1.20+
* **Concurrency**: Goroutines & WaitGroups
* **Integration**: Telegram Bot API
* **Environment**: Cross-platform (Windows/Linux/macOS)

---

## 📂 Project Structure

- `cmd/sentinel/main.go` - The Master Entry Point.
- `pkg/scanner/` - Network Discovery & Port Fingerprinting logic.
- `pkg/trading/` - Market data fetching & signal generation.
- `pkg/controller/` - Telegram Command Center (C2).
- `pkg/utils/` - Secure configuration & logging utilities.

---

## ⚙️ Setup & Installation

1. **Clone the repository**:
   ```bash
   git clone [https://github.com/Root-Aamir/omni-sentinel.git](https://github.com/Root-Aamir/omni-sentinel.git)
   cd omni-sentinel