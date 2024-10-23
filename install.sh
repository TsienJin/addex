#!/bin/zsh

# Set the name of your CLI tool
CLI_NAME="addex"

# Build the Go binary
echo "Building the Go CLI tool..."
go build -o ./dist/$CLI_NAME

# Check if the build was successful
if [ $? -ne 0 ]; then
  echo "Build failed. Exiting."
  exit 1
fi

# Move the binary to /usr/local/bin (requires sudo for permissions)
echo "Moving binary to /usr/local/bin..."
sudo mv ./dist/$CLI_NAME /usr/local/bin

# Check if the move was successful
if [ $? -eq 0 ]; then
  echo "$CLI_NAME installed successfully!"
else
  echo "Failed to move the binary. Exiting."
  exit 1
fi