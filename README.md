# Daytona Go

This project demonstrates how to run a Go application in a Daytona sandbox.

## Setup

1. Clone this repository
2. Create a `.env` file in the root directory with your Daytona API key:
   ```
   DAYTONA_API_KEY=your_daytona_api_key_here
   ```
3. Build the Go application:
   ```
   cd app
   CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .
   ```
4. Deploy the Go application to Daytona:

   ```
   python main.py
   ```

