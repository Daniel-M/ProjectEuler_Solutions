#!/usr/bin/python3


def fibonacci_n(Nth):
    sum=0
    fibonacci_number = 1;
    fibonacci_number_b = 1;
    fibonacci_seed = 1
    
    for i in range (1,Nth):
        fibonacci_number_b = fibonacci_number + fibonacci_seed
        fibonacci_seed = fibonacci_number
        fibonacci_number = fibonacci_number_b

    return fibonacci_number;


limit=4000000
counter =0
sum_tot=0

for i in range(0,limit):
    iBuffer = fibonacci_n(i)

    if(iBuffer%2==0 and iBuffer < limit):
        sum_tot+=iBuffer
        counter+=1

    if(iBuffer >= limit):
        break

print("The total sum of the first ",counter," Fibonacci even numbers is ",sum_tot)
