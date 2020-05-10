package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type dataSet struct {
	accleration, velocity, displacement, time float64
}
func main()  {	
	input := [4]string{"accleration", "velecity", "displacement", "time"}
	data := []float64{}
	reader := bufio.NewReader(os.Stdin)
	for _, qst := range input {
		fmt.Printf("Enter value for %s:", qst)
		line, err := reader.ReadString('\n')
		if (err != nil) {
			fmt.Println(err)
			os.Exit(1)
		}
		value, err := strconv.ParseFloat(strings.TrimRight(line, "\r\n"), 64)
		if (err != nil) {
			fmt.Println(err)
			os.Exit(1)
		}
		data = append(data, value)		
	}
	var dataSet = dataSet{
		accleration: data[0],
		velocity: data[1],
		displacement: data[2],
		time: data[3],
	}
	fn := GenDisplaceFn(dataSet.accleration, dataSet.velocity, dataSet.displacement)
	fmt.Println(fn(dataSet.time))
}
func GenDisplaceFn(accleration, velocity, displacement float64) (
	func (time float64) float64) {
		fn := func (time float64) float64 {
			return 0.5 * accleration * (time) *(time) + velocity*time + displacement
		}
		return fn
	}