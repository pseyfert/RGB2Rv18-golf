package main
import (b"bufio"
m"fmt"
"os"
s"strconv")
func main(){t,_:=b.NewReader(os.Stdin).ReadString('\n')
f,_:=s.ParseFloat(t[0:len(t)-1],64)
d:=s.FormatInt(int64(f*1728),12)
l:=len(d)
c:=1
for string(d[l-c])=="0"&&c<4{c+=1}
l-=3
if c<4{t="."}else{t=""}
m.Printf("%s%s%s",d[0:l],t,d[l:l-c+4])}
