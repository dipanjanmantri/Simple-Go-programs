package main

import (
"fmt"
"math"
)


func main(){
rect := rectangle{4,6}
c := circle{5}
fmt.Println("rectangle perimter =",getPerimeter(rect))
fmt.Println("circle perimeter =",getPerimeter(c))

}

type Shape interface{
	perimeter() float64
}

type rectangle struct{
	height,width float64
}

type circle struct{
	radius float64
}


func(r rectangle) perimeter() float64{
	return (2 * (r.width+r.height))
}

func (c circle) perimeter() float64{
	return (2*(math.Pi)*(c.radius))
}

func getPerimeter(shape Shape) float64{
	return shape.perimeter()
}