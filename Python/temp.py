#Calculate simple and compound interest by reading principal amount, rate of interest and period
from math import*
p = float(input('Enter the principal amount '))
t = int(input('Enter the time period in years '))
r = float(input('Enter the rate of interest '))
simple_interest = (p*t*r)/100
x = p * (pow((1 + r / 100), t))
compound_interest = x - p
print('Compund interest=',compound_interest)
print('simple interest=',simple_interest)
