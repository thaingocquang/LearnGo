package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/atomicgo/cursor"
)

var attribute map[int]string

var especialCard map[int]string

type Player struct {
	id    int
	cards []Card
}

type Card struct {
	point      int
	attributeC int
}

type MaxCard struct {
	playerId int
	maxC     Card
}

func initAttribute() {
	attribute = make(map[int]string)
	attribute[1] = "co"
	attribute[2] = "ro"
	attribute[3] = "chuon"
	attribute[4] = "bich"
}

func initPackOfCard(PackOfCard *[]Card) {
	for i := 1; i <= 13; i++ {
		for j := 1; j <= len(attribute); j++ {
			card := Card{i, j}
			*PackOfCard = append(*PackOfCard, card)
		}
	}
}

func initEspecialCard() {
	especialCard = make(map[int]string)
	especialCard[1] = "A"
	especialCard[11] = "J"
	especialCard[12] = "Q"
	especialCard[13] = "K"
}

func checkIfElementInArray(ele int, arr []int) (indexWinner int, isHave bool) {
	for i := 0; i < len(arr); i++ {
		if ele == arr[i] {
			return i, true
		}
	}
	return 0, false
}

func ChiaBai(playerNums int, players []Player) {
	for k := 0; k < 3; k++ {
		cursor.Up(playerNums)
		if k == 0 {
			cursor.Right((k + 1) * 15)
		} else {
			cursor.Right(20)
		}
		for i := 0; i < playerNums; i++ {
			point := especialCard[players[i].cards[k].point]
			if point != "" {
				fmt.Printf("%2s%6s", point, attribute[players[i].cards[k].attributeC])
			} else {
				fmt.Printf("%2d%6s", players[i].cards[k].point, attribute[players[i].cards[k].attributeC])
			}

			time.Sleep(time.Millisecond * 50)
			cursor.Down(1)
			cursor.Left(8)
		}
	}
	cursor.Bottom()
	cursor.Left(30)
}

func TinhDiem(res []int, players []Player) {
	for i := 0; i < len(players); i++ {
		sumPoint := 0
		for j := 0; j < 3; j++ {
			point := 0
			if players[i].cards[j].point >= 10 {
				point = 10
			} else {
				point = players[i].cards[j].point
			}
			sumPoint += point
		}
		res[i] = sumPoint % 10
	}
}

func FindFinalWinnerInWinner(winner []int, players []Player, res []int) {
	if len(winner) == 1 {
		for i := 0; i < len(players); i++ {
			if i == winner[0] {
				fmt.Println("Player", i+1, ":", res[i], "diem -> Winner")
			} else {
				fmt.Println("Player", i+1, ":", res[i], "diem")
			}

		}
	} else if len(winner) > 1 {

		// tim card lon nhat cua tung thang
		var maxCards []MaxCard
		for i := 0; i < len(winner); i++ {
			maxCard := MaxCard{winner[i], players[winner[i]].cards[0]}
			// fmt.Println("check khoi tao max card", maxCard)
			for j := 1; j < len(players[winner[i]].cards); j++ {
				//so sanh chat
				if players[winner[i]].cards[j].attributeC == maxCard.maxC.attributeC {
					diemChat1 := players[winner[i]].cards[j].point
					diemChat2 := maxCard.maxC.point
					if diemChat1 == 1 {
						diemChat1 += 15
					}
					if diemChat2 == 1 {
						diemChat2 += 15
					}
					if diemChat1 >= diemChat2 {
						maxCard = MaxCard{winner[i], players[winner[i]].cards[j]}
					}

				} else if players[winner[i]].cards[j].attributeC < maxCard.maxC.attributeC {
					maxCard = MaxCard{winner[i], players[winner[i]].cards[j]}
				}
			}

			maxCards = append(maxCards, maxCard)

		}

		finalWinner := maxCards[0]
		for _, v := range maxCards {
			if v.maxC.attributeC == finalWinner.maxC.attributeC {
				if v.maxC.point >= finalWinner.maxC.point {
					finalWinner = v
				}
			} else if v.maxC.attributeC < finalWinner.maxC.attributeC {
				finalWinner = v
			}
		}
		// fmt.Println("Nguoi chien thang cuoi cung:", finalWinner)
		for i := 0; i < len(players); i++ {
			isWinner := false
			indexWinner, isHave := checkIfElementInArray(i, winner)

			if isHave {
				if i == finalWinner.playerId {
					point := especialCard[finalWinner.maxC.point]
					if point != "" {
						fmt.Println("Player", i+1, ":", res[i], "diem", "(", point, attribute[finalWinner.maxC.attributeC], ")", " -> Winner")
					} else {
						fmt.Println("Player", i+1, ":", res[i], "diem", "(", finalWinner.maxC.point, attribute[finalWinner.maxC.attributeC], ")", " -> Winner")
					}

					isWinner = true
				} else {
					point := especialCard[maxCards[indexWinner].maxC.point]
					if point != "" {
						fmt.Println("Player", i+1, ":", res[i], "diem", "(", point, attribute[maxCards[indexWinner].maxC.attributeC], ")")
					} else {
						fmt.Println("Player", i+1, ":", res[i], "diem", "(", maxCards[indexWinner].maxC.point, attribute[maxCards[indexWinner].maxC.attributeC], ")")
					}

					isWinner = true
				}
			}
			if !isWinner {
				fmt.Println("Player", i+1, ":", res[i], "diem")
			}
		}
		// fmt.Println(maxCards)

	}
}

