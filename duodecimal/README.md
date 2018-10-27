# SHELL (zsh)

NB: the submitted script is the invisible `.sh`. Submit with:

## without using extendedglobs

```
# read from stdin
read l;

# multiply the number by 12^6 (base change can't handle stuff after the decimal point)
# print the result to the variable r
# print it in base 12 without the leading base modifier
r=$(([##12]l*12**6));

# -n do not print a trailing line feed
# -- do not interpret the result as options (in case of negative numbers
# (L) convert letters to lower case
# [1,-7] and [-6,-1] pick all but the last 6 characters, and the last 6 characters, respectively
# . print a '.' between them
# first sed: remove trailing 0s
# second sed: remove trailing decimal point
print -n -- ${(L)r[1,-7]}.${(L)r[-6,-1]}|sed -e"s/0*$//" -e"s/\.$//"
```

run with

```sh
duodecimal -run 'zsh %s' .sh
```

## with using extendedglobs

```
# read from stdin
read l;

# multiply the number by 12^6 (base change can't handle stuff after the decimal point)
# print the result to the variable r
# print it in base 12 without the leading base modifier
r=$(([##12]l*12**6));

# -- do not interpret the result as options (in case of negative numbers
# (L) convert letters to lower case
# [1,-7] and [-6,-1] pick all but the last 6 characters, and the last 6 characters, respectively
# %% remove the largest trailing match ...
# 0## ... of an arbitrary amout of zeroes
# . print a . between them
s=${(L)r[1,-7]}.${(L)r[-6,-1]%%0##}

# -n do not print a trailing line feed
# %. remove trailing .
print -n -- ${s%.}
```

run with

```sh
duodecimal -run 'zsh --extendedglob %s' .sh
```

# C++

```cpp
#include <iostream>
using namespace std;

int main(){
  string s;

  // read from stdin
  getline(cin,s);

  // convert to double
  auto i=atof(s.c_str());
  auto t=1.;

  // if input is smaller than 0, print a minus and flip sign
  if (i<0) {
    cout<<"-";
    i*=-1;
  }

  // increment t by a factor 12 until it is less than a factor 12 smaller than the input
  while ((int)i/t>12) t*=12;

  // iterate to the input until there's less than 1e-5 missing
  while (i>1e-5) {
    // print current digit as hexadecimal (0 to b)
    printf("%x",(int)(i/t));
    // remove what just got computed
    i-=(int)(i/t)*t;
    // when flipping past the decimal point, print it, if needed
    if (t==1&&i>1e-5) cout<<".";
    // go to the next less significant digit
    t/=12;
  }
  // add trailing zeroes
  while (t>.5) {
    cout<<"0";
    t/=12;
  }
}
```
