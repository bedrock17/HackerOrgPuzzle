import requests
import os

s = requests.session()

userid = input("ID : ")
userpwd = input("PWD : ")

payload = {'username': userid, 'password': userpwd,
"redirect":"", "login":"Log+in"}

url = 'http://www.hacker.org/forum/login.php'
ret = s.post(url, data=payload)

#print(ret.text)

now = input("level : ")

maxlevel = 300
for level in range(int(now), maxlevel + 1):
  level = str(level)


  req = s.get("http://www.hacker.org/coil/index.php?gotolevel="+level+"&go=Go+To+Level") 
  # print(req.text)


  pre="<param name=\"FlashVars\" value=\""
  value=req.text[req.text.find(pre)+len(pre):]
  value=value[:value.find("\"")]

  print("Value= "+value)

  arr = value.split("&")

  #print(arr)

  cmd = "goapp " + arr[0][2:] + " " + arr[1][2:] + " " + arr[2][6:]

  print("cube golang app command : " + cmd)
  os.system(cmd)

  urlf = open("outurl", "rt")
  qpathurl = urlf.read(1 << 15)
  urlf.close()

  print("tarter url => ", qpathurl)
  req = s.get(qpathurl)

  # print(req.text)

