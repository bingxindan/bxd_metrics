package Parse

import (
	"context"
	"testing"
)

func TestTomlParse(t *testing.T) {
	parse := new(Parse)
	parse.ParseToml(context.Background())
}
