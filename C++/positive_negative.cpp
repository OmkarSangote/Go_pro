#include <iostream>
int main ()
{
    int psum=0,pcount=0,nsum=0,ncount=0,arr[100],n,i;
    std::cout << "Enter the number of element of array : ";
    std::cin >> n;
    for (i=0;i<n;i++)
    {
        std::cout << "Enter the elements for the array : ";
        std::cin >> arr[i];
    }

    for (i=0;i<n;i++)
    {
        if (arr[i] > 0) 
        {
            psum = arr[i] + psum;
            pcount++;
        }

        if (arr[i] < 0)
        {
            nsum = nsum + arr[i];
            ncount++;
        }

    }
    std::cout << "Number of postive numbers = " << pcount << " and ther sum is = " << psum << "\n";
    std::cout << "Number of negative numbers = " << ncount << " and ther sum is = " << nsum << "\n";
}