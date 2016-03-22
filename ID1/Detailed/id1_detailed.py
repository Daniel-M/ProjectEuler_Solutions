#!/usr/bin/python3


def sum_total(values):
    sum=0
    for i in range(0,len(values)):
        sum+=values[i]
    return sum


limit=1000

multiples_of_3 = [0]
multiples_of_5 = [0]
multiples_of_3_and_5 = [0]

for i in range(0,limit):
    if(i%3==0):
        multiples_of_3.append(i)
    if(i%5==0):
        multiples_of_5.append(i)
    if(i%3==0 or i%5==0):
        multiples_of_3_and_5.append(i)

print("The total sum of multiples of 3 is ",sum_total(multiples_of_3))
print("The total sum of multiples of 5 is ",sum_total(multiples_of_5))
print("The total sum of multiples of 3 or 5 is ",sum_total(multiples_of_3_and_5))
