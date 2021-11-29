// package main

// import (
// 	"fmt"
// 	"gowork2/week13/calendar"
// 	"log"
// )

// func main() {
// event := calendar.Event{}
// event.Mydate.SetYear(2021)
// event.Mydate.SetMonth(2021)
// event.Mydate.SetDay(2021)
// fmt.Println(event.Mydate.GetYear())

//익명 구조체 사용
//event := calendar.Event{}

//event.Title = "전역일"

// err := event.SetTitle("눈 오는 날")
// if err != nil {
// 	log.Fatal(err)
// }

// event.SetYear(2019)
// event.SetMonth(12)
// event.SetDay(9)

// fmt.Println(event.GetTitle(), event.GetYear(), event.GetMonth(), event.GetDay())
//}

package main

// type 인터페이스명 interface {
// 	메서드(타입)
// }
import (
	"gowork2/week13/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

//func playList(device gadget.TapeRecorder, songs []string) {
func playList(device Player, songs []string) { //인터페이스 타입으로 받음
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()

	//device.Record() 에러
	//Player 인터페이스는 Play(), Stop()만 구현이 되어 있어서 Record()사용 불가
}

func main() {
	player := gadget.TapeRecorder{}
	mixtape := []string{"난 알아요", "거짓말", "그우사우"}

	playList(player, mixtape)
}
