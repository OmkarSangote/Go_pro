#include<iostream>
int main()
{
    int a[10],i,n,ele,pos;
    std::cout << "Enter the number of elements\n";
    std::cin >> n;
    std::cout << "Enter the elements of the array\n";
    for (i=0;i<n;i++)
    {
        std::cin >> a[i];
    }
    std::cout << "Enter the element to be added\n";
    std::cin >> ele;
    std::cout << "Enter the pos to be added\n";
    std::cin >> pos;
    if (pos > n)
    {
        std::cout << "Invalid Position";
    }
    else
    {
        for (i=n-1;i>=pos;i--)
        {
            a[i+1] = a[i];
        }
            a[pos] = ele;
            n = n+1;
            std::cout << "The elements of the array\n";
            for (i=0;i<n;i++)
            {
                std::cout << a[i] << "\n";
            }
    }
    
}
