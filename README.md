# 100 Days of Code Challenge

A language-agnostic approach to mastering programming through 100 days of consistent practice and project-based learning.

## Overview

This repository contains a structured 100-day coding challenge that can be completed using **any programming language** of your choice. Whether you're learning Python, JavaScript, Java, Go, Rust, or any other language, these requirements will guide you through essential programming concepts from beginner to professional level.

### Key Features
✅ **100 Language-Agnostic Requirements** - Each day has clear, language-independent goals

✅ **Multi-Language Support** - Organized folders for different programming languages

✅ **Progressive Difficulty** - From beginner to professional portfolio projects

✅ **Flexible Learning** - Use one language or try multiple languages

✅ **Project-Based** - Learn by building real applications

## Project Structure

```
100-days-of-code/
├── README.md              # This file
├── LICENSE                # Project license
├── .gitignore             # Git ignore patterns
├── requirements/          # Daily requirements (language-agnostic)
│   ├── day001.md         # Day 1 requirements
│   ├── day002.md         # Day 2 requirements
│   └── ...               # 100 days total
├── python/               # Python implementations (optional)
│   ├── 001/
│   ├── 002/
│   └── ...
├── js/                   # JavaScript implementations (optional)
│   ├── 001/
│   ├── 002/
│   └── ...
├── go/                   # Go implementations (optional)
│   ├── 001/
│   ├── 002/
│   └── ...
├── bash/                 # Bash implementations (optional)
│   ├── 001/
│   ├── 002/
│   └── ...
├── java/                 # Java implementations (optional)
│   ├── 001/
│   ├── 002/
│   └── ...
└── [your-language]/      # Your chosen language folder
    ├── 001/
    ├── 002/
    └── ...
```

## How It Works

### 1. Daily Requirements
Each day has a corresponding requirement file in the `requirements/` folder (e.g., `day001.md`, `day002.md`). These files outline:
- **Concepts to Practice**: Core programming concepts for that day
- **Project**: The main project to build (when applicable)
- **Level**: Difficulty level with visual indicators

### 2. Learning Progression

The challenge is divided into four progressive levels:

#### 🔰 Beginner (Days 1-14)
- Variables and Data Types
- Control Flow and Logical Operators
- Lists and Randomization
- Loops and Functions
- Basic Projects: Hangman, Calculator, Higher Lower Game

#### 📚 Intermediate (Days 15-31)
- Object-Oriented Programming
- GUI Development
- File Handling
- CSV Data Manipulation
- Projects: Snake Game, Pong, Password Manager, Flash Cards

#### 👨‍💻 Intermediate+ (Days 32-58)
- APIs and Web Scraping
- Email Automation
- HTML & CSS
- Web Frameworks (Flask/Express/etc.)
- Browser Automation
- Projects: Flight Finder, Spotify Playlist Creator, Web Applications

#### 🏆 Advanced (Days 59-80)
- Database Integration and ORM
- Authentication and Security
- Data Science and Visualization
- Statistical Analysis
- Deployment
- Projects: Blog Platform, Movie Database, Data Analysis

#### ⚔ Professional Portfolio Projects (Days 81-100)
- Real-world portfolio projects
- Full-stack applications
- Machine Learning applications
- Projects: Portfolio Website, E-commerce, Space Race Analysis, Earnings Prediction

## Quick Start

1. **Fork or clone this repository**
2. **Choose your programming language** (Python, JavaScript, Go, etc.)
3. **Read `requirements/day001.md`**
4. **Create your language folder** (e.g., `python/001/`, `js/001/`)
5. **Start coding!**

## Getting Started

### 1. Choose Your Language
Decide which programming language you want to use for the challenge. Examples:
- Python
- JavaScript/TypeScript
- Java
- C#
- Go
- Rust
- Ruby
- Any language you want to learn!

