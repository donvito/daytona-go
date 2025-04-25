# Daytona Go

This project uses the Daytona SDK to create and manage sandboxes securely.

## Setup

1. Clone this repository
2. Create a `.env` file in the root directory with your Daytona API key:
   ```
   DAYTONA_API_KEY=your_daytona_api_key_here
   ```
3. Install required dependencies:
   ```
   pip install -r requirements.txt
   ```

## Usage

Run the main script:

```
python main.py
```

## Security Notes

- The API key is stored in an environment variable to keep it out of the codebase
- The `.env` file is included in `.gitignore` to prevent accidental commits
- For production use, consider using a more secure secret management system 