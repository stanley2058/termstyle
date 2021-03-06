package foreground

import (
	"fmt"

	"github.com/stanley2058/termstyle/utils"
)

// Foreground colors
const (
	Black   string = "30"
	Red     string = "31"
	Green   string = "32"
	Yellow  string = "33"
	Blue    string = "34"
	Magenta string = "35"
	Cyan    string = "36"
	White   string = "37"

	BrightBlack   string = "90"
	BrightRed     string = "91"
	BrightGreen   string = "92"
	BrightYellow  string = "93"
	BrightBlue    string = "94"
	BrightMagenta string = "95"
	BrightCyan    string = "96"
	BrightWhite   string = "97"
)

// TrueColor construct terminal usable true color string, not all terminal support this.
func TrueColor(r uint8, g uint8, b uint8) string {
	return fmt.Sprintf("38;%d;%d;%d", r, g, b)
}

// HSL return converted true color string (0<= h <360, 0<= s,l <=1)
func HSL(h float64, s float64, l float64) string {
	return TrueColor(utils.HslToRgb(h, s, l))
}

// Fg256 256 colors for foreground
//   0-7:      standard colors (Black, Red, Green, Yellow, Blue, Magenta, Cyan, White)
//   8-15:     high intensity colors
//   16-231:   16 + 36×r + 6×g + b (0 ≤ r, g, b ≤ 5)
//   232-255:  grayscale from black to white in 24 steps
//   see: "https://jonasjacek.github.io/colors"
func Fg256(num uint8) string {
	return fmt.Sprintf("38;5;%d", num)
}
