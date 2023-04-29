# %%

# HARDWARE - REGISTER BASED MACHINE

registers = {f"R{i}": None for i in range(3)}

clock = 0

debug = lambda: print("CLK:", clock, "REG:", registers)


def run(program):
    global clock
    for cmd in program:
        debug()
        clock += 1
        oper = cmd[0:3]
        if oper == "PUT":
            p1, p2 = tuple(cmd[4:].split(","))
            p2 = int(p2)
            registers[p1] = p2
        elif oper == "ADD":
            p1, p2, p3 = tuple(cmd[4:].split(","))
            print(p1, p2, p3)
            registers[p1] = registers[p2] + registers[p3]
        elif oper == "PRN":
            p1 = cmd[4:]
            print("PRN", registers[p1])
        else:
            raise "???"
    return


# SOFTWARE

program = ["PUT R0,1", "PUT R1,2", "ADD R2,R0,R1", "PRN R2"]

run(program)

# %%