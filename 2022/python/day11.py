from typing import Self


class Monkey:
    def __init__(
        self,
        items: list[int],
        op: str,
        op_val: int | str,
        test_val: int,
        pass_to: list[int],
    ) -> None:
        self.items = items
        self.op = op
        if op_val != "old":
            self.op_val = int(op_val)
        else:
            self.op_val = op_val

        self.test_val = test_val
        self.pass_to = pass_to

        self.inspections = 0

    def operate(self, a) -> int:
        op_val = self.op_val
        if self.op_val == "old":
            op_val = a

        match self.op:
            case "+":
                return a + op_val
            case "-":
                return a - op_val
            case "*":
                return a * op_val
            case "/":
                return a / op_val

        raise Exception("this shouldn't happen")

    def test_and_throw(self, a, monkeys: list[Self]):
        a = self.operate(a) // 3

        self.inspections += 1
        if a % self.test_val == 0:
            monkeys[self.pass_to[0]].items.append(a)

        else:
            monkeys[self.pass_to[1]].items.append(a)

    def __str__(self) -> str:
        return f"Items: {self.items} | Op: {self.op_val} | Test: {self.test_val}, {self.pass_to}"


def get_data(pth="../data/day_11.txt") -> list[Monkey]:
    with open(pth) as fin:
        x = fin.read()

    x = [[j.strip() for j in i.split("\n")][1:] for i in x.split("\n\n")]
    x = [[j.split(":")[1].strip() for j in i if len(j) > 0] for i in x]
    monkeys = []
    for m in x:
        # init items
        items = [int(i.strip()) for i in m[0].split(",")]

        # Operation
        op = m[1].split()
        operation = op[-2]
        val = op[-1]

        # Test
        test_val = int(m[2].split()[-1])
        pass_to = [int(m[3].split()[-1]), int(m[4].split()[-1])]

        monk = Monkey(items, operation, val, test_val, pass_to)
        monkeys.append(monk)
        print(monk)

    return monkeys


def run():
    # pth = "test.txt"
    # monkeys = get_data(pth)
    monkeys = get_data()
    print()

    for _ in range(20):
        for m in monkeys:
            while len(m.items) > 0:
                i = m.items[0]
                m.items = m.items[1:]
                m.test_and_throw(i, monkeys)
        # break

    print()
    ans = []
    for m in monkeys:
        ans.append(m.inspections)

    ans = sorted(ans, reverse=True)
    ans = ans[0] * ans[1]

    print(ans)
