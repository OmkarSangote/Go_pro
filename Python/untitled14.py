# -*- coding: utf-8 -*-
"""
Find smallest of four numbers accept numbers from keyboard
"""
p,q,r,s=input("Enter the numbers\n").split( )
p,q,r,s=int(p),int(q),int(r),int(s)
if p<q and p<r and p<s:
    print("The smallest number among",p,q,r,s,"is",p)
elif q<r and q<s:     
    print("The smallest number among",p,q,r,s,"is",q)
elif r<s:
    print("The smallest number among",p,q,r,s,"is",r)
else:
    print("The smallest number among",p,q,r,s,"is",s)