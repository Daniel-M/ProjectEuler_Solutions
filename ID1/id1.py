#!/usr/bin/python3

limit=1000
sum_tot = 0

for i in range(0,limit):
    if(i%3==0 or i%5==0):
        sum_tot+=i

print("The total sum of multiples of 3 and 5 is ",sum_tot)
