package main 

type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForRestangle(*restangle )
}