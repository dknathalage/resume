find resumes -type f -name "*.yaml" | while read yaml_file; do
    md_file="${yaml_file%.yaml}.md"  # Replace .yaml with .md, keeping the same directory structure
    go run cmd/main.go "$yaml_file"  # Run the Go program to generate the markdown
    mv resume.md "$md_file"  # Move the generated file to the correct location
    echo "Generated: $md_file"
done
