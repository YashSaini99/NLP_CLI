# NLP CLI Task Runner

A command-line tool that translates natural language instructions into safe shell commands using advanced NLP. This project integrates with the Google Gemini API to process natural language input, then validates and executes (or simulates) the resulting shell commands.

## Features

- **Natural Language Processing:**  
  Converts user instructions into shell commands using a generative AI API.

- **Command Validation:**  
  Validates generated commands against a whitelist to ensure safe execution.

- **CLI-Centric Interface:**  
  Fully command-line based, with interactive prompts and command-line flags.

- **Configurable Storage & Environment:**  
  Uses environment variables (loaded from a `.env` file) to store API keys and project settings.

- **Modular Design:**  
  Organized into modules for configuration, NLP processing, command translation, validation, and command execution.

## Prerequisites

- **Go:** Version 1.20 or later  
- **Docker (optional):** For containerized deployment
- **Environment:** Linux (or any system that supports Go)

## Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/YashSaini99/NLP_CLI.git
   cd NLP_CLI
   ```

2. **Set Up Environment Variables:**

   Create a `.env` file in the project root with the following content (replace the placeholder values):

   ```ini
   GOOGLE_GEMINI_API_KEY=your_google_gemini_api_key_here
   GOOGLE_GEMINI_MODEL=gemini-model-default
   ```

3. **Download Dependencies:**

   ```bash
   go mod download
   ```

4. **Build the Project:**

   ```bash
   go build -o nlp-task-runner
   ```

## Usage

Run the application using one of the following methods:

### Interactive Mode

Simply run the binary without any flags:

```bash
./nlp-task-runner
```

You will be prompted to enter a natural language instruction. For example:

```
Enter your instruction: archive all logs modified yesterday and compress them
```

### Command-Line Flags

You can also pass the instruction directly as a flag and enable verbose output:

```bash
./nlp-task-runner -instruction="archive all logs modified yesterday and compress them" -verbose
```

### Dry-Run Mode

To simulate the command execution without making any system changes, you can use the `-dry-run` flag:

```bash
./nlp-task-runner -instruction="archive all logs modified yesterday and compress them" -dry-run
```

## Troubleshooting

- **Configuration Issues:**  
  Ensure your `.env` file is in the project root and contains valid values.

- **API Errors:**  
  If you encounter API errors, verify your API key and model configuration. Refer to the official documentation.

- **Command Validation:**  
  The generated shell command is validated against a whitelist in `validator/validator.go`. If a valid command is being rejected, adjust the whitelist regex patterns accordingly.
