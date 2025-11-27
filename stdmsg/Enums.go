package stdmsg

type YesNoEnum uint8
type YesNoValues struct {
	No        YesNoEnum
	Yes       YesNoEnum
	NullValue YesNoEnum
}

var YesNo = YesNoValues{0, 1, 255}

func (m *YesNoEnum) PPrint(i int) {
	switch *m {
	case YesNo.No:
		PPrintlnInd(i, "no")
	case YesNo.Yes:
		PPrintlnInd(i, "yes")
	default:
		PPrintlnInd(i, "null")
	}
}

func (m *YesNoEnum) PPrintCustom(i int, prep, noStr, yesStr string) {
	switch *m {
	case YesNo.No:
		PPrintlnInd(i, prep, noStr)
	case YesNo.Yes:
		PPrintlnInd(i, prep, yesStr)
	default:
		PPrintlnInd(i, prep, "null")
	}
}

type DirectionEnum uint8
type DirectionValues struct {
	Buy          DirectionEnum
	Sell         DirectionEnum
	No_direction DirectionEnum
	NullValue    DirectionEnum
}

var Direction = DirectionValues{0, 1, 2, 255}

func (m *DirectionEnum) PPrint(i int) {
	switch *m {
	case Direction.Buy:
		PPrintlnInd(i, "Side: buy")
	case Direction.Sell:
		PPrintlnInd(i, "Side: sell")
	case Direction.No_direction:
		PPrintlnInd(i, "Side: no direction")
	default:
		PPrintlnInd(i, "Side: null")
	}
}

type BookSideEnum uint8
type BookSideValues struct {
	Ask       BookSideEnum
	Bid       BookSideEnum
	NullValue BookSideEnum
}

var BookSide = BookSideValues{0, 1, 255}

func (m *BookSideEnum) PPrint(i int) {
	switch *m {
	case BookSide.Ask:
		PPrintlnInd(i, "Side: ask")
	case BookSide.Bid:
		PPrintlnInd(i, "Side: bid")
	default:
		PPrintlnInd(i, "Side: null")
	}
}

type BookChangeEnum uint8
type BookChangeValues struct {
	Created   BookChangeEnum
	Changed   BookChangeEnum
	Deleted   BookChangeEnum
	NullValue BookChangeEnum
}

var BookChange = BookChangeValues{0, 1, 2, 255}

func (m *BookChangeEnum) PPrint(i int) {
	switch *m {
	case BookChange.Created:
		PPrintlnInd(i, "Change: Created")
	case BookChange.Changed:
		PPrintlnInd(i, "Change: Changed")
	case BookChange.Deleted:
		PPrintlnInd(i, "Change: Deleted")
	default:
		PPrintlnInd(i, "Change: null")
	}
}
