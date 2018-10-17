package chess

import "fmt"

func ExampleBits_String() {
	x := Bits(0x0000000F)

	for i := uint8(0); i < 64; i++ {
		fmt.Print(x.At(i))
	}

	//Output:
	//1111000000000000000000000000000000000000000000000000000000000000
}

func ExampleBitboard_SetTo() {
	var bb BitBoard
	bb.SetTo(NewBoard().WithMove(MoveP(RC(1, 4), RC(3, 4)))) // King's Pawn
	fmt.Println(&bb)

	//Output:
	//xxxxxxxx
	//xxxxxxxx
	//........
	//........
	//....x...
	//........
	//xxxx.xxx
	//xxxxxxxx

}
