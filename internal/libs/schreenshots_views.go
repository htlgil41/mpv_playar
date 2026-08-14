package libs

import (
	"fmt"

	"github.com/kbinani/screenshot"
)

func ScreenShotsView() int {
	n := screenshot.NumActiveDisplays()
	fmt.Printf("Número de monitores conectados: %d\n", n)

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		fmt.Printf("Monitor %d: resolución %dx%d en posición (%d, %d)\n",
			i, bounds.Dx(), bounds.Dy(), bounds.Min.X, bounds.Min.Y)
	}

	if n >= 2 {
		return 1
	}

	return 0
}
