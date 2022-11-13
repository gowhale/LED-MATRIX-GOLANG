// Package gui is responsible for visual output
// File letters.go stores information on how to represent letters with lights
package gui

const (
	//LetterWidth is how many LED's wide a single letter is
	LetterWidth = 5
)

// type letterMatrix [5][5]int

//LetterMap maps a char to a LetterMatrix
var LetterMap = map[string][][]int{
	" ": space,
	"a": letterA,
	"b": letterB,
	"c": letterC,
	"d": letterD,
	"e": letterE,
	"f": letterF,
	"g": letterG,
	"h": letterH,
	"l": letterL,
	"m": letterM,
	"o": letterO,
	"p": letterP,
	"r": letterR,
	"s": letterS,
	"t": letterT,
	"u": letterU,
	"w": letterW,
	"y": letterY,
}

var fullRow = []int{LEDOn, LEDOn, LEDOn, LEDOn, LEDOn}
var emptyRow = []int{LEDOff, LEDOff, LEDOff, LEDOff, LEDOff}

var space = [][]int{
	emptyRow,
	emptyRow,
	emptyRow,
	emptyRow,
	emptyRow,
}

var letterA = [][]int{
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOn, LEDOff, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	fullRow,
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
}

var letterC = [][]int{
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
}

var letterD = [][]int{
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
}

var letterP = [][]int{
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
}

var letterE = [][]int{
	fullRow,
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	fullRow,
}

var letterF = [][]int{
	fullRow,
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	{LEDOn, LEDOn, LEDOn, LEDOff, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
}

var letterG = [][]int{
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOn, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOn},
}

var letterL = [][]int{
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	fullRow,
}

var letterO = [][]int{
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
}

var letterU = [][]int{
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
}

var letterM = [][]int{
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOn, LEDOff, LEDOn, LEDOn},
	{LEDOn, LEDOff, LEDOn, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
}

var letterR = [][]int{
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
}

var letterB = [][]int{
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
}

var letterT = [][]int{
	fullRow,
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
}

var letterH = [][]int{
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	fullRow,
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
}

var letterS = [][]int{
	fullRow,
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	fullRow,
	{LEDOff, LEDOff, LEDOff, LEDOff, LEDOn},
	fullRow,
}

var letterW = [][]int{
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOn, LEDOff, LEDOn},
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
}

var letterY = [][]int{
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOff, LEDOn, LEDOff, LEDOn, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
}
