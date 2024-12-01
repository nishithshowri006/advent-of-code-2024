with open("day1.txt", "r") as f:
    a = []
    b = []
    for line in f.readlines():
        o,s = line.split()
        a.append(int(o))
        b.append(int(s))
a.sort()
b.sort()
print(sum([abs(i-j) for i,j in zip(a,b)]))


