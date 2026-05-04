//go:build !js

package main

import (
	"flag"
	"fmt"
)

func main() {
	lines := flag.Int("lines", 4, "回線数")
	minGB := flag.Int("min", 25, "最低容量(GB)")
	maxGB := flag.Int("max", 45, "最高容量(GB)")
	all := flag.Bool("all", false, "min〜maxを1GB単位で一覧表示")
	flag.Parse()

	if *all {
		fmt.Printf("%-8s %-14s %-12s %s\n", "合計GB", "最安値(割引後)", "単価(/GB)", "組み合わせ")
		fmt.Println("---------------------------------------------------------------")
		for gb := *minGB; gb <= *maxGB; gb++ {
			result := findCheapestCombos(*lines, gb, gb)
			if result.BestCost == -1 {
				fmt.Printf("%-8s %-14s\n", fmt.Sprintf("%dGB", gb), "組み合わせなし")
				continue
			}
			perGB := float64(result.FinalCost) / float64(gb)
			for i, combo := range result.Combos {
				label := ""
				for idx, p := range combo {
					if idx > 0 {
						label += " + "
					}
					label += fmt.Sprintf("%dGB(¥%d)", p.GB, p.Price)
				}
				if i == 0 {
					fmt.Printf("%-8s ¥%-13d ¥%-11.1f %s\n", fmt.Sprintf("%dGB", gb), result.FinalCost, perGB, label)
				} else {
					fmt.Printf("%-8s %-14s %-12s %s\n", "", "", "", label)
				}
			}
		}
		return
	}

	result := findCheapestCombos(*lines, *minGB, *maxGB)
	if result.BestCost == -1 {
		fmt.Println("該当する組み合わせがありません")
		return
	}

	fmt.Printf("割引: ¥%d/月 (%d回線 × ¥%d)\n", result.Discount, *lines, discountPerLine)
	fmt.Printf("最安値(割引後): ¥%d/月\n\n", result.FinalCost)
	for _, combo := range result.Combos {
		totalGB := 0
		for _, p := range combo {
			totalGB += p.GB
		}
		pricePerGB := float64(result.FinalCost) / float64(totalGB)
		fmt.Printf("合計 %dGB (¥%.1f/GB): ", totalGB, pricePerGB)
		for idx, p := range combo {
			if idx > 0 {
				fmt.Print(" + ")
			}
			fmt.Printf("%dGB(¥%d)", p.GB, p.Price)
		}
		fmt.Println()
	}
}
