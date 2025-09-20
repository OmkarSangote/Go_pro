"""
Read a three book names and check a book ‘Data structures using C’ in a read book names
"""
a=input("Enter the first book name")
b=input("Enter the second book name")
c=input("Enter the third book name")
s=('Data structures using C')
s=s.lower()
if s in a:
    print("The book is present in a")
elif s in b:
    print("The book is present in b")
elif si in c:
    print("The book is present in c")
else:
    print("The book is not in the given list")
    