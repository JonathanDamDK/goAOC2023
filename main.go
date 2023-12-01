package main

func main() {
	dayArray := []Day{
		DayOne{inputPath: "./inputs/dayOne.txt"},
		DayTwo{inputPath: "./inputs/dayTwo.txt"},
		//new days here
	}
	for _, day := range dayArray {
		day.Solve()
	}
}
