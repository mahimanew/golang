- install hasicorp API in your project folder
```
go get github.com/hashicorp/vault/api
```
- main.go
```


package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/vault/api"
)

func main() {
	config := api.DefaultConfig()
	config.Address = "http://127.0.0.1:8200" // Replace with your actual Vault server URL

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Vault client: %v", err)
	}

	client.SetToken("xxxxxxxxxxxxxxxxxxxxxxxxxx") // Replace with your Vault token

	// Define the API path to your secret
	secretPath := "kv/data/Keys"

	// Retrieve the secret from Vault
	secret, err := client.Logical().Read(secretPath)
	if err != nil {
		log.Fatalf("Failed to read secret: %v", err)
	}

	// Extract the WebKey value from the nested data
	if secret != nil {
		if data, ok := secret.Data["data"].(map[string]interface{}); ok {
			if webKey, ok := data["WebKey"].(string); ok {
				fmt.Printf("Retrieved WebKey: %s\n", webKey)
			} else {
				log.Println("WebKey not found or is not a string in the data field")
			}
		} else {
			log.Println("Data field missing or in unexpected format")
		}
	} else {
		log.Println("No secret found at specified path")
	}
}

```

```
go run main.go

```
