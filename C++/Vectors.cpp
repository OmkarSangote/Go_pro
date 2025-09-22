 #include <iostream>
 #include <vector>

using namespace std;

int main()
{
    vector<int> g;

    for (int i = 1; i <= 5; i++)
    g.push_back(i);

    for (auto i = g.begin(); i != g.end(); i++)
    cout << *i << " " ;

// Adding element
    g.push_back(6);
    g.push_back(7);
    cout <<"\nresized array is:" << endl;

    for (auto i = g.begin(); i != g.end(); i++)
    cout << *i << " ";

//accessing element

cout << "\nelement at index 0 is:" << g.at(0);
cout << "\nelement at index 1 is:" << g[1];

//changing vector element
g.at(0)= 10;
g[5]=50;
  cout <<"\nupdated array is:" << endl;

    for (auto i = g.begin(); i != g.end(); i++)
    cout << *i << " ";

//last element
 vector<int>::iterator iter;
  iter= g.end()-1;

 cout << "\nlast element is:" << *iter ;

// Delete element 
g.pop_back();

cout <<"\nafter delete updated array is:" << endl;

    for (auto i = g.begin(); i != g.end(); i++)
    cout << *i << " ";
 
//print using iterator
cout << "\nprint using iterators: \n";
for (iter = g.begin(); iter != g.end(); ++iter)
cout << *iter << " ";

//delete first element

g.erase(g.begin());
cout <<"\nafter removing first element updated array is:" << endl;

    for (auto i = g.begin(); i != g.end(); i++)
    cout << *i << " ";


// insert middle

g.insert(g.begin()+2, 68);
cout <<"\nafter inserting missdle element updated array is:" << endl;

    for (auto i = g.begin(); i != g.end(); i++)
    cout << *i << " ";


  return 0;

}




// // C++ program to illustrate the
// // iterators in vector
// #include <iostream>
// #include <vector>
  
// using namespace std;
  
// int main()
// {
//     vector<int> g1;
  
//     for (int i = 1; i <= 5; i++)
//         g1.push_back(i);
  
//     cout << "Output of begin and end: ";
//     for (auto i = g1.begin(); i != g1.end(); ++i)
//         cout << *i << " ";
  
//     cout << "\nOutput of cbegin and cend: ";
//     for (auto i = g1.cbegin(); i != g1.cend(); ++i)
//         cout << *i << " ";
  
//     cout << "\nOutput of rbegin and rend: ";
//     for (auto ir = g1.rbegin(); ir != g1.rend(); ++ir)
//         cout << *ir << " ";
  
//     cout << "\nOutput of crbegin and crend : ";
//     for (auto ir = g1.crbegin(); ir != g1.crend(); ++ir)
//         cout << *ir << " ";
  
//     return 0;
// }