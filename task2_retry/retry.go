// Package retry містить завдання 2 домашньої роботи: retry-обгортку
// навколо "нестабільної" операції, що використовує sentinel error
// ErrTemporary для розпізнавання відновлюваних відмов.
package retry

import (
	"errors"
	"fmt"
	"time"
)

// ErrTemporary — sentinel error, що позначає тимчасову (відновлювану) відмову
// операції. Обгортайте ErrTemporary через %w у своїй "нестабільній" функції,
// коли відмова є тимчасовою і варта повторної спроби.
var ErrTemporary = errors.New("retry: temporary failure")

// ErrNilOperation is returned when Do receives no operation to execute.
var ErrNilOperation = errors.New("retry: nil operation")

// Operation — довільна операція, що може повернути помилку.
type Operation func() (string, error)

// Do виконує op до maxAttempts разів. Між невдалими спробами робить паузу
// тривалістю backoff. Повторює спробу лише тоді, коли помилка є тимчасовою
// (errors.Is(err, ErrTemporary)) — для інших помилок Do має одразу
// повернути обгорнуту помилку без повторних спроб.
//
// Якщо всі спроби вичерпано, Do повертає чітко обгорнуту фінальну помилку
// (з інформацією про кількість спроб).
//
// TODO(Завдання 2): реалізуйте функцію Do.
// Вимоги:
//   - maxAttempts має бути >= 1; якщо операція вдається одразу — повторів немає
//   - між спробами (окрім останньої) чекайте backoff перед наступною спробою
//   - якщо помилка НЕ є errors.Is(err, ErrTemporary) — не повторюйте спробу,
//     одразу поверніть обгорнуту помилку
//   - якщо всі спроби вичерпано — поверніть обгорнуту фінальну помилку,
//     яка через errors.Is все ще розпізнається як ErrTemporary
func Do(op Operation, maxAttempts int, backoff time.Duration) (string, error) {
	if op == nil {
		return "", ErrNilOperation
	}
	if maxAttempts < 1 {
		return "", fmt.Errorf("maxAttempts must be >= 1, got %d", maxAttempts)
	}

	var lastErr error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		result, err := op()
		if err == nil {
			return result, nil // успішне виконання
		}

		if !errors.Is(err, ErrTemporary) {
			return "", fmt.Errorf("operation failed on attempt %d: %w", attempt, err)
		}

		lastErr = fmt.Errorf("operation failed on attempt %d: %w", attempt, err)

		if attempt < maxAttempts {
			time.Sleep(backoff) // пауза перед наступною спробою
		}
	}

	return "", fmt.Errorf("all %d attempts failed: %w", maxAttempts, lastErr)
}

// NewFlakyOperation — допоміжна функція для тестів/демонстрації: повертає
// Operation, що падає з ErrTemporary перші failuresBeforeSuccess разів,
// а потім повертає successValue без помилки.
func NewFlakyOperation(failuresBeforeSuccess int, successValue string) Operation {
	attempt := 0
	return func() (string, error) {
		attempt++
		if attempt <= failuresBeforeSuccess {
			return "", fmt.Errorf("flaky op: attempt %d failed: %w", attempt, ErrTemporary)
		}
		return successValue, nil
	}
}
