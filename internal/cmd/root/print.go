package root

import (
	"fmt"
	"github.com/fatih/color"
	youdao "github.com/footgod368/translator-sdk"
	"strings"
)

func printResult(sourceText string, resp *youdao.TranslateResponse) {
	fmt.Println(sourceText)
	if resp.Ukphone != "" {
		fmt.Printf("英音：[%s]\n", resp.Usphone)
	}
	if resp.Usphone != "" {
		fmt.Printf("美音：[%s]\n", resp.Ukphone)
	}
	if len(resp.WebTranslations) > 0 {
		fmt.Printf("网络释义: %s\n", strings.Join(resp.WebTranslations, "；"))
	}
	for _, translation := range resp.Translations {
		color.Red(translation)
	}
	if len(resp.WordForms) > 0 {
		fmt.Printf("(%s)\n", strings.Join(resp.WordForms, "，"))
	}
	if len(resp.Etymologies) > 0 {
		color.Green("词源：")
		for i, etymology := range resp.Etymologies {
			color.Green("%d. %s\n%s\n", i+1, etymology.Desc, etymology.Value)
		}
	}
	if len(resp.EgSentences) > 0 {
		fmt.Println("例句：")
		for i, egSentence := range resp.EgSentences {
			fmt.Printf("%d. %s\n%s\n", i+1, egSentence.Sentence, egSentence.Translation)
		}
	}
	if Discriminate && len(resp.Discrimination) > 0 {
		color.Blue("辨析：")
		for i, discriminateUsage := range resp.Discrimination {
			color.Blue("%d. %s\n%s\n", i+1, discriminateUsage.Headword, discriminateUsage.Usage)
		}
	}
}
