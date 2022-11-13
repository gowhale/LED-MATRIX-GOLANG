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


## How it works?

### LED MATRIX

The first 2 pages of the following document give an overview of what an LED matrix is: https://docs.broadcom.com/doc/AV02-3697EN 

My summary: An LED is a Light Emitting Diode. A diode is that only takes electricity in one direction. A LED is therefore a light source which only takes electricity one way. Because the LED only takes electricity one way we can wire up the LED's into a matrix to control lots of LEDS.

In my setup I have wired up 64 LED's and control them only using 16 LED's. Note my code can only turn on one LED at a time. I use a technique called multiplexing to flash many LED's quickly that the HUMAN eye cannot percieve them turning off. Meaning that we can mock the behaviour of controlling many LED's at once. 

For example: 