func InitCardForPlayer(players []Player, copyPackOfCard []Card) {
	remainCards := 52
	for i := 0; i < len(players); i++ {
		players[i].id = i + 1
		players[i].cards = make([]Card, 3)
		for j := 0; j < 3; j++ {
			randIndex := rand.Intn(remainCards)
			players[i].cards[j] = copyPackOfCard[randIndex]
			copyPackOfCard = append(copyPackOfCard[:randIndex], copyPackOfCard[randIndex+1:]...)
			remainCards = remainCards - 1
		}
	}
}

func FindWinnerEqualPoint(res []int, winner *[]int) {
	max := 0
	for _, v := range res {
		if v >= max {
			max = v
		}
	}
	// tim nhung thang co diem bang nhau
	for i, v := range res {
		if v == max {
			*winner = append(*winner, i)
		}
	}
}

func main() {

	for {

		// init chat, JQKA, bo bai
		rand.Seed(time.Now().UnixNano())
		PackOfCard := make([]Card, 0)
		initAttribute()
		initEspecialCard()
		initPackOfCard(&PackOfCard)

		// nhap so nguoi choi
		var playerNums int
		fmt.Print("Nhap so nguoi choi: ")
		_, err := fmt.Scanf("%d\n", &playerNums)
		if err == nil {
			fmt.Println("So nguoi choi:", playerNums)
		}
		if playerNums < 3 || playerNums > 17 {
			fmt.Println("So nguoi choi phai lon hon 3 va nho hon 18")
			return
		}

		//khoi tao bai cho nguoi choi
		copyPackOfCard := PackOfCard
		players := make([]Player, playerNums)
		InitCardForPlayer(players, copyPackOfCard)

		fmt.Print("\033[H\033[2J")
		for i := 0; i < playerNums; i++ {
			fmt.Println("Player", i+1, ":")
		}

		// Chia bai cho player
		ChiaBai(playerNums, players)
		fmt.Println("KET QUA:")

		// tinh diem luu vao mang ket qua
		var res = make([]int, len(players))
		TinhDiem(res, players)

		var winner []int
		// tim nguoi choi co diem bang nhau
		FindWinnerEqualPoint(res, &winner)
		// tim nguoi chien thang cuoi cung
		FindFinalWinnerInWinner(winner, players, res)

		fmt.Println()
		fmt.Print("Wanna play a gain? (y-n)")
		var isContinue string
		_, err2 := fmt.Scanf("%s\n", &isContinue)
		if err2 != nil {
			break
		} else if isContinue != "y" {
			break
		}
	}
}
