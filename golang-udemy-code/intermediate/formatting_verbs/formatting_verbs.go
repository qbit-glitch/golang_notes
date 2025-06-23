package main

import "fmt"

func main() {
	i := 8_295.5
	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%v%%\n", i)

	str := "Hello World"
	fmt.Printf("%v\n", str)
	fmt.Printf("%#v\n", str)
	fmt.Printf("%T\n", str)

	inte := 255
	fmt.Printf("%b\n", inte)
	fmt.Printf("%d\n", inte)
	fmt.Printf("%+d\n", inte)
	fmt.Printf("%o\n", inte)
	fmt.Printf("%O\n", inte)
	fmt.Printf("%x\n", inte)
	fmt.Printf("%X\n", inte)
	fmt.Printf("%#x\n", inte)
	fmt.Printf("%4d\n", inte)
	fmt.Printf("%-4d\n", inte)
	fmt.Printf("%04d\n", inte)

	txt := "World"
	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	t := true
	f := false

	fmt.Printf("%t\n", t)
	fmt.Printf("%t\n", f)

	flt := 9129876545678.18987654
	fmt.Printf("%e\n", flt)
	fmt.Printf("%f\n", flt)
	fmt.Printf("%.2f\n", flt)
	fmt.Printf("%10.2f\n", flt)
	fmt.Printf("%g\n", flt)


}
