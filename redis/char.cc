#include <iostream>
#include <sstream>
#include <string>
#include <vector>

using namespace std;

void deal(vector<int> vec)
{
  
}
int main() {
  const char *buf = "asc";

  cout << buf;

  vector<string> ss;
  ss.push_back("sss");
  ss.push_back("eee");

  for (int i = 0; i < ss.size(); i++) {
    cout << ss[i] << ",";
  }

  // for (auto i : ss) {
  //   const char* tmp = i;
  //   cout << tmp << "\n";
  // }
  // // string s = "hello";

  // const char * b = s.data();
  // cout << b;
  // const char *c = s.c_str();
  // cout << c;
  vector<int> vec;
  deal(vec);
  deal(vector<int>());
  return 0;
}

