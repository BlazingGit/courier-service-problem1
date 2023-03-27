# courier-service-problem1

The project is written with Go.

You can run the program by opening the executable main.exe.

If Go is installed in your machine, you can run below command to execute the program:
go run ./main.go

At the main_test.go file, function getTestData() is where I put the test input and output list.
Run 'go test -v' to execute the test function.

I also attached a sample screenshot of the program. (SampleOutput.PNG)

Sample Input:
100 3
PKG1 5 5 OFR001
PKG2 15 5 OFR002
PKG3 10 100 OFR003

Sample Output:
PKG1 0 175 0
PKG2 0 275 0
PKG3 35 665 0

Thank you~
