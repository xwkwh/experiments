#include <iostream>
#include <sstream>
#include <string>
#include <vector>


int main() {
//   std::string a = ("a" + "b");
   std::stringstream buf;
   buf << "a" << "b";

   std::cout << buf.str() ;


   return 0;
}