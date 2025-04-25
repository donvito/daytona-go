import os
from dotenv import load_dotenv
from daytona_sdk import Daytona, DaytonaConfig, CreateSandboxParams, SessionExecuteRequest

load_dotenv()
# Run the code securely inside the sandbox
try:
    # Initialize the Daytona client with proper configuration

    api_key = os.getenv("DAYTONA_API_KEY")
    if api_key is None:
        raise Exception("DAYTONA_API_KEY environment variable not set.")

    config = DaytonaConfig(
        api_key=api_key,  # Replace with your actual API key
        api_url="https://app.daytona.io/api",
        target="us"
    )
    daytona = Daytona(config)

    # Create a sandbox with CreateWorkspaceParams class
    params = CreateSandboxParams(   
        # image="alpine:3.21.3",
        language="python",  # Using Python image   
        public=True
    )
    sandbox = daytona.create(params)

    root_dir = sandbox.get_user_root_dir()
    print(root_dir)

    with open("app/app", "rb") as f:
        content = f.read()
    sandbox.fs.upload_file(root_dir + "/app", content)
    sandbox.fs.set_file_permissions(root_dir + "/app", "744")   
    
    session_id = f"{root_dir}/app".replace("/", "-")
    print(session_id)
    sandbox.process.create_session(session_id)        
    sandbox.process.execute_session_command(session_id, SessionExecuteRequest(command=f"{root_dir}/app", var_async=True))     
    session = sandbox.process.get_session(session_id)
    for command in session.commands:
        print(f"Command: {command.command}, Exit Code: {command.exit_code}")    
            
except Exception as e:
    print(f"An error occurred: {e}")