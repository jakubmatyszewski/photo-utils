#!/bin/bash

# Find all DNG files in the current directory and subdirectories
find . -type f -name "*.DNG" | while read -r file; do
    # Convert the DNG to PNG
    convert "$file" "${file%.DNG}.png"
    
    # If conversion was successful, delete the original DNG file
    if [ $? -eq 0 ]; then
        rm "$file"
    else
        echo "Conversion failed for $file"
    fi
done

