package main

import (
	"html/template"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Resume struct {
	Name       string `yaml:"name"`
	Email      string `yaml:"email"`
	Phone      string `yaml:"phone"`
	LinkedIn   string `yaml:"linkedin"`
	GitHub     string `yaml:"github"`
	Summary    string `yaml:"summary"`
	Experience []struct {
		Company  string   `yaml:"company"`
		Position string   `yaml:"position"`
		Duration string   `yaml:"duration"`
		Details  []string `yaml:"details"`
	} `yaml:"experience"`
	Education []struct {
		Institution string `yaml:"institution"`
		Degree      string `yaml:"degree"`
		Year        string `yaml:"year"`
	} `yaml:"education"`
	Skills   []string `yaml:"skills"`
	Projects []struct {
		Name         string   `yaml:"name"`
		Description  string   `yaml:"description"`
		Technologies []string `yaml:"technologies"`
	} `yaml:"projects"`
}

const resumeTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ .Name }}'s Resume</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        h1, h2 { color: #333; }
        .contact { margin-bottom: 20px; }
        .section { margin-bottom: 30px; }
        ul { padding-left: 20px; }
    </style>
</head>
<body>
    <h1>{{ .Name }}</h1>
    <div class="contact">
        <p>Email: {{ .Email }}</p>
        <p>Phone: {{ .Phone }}</p>
        <p>LinkedIn: <a href="{{ .LinkedIn }}">{{ .LinkedIn }}</a></p>
        <p>GitHub: <a href="{{ .GitHub }}">{{ .GitHub }}</a></p>
    </div>

    <div class="section">
        <h2>Summary</h2>
        <p>{{ .Summary }}</p>
    </div>

    <div class="section">
        <h2>Experience</h2>
        {{ range .Experience }}
        <h3>{{ .Position }} - {{ .Company }} ({{ .Duration }})</h3>
        <ul>
            {{ range .Details }}<li>{{ . }}</li>{{ end }}
        </ul>
        {{ end }}
    </div>

    <div class="section">
        <h2>Education</h2>
        {{ range .Education }}
        <h3>{{ .Degree }} - {{ .Institution }} ({{ .Year }})</h3>
        {{ end }}
    </div>

    <div class="section">
        <h2>Skills</h2>
        <ul>
            {{ range .Skills }}<li>{{ . }}</li>{{ end }}
        </ul>
    </div>

    <div class="section">
        <h2>Projects</h2>
        {{ range .Projects }}
        <h3>{{ .Name }}</h3>
        <p>{{ .Description }}</p>
        <p>Technologies: {{ range .Technologies }}{{ . }}, {{ end }}</p>
        {{ end }}
    </div>
</body>
</html>
`

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <path_to_yaml>")
	}

	yamlFile := os.Args[1]
	data, err := os.ReadFile(yamlFile)
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	var resume Resume
	if err := yaml.Unmarshal(data, &resume); err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	tmpl, err := template.New("resume").Parse(resumeTemplate)
	if err != nil {
		log.Fatalf("Failed to parse HTML template: %v", err)
	}

	if err := tmpl.Execute(os.Stdout, resume); err != nil {
		log.Fatalf("Failed to render template: %v", err)
	}
}
