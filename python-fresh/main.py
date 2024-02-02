#!/usr/bin/env python3
calculation_to_units = 24
name_of_unit = "hours"
num_of_days = 365

print(f"20 days are {20 * calculation_to_units} {name_of_unit}")

def days_to_units(num_of_days):
    return f"{num_of_days} days are {num_of_days * calculation_to_units} {name_of_unit}"

print(days_to_units(num_of_days))

# Path: python-fresh/main.py
