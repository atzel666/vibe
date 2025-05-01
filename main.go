package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"google.golang.org/genai"
)

func main() {
	// 1) CLI flags
	prompt := flag.String("prompt", "", "Input specification for generation")
	generator := flag.String("generator", "proto", "Type of file to generate: proto, prisma, or both")
	protoOut := flag.String("o", "example.proto", "Output .proto file name")
	schemaOut := flag.String("schema", "schema.prisma", "Output Prisma schema file name")
	model := flag.String("model", "gemini-2.0-flash", "Gemini model to use")
	flag.Parse()

	if *prompt == "" {
		log.Fatal("Please provide your prompt using -prompt flag")
	}

	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		log.Fatal("GOOGLE_API_KEY environment variable is not set")
	}

	// 2) Create GenAI client
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatalf("Failed to create GenAI client: %v", err)
	}

	// Helper to extract code from markdown
	extract := func(text string) string {
		re := regexp.MustCompile("(?s)```(?:[a-z]+)?\\s*\\n(.*?)```")
		m := re.FindStringSubmatch(text)
		if len(m) >= 2 {
			return strings.TrimSpace(m[1])
		}
		return strings.TrimSpace(text)
	}

	// Function to invoke chat and write file
	run := func(desc, promptText, outFile string) {
		chat, err := client.Chats.Create(ctx, *model, nil, nil)
		if err != nil {
			log.Fatalf("Failed to create chat session for %s: %v", desc, err)
		}
		resp, err := chat.SendMessage(ctx, genai.Part{Text: promptText})
		if err != nil {
			log.Fatalf("Failed to send message for %s: %v", desc, err)
		}
		content := extract(resp.Text())
		if err := os.WriteFile(outFile, []byte(content), 0644); err != nil {
			log.Fatalf("Failed to write %s file %s: %v", desc, outFile, err)
		}
		fmt.Printf("%s saved to %s\n", desc, outFile)
	}

	// Generate based on generator flag
	switch strings.ToLower(*generator) {
	case "proto":
		protoPrompt := fmt.Sprintf(
			"You are a Go developer specializing in Protocol Buffers and gRPC-Gateway.\n"+
				"Generate a complete .proto file with syntax declaration, package, imports (including google/api/annotations.proto), message and service definitions with HTTP annotations.\n"+
				"RETURN ONLY THE .PROTO CONTENT.\n\nInput specification:\n%s\n", *prompt)
		run("Protobuf", protoPrompt, *protoOut)

	case "prisma":
		schemaPrompt := fmt.Sprintf(
			"You are a backend developer specializing in Prisma ORM.\n"+
				"Generate a complete Prisma schema file (schema.prisma) defining models, fields, relations, and generator/config blocks.\n"+
				"Include exactly this generator block at the top of the schema:\n```prisma\ngenerator client {\n  provider = \"go run github.com/steebchen/prisma-client-go\"\n}\n```\n"+
				"RETURN ONLY THE PRISMA SCHEMA CONTENT.\n\nInput specification:\n%s\n", *prompt)
		run("Prisma schema", schemaPrompt, *schemaOut)

	case "both":
		// Protobuf
		protoPrompt := fmt.Sprintf(
			"You are a Go developer specializing in Protocol Buffers and gRPC-Gateway.\n"+
				"Generate a complete .proto file with syntax declaration, package, imports (including google/api/annotations.proto), message and service definitions with HTTP annotations.\n"+
				"RETURN ONLY THE .PROTO CONTENT.\n\nInput specification:\n%s\n", *prompt)
		run("Protobuf", protoPrompt, *protoOut)

		// Prisma
		schemaPrompt := fmt.Sprintf(
			"You are a backend developer specializing in Prisma ORM.\n"+
				"Generate a complete Prisma schema file (schema.prisma) defining models, fields, relations, and generator/config blocks.\n"+
				"Include exactly this generator block at the top of the schema:\n```prisma\ngenerator client {\n  provider = \"go run github.com/steebchen/prisma-client-go\"\n}\n```\n"+
				"RETURN ONLY THE PRISMA SCHEMA CONTENT.\n\nInput specification:\n%s\n", *prompt)
		run("Prisma schema", schemaPrompt, *schemaOut)

	default:
		log.Fatalf("Unknown generator type: %s. Use 'proto', 'prisma' or 'both'", *generator)
	}
}
