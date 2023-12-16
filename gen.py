import requests

day = input("What day is it? ")

url = f'https://adventofcode.com/2023/day/{day}/input'

with open(".env", "r") as env:
    cookie = env.read().split("=")[1]
    session = requests.session()
    session.cookies.update({
        'session': cookie
    })

    with open("./days/template.go", "r") as temp:
        with open(f"./data/day{day}", "w") as data:
            data.write(session.get(url).text)
        
        with open(f"./days/day{day}.go", "w") as new_day:
            new_day_content = temp.read().replace("_DAY_", day)
            new_day_content = new_day_content.replace("partOne", f"{day}partOne")
            new_day_content = new_day_content.replace("partTwo", f"{day}partTwo")
            new_day.write(new_day_content)

        with open(f"main.go", "r") as main_r:
            text = main_r.read()

            with open(f"main.go", "w") as main:
                main.write(text[:-2] + f"\tdays.Day{day}()\n" + "}\n")
