package termcolor

import (
	"strings"
	"testing"
)

func TestJoin(t *testing.T) {
	got := join([]style{})
	want := ""
	if got != want {
		t.Errorf("expected '%s' but got '%s'", want, got)
	}

	got = join([]style{Bold})
	want = "1"
	if got != want {
		t.Errorf("expected '%s' but got '%s'", want, got)
	}

	got = join([]style{Bold, ForegroundBlack})
	want = "1;30"
	if got != want {
		t.Errorf("expected '%s' but got '%s'", want, got)
	}
}

func TestActivate(t *testing.T) {
	got := Activate()
	want := startMarker + endMarker
	if got != want {
		t.Errorf("expected '%s' but got '%s'", want, got)
	}

	got = Activate(Bold)
	want = startMarker + string(Bold) + endMarker
	if got != want {
		t.Errorf("expected '%s' but got '%s'", want, got)
	}

}

func TestApplyStyles(t *testing.T) {
	got := ApplyStyles("foo")
	want := "foo"
	if got != want {
		t.Errorf("expected '%s' but got '%s'", want, got)
	}

	got = ApplyStyles("foo", Bold)
	want = startMarker + string(Bold) + endMarker + "foo" + startMarker + string(Reset) + endMarker
	if got != want {
		t.Errorf("expected '%s' but got '%s'", want, got)
	}
}

func TestPrinter(t *testing.T) {
	var buf strings.Builder
	p := New(&buf, false)

	p.Print("hello, ", Bold)
	p.Println("world!", ForegroundBlack)
	p.Printf("hello, %s!", p.Styled("world", ForegroundCyan), Bold)

	got := buf.String()
	want := "hello, world!\nhello, world!"

	if want != got {
		t.Errorf("expected '%s' but got '%s'", want, got)
	}

	buf.Reset()
	p = New(&buf, true)

	p.Print("hello, ", Bold)
	p.Println("world!", ForegroundBlack)
	p.Printf("hello, %s!", p.Styled("world", ForegroundCyan), Bold)

	got = buf.String()
	want = ApplyStyles("hello, ", Bold) + ApplyStyles("world!\n", ForegroundBlack) +
		ApplyStyles("hello, "+startMarker+string(ForegroundCyan)+endMarker+"world", Bold) +
		"!" + startMarker + string(Reset) + endMarker

	if want != got {
		t.Errorf("expected '%s' but got '%s'", want, got)
	}
}
