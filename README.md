# Go SOCKS5 Proxy Server

A lightweight, secure, and feature-rich SOCKS5 proxy server implementation in Go with flexible authentication and multi-platform support.

## Features

- 🚀 **SOCKS5 Protocol Support** - Full SOCKS5 proxy implementation
- 🔐 **Flexible Authentication** - File-based and command-line user management
- 🐳 **Docker Ready** - Multi-architecture Docker images (amd64, arm64)
- ⚙️ **Command Line Interface** - Rich CLI with multiple configuration options
- 🌍 **Multi-Platform** - Binaries for Linux, Windows (32/64-bit), and ARM
- 📦 **No Authentication Mode** - Optional anonymous access
- 🔄 **Auto-Releases** - Automated GitHub Actions for building and releasing
- 🛡️ **Secure** - Support for multiple users with different passwords
- 📝 **Systemd Service** - Easy system service integration

## Installation

### Option 1: Download Pre-built Binaries

Download the latest release for your platform from [GitHub Releases](https://github.com/ariadata/go-socks5-proxy/releases):

- **Linux (64-bit):** `go-socks5-proxy-linux-amd64.tar.gz`
- **Linux (ARM64):** `go-socks5-proxy-linux-arm64.tar.gz`
- **Linux (ARM 32-bit):** `go-socks5-proxy-linux-arm.tar.gz`

```bash
# Download and extract (example for Linux amd64)
wget https://github.com/ariadata/go-socks5-proxy/releases/latest/download/go-socks5-proxy-linux-amd64.tar.gz
tar -xzf go-socks5-proxy-linux-amd64.tar.gz
chmod +x go-socks5-proxy-linux-amd64
```

### Option 2: Using Docker

```bash
# Pull the latest image
docker pull ghcr.io/ariadata/go-socks5-proxy:latest

# Run with Docker
docker run -d -p 1080:1080 \
  -v $(pwd)/users.conf:/app/users.conf \
  --name socks5-proxy \
  ghcr.io/ariadata/go-socks5-proxy:latest
```

### Option 3: Build from Source

```bash
git clone https://github.com/ariadata/go-socks5-proxy.git
cd go-socks5-proxy
go build -o socks5-server main.go
```

## Usage

### Command Line Options

```bash
./socks5-server [OPTIONS]
```

**Available Options:**
- `--host HOST` - Host to bind to (default: 0.0.0.0)
- `--port PORT` - Port to listen on (default: 1080)
- `--users FILE` - Path to users configuration file
- `--user USER` - User credentials in format username:password (can be used multiple times)
- `--version` - Show version information
- `--help` - Show help message

### Authentication Modes

#### 1. No Authentication (Anonymous Access)
```bash
./socks5-server --port 1080
```

#### 2. File-based Authentication
Create a `users.conf` file:
```plaintext
# SOCKS5 Proxy User Configuration
# Format: username:password
# Lines starting with # are comments

user1:secure_password_123
user2:another_secure_password
admin:very_secure_admin_password
```

Run with file authentication:
```bash
./socks5-server --users users.conf --port 1080
```

#### 3. Command-line Authentication
```bash
./socks5-server --user "user1:pass1" --user "user2:pass2" --port 1080
```

#### 4. Mixed Authentication (File + Command-line)
```bash
./socks5-server --users users.conf --user "extrauser:extrapass" --port 1080
```

### Environment Variables

You can also configure the server using environment variables:
- `SOCKS5_PORT` - Port to listen on (default: :1080)
- `SOCKS5_CONFIG` - Path to users configuration file

```bash
export SOCKS5_PORT=":8080"
export SOCKS5_CONFIG="/etc/socks5/users.conf"
./socks5-server
```

## Systemd Service Setup

Create a systemd service for automatic startup:

```bash
# Copy binary to system location
sudo cp socks5-server /usr/local/bin/
sudo chmod +x /usr/local/bin/socks5-server

# Create users configuration
sudo mkdir -p /etc/socks5
sudo cp users.conf /etc/socks5/

# Create systemd service
sudo tee /etc/systemd/system/socks5-server.service > /dev/null << 'EOF'
[Unit]
Description=SOCKS5 Proxy Server
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/socks5-server --users /etc/socks5/users.conf --port 1080
Restart=always
RestartSec=5
User=nobody
Group=nogroup
WorkingDirectory=/etc/socks5

[Install]
WantedBy=multi-user.target
EOF

# Enable and start the service
sudo systemctl daemon-reload
sudo systemctl enable socks5-server
sudo systemctl start socks5-server

# Check status
sudo systemctl status socks5-server
```

## Docker Deployment

### Using Docker Run

```bash
# With authentication
docker run -d \
  --name socks5-proxy \
  -p 1080:1080 \
  -v $(pwd)/users.conf:/app/users.conf \
  --restart unless-stopped \
  ghcr.io/ariadata/go-socks5-proxy:latest

# Without authentication (anonymous access)
docker run -d \
  --name socks5-proxy \
  -p 1080:1080 \
  --restart unless-stopped \
  ghcr.io/ariadata/go-socks5-proxy:latest
```

### Using Docker Compose

Create `docker-compose.yml`:

```yaml
version: "3.8"
services:
  socks5-proxy:
    image: ghcr.io/ariadata/go-socks5-proxy:latest
    container_name: socks5-proxy
    restart: unless-stopped
    ports:
      - "1080:1080"
    volumes:
      - ./users.conf:/app/users.conf
    # Optional: Override default command
    # command: ["./socks5-server", "--users", "/app/users.conf", "--port", "1080"]
```

Start the service:
```bash
docker-compose up -d
```

## Testing the Proxy

### Using curl

```bash
# Test with authentication
curl -x socks5://username:password@127.0.0.1:1080 https://httpbin.org/ip

# Test without authentication (if running in anonymous mode)
curl -x socks5://127.0.0.1:1080 https://httpbin.org/ip

# Test SOCKS5h (hostname resolution through proxy)
curl -x socks5h://username:password@127.0.0.1:1080 https://httpbin.org/ip
```

### Using Browser

Configure your browser's SOCKS5 proxy settings:
- **Proxy Type:** SOCKS5
- **Host:** 127.0.0.1 (or your server IP)
- **Port:** 1080
- **Username/Password:** As configured

## Configuration Examples

### High Security Setup
```bash
# Create secure users file with strong passwords
echo "admin:$(openssl rand -base64 32)" > users.conf
echo "user1:$(openssl rand -base64 32)" >> users.conf

# Set secure file permissions
chmod 600 users.conf

# Run server
./socks5-server --users users.conf --host 127.0.0.1 --port 1080
```

### Multi-User Corporate Setup
```bash
# users.conf
admin:AdminSecurePass123!
sales_team:SalesPass456@
dev_team:DevSecurePass789#
guest:GuestTempPass000$

# Run server
./socks5-server --users users.conf --port 1080
```

## Development & Building

### GitHub Actions Workflows

This project includes automated workflows:

1. **Docker Build** (`build.yml`) - Builds and pushes Docker images to GHCR on every push to main
2. **Multi-Platform Release** (`release-multiplatform.yml`) - Creates binaries for all platforms on release

### Manual Building

```bash
# Build for current platform
go build -o socks5-server main.go

# Cross-compile for different platforms
GOOS=linux GOARCH=amd64 go build -o socks5-server-linux-amd64 main.go
GOOS=windows GOARCH=amd64 go build -o socks5-server-windows-amd64.exe main.go
GOOS=linux GOARCH=arm64 go build -o socks5-server-linux-arm64 main.go
```

## Security Considerations

1. **Strong Passwords** - Use complex passwords (minimum 12 characters)
2. **File Permissions** - Set `users.conf` to 600 (`chmod 600 users.conf`)
3. **Network Security** - Use firewall rules to restrict access
4. **Regular Updates** - Keep the proxy server updated
5. **Monitoring** - Monitor logs for suspicious activity
6. **User Management** - Regularly rotate credentials

## Troubleshooting

### Common Issues

1. **Connection Refused**
   ```bash
   # Check if server is running
   netstat -tlnp | grep 1080
   
   # Check logs
   journalctl -u socks5-server -f
   ```

2. **Authentication Failed**
   - Verify credentials in `users.conf`
   - Check file permissions
   - Ensure no extra spaces in username:password format

3. **Docker Issues**
   ```bash
   # Check container logs
   docker logs socks5-proxy
   
   # Check if users.conf is mounted correctly
   docker exec socks5-proxy ls -la /app/
   ```

### Logs and Monitoring

```bash
# View systemd service logs
sudo journalctl -u socks5-server -f

# View Docker container logs
docker logs -f socks5-proxy

# Check service status
sudo systemctl status socks5-server
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- 📖 **Documentation:** Check this README and inline help (`--help`)
- 🐛 **Issues:** [GitHub Issues](https://github.com/ariadata/go-socks5-proxy/issues)
- 💬 **Discussions:** [GitHub Discussions](https://github.com/ariadata/go-socks5-proxy/discussions)

## Acknowledgments

- Built with [go-socks5](https://github.com/armon/go-socks5) library
- Inspired by the need for a simple, secure SOCKS5 proxy solution