package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Resume struct {
	Name      string   `yaml:"name"`
	Email     string   `yaml:"email"`
	Phone     string   `yaml:"phone"`
	LinkedIn  string   `yaml:"linkedin"`
	GitHub    string   `yaml:"github"`
	Summary   string   `yaml:"summary"`
	KeySkills []string `yaml:"key_skills"`
	Education []struct {
		Institution  string   `yaml:"institution"`
		Degree       string   `yaml:"degree"`
		Year         string   `yaml:"year"`
		Achievements []string `yaml:"achievements"`
	} `yaml:"education"`
	Experience []struct {
		Company  string   `yaml:"company"`
		Position string   `yaml:"position"`
		Duration string   `yaml:"duration"`
		Details  []string `yaml:"details"`
	} `yaml:"experience"`
	Projects []struct {
		Name         string   `yaml:"name"`
		Technologies []string `yaml:"technologies"`
		Details      []string `yaml:"details"`
	} `yaml:"projects"`
	TechnicalSkills []struct {
		Category string   `yaml:"category"`
		Skills   []string `yaml:"skills"`
	} `yaml:"technical_skills"`
	Community []struct {
		Role     string   `yaml:"role"`
		Duration string   `yaml:"duration"`
		Details  []string `yaml:"details"`
	} `yaml:"community"`
	References string `yaml:"references"`
}

const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ .Name }}'s Resume</title>
    <style>
        body { font-family: "Times New Roman", serif; margin: 40px; }
        h1, h2 { color: #333; }
        .contact { margin-bottom: 20px; }
        .section { margin-bottom: 30px; }
        ul { padding-left: 20px; }
    </style>
</head>
<body>
    <h1>{{ .Name }}</h1>
    <div class="contact">
        <p>Email: <a href="mailto:{{ .Email }}">{{ .Email }}</a></p>
        <p>Phone: {{ .Phone }}</p>
        <p>LinkedIn: <a href="{{ .LinkedIn }}">{{ .LinkedIn }}</a></p>
        <p>GitHub: <a href="{{ .GitHub }}">{{ .GitHub }}</a></p>
    </div>

    <div class="section">
        <h2>Summary</h2>
        <p>{{ .Summary }}</p>
    </div>

    <div class="section">
        <h2>Key Skills</h2>
        <ul>{{ range .KeySkills }}<li>{{ . }}</li>{{ end }}</ul>
    </div>

    <div class="section">
        <h2>Education</h2>
        {{ range .Education }}
        <h3>{{ .Degree }} - {{ .Institution }} ({{ .Year }})</h3>
        <ul>{{ range .Achievements }}<li>{{ . }}</li>{{ end }}</ul>
        {{ end }}
    </div>

    <div class="section">
        <h2>Experience</h2>
        {{ range .Experience }}
        <h3>{{ .Position }} - {{ .Company }} ({{ .Duration }})</h3>
        <ul>{{ range .Details }}<li>{{ . }}</li>{{ end }}</ul>
        {{ end }}
    </div>

    <div class="section">
        <h2>Projects</h2>
        {{ range .Projects }}
        <h3>{{ .Name }}</h3>
        <p><strong>Technologies:</strong> {{ range .Technologies }}{{ . }}, {{ end }}</p>
        <ul>{{ range .Details }}<li>{{ . }}</li>{{ end }}</ul>
        {{ end }}
    </div>

    <div class="section">
        <h2>References</h2>
        <p>{{ .References }}</p>
    </div>
</body>
</html>`

func generateHTML(resume Resume, outputFile string) {
	tmpl, err := template.New("resume").Parse(htmlTemplate)
	if err != nil {
		log.Fatalf("Failed to parse HTML template: %v", err)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Failed to create HTML file: %v", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, resume); err != nil {
		log.Fatalf("Failed to render HTML: %v", err)
	}

	fmt.Println("HTML resume generated:", outputFile)
}

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

	outputPath := filepath.Join("docs", filepath.Base(yamlFile[:len(yamlFile)-5])+".html")
	generateHTML(resume, outputPath)
}
