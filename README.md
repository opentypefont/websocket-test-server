# WebSocket Test Server
Simple WebSocket test server for applications
## Setup
### Use Docker
- With default environment value(HOST=0.0.0.0, PORT=8000)
```sh
docker run -p <port>:8000 opentypefont/websocket-test-server
```
- With custom environment value
```sh
docker run -e HOST=<host> -e PORT=<port> -p <port>:<port> opentypefont/websocket-test-server
```
### Or
- Linux/macOS
```sh
go build -o server main.go
HOST=<host> PORT=<port> ./server
```
- Windows(PowerShell)
```powershell
go build -o server.exe main.go
$env:HOST = <HOST> | $env:PORT = <PORT>
./server
```
## Routes
### `/echo`
> Server resend client's message
### `/receive`
> Server send `Message <count>` forever