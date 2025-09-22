#include <iostream>
int main ()
{
    int i,arr[20],largest,pos,n;
    std::cout << "Enter the number of values: ";
    std::cin >> n;
    for (i=0;i<n;i++)
    {
        std::cin >> arr[i];
    }

    largest = arr[0];
    pos = 0;
    for (i=1;i<n;i++)
    {
        if (largest < arr[i])
        {
            largest = arr[i];
            pos = i;
        }      
    }
    std::cout << "Largest = " << largest << " and its Position is " << pos+1 << "\n";
}