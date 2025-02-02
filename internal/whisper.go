// Maintainer 2025 captions Pedro G. Branquinho
package transcribe

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var WhisperCmd = &Z.Cmd{
	Name:     `whisper`,
	Aliases:  []string{`third-party`, `w`},
	Usage:    `<command>`,
	Summary:  `Whisper-related service utilities; start whisper-binding API; download models; etc.`,
	NumArgs:  0,
	Commands: []*Z.Cmd{help.Cmd, WdownloadModelsCmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		return nil
	},
}

var WdownloadModelsCmd = &Z.Cmd{
	Name:     `download`,
	Aliases:  []string{`d`},
	Usage:    `<[]models>`,
	Summary:  `Download whisper models;`,
	NumArgs:  0,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		url := "http://localhost:8080/v1/models?stream=true"
		for _, model := range args {
			// Define the URL and JSON payload
			payload := map[string]string{
				"Path": fmt.Sprint(model),
			}

			// Marshal the payload into JSON
			jsonData, err := json.Marshal(payload)
			if err != nil {
				fmt.Println("Error marshaling JSON:", err)
				return err
			}

			// Create a new HTTP POST request
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
			if err != nil {
				fmt.Println("Error creating request:", err)
				return err
			}

			// Set the Content-Type header
			req.Header.Set("Content-Type", "application/json")

			// Send the request using the default HTTP client
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return fmt.Errorf("error sending request for model %s: %v", model, err)
			}
			defer resp.Body.Close()

			// Check the response status code
			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("request for model %s failed with status code: %d", model, resp.StatusCode)
			}

			// Handle streaming response
			fmt.Printf("Streaming response for model %s:\n", model)
			reader := bufio.NewReader(resp.Body)
			for {
				// Read a line from the stream
				line, err := reader.ReadBytes('\n')
				if err != nil {
					if err == io.EOF {
						break // End of stream
					}
					return fmt.Errorf("error reading stream for model %s: %v", model, err)
				}

				// Process the line (e.g., print it or decode it as JSON)
				fmt.Print(string(line))
			}

			fmt.Printf("Finished streaming for model: %s\n", model)
		}

		return nil
	},
}
