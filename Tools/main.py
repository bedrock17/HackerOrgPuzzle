import requests

s = requests.session()

level = input("level : ");

req = s.get("http://www.hacker.org/coil/index.php?gotolevel=8&go=Go+To+Level", cookies={"PHPSESSID":"vtaJUtPy3mTed7bESEcSZ1", "phpbb2mysql_sid":"be01423027e7e26f871a81ee368bb2b5"})

print(req.text)


