
# Key Guardian 🛡️  
**AI-Powered Parental Control for Safer Digital Experiences**

## Overview

**Key Guardian** is a parental control application designed to help parents monitor their child's online activity and ensure a safer digital environment.  
The application is installed on the child's laptop, where it **detects and logs keystrokes** in real-time.  
The captured data is sent to an **AI-based backend** that analyzes the keystrokes to detect the use of **abusive language** or **NSFW content**.

If any inappropriate content is detected, the results are logged onto a **live dashboard**, enabling parents to stay updated on their child's activities instantly.

---

## Features
- 🎯 **Real-time Keystroke Detection**  
- 🧠 **AI-based Content Analysis** (Abusive/NSFW language detection)
- 📊 **Live Parent Dashboard**  
- 🛡️ **Secure Data Logging**  
- ⚡ **Fast and Lightweight**  

---

## Tech Stack

| Area         | Technology Used                  |
| ------------ | --------------------------------- |
| Frontend     | React.js, TailwindCSS, shadcn UI  |
| Backend API  | Golang (Go)                       |
| Keylogger App| Electron.js, React.js             |

---

## Installation

### Prerequisites
- Node.js
- Golang
- npm or yarn
- Docker (optional for backend)

### Frontend Setup
```bash
cd frontend
npm install
npm run dev
```

### Backend Setup
```bash
cd backend
go mod tidy
go run main.go
```

### Keylogger (Electron App) Setup
```bash
cd keylogger
npm install
npm run dev
```

---

## How It Works

1. **Keystroke Monitoring**:  
   The Electron app runs in the background and logs keystrokes from the child's laptop.

2. **AI Analysis**:  
   Logged keystrokes are sent securely to the backend API, where an AI model analyzes them for abusive and NSFW content.

3. **Dashboard Updates**:  
   If inappropriate content is detected, the parent's dashboard receives live updates and logs the incident for review.

---

## Future Improvements 🚀
- Advanced AI models with context understanding
- Notification system (Email/SMS alerts)
- Mobile dashboard app for parents
- End-to-end encryption for keystroke data
- Whitelist/Blacklist customization for parents

---

## Built By
[Vincent Samuel Paul](https://github.com/VincentSamuelPaul)

---

## License

This project is licensed under the [MIT License](LICENSE).

---

