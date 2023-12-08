package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

const (
	HighCard  int = 1
	OnePair       = 2
	TwoPair       = 3
	ThreeKind     = 4
	FullHouse     = 5
	FourKind      = 6
	FiveKind      = 7
)

type DaySeven struct {
	inputPath string
	input     []string
	outputpt1 int
	outputpt2 int
}

// a hand in camelCard
type camelHand struct {
	cards   [5]int
	bid     int
	handVal int
}

type handArray []camelHand

func (hands handArray) Len() int {
	return len(hands)
}
func (hands handArray) Less(i, j int) bool {
	// You can define your custom comparison logic here.
	// For example, let's sort people by age in ascending order.
	if hands[i].handVal == hands[j].handVal {
		index := 0
		for hands[i].cards[index] == hands[j].cards[index] {
			index++
		}
		return hands[i].cards[index] < hands[j].cards[index]
	}
	return hands[i].handVal < hands[j].handVal
}
func (hands handArray) Swap(i, j int) {
	hands[i], hands[j] = hands[j], hands[i]
}

func (day DaySeven) Solve() {
	//A, K, Q, J, T
	regex := regexp.MustCompile("(\\d|A|K|Q|J|T)(\\d|A|K|Q|J|T)(\\d|A|K|Q|J|T)(\\d|A|K|Q|J|T)(\\d|A|K|Q|J|T) (\\d+)")
	lines := MapFileToStringArr(day.inputPath)
	var hands handArray
	for _, line := range lines {
		match := regex.FindStringSubmatch(line)
		var hand camelHand
		for i := 0; i < 5; i++ {
			hand.cards[i] = getNumericValue(match[i+1])
		}
		num, _ := strconv.Atoi(match[6])
		hand.bid = num
		sortedArr := hand.cards
		sort.Ints(sortedArr[:])
		hands = append(hands, hand)
	}

	for i, hand := range hands {
		sortedArr := hand.cards
		sort.Ints(sortedArr[:])
		score := 0
		prev := sortedArr[0]
		count := 1
		for i := 1; i < 5; i++ {
			if prev != sortedArr[i] {
				switch count {
				case 1:
					if HighCard > score {
						score = HighCard
					}
				case 2: // pair
					if score == OnePair {
						if TwoPair > score {
							score = TwoPair
						}
					} else if score == ThreeKind {
						if FullHouse > score {
							score = FullHouse
						}
					} else {
						if OnePair > score {
							score = OnePair
						}
					}
				case 3:
					if score == OnePair {
						if FullHouse > score {
							score = FullHouse
						}
					} else {
						if ThreeKind > score {
							score = ThreeKind
						}
					}
				case 4:
					if FourKind > score {
						score = FourKind
					}
				case 5:
					if FiveKind > score {
						score = FiveKind
					}
				}
				count = 1
				prev = sortedArr[i]
			} else {
				count++
			}
		}
		if sortedArr[4] == sortedArr[3] {
			switch count {
			case 1:
				if HighCard > score {
					score = HighCard
				}

			case 2: // pair
				if score == OnePair {
					if TwoPair > score {
						score = TwoPair
					}
				} else if score == ThreeKind {
					if FullHouse > score {
						score = FullHouse
					}
				} else {
					if OnePair > score {
						score = OnePair
					}
				}
			case 3:
				if score == OnePair {
					if FullHouse > score {
						score = FullHouse
					}
				} else {
					if ThreeKind > score {
						score = ThreeKind
					}
				}
			case 4:
				if FourKind > score {
					score = FourKind
				}
			case 5:
				if FiveKind > score {
					score = FiveKind
				}
			}
		}
		hand.handVal = score
		hands[i] = hand
	}
	sort.Sort(hands)
	for i, hand := range hands {
		day.outputpt1 += (i + 1) * hand.bid
	}

	//part 2
	for i, hand := range hands {
		for j, card := range hand.cards {
			if card == 11 {
				hands[i].cards[j] = -1
				hand.cards[j] = -1
			}
		}
		sortArr := hand.cards
		sort.Ints(sortArr[:])
		prev := 0
		currCount := 1
		count_1 := 1
		joker_count := 0
		count_2 := 1
		for _, card := range sortArr {
			if card == -1 {
				joker_count += 1
			} else {
				if card != prev {
					if currCount > count_1 {
						count_1 = currCount
					} else if currCount > count_2 {
						count_2 = currCount
					}
					currCount = 1
					prev = card
				} else {
					currCount++
					prev = card
				}
			}
		}
		if sortArr[4] == sortArr[3] {
			if currCount > count_1 {
				count_1 = currCount
			} else if currCount > count_2 {
				count_2 = currCount
			}
		}
		//case time
		//Two pair
		if count_1 == 2 && count_2 == 2 {
			if joker_count == 1 {
				hands[i].handVal = FullHouse
			}
		} else {
			if joker_count > 0 {
				count_1 = joker_count + count_1
				score := hand.handVal
				switch count_1 {
				case 2:
					score = OnePair
				case 3:
					score = ThreeKind
				case 4:
					score = FourKind
				case 5:
					score = FiveKind
				}
				hands[i].handVal = score
			}
		}
	}
	sort.Sort(hands)
	for i, hand := range hands {
		day.outputpt2 += (i + 1) * hand.bid
	}
	fmt.Printf("day 7 part 1 : %d part 2: %d \n", day.outputpt1, day.outputpt2)
}

func getNumericValue(val string) int {
	switch val {
	case "T":
		return 10
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	default:
		num, _ := strconv.Atoi(val)
		return num
	}
}
