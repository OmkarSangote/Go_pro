#include <iostream>
using namespace std;



template < class T >
void bubblesort( T a[], int n)
{
    // int swap(int a, int b)
// {
//     int c = a;
//     a = b;
//     b = c;
//     return 0;
// }
    for (int i = 0; i < n-i; i++)
    {
       for ( int j = n-1; i < j; j--)
       {
           if(a[j] < a[j-1])
           {
               int temp =  a[j];
               a[j] = a[j-1];
               a[j-1] = temp;
           }
           
       } 
    }
}


int main ()
{
    int arr[5] = {10,85,20,96,100};
    int n = sizeof(arr)/ sizeof(arr[0]);
    cout << "Bubble sorted array is" << endl;
    bubblesort<int>(arr, n);
    for ( int i = 0; i < n; i++)
    {
        cout << arr[i] << endl;
    }
    return 0;
}

