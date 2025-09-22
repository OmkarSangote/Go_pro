#include<iostream>
int main()
{
    int a[10],i,n,m,loc,b,e,pos,ele;
    std::cout << "Enter the number of elements";
    std::cin >> n;
    std::cout << "Enter the elements of the array";
    for (i=0;i<n;i++)
    {
        std::cin >> a[i];
    }
    std::cout << "Enter the element to be serached";
    std::cin >> ele;
    pos = -1;
    b = 0;
    e = n - 1;
    while (b <= e)
    {
        m = (b+e)/2;
        if(ele == a[m])
        {
            pos = m;
            break;
        }
        else
        {
            if(ele < a[m])
            {
                e = m - 1;
            }
            else
            {
                b = m + 1;
            }
        }
    }
    if (pos >= 0)
    {
        std::cout << "Positon = " << pos;
    }
    else 
    {
        std::cout << "Search unsuccessful";
    }

}