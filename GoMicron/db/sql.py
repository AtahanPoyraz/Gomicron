import sqlite3

con = sqlite3.connect("users.db")
cur = con.cursor()

cur.execute("SELECT * FROM users")
res = cur.fetchall()

for row in res:
    print(row)

con.close()