### 2. Set Up Your Environment
Create a language folder (if it doesn't exist) and organize your daily work within it:
```bash
# Choose your language folder (e.g., python, js, go, rust, java, etc.)
mkdir -p python/001
cd python/001

# Or for JavaScript
mkdir -p js/001
cd js/001

# Or for Bash
mkdir -p bash/001
cd bash/001

# Or create your own language folder
mkdir -p rust/001
cd rust/001
```

### 3. Read the Daily Requirement
```bash
# From your language/day folder
cat ../../requirements/day001.md

# Or from the root
cat requirements/day001.md
```

### 4. Implement the Project
Build the day's project using your chosen language, applying the concepts outlined in the requirements.

### 5. Document Your Work
Create a README.md in each day's folder explaining:
- What you learned
- Challenges you faced
- Solutions you implemented
- Any notes for future reference

### 6. Track Your Progress
Commit your work daily to maintain consistency and track your learning journey.

## Example Workflow

### Python Example
```bash
# Day 1
mkdir -p python/001
cd python/001
cat ../../requirements/day001.md

# Read requirements and implement Band Name Generator
# Create main.py
# Write your code
# Test your implementation
# Create README.md with learnings

git add .
git commit -m "Day 1: Band Name Generator - Variables and Input (Python)"
```

### JavaScript Example
```bash
# Day 1
mkdir -p js/001
cd js/001
cat ../../requirements/day001.md

# Read requirements and implement Band Name Generator
# Create index.js
# Write your code
# Test your implementation
# Create README.md with learnings

git add .
git commit -m "Day 1: Band Name Generator - Variables and Input (JavaScript)"
```

### Multi-Language Approach
You can even complete the challenge in multiple languages simultaneously:
```bash
# Implement the same day in different languages
python/001/main.py
js/001/index.js
go/001/main.go
```

## Language Folders

The repository supports multiple programming languages organized in separate folders:

- **`python/`** - Python implementations
- **`js/`** - JavaScript/TypeScript implementations
- **`go/`** - Go implementations
- **`bash/`** - Bash/Shell script implementations
- **`java/`** - Java implementations
- **Add your own!** - Create folders for any language (rust/, csharp/, cpp/, ruby/, etc.)

### Multi-Language Learning
You can:
1. **Single Language**: Focus on one language throughout all 100 days
2. **Multiple Languages**: Implement each day in different languages to compare approaches
3. **Progressive**: Start with one language, switch to another midway to learn differences
4. **Collaborative**: Work with others, each using different languages

**Note**: Start your implementation from scratch before looking at other implementations. Learning comes from solving problems yourself!

## Language-Agnostic Approach

All requirements have been crafted to be language-agnostic. Here's how concepts translate across different languages:

| Requirement | Python | JavaScript | Go | Java | Rust |
|-------------|--------|------------|-----|------|------|
| GUI Frameworks | Tkinter, PyQt | Electron, React | Fyne, Gio | JavaFX, Swing | egui, iced |
| Web Frameworks | Flask, Django | Express, Next.js | Gin, Echo | Spring Boot | Actix, Rocket |
| Data Analysis | Pandas, NumPy | D3.js, Danfo.js | gonum | Apache Commons | ndarray, polars |
| Testing | pytest, unittest | Jest, Mocha | testing pkg | JUnit | cargo test |
| HTTP Requests | requests | axios, fetch | net/http | HttpClient | reqwest |
| Database ORM | SQLAlchemy | Sequelize, Prisma | GORM | Hibernate | Diesel, SeaORM |

## Tips for Success

### 1. Consistency Over Intensity
- Code every day, even if just for 30 minutes
- Small daily progress compounds over time

### 2. Document Your Journey
- Write notes about what you learned
- Keep track of challenges and solutions
- Build a portfolio of your work

### 3. Adapt to Your Language
- Use idiomatic patterns for your chosen language
- Leverage language-specific libraries and tools
- Follow community best practices

### 4. Don't Skip Days
- If you miss a day, don't give up
- Get back on track the next day
- Consistency builds habits

### 5. Share Your Progress
- Post updates on social media with #100DaysOfCode
- Write blog posts about your learnings
- Connect with others doing the challenge

## Project Ideas by Language

### Python
- Use matplotlib for data visualization
- Flask/Django for web projects
- Tkinter/PyQt for GUI applications

### JavaScript
- Use Chart.js for visualizations
- Express/Next.js for web projects
- Electron for desktop applications

### Java
- Use JFreeChart for visualizations
- Spring Boot for web projects
- JavaFX for GUI applications

### Go
- Use gonum/plot for visualizations
- Gin/Echo for web projects
- Fyne for GUI applications

### Rust
- Use plotters for visualizations
- Actix/Rocket for web projects
- egui/iced for GUI applications

### TypeScript
- Use Chart.js/D3.js for visualizations
- Next.js/NestJS for web projects
- Electron for desktop applications

### Bash
- Use gnuplot for visualizations
- CGI scripts for simple web servers
- Dialog/Whiptail for terminal UIs

## Resources

### General Programming
- [GitHub](https://github.com) - Version control and collaboration
- [Stack Overflow](https://stackoverflow.com) - Q&A community
- [MDN Web Docs](https://developer.mozilla.org) - Web development reference

### Language-Specific
- **Python**: [Python.org](https://python.org), [Real Python](https://realpython.com), [PyPI](https://pypi.org)
- **JavaScript**: [JavaScript.info](https://javascript.info), [MDN](https://developer.mozilla.org), [NPM](https://npmjs.com)
- **TypeScript**: [TypeScript Handbook](https://www.typescriptlang.org/docs/), [TypeScript Playground](https://www.typescriptlang.org/play)
- **Java**: [Oracle Docs](https://docs.oracle.com/javase/), [Baeldung](https://baeldung.com), [Maven Central](https://search.maven.org)
- **Go**: [Go.dev](https://go.dev), [Go by Example](https://gobyexample.com), [Awesome Go](https://awesome-go.com)
- **Bash**: [GNU Bash Manual](https://www.gnu.org/software/bash/manual/), [ShellCheck](https://www.shellcheck.net), [Bash Guide](https://mywiki.wooledge.org/BashGuide)
- **Rust**: [Rust Book](https://doc.rust-lang.org/book/), [Rust by Example](https://doc.rust-lang.org/rust-by-example/), [Crates.io](https://crates.io)
- **C#**: [Microsoft Docs](https://docs.microsoft.com/dotnet/csharp/), [C# Corner](https://www.c-sharpcorner.com)
- **Ruby**: [Ruby Docs](https://ruby-doc.org), [RubyGems](https://rubygems.org)

## FAQ

### Can I use multiple languages?
Yes! You can implement each day in multiple languages by creating separate folders (e.g., `python/001/`, `js/001/`, `go/001/` for Day 1).

### Do I have to start from Day 1?
While it's recommended to follow the progression, you can adapt the challenge to your skill level. Beginners should start at Day 1, while experienced developers might skip to Intermediate or Advanced sections.

### What if I miss a day?
Don't give up! Just continue from where you left off. The key is consistency over time, not perfection.

### Can I modify the requirements?
Absolutely! These are guidelines. Feel free to adapt projects to your interests or add your own twist.

### How do I share my progress?
- Post on social media with **#100DaysOfCode**
- Create a blog documenting your journey
- Share your GitHub repository
- Join the 100 Days of Code community

### What language should I choose?
Choose based on your goals:
- **Web Development**: JavaScript/TypeScript, Python
- **Systems Programming**: Rust, Go, C++
- **Mobile**: Swift, Kotlin, Dart
- **Enterprise**: Java, C#
- **Data Science**: Python, R
- **Learning**: Any language you want to master!

## Repository Organization

This repository supports a **multi-language, collaborative approach**:

```bash
# Single developer, single language
python/001/, python/002/, ... python/100/

# Single developer, multiple languages
python/001/, js/001/, go/001/ (Day 1 in 3 languages)
python/002/, js/002/, go/002/ (Day 2 in 3 languages)

# Multiple developers, different languages
python/001/  # Developer A using Python
js/001/      # Developer B using JavaScript
go/001/      # Developer C using Go
```

## Credits

This project structure is inspired by the [100 Days of Code: The Complete Python Pro Bootcamp](https://www.udemy.com/course/100-days-of-code/) course by Angela Yu, adapted to be language-agnostic and accessible to developers working with any programming language.

## License

This project is open source and available for educational purposes. Feel free to fork, modify, and use it for your own learning journey.

---

**Ready to start?** Check out `requirements/day001.md` and begin your 100-day journey today!

**#100DaysOfCode** | **#CodeNewbie** | **#LearnToCode**
