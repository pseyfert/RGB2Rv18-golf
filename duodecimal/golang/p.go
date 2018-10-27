package main
import (b"bufio"
m"fmt"
"os"
s"strconv")
func main(){t,_:=b.NewReader(os.Stdin).ReadString('\n')
f,_:=s.ParseFloat(t[0:len(t)-1],64)
t=s.FormatInt(int64(f*1728),12)
l:=len(t)
c:=3
for string(t[l-1])=="0"&&c>0{l-=1
t=t[0:l]
c-=1}
if c>0{m.Printf("%s.%s",t[0:l-c],t[l-c:l])}else{m.Printf("%s",t[0:l])}}
