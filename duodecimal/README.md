# preface

The task is to convert decimal number from stdin to duodecimal.  Until a patch
it had to be printed without trailing line feed.  The number is known to have a
finite number of non-zero digits. It can have digits after the (duo)decimal
point, in both bases a finite number of them. The number may be negative.
Duodecimal digits shall be:
0, 1, 2, 3, 4, 5, 6, 7, 8, 9, a, b

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

# originally `print -n --`, later `echo - `
# print:
# -n do not print a trailing line feed
# -- do not interpret the result as options (in case of negative numbers)
# echo:
# - do not interpret the result as options (in case of negative numbers)
# (L) convert letters to lower case
# [1,-7] and [-6,-1] pick all but the last 6 characters, and the last 6 characters, respectively
# . print a '.' between them
# first sed: remove trailing 0s
# second sed: remove trailing decimal point
echo - ${(L)r[1,-7]}.${(L)r[-6,-1]}|sed -e's/0*$//' -e's/\.$//'
```

run with

```sh
duodecimal -run 'zsh *' .sh
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

# echo/print same as above
# %. remove trailing .
echo - ${s%.}
```

run with

```sh
duodecimal -run 'zsh --extendedglob *' .sh
```

# C++

```cpp
#include <iostream>
using namespace std;

int main(){
  string s;

  // read from stdin
  getline(cin, s);

  // convert to double
  auto i = atof(s.c_str());
  auto t = 1.;

  // if input is smaller than 0, print a minus and flip sign
  if (i<0) {
    cout << "-";
    i = -i;
  }

  // increment t by a factor 12 until it is less than a factor 12 smaller than the input
  while ((int)i / t > 12) t *= 12;

  // iterate to the input until there's less than 1e-5 missing
  // or until the 12^0 digit is reached (print trailing zeroes until the decimal point, if integer input)
  while (i > 1e-5 || t > 0.5) {
    // print current digit as hexadecimal (0 to b)
    printf("%x", (int)(i / t));
    // remove what just got computed
    i -= (int)(i / t) * t;
    // when flipping past the decimal point, print it, if needed
    if (t == 1 && i > 1e-5) cout << ".";
    // go to the next less significant digit
    t /= 12;
  }
}
```

# python3

```py
# shorthand for printing without trailing space or line feed
def p(a):
    print(a, end='')

# read from stdin and convert to float
i = float(input())
t = 1.
x = 1e-5
# print leading minus (if input negative) and make input positive
if i < 0:
    p('-')
    i = -i
# increment t until int(i/t) is a single duodecimal digit
while int(i/t) > 12:
    t *= 12

# iterate down to the 12^0 and if needed beyond until iteration better than x
while i > x or t > .5:
    # use f-string printing as hexadecimal of int(i/t)
    # (using this in the winning solution by nicolas would've beaten the winning solution)
    p(f"{int(i/t):x}")
    # subtract from i what just got printed
    i -= int(i/t)*t
    # check if decimal point is needed now
    if t == 1 and i > x:
        p(".")
    # go to the next less significant digit
    t /= 12
```

# go

```go
package main

// imports with shorthand module name (when beneficial)
import (
	b "bufio"
	m "fmt"
	"os"
	s "strconv"
)

func main() {
	// read from stdin until line feed
	t, _ := b.NewReader(os.Stdin).ReadString('\n')
	// strip trailing line feed and convert to 64-bit float
	f, _ := s.ParseFloat(t[0:len(t)-1], 64)
	// shift duodecimal point by three digits to the right
	// and convert the result in base 12 string
	d := s.FormatInt(int64(f*1728), 12)
	// get number of current digits
	l := len(d)
	// count how many trailing zeroes need to get cut off
	// counter off by one to save characters
	c := 1
	// maximum cut off is three (otherwise we reach the 12^0 digit)
	for string(d[l-c]) == "0" && c < 4 {
		c += 1
	}
	// reduce l to save a few characters in the slices below
	l -= 3
	// recycle the t variable (saves declaration) and put the duodecimal point
	// for printing into it, and leave out if it won't be needed (i.e. when all
	// post-duodecimal point digits will get supressed
	if c < 4 {
		t = "."
	} else {
		t = ""
	}
	// print all except the last three digits (`l` is length minus three) of the
	// string version of the duodecimal representation of input*12^3, then
	// the decimal point/empty string if needed, then the remaining digits except
	// trailing zeroes
	m.Printf("%s%s%s", d[0:l], t, d[l:l-c+4])
}
```
