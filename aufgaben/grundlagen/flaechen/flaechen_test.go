package main

import "fmt"

// Erwartet zwei Seitenlängen a und b.
// Liefert die Fläche des entsprechenden Rechtecks.
func ExampleAreaRectangle() {
	fmt.Println(AreaRectangle(3, 4))
	fmt.Println(AreaRectangle(2.5, 10))
	fmt.Println(AreaRectangle(1.5, 25))

	// Output:
	// 12
	// 25
	// 37.5
}

// Erwartet eine Seitenlänge a.
// Liefert die Fläche des entsprechenden Quadrats.
func ExampleAreaSquare() {
	fmt.Println(AreaSquare(0))
	fmt.Println(AreaSquare(1))
	fmt.Println(AreaSquare(2))
	fmt.Println(AreaSquare(3))
	fmt.Println(AreaSquare(4))
	fmt.Println(AreaSquare(5))

	// Output:
	// 0
	// 1
	// 4
	// 9
	// 16
	// 25
}

// Erwartet zwei Seitenlängen a und b.
// Liefert die Fläche eines rechtwinkligen Dreiecks, dessen Katheten a und b sind.
func ExampleAreaRightTriangle() {
	fmt.Println(AreaRightTriangle(3, 4))
	fmt.Println(AreaRightTriangle(2.5, 10))
	fmt.Println(AreaRightTriangle(1.5, 25))

	// Output:
	// 6
	// 12.5
	// 18.75
}

// Erwartet zwei Seitenlängen a und b.
// Liefert den Umfang des entsprechenden Rechtecks.
func ExamplePerimeterRectangle() {
	fmt.Println(PerimeterRectangle(3, 4))
	fmt.Println(PerimeterRectangle(2.5, 10))
	fmt.Println(PerimeterRectangle(1.5, 25))

	// Output:
	// 14
	// 25
	// 53
}

// Erwartet eine Seitenlänge a.
// Liefert den Umfang des entsprechenden Quadrats.
func ExamplePerimeterSquare() {
	fmt.Println(PerimeterSquare(0))
	fmt.Println(PerimeterSquare(1))
	fmt.Println(PerimeterSquare(2))
	fmt.Println(PerimeterSquare(3))
	fmt.Println(PerimeterSquare(4))
	fmt.Println(PerimeterSquare(5))

	// Output:
	// 0
	// 4
	// 8
	// 12
	// 16
	// 20
}

// Erwartet zwei Seitenlängen a und b.
// Liefert den Umfang eines rechtwinkligen Dreiecks, dessen Katheten a und b sind.
func ExamplePerimeterRightTriangle() {
	fmt.Println(PerimeterRightTriangle(3, 4))
	fmt.Println(PerimeterRightTriangle(2.5, 10))
	fmt.Println(PerimeterRightTriangle(1.5, 25))

	// Output:
	// 12
	// 22.80776406404415
	// 51.54495957273639
}

// Erwartet die Längen der Katheten eines rechtwinkligen Dreiecks.
// Liefert die Länge der Hypotenuse.
func ExampleHypotenuse() {
	fmt.Println(Hypotenuse(1, 1))
	fmt.Println(Hypotenuse(3, 4))

	// Output:
	// 1.4142135623730951
	// 5
}
