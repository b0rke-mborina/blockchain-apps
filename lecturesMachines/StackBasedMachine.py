# %%

# HARDWARE - STACK BASED MACHINE

stack = []

clock = 0
debug = lambda: print("CLK:", clock, "STK:", stack)


def run(program):
    global clock
    for cmd in program:
        debug()
        clock += 1
        oper = cmd[0:3]
        if oper == "PUT":
            p1 = int(cmd[4:])
            stack.append(p1)
        elif oper == "ADD":
            p1 = stack.pop()
            p2 = stack.pop()
            stack.append(p1 + p2)
        elif oper == "CMP":
            p1 = stack.pop()
            p2 = stack.pop()
            stack.append(1 if p1 == p2 else 0)
        elif oper == "PRN":
            print("PRN", stack[-1])
        else:
            raise "???"
    return


# SOFTWARE

program = [
    "PUT 6",
    "PUT 4",
    "PUT 2",
    "ADD",
    "CMP",
    "PRN",
]

run(program)

# %%