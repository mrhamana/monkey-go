import random

mylis = ["rock", "paper", "scissor"]

mydic = {"rock": "scissors", "paper": "rock", "scissors": "paper"}

win, loss = 0, 0
enter = ""
while True:
    enter = input("rock, paper , scissors :")

    if enter in mydic:
        comp = mylis[random.randint(0, 2)]

        if mydic[enter] == comp:
            win += 1
            print("WONNN. The score is {} win and {} losses".format(win, loss))
        elif enter == comp:
            print("Drawww")
        else:

            loss += 1
            print("LOSTT. The score is {} win and {} losses".format(win, loss))

    else:
        print("Enter valid input")
