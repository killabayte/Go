calculation_to_units = 24
name_of_unit = "hours"

print(f"20 days are {20 * calculation_to_units} {name_of_unit}")

def days_to_units(num_of_days):
    return f"{num_of_days} days are {num_of_days * calculation_to_units} {name_of_unit}"