#!/bin/bash

# Extract the tarball
tar -xzvf bigmile-cli.tar.gz

# Move the binary to /usr/local/bin
sudo mv your-cli-tool /usr/local/bin

# Make sure it's executable
sudo chmod +x /usr/local/bin/bigmile-cli

echo "Installation complete. You can now use 'bigmile' from the command line."
