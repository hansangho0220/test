package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println(song, "재생중...")
}

func (t TapePlayer) Stop() {
	fmt.Println("중지!")
}

///////////////////////////////////

type TapeRecorder struct {
	Micorphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println(song, "재생중...")
}

func (t TapeRecorder) Record() {
	fmt.Println("녹음중...")
}

func (t TapeRecorder) Stop() {
	fmt.Println("중지!")
}
