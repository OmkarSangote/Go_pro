"""
Read a complex number and perform addition and subtraction operations on them and print result in a single
line separated by character 
"""
x_real = int(input("Enter the real part"))
x_img= int(input("Enter the imaginary part"))
y_real= int(input("Enter the real number"))
y_img= int(input("Enter the imaginary number"))
a=complex(x_real,x_img)
b=complex(y_real,y_img)
print(a)
print(b)
print('sum=',(a.real+b.real),'+j', (a.imag+b.imag),'difference=',(a.real-b.real),'+j', (a.imag-b.imag))

