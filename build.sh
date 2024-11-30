#!/bin/bash

# Define the path to the bin folder
BIN_FOLDER="./bin"

# Check if the bin folder exists
if [ ! -d "$BIN_FOLDER" ]; then
	echo "Bin folder does not exist. Creating $BIN_FOLDER..."

	# Create the bin folder
	mkdir -p "$BIN_FOLDER"

	# Check if the folder was created successfully
	if [ $? -eq 0 ]; then
		echo "Bin folder created successfully."
	else
		echo "Error: Failed to create bin folder."
		exit 1
	fi
else
	echo "Bin folder already exists at $BIN_FOLDER"
fi

exit 0
