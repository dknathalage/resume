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
		Achievements []string `yaml:"details"`
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

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <script src="https://cdn.tailwindcss.com"></script>
    <title>{{ .Name }} - Resume</title>
    <link rel="stylesheet" type="text/css" href="https://cdn.jsdelivr.net/gh/bitmaks/cm-web-fonts@latest/fonts.css">
    <style>
        body {
        font-family: "Computer Modern Serif", serif;
        }
    </style>
</head>
<body class="max-w-3xl mx-auto my-5 text-gray-800 text-sm">
    <header class="text-center mb-3">
        <h1 class="text-2xl font-bold">{{ .Name }}</h1>
        <p class="">
            <a href="mailto:{{ .Email }}" class="text-blue-600 hover:underline">{{ .Email }}</a> | 
            {{ .Phone }} | 
            <a href="{{ .LinkedIn }}" class="text-blue-600 hover:underline">LinkedIn</a> | 
            <a href="{{ .GitHub }}" class="text-blue-600 hover:underline">GitHub</a>
        </p>
    </header>

    <section class="section">
        <h2 class="text-lg font-semibold mt-2 border-b">Summary</h2>
        <p class="text-md my-1">{{ .Summary }}</p>
    </section>

    {{ if .KeySkills }}
    <section class="section">
        <h2 class="text-lg font-semibold mt-2 border-b">Key Skills</h2>
        <ul class="list-disc grid grid-cols-2 my-1 pl-4">
            {{ range .KeySkills }}<li>{{ . }}</li>{{ end }}
        </ul>
    </section>
    {{ end }}

    <section class="section">
        <h2 class="text-lg font-semibold mt-2 border-b">Education</h2>
        {{ range .Education }}
        <div class="my-1">
            <div class="flex justify-between">
                <span class="font-bold ">{{ .Degree }}</span>
                <span class="text-gray-600">{{ .Year }}</span>
            </div>
            <p class="text-gray-700 italic">{{ .Institution }}</p>
            <ul class="list-disc pl-4 text-gray-700">
                {{ range .Achievements }}<li>{{ . }}</li>{{ end }}
            </ul>
        </div>
        {{ end }}
    </section>

    <section class="section">
        <h2 class="text-lg font-semibold mt-2 border-b">Experience</h2>
        {{ range .Experience }}
        <div class=" my-1">
            <div class="flex justify-between">
                <span class="font-bold">{{ .Position }}</span>
                <span class="text-gray-600">{{ .Duration }}</span>
            </div>
            <p class="text-gray-700 italic">{{ .Company }}</p>
            <ul class="list-disc pl-4 text-gray-700">
                {{ range .Details }}<li>{{ . }}</li>{{ end }}
            </ul>
        </div>
        {{ end }}
    </section>

    {{ if .Projects }}
    <section class="section">
        <h2 class="text-lg font-semibold mt-2 border-b">Projects</h2>
        {{ range .Projects }}
        <div class="my-1">
            <h3 class="font-bold ">{{ .Name }}</h3>
            <p class="text-gray-700"><strong>Technologies:</strong> {{ range .Technologies }}{{ . }}, {{ end }}</p>
            <ul class="list-disc pl-4 text-gray-700">
                {{ range .Details }}<li>{{ . }}</li>{{ end }}
            </ul>
        </div>
        {{ end }}
    </section>
    {{ end }}

    {{ if .References }}
    <section class="section">
        <h2 class="text-lg font-semibold mt-2 border-b">References</h2>
        <p class="my-1">{{ .References }}</p>
    </section>
    {{ end }}

</body>
</html>


`

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
