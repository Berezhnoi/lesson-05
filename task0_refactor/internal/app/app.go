package app

import (
	"errors"
	"fmt"
	"io"

	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/calculator"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/textanalyzer"
)

// Run executes the command workflow and writes its output to output.
func Run(output io.Writer) error {
	sum := calculator.Add(2, 3)
	if _, err := fmt.Fprintln(output, "2 + 3 =", sum); err != nil {
		return fmt.Errorf("write sum: %w", err)
	}

	quotient, err := calculator.Divide(10, 0)
	if err != nil {
		if errors.Is(err, calculator.ErrDivisionByZero) {
			if _, writeErr := fmt.Fprintln(output, "division by zero was correctly detected:", err); writeErr != nil {
				return fmt.Errorf("write division-by-zero result: %w", writeErr)
			}
		} else {
			return fmt.Errorf("calculator.Divide: %w", err)
		}
	} else if _, err := fmt.Fprintln(output, "10 / 0 =", quotient); err != nil {
		return fmt.Errorf("write quotient: %w", err)
	}

	words, err := textanalyzer.WordCount("the quick brown fox")
	if err != nil {
		return fmt.Errorf("textanalyzer.WordCount: %w", err)
	}
	if _, err := fmt.Fprintln(output, "word count:", words); err != nil {
		return fmt.Errorf("write word count: %w", err)
	}
	return nil
}
