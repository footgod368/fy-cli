package root

import (
	"context"
	"fmt"
	youdao "github.com/footgod368/translator-sdk"
	"github.com/footgod368/translator-sdk/utils"
	"strings"
)

func RunE(args []string) error {
	ctx := context.Background()
	resp, err := youdao.Query(ctx, strings.Join(args, " "))
	if err != nil {
		return err
	}
	fmt.Println(utils.MarshalJSONToString(resp))
	return nil
}
