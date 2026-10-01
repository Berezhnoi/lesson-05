// Command app — тонкий main.go, що лише "з'єднує" (wiring) пакунки
// calculator та textanalyzer. Уся бізнес-логіка живе в internal/.
package main

import (
	"log"
	"os"

	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/app"
)

func main() {
	if err := app.Run(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
