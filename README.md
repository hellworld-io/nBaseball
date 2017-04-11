# nBaseball

## What is nBaseball?

### nBaseball is a Number Baseball Game which is written by Go.

### This project is that looks for correct number.

## How to use it?

* Here are options.
	* maxL : max number length(3 ~ 5). Default 3.
	* lOpt : 1, 2, 3
		* 1 : All number do not use 0.
		* 2 : First number does no use 0 only.
		* 3 : 0 to 9 numbers use (default "1")
```
	go run nBaseball.go -maxL 3 -lOpt 1
```

* When you start, computer will make random number.
* When you inputed numbers, you can check result.
	* Strikes : same position and same number
	* Ball : different position same number.
* When you find correct number, you will see your result.

```
	$ go run nBaseball.go -maxL 3 -lOpt "1"
    Computer is ready.
    If you want to exit, you will input exit.
    When you started, It will be recoded time.
    321
    "321" result is "0Strikes 0Balls"
    987
    "987" result is "1Strikes 0Balls"
    965
    "965" result is "2Strikes 0Balls"
    964
    Your result is 0D 0H 0M 22S, and you try 4
    If you want to see results, press enter.
```



