#include<iostream>
int main()
{
    int arr[100],i,n,asc=1,dsc=1;
    std::cout << "Enter the number of elements for the series : ";
    std::cin >> n;
    for (i=0;i<n;i++)
    {
        std::cout << "Enter the elements of the array: \n";
        std::cin >> arr[i];
    }

    for (i=0;i<n;i++)
    {
        if (arr[i] < arr[i+1]) asc = 0;
        if (arr[i] > arr[i+1]) dsc = 0;
    }

    if (asc==1) std::cout << "Array is in ascending order  \n";
    else if (dsc==1) std::cout << "Array is in descending order  \n";
    else std::cout << "Array is not unsorted  \n";

}