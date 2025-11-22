package root

import (
	"context"
	youdao "github.com/footgod368/translator-sdk"
	"strings"
)

func RunE(args []string) error {
	ctx := context.Background()
	sourceText := strings.Join(args, " ")
	resp, err := youdao.Query(ctx, sourceText)
	if err != nil {
		return err
	}
	printResult(sourceText, resp)
	return nil
}
