#include <iostream>
#include <queue>
using namespace std;

int main()
{
 queue<int> q;
 q.push(54);
 q.push(85);
 q.push(40);

 cout<< "Queue is:" << endl;
 while(!q.empty())
 {
     cout << q.top() << " ";
     q.pop();
 }
 
 return 0;
}

