package question2

import "testing"

func Test_Interface_Check_Circle(t *testing.T) {
	circleInstance := Circle{0, 0, 5}
    circlePerimeter := returnPerimeter(circleInstance)
    if circlePerimeter!=31.41592653589793 {
        t.Error("Circle does not implement Shape interface and does not return perimeter")
    }
}

func Test_Interface_Check_Rectangle(t *testing.T) {
    rectangleInstance := Rectangle{0,0,5,5}
    rectanglePerimeter := returnPerimeter(rectangleInstance)
    if rectanglePerimeter!=10 {
        t.Error("Rectangle does not implement Shape interface and does not return perimeter")
    }
}

func Test_Rectangle_area(t *testing.T) {
    rec := Rectangle{0,0,5,5}
    result := rec.area()
    if result!=25 {
        t.Error("Expected 25, got ",result)
    }
}

func Test_Rectangle_perimeter(t *testing.T) {
    rec := Rectangle{0,0,5,5}
    result := rec.perimeter()
    if result!=10 {
        t.Error("Expected 10, got ",result)
    }
}

func Test_Circle_perimeter(t *testing.T) {
    cir := Circle{0, 0, 5}
    result := cir.perimeter()
    if result!=31.41592653589793 {
        t.Error("Expected 31.41592653589793, got ",result)
    }
}

func Test_Circle_area(t *testing.T) {
    cir := Circle{0, 0, 5}
    result := cir.area()
    if result!=78.53981633974483 {
        t.Error("Expected 78.53981633974483, got ",result)
    }
}