with open("day1.txt", "r") as f:
    a = {}
    b = {}
    for line in f.readlines():
        o,s = line.split()
        if a.get(int(o)):
            a[int(o)] +=1
        else:
            a[int(o)] =1
        if b.get(int(s)):
            b[int(s)] +=1
        else:
            b[int(s)] =1
ans = 0
for key in a:
    ans += key*a[key]*b.get(key,0)
print(ans)

