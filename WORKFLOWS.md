# GitHub Actions Workflows

This repository includes GitHub Actions workflows for automated building, testing, and releasing.

## Workflows Overview

### 1. Build and Push Docker Image (`build.yml`)
**Trigger:** Push to `main` branch or Pull Requests
**Purpose:** Builds and pushes Docker images to GitHub Container Registry (GHCR)

**Features:**
- Builds multi-architecture images (linux/amd64, linux/arm64)
- Uses Docker layer caching for faster builds
- Only pushes images on main branch (not on PRs)
- Tags images with branch name, commit SHA, and `latest`

**Image Location:** `ghcr.io/YOUR_USERNAME/go-socks5-proxy`

### 2. Multi-Platform Release (`release-multiplatform.yml`)
**Trigger:** When a GitHub release is created
**Purpose:** Builds binaries for multiple platforms and architectures

**Supported Platforms:**
- Linux (amd64, arm64, arm)
- Windows (386/32-bit, amd64/64-bit)

**Features:**
- Creates compressed archives (.tar.gz for Linux, .zip for Windows)
- Generates checksums for each platform
- Injects version information
- Optimized binaries with stripped symbols

## How to Use

### Creating a Release

1. **Create a Git Tag:**
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

2. **Create GitHub Release:**
   - Go to your repository on GitHub
   - Click "Releases" → "Create a new release"
   - Choose your tag (v1.0.0)
   - Add release notes
   - Click "Publish release"

3. **Automatic Build:**
   - The release workflow will trigger automatically
   - Binaries will be built and attached to the release
   - Docker images will be built and pushed to GHCR

### Using Docker Images

**Pull the latest image:**
```bash
docker pull ghcr.io/YOUR_USERNAME/go-socks5-proxy:latest
```

**Run the container:**
```bash
docker run -d -p 1080:1080 \
  -v $(pwd)/users.conf:/app/users.conf \
  ghcr.io/YOUR_USERNAME/go-socks5-proxy:latest
```

### Environment Variables

The workflows support these environment variables:
- `SOCKS5_PORT`: Port to listen on (default: :1080)
- `SOCKS5_CONFIG`: Path to users configuration file (default: users.conf)

## Self-Hosted Runners

These workflows are configured to use self-hosted runners with:
```yaml
runs-on:
  group: self-hosted
```

If you want to use GitHub-hosted runners instead, change this to:
```yaml
runs-on: ubuntu-latest
```

## Security Notes

1. **GITHUB_TOKEN:** Automatically provided by GitHub Actions
2. **GHCR Access:** Uses the repository's GITHUB_TOKEN for authentication
3. **Permissions:** Workflows have minimal required permissions
4. **Secrets:** No additional secrets required for basic functionality

## Troubleshooting

### Build Failures
- Check Go version compatibility (currently set to 1.24.3)
- Ensure all dependencies are available
- Check for syntax errors in Go code

### Docker Push Failures
- Verify GHCR permissions
- Check if the repository is public or if you have package write permissions
- Ensure Docker login is successful

### Release Upload Failures
- Check repository permissions
- Verify the release was created properly
- Ensure the GITHUB_TOKEN has write access to releases 