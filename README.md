# LED MATRIX GOLANG

## Summary

This repository contains code which can display text on a 2 dimensional LED matrix. 

## Raspberry Pi Demo 

If an LED segment display has been wired up to the Raspberry Pi correctly. After running the following command: 
`go run .`

https://user-images.githubusercontent.com/32711718/201545693-e95054c0-2adf-431f-a8f4-d615a9b160d9.mov

## Terminal Demo

To run the code without using an Raspberry Pi and to get the output to print onto the terminal use the following command:

`go run . -debug` The debug flag specifies to just print on the terminal. 

https://user-images.githubusercontent.com/32711718/201545908-18fe8b8d-04c7-478e-a7ae-63ae89c2b968.mov

## Available characters

The system currently supports the following chars:

![image](https://user-images.githubusercontent.com/32711718/210154944-e522fb2f-626d-4ad0-a933-b0437dae911f.png)

The above diagram was generated in the ./cmd/possible-letters main prog. Thanks to this repo: https://github.com/gowhale/go-circuit-diagram

## How it works?

### LED MATRIX

The first 2 pages of the following document give an overview of what an LED matrix is: https://docs.broadcom.com/doc/AV02-3697EN 

My summary: An LED is a Light Emitting Diode. A diode is that only takes electricity in one direction. A LED is therefore a light source which only takes electricity one way. Because the LED only takes electricity one way we can wire up the LED's into a matrix to control lots of LEDS.

In my setup I have wired up 64 LED's and control them only using 16 GPIO pins. Note my code can only turn on one LED at a time. I use a technique called multiplexing to flash many LED's quickly that the HUMAN eye cannot percieve them turning off. Meaning that we can mock the behaviour of controlling many LED's at once. 

### Configuration

To create a custom config for different sized matrixs you must create a config.json file to specify what pins control the rows and columns of the LED Matrix. This configuration file is also what the terminal output works off. So if you specify a 10 by 50 matrix in a .json and use `go run . -debug` it will show that size matrix in the terminal.  

## Developing High Quality Software

To develop High Quality Software I have followed the following best practices:
1. ensure all pkg's have code coverage of over 75%
2. Enforced the golangci-lint linter on all Pull Requests (PR)
3. Enforced tests to pass on PR's
4. Enforced the revive linter on all PR's

This being said the code is far from perfect. If you can think of any features, refactors or bugs feel free to raise an Issue or even fork the repo and create a PR! 

## Test Coverage

One of my GitHub actions deploys a website which shows what code is not covered: https://gowhale.github.io/LED-MATRIX-GOLANG/
