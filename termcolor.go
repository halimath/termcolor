// termcolor defines types and functions to output colorized console output.
//
// termcolor provides a Printer that wraps an io.Writer to print colorized messages. It supports a guard that
// suppresses all color information if the Printer's underlying io.Writer is no TTY. Thus, applications can
// use a single, colorized output API but have standard text being printed if outout is redirected.
package termcolor

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// style defines a type to represent style tags.
type style string

const (
	// startMarker is the ANSI terminal escape sequence to start color instructions.
	startMarker = "\033["
	// endMarker terminates a ANSI terminal escape sequence.
	endMarker = "m"
	// separator separates style instructions.
	separator = ';'

	// DisableCursor = "\033[?25l"
	// EnableCursor  = "\033[?25h"

	// Reset rests all styles to their default.
	Reset style = "0"

	// Bold renders text in bold (somewhat brighter).
	Bold style = "1"

	// Default applies the default styles.
	Default style = "22"

	// Foreground color black
	ForegroundBlack style = "30"
	// Foreground color red
	ForegroundRed style = "31"
	// Foreground color green
	ForegroundGreen style = "32"
	// Foreground color yellow
	ForegroundYellow style = "33"
	// Foreground color blue
	ForegroundBlue style = "34"
	// Foreground color magenta
	ForegroundMagenta style = "35"
	// Foreground color cyan
	ForegroundCyan style = "36"
	// Foreground color white
	ForegroundWhite style = "37"

	// Background color back
	BackgroundBlack style = "40"
	// Background color red
	BackgroundRed style = "41"
	// Background color green
	BackgroundGreen style = "42"
	// Background color yellow
	BackgroundYellow style = "43"
	// Background color blue
	BackgroundBlue style = "44"
	// Background color magenta
	BackgroundMagenta style = "45"
	// Background color cyan
	BackgroundCyan style = "46"
	// Background color white
	BackgroundWhite style = "47"
)

// Printer defines a type that provides methods to produce colorized output. Output is written to w if isTTY
// is set to true, otherwise output is written unstyled.
type Printer struct {
	w     io.Writer
	isTTY bool
}

// New creates a new Printer writing to w and applying coloring if tty is set to true.
func New(w io.Writer, tty bool) *Printer {
	return &Printer{
		w:     w,
		isTTY: tty,
	}
}

// NewFromFile creates a new Printer writing to f and applying colors if f is character device.
func NewFromFile(f *os.File) *Printer {
	return New(f, isTerminal(f))
}

// Stdout creates a Printer for os.Stdout applying colors if stout is a console device.
func Stdout() *Printer {
	return NewFromFile(os.Stdout)
}

// Stderr creates a Printer for os.Stderr applying colors if stout is a console device.
func Stderr() *Printer {
	return NewFromFile(os.Stderr)
}

// Printf prints formatted and styled output to p's underlying io.Writer. It works almost the same as
// fmt.Printf with the exception that any value given in argsAndStyles that is of type style will be excluded
// from formatting and will be used to style the overall string in stead. So
//
//	p.Printf("hello, %s", "world", termcolor.ForegroundRed)
//
// will print "hello, world" in red color, if p is a TTY.
//
// It returns any error returned from writing to p's io.Writer.
func (p *Printer) Printf(format string, argsAndStyles ...interface{}) error {
	args := make([]interface{}, 0, len(argsAndStyles))
	styles := make([]style, 0, len(argsAndStyles))

	for _, a := range argsAndStyles {
		if s, ok := a.(style); ok {
			styles = append(styles, s)
		} else {
			args = append(args, a)
		}
	}

	return p.Print(fmt.Sprintf(format, args...), styles...)
}

// Print prints msg with all styles applied. It returns any error returned from writing to p's io.Writer.
func (p *Printer) Print(msg string, styles ...style) error {
	if !p.isTTY {
		_, err := fmt.Fprint(p.w, msg)
		return err
	}

	_, err := fmt.Fprint(p.w, ApplyStyles(msg, styles...))
	return err
}

// Println prints msg with all styles applied followed by a single newline. It returns any error returned from
// writing to p's io.Writer.
func (p *Printer) Println(msg string, styles ...style) error {
	return p.Print(msg+"\n", styles...)
}

// Styled applies styles to s and returns that string if p's underlying io.Writer is a TTY. Otherwise s is
// returned unchanged.
func (p *Printer) Styled(s string, styles ...style) string {
	if !p.isTTY {
		return s
	}

	return ApplyStyles(s, styles...)
}

// ApplyStyles applies all styles to s and returns the resulting string.
func ApplyStyles(s string, styles ...style) string {
	if len(styles) == 0 {
		return s
	}
	return Activate(styles...) + s + Activate(Reset)
}

// Activate returns a string that activates styles.
func Activate(styles ...style) string {
	return startMarker + join(styles) + endMarker
}

// join joins styles using the defined separator and returns them as a single string.
func join(styles []style) string {
	if len(styles) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString(string(styles[0]))

	for i := range styles[1:] {
		b.WriteRune(separator)
		b.WriteString(string(styles[i+1]))
	}

	return b.String()
}

// isTerminal checks whether f is a character device or not.
func isTerminal(f *os.File) bool {
	if fileInfo, _ := f.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		return true
	}
	return false
}
