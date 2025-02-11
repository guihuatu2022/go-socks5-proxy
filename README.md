# Go SOCKS5 Proxy Server

A lightweight and secure SOCKS5/SOCKS5h proxy server implementation in Go with multiple user authentication support.

## Features

- SOCKS5 and SOCKS5h proxy protocols support
- Multiple user authentication
- Docker support
- Simple configuration
- Lightweight and fast
- Secure authentication

## Prerequisites

- Docker and Docker Compose (for containerized deployment)
- A valid `users.conf` file with user credentials

## Configuration

### Setting Up Users

Create a `users.conf` file in your project directory with the following format:

```plaintext
username1:password1
username2:password2
```

Example:
```plaintext
myuser:mypass
myuser2:myotherpass
```

Each line should contain one user's credentials in the format `username:password`.

## Deployment

### Using Docker Run

```bash
docker run -p 1081:1080 -v ./users.conf:/app/users.conf ghcr.io/ariadata/go-socks5-proxy
```

### Using Docker Compose

1. Create a `docker-compose.yaml` file:

```yaml
version: "3.8"
services:
  socks5-proxy:
    image: 'ghcr.io/ariadata/go-socks5-proxy:latest'
    container_name: go-socks5-proxy
    restart: unless-stopped
    ports:
      - '1081:1080'
    volumes:
      - ./users.conf:/app/users.conf
```

2. Start the service:

```bash
docker-compose up -d
```

## Testing the Connection

You can test your proxy connection using curl:

### SOCKS5
```bash
curl -x socks5://username:password@127.0.0.1:1081 https://myip4.ir
```

### SOCKS5h
```bash
curl -x socks5h://username:password@127.0.0.1:1081 https://myip4.ir
```

## Port Configuration

- Default internal port: 1080
- Default exposed port: 1081 (configurable)

## Security Considerations

- Always use strong passwords in your `users.conf` file
- Regularly rotate credentials
- Monitor proxy access logs
- Keep the container and base image updated

## Troubleshooting

1. If connection fails, verify:
   - The proxy server is running (`docker ps`)
   - Correct credentials are being used
   - Ports are correctly mapped
   - `users.conf` file is properly mounted

2. Check logs:
```bash
docker logs go-socks5-proxy
```

## License

[Add your license information here]

## Contributing

[Add contribution guidelines if applicable]

## Support

For issues and feature requests, please [open an issue](https://github.com/yourusername/go-socks5-proxy/issues).