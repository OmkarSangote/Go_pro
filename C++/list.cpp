#include <iostream>
#include <list>
using namespace std;

int main ()
{
    list<int> l;
    list<int> l1;
    list<int>::iterator i;

    for(int a = 1 ; a <= 5; a++)
    {
     l.push_back(a);
    }

    cout << "the list is:\n";
    for(i = l.begin(); i != l.end(); ++i)
    {
        cout << *i << " ";
    }
  l.push_front(8);
  cout << "\nthe list after push front is:\n";
    for(i = l.begin(); i != l.end(); ++i)
    {
        cout << *i << " ";
    }

    l.sort();
      cout << "\nthe list after sorting is:\n";
    for(i = l.begin(); i != l.end(); ++i)
    {
        cout << *i << " ";
    }
    l.pop_front();
    cout << "\nthe list after pop is:\n";
    for(i = l.begin(); i != l.end(); ++i)
    {
        cout << *i << " ";
    }
    for(int a = 1 ; a <= 5; a++)
    {
     l1.push_back(a);
    }
    l1.splice(l1.begin(), l);
      cout << "\nthe list after splice is:\n";
    for(i = l1.begin(); i != l1.end(); ++i)
    {
        cout << *i << " ";
    }

    return 0;

}