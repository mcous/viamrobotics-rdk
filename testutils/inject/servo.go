package inject

import (
	"context"

	"go.viam.com/core/component/servo"
)

// Servo is an injected servo.
type Servo struct {
	servo.Servo
	MoveFunc    func(ctx context.Context, angle uint8) error
	CurrentFunc func(ctx context.Context) (uint8, error)
}

// Move calls the injected Move or the real version.
func (s *Servo) Move(ctx context.Context, angle uint8) error {
	if s.MoveFunc == nil {
		return s.Servo.Move(ctx, angle)
	}
	return s.MoveFunc(ctx, angle)
}

// Current calls the injected Current or the real version.
func (s *Servo) Current(ctx context.Context) (uint8, error) {
	if s.CurrentFunc == nil {
		return s.Servo.Current(ctx)
	}
	return s.CurrentFunc(ctx)
}
