package main;


func main() {
	dayArray := []Day{
		DayOne{inputPath : "./inputs/dayOne.txt"},
		//new days here
	};
	for _,day := range dayArray{
		day.Solve()
	}
}
