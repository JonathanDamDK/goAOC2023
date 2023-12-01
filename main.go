package main;


func main() {
	dayArray := []Day{
		DayOne{inputPath : "./inputs/dayOneDebug.txt"},
		//new days here
	};
	for _,day := range dayArray{
		day.Solve()
	}
}
