import pandas as pd

# Replace 'data.json' with the path to your JSON file
json_file = 'security_groups.json'

# Read the JSON data into a DataFrame
df = pd.read_json(json_file)

# Replace 'output.xlsx' with the desired output Excel file name
excel_file = 'output.xlsx'

# Write the DataFrame to an Excel file
df.to_excel(excel_file, index=False, engine='openpyxl')
