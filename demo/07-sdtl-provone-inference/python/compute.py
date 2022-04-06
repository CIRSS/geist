import pandas as pd

print('Load two variables A and B from df.csv', '\n')
df = pd.read_csv('../data/df.csv')
print(df, '\n')

print('Overwrite variable with 3', '\n')
df['A'] = 3
print(df, '\n')

print('Add 6.5 to B', '\n')
df['B'] = df.B + 6.5
print(df, '\n')

print('Append new variable C with representing difference between A and B', '\n')
df['C'] = df.A - df.B
print(df, '\n')

print('Load Fahrenheit variable from temps.csv', '\n')
temps = pd.read_csv('../data/temps.csv')
print(temps, '\n')

print('Append new variable Celcius', '\n')
temps = temps.assign(Celsius=((temps.Fahrenheit - 32) * 5 / 9))
print(temps, '\n')

print('Append new variable Kelvin', '\n')
temps = temps.assign(Kelvin=(temps.Celsius + 273))
print(temps, '\n')

print('Write A and B variables to df_updated.csv', '\n')
df.to_csv('products/df_updated.csv', index=False)

print('Write Fahrenheit, Celsius, and Kelvin variables to temps_updated.csv', '\n')
temps.to_csv('products/temps_updated.csv', index=False)
