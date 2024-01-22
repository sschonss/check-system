# System Monitor Program Documentation

## Overview

The System Monitor program is designed to monitor and display real-time system information such as CPU usage, memory usage, and more. It utilizes a WebSocket for real-time updates and provides an API endpoint to fetch the latest system information.

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
    - [Web Interface](#web-interface)
    - [WebSocket](#websocket)
    - [API](#api)
- [File Structure](#file-structure)
- [License](#license)

## Requirements

- Go (version 1.16 or higher)
- BoltDB (embedded key/value database)
- Gorilla WebSocket (for WebSocket functionality)
- Gin (for API functionality)

## Installation

1. Clone the repository
2. Navigate to the root directory of the repository
3. Install dependencies with `go mod download`
4. Build the program with `go build -o system-monitor`
5. Run the program with `./system-monitor`

## Usage

### Web Interface

The web interface is available at `http://localhost:8081/static`. It displays the latest system information and updates in real-time.

### WebSocket

The WebSocket is available at `ws://localhost:8081/ws`. It sends the latest system information in JSON format every second.

### API

The API is available at `http://localhost:8081/api/system-info`. It returns the latest system information in JSON format.

## File Structure

```lua
.
|-- main.go
|-- static/
|   |-- index.html
|-- api/
|   |-- api.go
|-- websocket/
|   |-- websocket.go
|-- database/
|   |-- database.go
|-- models/
|   |-- models.go
|-- utils/
|   |-- utils.go
|-- websocket/
|   |-- websocket.go
|-- webserver/
|   |-- webserver.go
|-- out/
|   |-- system-monitor.log
|-- database.db
|-- go.mod
|-- go.sum
|-- README.md

```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.