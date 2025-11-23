package root

import (
	"context"
	"github.com/footgod368/fy-cli/internal/cmd/common"
	"github.com/footgod368/fy-cli/utils"
	youdao "github.com/footgod368/translator-sdk"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strings"
)

func Complete(toComplete string) ([]string, cobra.ShellCompDirective) {
	outputs, _ := youdao.Suggest(toComplete, 10)
	return outputs, cobra.ShellCompDirectiveNoFileComp
}

func PreRunE(args []string) error {
	if common.Verbose {
		utils.Logger.SetLevel(logrus.DebugLevel)
	}
	return nil
}

func RunE(args []string) error {
	ctx := context.Background()
	sourceText := strings.Join(args, " ")
	resp, err := youdao.Query(ctx, sourceText)
	if err != nil {
		return err
	}
	if common.Verbose {
		utils.Logger.Debugf("youdao resp is: %s", utils.MarshalJSONToString(resp))
	}
	printResult(sourceText, resp)
	return nil
}
