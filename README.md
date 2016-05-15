# ProjectEuler_Solutions  
### Daniel Mejía R ([danielmejia55@gmail.com](mailto:danielmejia55@gmail.com))  

## Description
The implementation of my solutions of the project euler problems in `Go`.        

# Installing the package
If you have `go` installed just run  

`go get github.com/Daniel-M/ProjectEuler_Solutions`  

Otherwise clone this repository and compile it with `go build`   

# Running the solutions

Call the executable `ProjectEuler_Solutions` like this,    

`ProjectEuler_Solutions --id ##`

Where `##` is the number of the problem whose solution is gonna be implemented.   

To see more features invoke

`ProjectEuler_Solutions --help`

## About the structure of this project

The main source `preuler.go` imports the solutions provided in each `idx` sub directory.   
The folder `utilities` has several packages that are of common use when solving the problems.   
They are still undocumented but I'll work on that for future reference.   
