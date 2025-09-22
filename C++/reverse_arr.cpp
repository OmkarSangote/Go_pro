#include <iostream>
int main()
{
    int i,arr[100],temp,n,r;
    std::cout << "Input number of elements : ";
    std::cin >> n;
    for (i=0;i<n;i++)
    {
        std::cout << "Enter the elements of the array : ";
        std::cin >> arr[i];
    }

    for (i=0,r=n-1;i<n/2;i++,r--)
    {
        temp = arr[i];
        arr[i] = arr[r];
        arr[r] = temp;
    }

    std::cout << "Reversed array elements are : \n";
    for(i=0;i<n;i++)
    {
        std::cout << arr[i] << "\n";
    }
}