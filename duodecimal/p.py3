def p(a):print(a,end='')
i=float(input());t=1.;x=1e-5
if i<0:p('-');i=-i
while int(i/t)>12:t*=12
while i>x or t>.5:
 p(f"{int(i/t):x}");i-=int(i/t)*t
 if t==1 and i>x:p(".")
 t/=12
