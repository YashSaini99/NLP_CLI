package command

import (
	"log"
	"strings"

	"github.com/yourusername/go-nlp-task-runner/config"
	"github.com/yourusername/go-nlp-task-runner/nlp"
	"github.com/yourusername/go-nlp-task-runner/validator"
)

// extractCommand removes markdown code blocks from the NLP response.
func extractCommand(response string) string {
	if strings.Contains(response, "```bash") {
		start := strings.Index(response, "```bash") + len("```bash")
		end := strings.Index(response[start:], "```")
		if end != -1 {
			return strings.TrimSpace(response[start : start+end])
		}
	}
	return strings.TrimSpace(response) // Fallback: use the whole response if no code block is found.
}

// TranslateInstruction processes the instruction and prepares it for execution.
func TranslateInstruction(cfg *config.Config, instruction string) (string, error) {
	shellCommand, err := nlp.ParseInstruction(cfg.GeminiAPIKey, cfg.GeminiModel, instruction)
	if err != nil {
		return "", err
	}
	// Log the raw command for debugging.
	log.Printf("Raw generated command: %q", shellCommand)
	// Extract the command from markdown.
	shellCommand = extractCommand(shellCommand)
	// Log the extracted command.
	log.Printf("Extracted command: %q", shellCommand)
	// Sanitize: remove "Command:" prefix and trim whitespace.
	shellCommand = strings.TrimPrefix(shellCommand, "Command:")
	shellCommand = strings.TrimSpace(shellCommand)
	// Log the sanitized command.
	log.Printf("Sanitized command for validation: %q", shellCommand)
	// Validate against the whitelist.
	if err := validator.ValidateCommand(shellCommand); err != nil {
		return "", err
	}
	return shellCommand, nil
}
