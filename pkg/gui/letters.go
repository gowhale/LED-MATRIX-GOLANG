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
	"i": letterI,
	"j": letterJ,
	"k": letterK,
	"l": letterL,
	"m": letterM,
	"n": letterN,
	"o": letterO,
	"p": letterP,
	"q": letterQ,
	"r": letterR,
	"s": letterS,
	"t": letterT,
	"u": letterU,
	"v": letterV,
	"w": letterW,
	"x": letterX,
	"y": letterY,
	"z": letterZ,

	"0": number0,
	"1": number1,
	"2": number2,
	"3": number3,
	"4": number4,
	"5": number5,
	"6": number6,
	"7": number7,
	"8": number8,
	"9": number9,
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

var letterB = [][]int{
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
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

var letterH = [][]int{
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	fullRow,
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
}

var letterI = [][]int{
	fullRow,
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	fullRow,
}

var letterJ = [][]int{
	fullRow,
	{LEDOff, LEDOff, LEDOff, LEDOn, LEDOff},
	{LEDOff, LEDOff, LEDOff, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOn, LEDOff},
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
}

var letterK = [][]int{
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOn, LEDOff},
	{LEDOn, LEDOn, LEDOn, LEDOff, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
}

var letterL = [][]int{
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	fullRow,
}

var letterM = [][]int{
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOn, LEDOff, LEDOn, LEDOn},
	{LEDOn, LEDOff, LEDOn, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
}

var letterN = [][]int{
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOn, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOn, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOn, LEDOn, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
}

var letterO = [][]int{
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
}

var letterP = [][]int{
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
}

var letterQ = [][]int{
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOn, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOn, LEDOff},
	{LEDOff, LEDOn, LEDOn, LEDOff, LEDOn},
}

var letterR = [][]int{
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
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

var letterT = [][]int{
	fullRow,
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
}

var letterU = [][]int{
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
}

var letterV = [][]int{
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOff, LEDOn, LEDOff, LEDOn, LEDOff},
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
}

var letterW = [][]int{
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOff, LEDOn, LEDOff, LEDOn},
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
}

var letterX = [][]int{
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOff, LEDOn, LEDOff, LEDOn, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOn, LEDOff, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
}

var letterY = [][]int{
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOff, LEDOn, LEDOff, LEDOn, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
}

var letterZ = [][]int{
	fullRow,
	{LEDOff, LEDOff, LEDOff, LEDOn, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOn, LEDOff, LEDOff, LEDOff},
	fullRow,
}

// numbers

//  000
// 0  00
// 0 0 0
// 00  0
//  000
var number0 = [][]int{
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOn, LEDOn},
	{LEDOn, LEDOff, LEDOn, LEDOff, LEDOn},
	{LEDOn, LEDOn, LEDOff, LEDOff, LEDOn},
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
}

//  00
//   0
//   0
//   0
//  000
var number1 = [][]int{
	{LEDOff, LEDOn, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
}

// 0000
//     0
//  000
// 0
// 00000
var number2 = [][]int{
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOff, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	fullRow,
}

// 0000
//     0
//   00
//     0
// 0000
var number3 = [][]int{
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOff, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOff, LEDOff, LEDOn, LEDOn, LEDOff},
	{LEDOff, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
}

//   00
//  0 0
// 0  0
// 00000
//    0
var number4 = [][]int{
	{LEDOff, LEDOff, LEDOn, LEDOn, LEDOff},
	{LEDOff, LEDOn, LEDOff, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOn, LEDOff},
	fullRow,
	{LEDOff, LEDOff, LEDOff, LEDOn, LEDOff},
}

// 00000
// 0
// 0000
// 	   0
// 0000
var number5 = [][]int{
	fullRow,
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOff, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
}

//  0000
// 0
// 0000
// 0   0
//  000
var number6 = [][]int{
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOn},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOff},
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
}

// 00000
//     0
//    0
//   0
//  0
var number7 = [][]int{
	fullRow,
	{LEDOff, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOff, LEDOff, LEDOff, LEDOn, LEDOff},
	{LEDOff, LEDOff, LEDOn, LEDOff, LEDOff},
	{LEDOff, LEDOn, LEDOff, LEDOff, LEDOff},
}

//  000
// 0   0
//  000
// 0   0
//  000
var number8 = [][]int{
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
}

//  000
// 0   0
//  0000
//     0
// 0000
var number9 = [][]int{
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOff},
	{LEDOn, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOff, LEDOn, LEDOn, LEDOn, LEDOn},
	{LEDOff, LEDOff, LEDOff, LEDOff, LEDOn},
	{LEDOn, LEDOn, LEDOn, LEDOn, LEDOff},
}
