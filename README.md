# Resume Generator

A Go-based tool that converts YAML resume files into beautiful HTML resumes with modern styling using Tailwind CSS.

## Features

- Convert YAML resume data to professional HTML resumes
- Clean, modern design with Tailwind CSS
- Responsive layout that works on all devices
- Automatic index page generation for multiple resumes
- Support for comprehensive resume sections:
  - Personal information and contact details
  - Professional summary
  - Key skills
  - Work experience
  - Education
  - Certifications
  - Projects
  - Technical skills
  - Community involvement
  - References

## Prerequisites

- Go 1.23.1 or higher
- Git (for cloning the repository)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/dknathalage/resumes.git
cd resumes
```

2. Install dependencies:
```bash
go mod tidy
```

## Usage

### Creating Your Resume

1. Create a new YAML file in the `resumes/` directory following the structure below:

```yaml
name: Your Name
email: your.email@example.com
phone: "+1 (555) 123-4567"
linkedin: https://www.linkedin.com/in/yourprofile/
github: https://github.com/yourusername
summary: |
  Your professional summary goes here. This should be a brief overview
  of your experience, skills, and career objectives.

key_skills:
  - Skill 1
  - Skill 2
  - Skill 3

education:
  - institution: University Name
    degree: Bachelor of Computer Science
    year: "2023"
    details:
      - GPA: 3.8/4.0
      - Relevant coursework: Data Structures, Algorithms

certifications:
  - institution: Certification Body
    certification: Certification Name
    year: "2023"
    details:
      - Description of certification

experience:
  - company: Company Name
    position: Software Engineer
    duration: "Jan 2023 - Present"
    details:
      - Developed and maintained web applications
      - Collaborated with cross-functional teams

projects:
  - name: Project Name
    technologies: ["Go", "React", "PostgreSQL"]
    details:
      - Built a full-stack web application
      - Implemented RESTful APIs

technical_skills:
  - category: Programming Languages
    skills: ["Go", "Python", "JavaScript"]
  - category: Frameworks
    skills: ["React", "Express.js", "Gin"]

community:
  - role: Volunteer Developer
    duration: "2022 - Present"
    details:
      - Contributed to open-source projects

references: Available upon request
```

### Generating a Single Resume

To generate an HTML resume from a YAML file:

```bash
go run cmd/main.go resumes/your-resume.yaml
```

This will create an HTML file in the `docs/` directory.

### Generating All Resumes

To generate HTML files for all YAML resumes and create an index page:

```bash
./generate-resumes.sh
```

This script will:
- Create a `docs/` directory if it doesn't exist
- Generate HTML files for all YAML files in the `resumes/` directory
- Create an `index.html` file listing all generated resumes
- Add a `.nojekyll` file for GitHub Pages compatibility

### Viewing Your Resume

After generation, you can:
1. Open the HTML files directly in your browser
2. Serve the `docs/` directory with a local web server
3. Deploy to GitHub Pages or any web hosting service

For local viewing with a simple web server:
```bash
# Using Python
cd docs && python -m http.server 8000

# Using Node.js (if you have npx)
cd docs && npx serve

# Then open http://localhost:8000 in your browser
```

## Project Structure

```
├── cmd/
│   └── main.go           # Main application logic
├── resumes/
│   └── *.yaml           # Your resume YAML files
├── docs/                # Generated HTML files (created automatically)
├── generate-resumes.sh  # Script to generate all resumes
├── go.mod               # Go module file
└── README.md           # This file
```

## YAML Schema

The resume YAML files support the following structure:

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | Yes | Your full name |
| `email` | string | No | Email address |
| `phone` | string | No | Phone number |
| `linkedin` | string | No | LinkedIn profile URL |
| `github` | string | No | GitHub profile URL |
| `summary` | string | Yes | Professional summary |
| `key_skills` | array | No | List of key skills |
| `education` | array | No | Education history |
| `certifications` | array | No | Professional certifications |
| `experience` | array | No | Work experience |
| `projects` | array | No | Personal/professional projects |
| `technical_skills` | array | No | Technical skills by category |
| `community` | array | No | Community involvement |
| `references` | string | No | References statement |

## Styling

The generated HTML uses:
- **Tailwind CSS** for styling (loaded via CDN)
- **Computer Modern Serif** font for a professional academic look
- Responsive design that works on desktop and mobile
- Clean, minimal layout optimized for readability

## Deployment

### GitHub Pages

1. Push your repository to GitHub
2. Run `./generate-resumes.sh` to generate the HTML files
3. Commit and push the `docs/` directory
4. Enable GitHub Pages in your repository settings, pointing to the `docs/` folder
5. Your resumes will be available at `https://yourusername.github.io/repository-name/`

### Other Hosting Services

Simply upload the contents of the `docs/` directory to any web hosting service.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test with sample resume data
5. Submit a pull request

## License

This project is open source. Please check the LICENSE file for details.

## Example

You can see an example resume by looking at `resumes/don-athalage-generic.yaml` and generating it with the provided tools.
