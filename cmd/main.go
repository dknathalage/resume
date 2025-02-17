package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

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

const markdownTemplate = `# {{ .Name }}

**Email:** [{{ .Email }}](mailto:{{ .Email }}) | **Phone:** {{ .Phone }}  
**LinkedIn:** [{{ .LinkedIn }}]({{ .LinkedIn }}) | **GitHub:** [{{ .GitHub }}]({{ .GitHub }})

---

## Summary
{{ .Summary }}

---

## Experience
{{ range .Experience }}
### {{ .Position }} - {{ .Company }} ({{ .Duration }})
<ul>
{{ range .Details }}<li>{{ . }}</li>{{ end }}
</ul>
{{ end }}

---

## Education
{{ range .Education }}
### {{ .Degree }} - {{ .Institution }} ({{ .Year }})
{{ end }}

---

## Skills
<ul>
{{ range .Skills }}<li>{{ . }}</li>{{ end }}
</ul>

---

## Projects
{{ range .Projects }}
### {{ .Name }}
{{ .Description }}

**Technologies:**  
{{ range .Technologies }}- {{ . }}
{{ end }}
{{ end }}
`

func generateMarkdown(resume Resume, outputFile string) {
	tmpl, err := template.New("resume").Parse(markdownTemplate)
	if err != nil {
		log.Fatalf("Failed to parse markdown template: %v", err)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Failed to create markdown file: %v", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, resume); err != nil {
		log.Fatalf("Failed to render markdown: %v", err)
	}

	fmt.Println("Markdown resume generated:", outputFile)
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

	generateMarkdown(resume, "resume.md")
}
