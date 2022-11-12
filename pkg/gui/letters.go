// Package gui is responsible for visual output
// File letters.go stores information on how to represent letters with lights
package gui

const (
	//LetterWidth is how many vapes wide a single letter is
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

var fullRow = []int{VapeOn, VapeOn, VapeOn, VapeOn, VapeOn}
var emptyRow = []int{VapeOff, VapeOff, VapeOff, VapeOff, VapeOff}

var space = [][]int{
	emptyRow,
	emptyRow,
	emptyRow,
	emptyRow,
	emptyRow,
}

var letterA = [][]int{
	{VapeOff, VapeOff, VapeOn, VapeOff, VapeOff},
	{VapeOff, VapeOn, VapeOff, VapeOn, VapeOff},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	fullRow,
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
}

var letterC = [][]int{
	{VapeOff, VapeOn, VapeOn, VapeOn, VapeOff},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOff},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOff, VapeOn, VapeOn, VapeOn, VapeOff},
}

var letterD = [][]int{
	{VapeOn, VapeOn, VapeOn, VapeOn, VapeOff},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOn, VapeOn, VapeOn, VapeOff},
}

var letterP = [][]int{
	{VapeOn, VapeOn, VapeOn, VapeOn, VapeOff},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOn, VapeOn, VapeOn, VapeOff},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOff},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOff},
}

var letterE = [][]int{
	fullRow,
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOff},
	{VapeOn, VapeOn, VapeOn, VapeOn, VapeOn},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOff},
	fullRow,
}

var letterF = [][]int{
	fullRow,
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOff},
	{VapeOn, VapeOn, VapeOn, VapeOff, VapeOff},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOff},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOff},
}

var letterG = [][]int{
	{VapeOff, VapeOn, VapeOn, VapeOn, VapeOn},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOff},
	{VapeOn, VapeOff, VapeOff, VapeOn, VapeOn},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOff, VapeOn, VapeOn, VapeOn, VapeOn},
}

var letterL = [][]int{
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOff},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOff},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOff},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOff},
	fullRow,
}

var letterO = [][]int{
	{VapeOff, VapeOn, VapeOn, VapeOn, VapeOff},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOff, VapeOn, VapeOn, VapeOn, VapeOff},
}

var letterU = [][]int{
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOff, VapeOn, VapeOn, VapeOn, VapeOff},
}

var letterM = [][]int{
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOn, VapeOff, VapeOn, VapeOn},
	{VapeOn, VapeOff, VapeOn, VapeOff, VapeOn},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
}

var letterR = [][]int{
	{VapeOn, VapeOn, VapeOn, VapeOn, VapeOff},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOn, VapeOn, VapeOn, VapeOff},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
}

var letterB = [][]int{
	{VapeOn, VapeOn, VapeOn, VapeOn, VapeOff},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOn, VapeOn, VapeOn, VapeOff},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOn, VapeOn, VapeOn, VapeOff},
}

var letterT = [][]int{
	fullRow,
	{VapeOff, VapeOff, VapeOn, VapeOff, VapeOff},
	{VapeOff, VapeOff, VapeOn, VapeOff, VapeOff},
	{VapeOff, VapeOff, VapeOn, VapeOff, VapeOff},
	{VapeOff, VapeOff, VapeOn, VapeOff, VapeOff},
}

var letterH = [][]int{
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	fullRow,
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
}

var letterS = [][]int{
	fullRow,
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOff},
	fullRow,
	{VapeOff, VapeOff, VapeOff, VapeOff, VapeOn},
	fullRow,
}

var letterW = [][]int{
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOn, VapeOff, VapeOn, VapeOff, VapeOn},
	{VapeOn, VapeOn, VapeOn, VapeOn, VapeOn},
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
}

var letterY = [][]int{
	{VapeOn, VapeOff, VapeOff, VapeOff, VapeOn},
	{VapeOff, VapeOn, VapeOff, VapeOn, VapeOff},
	{VapeOff, VapeOff, VapeOn, VapeOff, VapeOff},
	{VapeOff, VapeOff, VapeOn, VapeOff, VapeOff},
	{VapeOff, VapeOff, VapeOn, VapeOff, VapeOff},
}
