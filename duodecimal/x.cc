#include <iostream>
using namespace std;int main(){string s;getline(cin,s);auto i=atof(s.c_str());auto t=1.;if(i<0){cout<<"-";i=-i;}while((int)i/t>12)t*=12;while(i>1e-5||t>.5){printf("%x",(int)(i/t));i-=(int)(i/t)*t;if(t==1&&i>1e-5)cout<<".";t/=12;}}
