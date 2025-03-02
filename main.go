package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/YashSaini99/NLP_CLI/command"
	"github.com/YashSaini99/NLP_CLI/config"
	"github.com/YashSaini99/NLP_CLI/executor"
	"github.com/YashSaini99/NLP_CLI/utils"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file.
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Continuing with environment variables.")
	}

	// Initialize logger.
	utils.InitializeLogger()
	log.Println("Application starting (CLI Mode) with Google Gemini...")

	// Handle OS interrupts for graceful shutdown.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Received termination signal, shutting down gracefully...")
		os.Exit(0)
	}()

	// Load configuration.
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	// Define command-line flags.
	instructionPtr := flag.String("instruction", "", "Natural language instruction to translate to a shell command")
	dryRunPtr := flag.Bool("dry-run", false, "Perform a dry-run of the generated command without execution")
	verbosePtr := flag.Bool("verbose", false, "Enable verbose output")
	flag.Parse()

	var instruction string
	if *instructionPtr != "" {
		instruction = *instructionPtr
	} else {
		// Interactive mode: prompt for input.
		fmt.Print("Enter your instruction: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read input: %v", err)
		}
		instruction = strings.TrimSpace(input)
	}
	if instruction == "" {
		log.Fatalf("No instruction provided")
	}

	// Translate the instruction to a shell command using Google Gemini.
	shellCommand, err := command.TranslateInstruction(cfg, instruction)
	if err != nil {
		log.Fatalf("Error translating instruction: %v", err)
	}
	if *verbosePtr {
		fmt.Printf("Generated Shell Command: %s\n", shellCommand)
		log.Printf("Generated command: %s", shellCommand)
	} else {
		fmt.Printf("Proposed Shell Command: %s\n", shellCommand)
	}

	// If not already in dry-run mode via flag, ask for confirmation.
	dryRun := *dryRunPtr
	if !dryRun {
		fmt.Print("Do you want to execute this command? (yes/no/dry-run): ")
		reader := bufio.NewReader(os.Stdin)
		confirmation, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read confirmation: %v", err)
		}
		confirmation = strings.ToLower(strings.TrimSpace(confirmation))
		if confirmation == "no" {
			fmt.Println("Command execution aborted.")
			log.Println("User aborted execution.")
			os.Exit(0)
		} else if confirmation == "dry-run" {
			dryRun = true
		}
	}

	// Execute the command.
	output, err := executor.ExecuteCommand(shellCommand, dryRun)
	if err != nil {
		log.Printf("Error executing command: %v", err)
	}
	fmt.Printf("Command Output:\n%s\n", output)
	log.Println("Command executed. Exiting.")
}
