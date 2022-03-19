package opt

import (
	"testing"
)

func TestParsedDashDash(t *testing.T) {
	parser := New("")
	parser.Parse([]string{"--"})
	got := parser.parsedDashDash
	want := true
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func Test1(t *testing.T) {
	parser := New("")
	want := true
	verbose := parser.Flag("verbose", "v", "be verbose")
	parser.Parse([]string{"-v"})
	got := *verbose
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func Test2(t *testing.T) {
	parser := New("")
	want := true
	verbose := parser.Flag("verbose", "v", "be verbose")
	parser.Parse([]string{"--verbose"})
	got := *verbose
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func Test3(t *testing.T) {
	parser := New("")
	want := 69
	number := parser.Option(0, "number", "i", "a number", "<number>", false)
	parser.Parse([]string{"--number=69"})
	got := *number
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func Test4(t *testing.T) {
	parser := New("")
	want := 69
	number := parser.Option(0, "number", "i", "a number", "<number>", false)
	parser.Parse([]string{"--number", "69"})
	got := *number
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func Test5(t *testing.T) {
	parser := New("")
	want1 := 69
	want2 := true
	number := parser.Option(0, "number", "i", "a number", "<number>", false)
	verbose := parser.Flag("verbose", "v", "be verbose")
	parser.Parse([]string{"--number", "69", "-v"})
	got1 := *number
	got2 := *verbose
	if got1 != want1 {
		t.Errorf("got %v; want %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("got %v; want %v", got2, want2)
	}
}

func Test6(t *testing.T) {
	parser := New("")
	want := "word"
	text := parser.Option("", "text", "t", "a text", "<string>", false)
	parser.Parse([]string{"--text=word"})
	got := *text
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func Test7(t *testing.T) {
	parser := New("")
	want := "word"
	text := parser.Option("", "text", "t", "a text", "<string>", false)
	parser.Parse([]string{"--text", "word"})
	got := *text
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func Test8(t *testing.T) {
	parser := New("")
	want1 := 69
	want2 := true
	want3 := "word"
	number := parser.Option(0, "number", "i", "a number", "<number>", false)
	verbose := parser.Flag("verbose", "v", "be verbose")
	text := parser.Option("", "text", "t", "a text", "<string>", false)
	parser.Parse([]string{"--number", "69", "-v", "--text=word"})
	got1 := *number
	got2 := *verbose
	got3 := *text
	if got1 != want1 {
		t.Errorf("got %v; want %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("got %v; want %v", got2, want2)
	}
	if got3 != want3 {
		t.Errorf("got %v; want %v", got3, want3)
	}
}

func Test9(t *testing.T) {
	parser := New("")
	want1 := true
	want2 := true
	optionA := parser.Flag("optionA", "a", "")
	optionB := parser.Flag("optionB", "b", "")
	parser.Parse([]string{"-ab"})
	got1 := *optionA
	got2 := *optionB
	if got1 != want1 {
		t.Errorf("got %v; want %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("got %v; want %v", got2, want2)
	}
}

func Test10(t *testing.T) {
	parser := New("")
	want1 := true
	want2 := true
	want3 := true
	a := parser.Flag("aa", "a", "")
	b := parser.Flag("aa", "b", "")
	c := parser.Flag("bb", "c", "")
	parser.Parse([]string{"-abc"})
	got1 := *a
	got2 := *b
	got3 := *c
	if got1 != want1 {
		t.Errorf("got %v; want %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("got %v; want %v", got2, want2)
	}
	if got3 != want3 {
		t.Errorf("got %v; want %v", got3, want3)
	}
}

func Test11(t *testing.T) {
	parser := New("")
	want1 := "word"
	want2 := "word"
	want3 := "word"
	word := parser.Option("hello", "word", "", "a text", "<string>", false)
	text := parser.Option("hello", "text", "", "a text", "<string>", false)
	verb := parser.Option("hello", "", "v", "a text", "<string>", false)
	parser.Parse([]string{"--word", "word", "--text", "word", "-v=word"})
	got1 := *word
	got2 := *text
	got3 := *verb
	if got1 != want1 {
		t.Errorf("got %v; want %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("got %v; want %v", got2, want2)
	}
	if got3 != want3 {
		t.Errorf("got %v; want %v", got3, want3)
	}
}

func Test12(t *testing.T) {
	parser := New("")
	want1 := 1
	want2 := 2
	want3 := 3
	want4 := "str1"
	want5 := "str2"
	want6 := "str3"
	want7 := true
	want8 := true
	want9 := true
	want10 := true
	number1 := parser.Option(0, "number", "", "", "", false)
	number2 := parser.Option(0, "", "b", "", "", false)
	number3 := parser.Option(0, "", "c", "", "", false)
	str1 := parser.Option("empty", "str1", "", "", "", false)
	str2 := parser.Option("empty", "str2", "", "", "", false)
	str3 := parser.Option("empty", "", "s", "", "", false)
	bool1 := parser.Flag("", "x", "")
	bool2 := parser.Flag("", "y", "")
	bool3 := parser.Flag("", "z", "")
	bool4 := parser.Flag("bool", "", "")
	parser.Parse([]string{"-xyc", "3", "-zs", "str3", "--number=1", "-c=3", "--str1=str1", "--str2", "str2", "--bool", "-b=2"})
	got1 := *number1
	got2 := *number2
	got3 := *number3
	got4 := *str1
	got5 := *str2
	got6 := *str3
	got7 := *bool1
	got8 := *bool2
	got9 := *bool3
	got10 := *bool4
	if got1 != want1 {
		t.Errorf("got %v; want %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("got %v; want %v", got2, want2)
	}
	if got3 != want3 {
		t.Errorf("got %v; want %v", got3, want3)
	}
	if got4 != want4 {
		t.Errorf("got %v; want %v", got4, want4)
	}
	if got5 != want5 {
		t.Errorf("got %v; want %v", got5, want5)
	}
	if got6 != want6 {
		t.Errorf("got %v; want %v", got6, want6)
	}
	if got7 != want7 {
		t.Errorf("got %v; want %v", got7, want7)
	}
	if got8 != want8 {
		t.Errorf("got %v; want %v", got8, want8)
	}
	if got9 != want9 {
		t.Errorf("got %v; want %v", got9, want9)
	}
	if got10 != want10 {
		t.Errorf("got %v; want %v", got10, want10)
	}
}

func Test13(t *testing.T) {
	parser := New("")
	want1 := true
	want2 := true
	want3 := true
	want4 := 11
	want5 := 22
	a := parser.Flag("aa", "a", "")
	b := parser.Flag("aa", "b", "")
	c := parser.Flag("bb", "c", "")
	x := parser.Option(0, "", "x", "", "", false)
	y := parser.Option(0, "", "y", "", "", false)
	parser.Parse([]string{"-abx=11", "-cy", "22"})
	got1 := *a
	got2 := *b
	got3 := *c
	got4 := *x
	got5 := *y
	if got1 != want1 {
		t.Errorf("got %v; want %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("got %v; want %v", got2, want2)
	}
	if got3 != want3 {
		t.Errorf("got %v; want %v", got3, want3)
	}
	if got4 != want4 {
		t.Errorf("got %v; want %v", got4, want4)
	}
	if got5 != want5 {
		t.Errorf("got %v; want %v", got5, want5)
	}
}

func Test14(t *testing.T) {
	parser := New("test")
	want1 := "arg1"
	want2 := true
	verbose := parser.Flag("verbose", "v", "be verbose")
	parser.Parse([]string{"arg1", "-v"})
	got1 := parser.Args[0]
	got2 := *verbose
	if got1 != want1 {
		t.Errorf("got %v; want %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("got %v; want %v", got2, want2)
	}
}

func Test15(t *testing.T) {
	parser := New("test")
	want1 := "arg1"
	want2 := "-v"
	want3 := true
	verbose := parser.Flag("verbose", "v", "be verbose")
	parser.Parse([]string{"arg1", "-v", "--", "-v"})
	got1 := parser.Args[0]
	got2 := parser.Args[1]
	got3 := *verbose
	if got1 != want1 {
		t.Errorf("got %v; want %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("got %v; want %v", got2, want2)
	}
	if got3 != want3 {
		t.Errorf("got %v; want %v", got3, want3)
	}
}

func Test16(t *testing.T) {
	parser := New("test")
	want1 := 1
	want2 := 2
	num1 := parser.Option(0, "num", "", "", "", false)
	num2 := parser.Option(0, "number", "", "", "", false)
	parser.Parse([]string{"--num", "1", "--number", "2"})
	got1 := *num1
	got2 := *num2
	if got1 != want1 {
		t.Errorf("got %v; want %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("got %v; want %v", got2, want2)
	}
}

func Test17(t *testing.T) {
	parser := New("test")
	want1 := 1
	want2 := 2
	num1 := parser.Option(0, "num", "", "", "", false)
	num2 := parser.Option(0, "number", "", "", "", false)
	parser.Parse([]string{"--num=1", "--number=2"})
	got1 := *num1
	got2 := *num2
	if got1 != want1 {
		t.Errorf("got %v; want %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("got %v; want %v", got2, want2)
	}
}

func Test18(t *testing.T) {
	parser := New("test")
	want1 := "str1"
	want2 := "str2"
	str1 := parser.Option("", "str", "", "", "", false)
	str2 := parser.Option("", "string", "", "", "", false)
	parser.Parse([]string{"--str", "str1", "--string", "str2"})
	got1 := *str1
	got2 := *str2
	if got1 != want1 {
		t.Errorf("got %v; want %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("got %v; want %v", got2, want2)
	}
}

func Test19(t *testing.T) {
	parser := New("test")
	want1 := "str1"
	want2 := "str2"
	str1 := parser.Option("", "str", "", "", "", false)
	str2 := parser.Option("", "string", "", "", "", false)
	parser.Parse([]string{"--str=str1", "--string=str2"})
	got1 := *str1
	got2 := *str2
	if got1 != want1 {
		t.Errorf("got %v; want %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("got %v; want %v", got2, want2)
	}
}

func Test20(t *testing.T) {
	p := New("")
	want1 := true
	want2 := true
	want3 := true
	a := p.Flag("", "a", "")
	b := p.Flag("", "b", "")
	c := p.Flag("", "c", "")
	p.Parse([]string{"-abc"})
	got1 := *a
	got2 := *b
	got3 := *c
	if got1 != want1 {
		t.Errorf("got %v; want %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("got %v; want %v", got2, want2)
	}
	if got3 != want3 {
		t.Errorf("got %v; want %v", got3, want3)
	}
}
