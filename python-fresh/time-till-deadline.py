from datetime import datetime

user_input = input ("Enter your goal with a deadline separated by a colon\n")
input_list = user_input.split(":")

goal = input_list[0]
deadline = input_list[1]

deadline_date = datetime.strptime(deadline, "%d.%m.%Y")
today_date = datetime.today()

time_till_deadline = deadline_date - today_date
print(time_till_deadline)

print(f"The time remaining for your goal which is {input_list[0]}: " + str(time_till_deadline.days) + " days")