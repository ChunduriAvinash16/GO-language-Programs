package main

import "fmt"

func GenDisplaceFn(x,y,z float64)func (float64) float64 {
		fn:=func(t float64)float64{
			//fmt.Println(0.5*(x*t*t))
			return 0.5*(x*t*t)+(z*t)+(y)
		}
		return fn
}

func main(){
	var time,acce,intial_displacement,intial_velocity float64
	fmt.Println("Enter the acceleration")
	fmt.Scan(&acce)
	fmt.Println("Enter the intial displacement")
	fmt.Scan(&intial_displacement)
	fmt.Println("Enter the intial velocity")
	fmt.Scan(&intial_velocity)
	fn:=GenDisplaceFn(acce,intial_displacement,intial_velocity)
	fmt.Println("Enter the value of time")
	fmt.Scan(&time)
	fmt.Println("New Displacement is ")
	fmt.Println(fn(time))

}