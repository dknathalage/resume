find resumes -type f -name "*.yaml" | while read yaml_file; do
    # Create matching HTML path inside 'docs/'
    html_file="docs/${yaml_file#resumes/}"  # Replace 'resumes/' with 'docs/'
    html_file="${html_file%.yaml}.html"  # Change file extension from .yaml to .html

    # Ensure target directory exists
    mkdir -p "$(dirname "$html_file")"

    # Run Go program to generate HTML
    go run cmd/main.go "$yaml_file"

    echo "Generated: $html_file"
done
