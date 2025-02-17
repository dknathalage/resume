#!/bin/bash

# Ensure 'docs/' exists and create a .nojekyll file to disable Jekyll processing
mkdir -p docs
touch docs/.nojekyll

# Start building the index file
index_file="docs/index.html"
echo "<!DOCTYPE html>
<html lang='en'>
<head>
    <meta charset='UTF-8'>
    <meta name='viewport' content='width=device-width, initial-scale=1.0'>
    <title>Resume Index</title>
    <style>
        body { font-family: 'Times New Roman', serif; margin: 40px; }
        h1, h2 { color: #333; }
        ul { padding-left: 20px; }
        a { text-decoration: none; color: blue; }
        a:hover { text-decoration: underline; }
    </style>
</head>
<body>
    <h1>Resume Directory</h1>
    <ul>" > "$index_file"

# Process each YAML file
find resumes -type f -name "*.yaml" | while read yaml_file; do
    # Create matching HTML path inside 'docs/'
    html_file="docs/${yaml_file#resumes/}"  # Replace 'resumes/' with 'docs/'
    html_file="${html_file%.yaml}.html"  # Change file extension from .yaml to .html

    # Ensure target directory exists
    mkdir -p "$(dirname "$html_file")"

    # Run Go program to generate HTML
    go run cmd/main.go "$yaml_file"

    # Add the file to index.html
    relative_path="${html_file#docs/}"
    echo "        <li><a href='${relative_path}'>$(basename "${html_file}" .html)</a></li>" >> "$index_file"

    echo "Generated: $html_file"
done

# Close the index.html file
echo "    </ul>
</body>
</html>" >> "$index_file"

echo "Index file generated at $index_file"
