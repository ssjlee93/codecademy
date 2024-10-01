import enums.job as job
import enums.element as element
import character
import random

def validateJob():
    """
    Validates input string whether it belongs to Job enum

    Returns:
        string: the accepted name
    """
    accepted = False
    while accepted == False:
        try:
            test = job.Job(input("What is your job? \n"))
            accepted = True
            break
        except:
            print("Oops! we don't have that job in our world.")
    return test


def main():
    # start the game
    print("elcome to magic fights!")

    # intake player name
    player_name = input("What is your name? \n")
    print(f"Greetings! {player_name}!")

    # intake player job
    player_job = validateJob()
    print(f"You are now {player_job}")

    # randomize element
    random_element = random.choice(list(element.Element))
    print(f"You are chosen to wield {random_element}")

    player_character = character.Character(player_name, random_element, player_job)
    print(player_character)

# run python program
main()
