"""
Read a hexadecimal, octal and binary number from keyboard in a single line and find their product and print
result using format function.
"""
x,y,z=input("Enter the 3 numbers").split()
x=int(x)
y=int(y)
z=int(z)
print(hex(x),type(x))
print(oct(y),type(y))
print(bin(z),type(z))
product=x*y*z
print("The product of {0}*{1}*{2}".format(x,y,z),"=",product)