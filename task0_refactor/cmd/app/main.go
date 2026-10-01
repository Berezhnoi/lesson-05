// Command app — тонкий main.go, що лише "з'єднує" (wiring) пакунки
// calculator та textanalyzer. Уся бізнес-логіка живе в internal/.
package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/calculator"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/textanalyzer"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	sum := calculator.Add(2, 3)
	if _, err := fmt.Println("2 + 3 =", sum); err != nil {
		return fmt.Errorf("write sum: %w", err)
	}

	quotient, err := calculator.Divide(10, 0)
	if err != nil {
		if errors.Is(err, calculator.ErrDivisionByZero) {
			if _, writeErr := fmt.Println("division by zero was correctly detected:", err); writeErr != nil {
				return fmt.Errorf("write division-by-zero result: %w", writeErr)
			}
		} else {
			return fmt.Errorf("calculator.Divide: %w", err)
		}
	} else {
		if _, err := fmt.Println("10 / 0 =", quotient); err != nil {
			return fmt.Errorf("write quotient: %w", err)
		}
	}

	words, err := textanalyzer.WordCount("the quick brown fox")
	if err != nil {
		return fmt.Errorf("textanalyzer.WordCount: %w", err)
	}
	if _, err := fmt.Println("word count:", words); err != nil {
		return fmt.Errorf("write word count: %w", err)
	}
	return nil
}
