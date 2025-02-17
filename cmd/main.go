package main

import (
	"fmt"
	"log"
	"os"
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
		WAM          string   `yaml:"wam"`
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

const markdownTemplate = `# {{ .Name }}

<style>
body { font-family: "Times New Roman", serif; }
ul { padding-left: 20px; }
</style>

**Email:** [{{ .Email }}](mailto:{{ .Email }}) | **Phone:** {{ .Phone }}  
**LinkedIn:** [{{ .LinkedIn }}]({{ .LinkedIn }}) | **GitHub:** [{{ .GitHub }}]({{ .GitHub }})

---

## Summary
{{ .Summary }}

---

## Key Skills
<ul>
{{ range .KeySkills }}<li>{{ . }}</li>{{ end }}
</ul>

---

## Education
{{ range .Education }}
### {{ .Degree }} - {{ .Institution }} ({{ .Year }})
- WAM: {{ .WAM }}
<ul>
{{ range .Achievements }}<li>{{ . }}</li>{{ end }}
</ul>
{{ end }}

---

## Experience
{{ range .Experience }}
### {{ .Position }} - {{ .Company }} ({{ .Duration }})
<ul>
{{ range .Details }}<li>{{ . }}</li>{{ end }}
</ul>
{{ end }}

---

## Projects
{{ range .Projects }}
### {{ .Name }}
Technologies: *{{ range .Technologies }}{{ . }}, {{ end }}*
<ul>
{{ range .Details }}<li>{{ . }}</li>{{ end }}
</ul>
{{ end }}

---

## Technical Skills
{{ range .TechnicalSkills }}
### {{ .Category }}
<ul>
{{ range .Skills }}<li>{{ . }}</li>{{ end }}
</ul>
{{ end }}

---

## Community & Leadership
{{ range .Community }}
### {{ .Role }} ({{ .Duration }})
<ul>
{{ range .Details }}<li>{{ . }}</li>{{ end }}
</ul>
{{ end }}

---

## References
{{ .References }}
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
