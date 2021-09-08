
def contains(s, l):
    return s in l
    
file = open('./magical_suffix.go')
content = file.read()
file.close()

lines = content.split('\n')
lines = lines[11:757]

idx = 0
out = []
nums = []
for i in lines:
    s =i.split(":")
    n = s[0]
    nums.append(int(n))
    l = s[1]
    l = l.replace(',', '').replace(' ', '')
    if not contains(l, out):
        out.append(l)
    idx+=1

out2 = []

idx = 0
x = 0

while not idx == len(out):
    while not x == nums[idx]:
        out2.append('_')
        x+=1
    out2.append(out[idx])
    idx+=1
    x+=1


for i in out:
    print(i)

#file = open("./magic_prefix.go","w")
#file.write("package itemdata\n")
#file.write("const (\n")
#file.write("_ MagicalPrefix = iota\n")
#for i in out:
#    file.write(i)
#    file.write("\n")
#file.write(")")
#file.close()
