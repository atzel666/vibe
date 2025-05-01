# Vibe CLI

Welcome aboard! 🚀 Vibe is a Go-based CLI tool that leverages Google’s Gemini AI API to generate Protocol Buffer (`.proto`) files and Prisma schema (`.prisma`) files based on your specifications. Whether you’re architecting gRPC services, setting up a modern ORM layer, or exploring AI-driven code generation, Vibe has got your back.

## 🌟 Features

- **Flexible generation**: Create `.proto`, `schema.prisma`, or both in one go.
- **Customizable models**: Pick your favorite Gemini model (default: `gemini-2.0-flash`).
- **Smart extraction**: Automatically strips AI chat markdown, giving you clean code output.
- **Easy integration**: Drop it into your CI/CD pipeline or run locally.

## 🚀 Getting Started

### Prerequisites

- Go **1.18+** installed on your machine.
- A valid **Google API Key** with access to Gemini API (set as `GOOGLE_API_KEY`).

### Installation

1. **Install**
   ```bash
   go install github.com/Raezil/vibe
   ```

## 🛠️ Usage

```bash
vibe \
  -prompt "<YOUR SPECIFICATION>" \
  -generator <proto|prisma|both> \
  -o example.proto \
  -schema schema.prisma \
  -model gemini-2.0-flash
```

| Flag        | Description                                                        | Default             |
|-------------|--------------------------------------------------------------------|---------------------|
| `-prompt`   | **(Required)** Input specification for generation                  |                     |
| `-generator`| Type of file to generate (`proto`, `prisma`, or `both`)           | `proto`             |
| `-o`        | Output filename for `.proto`                                       | `example.proto`     |
| `-schema`   | Output filename for Prisma schema                                  | `schema.prisma`     |
| `-model`    | Gemini AI model to use                                             | `gemini-2.0-flash`  |

> **Tip:** If you forget the `-prompt` flag, Vibe will prompt you to include it—no more guesswork!

## 📦 Examples

- **Generate only a `.proto` file**
  ```bash
  vibe -prompt "Create blog where users create posts and comments" -generator proto
  ```
- **Generate only a Prisma schema**
  ```bash
  vibe -prompt "Create blog where users create posts and comments" -generator prisma
  ```
- **Generate both files**
  ```bash
  vibe -prompt "Create blog where users create posts and comments" -generator both
  ```

## 🔍 How It Works

1. **Flag parsing**: Reads your CLI flags and validates them.
2. **GenAI client setup**: Connects to Gemini AI using your `GOOGLE_API_KEY`.
3. **Prompt crafting**: Builds a tailored prompt for `.proto` and/or Prisma generation.
4. **AI chat session**: Sends the prompt to Gemini, receives AI-generated content.
5. **Code extraction**: Strips away Markdown fences to isolate pure code.
6. **File writing**: Saves the output to your specified filenames.

## 🤔 Troubleshooting

- **Missing `GOOGLE_API_KEY`**: Make sure the environment variable is set:
  ```bash
  export GOOGLE_API_KEY=YOUR_KEY_HERE
  ```
- **Invalid generator type**: Use one of `proto`, `prisma`, or `both`.
- **Go version issues**: Ensure you’re running Go 1.18 or newer.

---

Happy coding! 🎉 If you hit any snags or have feature requests, feel free to open an issue or submit a PR on the Vibe repo. Let’s build the future of AI-assisted development together!

