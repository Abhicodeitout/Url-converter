# Go URL Shortener

A simple URL shortener implemented in Go that allows you to create and manage shortened URLs.

## Features

- Shorten long URLs to make them easier to share.
- Redirect users to the original URL when they access the shortened URL.
- Easy-to-use web interface for URL shortening.

## Prerequisites

- Go 1.20.5 or later
- Docker (for running the application in a container)

## Installation and Usage

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/go-url-shortener.git
   cd go-url-shortener

2. Build and run the application:

    ```bash
    go build
    ./go-url-shortener

   or using Docker 

    docker build -t go-url-shortener .
    docker run -p 8080:8080 go-url-shortener

4.  Access the URL shortener web interface at http://localhost:8080 in your web browser.

# API Endpoints

-POST /shorturl: Shorten a URL via the web interface.
-POST /short: Shorten a URL via the API (requires an API key).

