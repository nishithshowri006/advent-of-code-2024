import re
def mul(a,b):
    return a*b
def part1(text):
    match = r"mul\(\d{1,3},\d{1,3}\)"
    a = re.findall(match,text)
    ans = 0
    for val in a:
        ans += eval(val)
    return ans

def part2(text):
    # match = r"don't\(\).*?(?:$|do\(\))
    match = r"(?:don't\(\)).*?(?:$|do\(\))"
    a = re.sub(match,"",text, flags=re.DOTALL)
    ans= part1(a)
    return ans
with open("inputs.txt","r") as f:
    text = f.read()
#part1
ans1 = part1(text)
#part2
ans2 = part2(text)
text = re.sub(r"don't\(\).*?(?:$|do\(\))", '', text, flags=re.DOTALL)
print('part 2', sum(int(a)*int(b) for a,b in re.findall(r'mul\((\d+),(\d+)\)', text)))
print(ans1,ans2, ans1-ans2)
