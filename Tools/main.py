import requests
import os

s = requests.session()

now = input("level : ")

for level in range(int(now), 101):
  level = str(level)

  #예제입니다 본인의 세션으로 변경해주세요
  qobj = {"PHPSESSID":"vtaJUtPy3md7bESEcSZ1", "phpbb2mysql_sid":"be01423027e7e26f871a81ee368b5"}

  req = s.get("http://www.hacker.org/coil/index.php?gotolevel="+level+"&go=Go+To+Level",
  cookies=qobj)

  # print(req.text)


  pre="<param name=\"FlashVars\" value=\""
  value=req.text[req.text.find(pre)+len(pre):]
  value=value[:value.find("\"")]

  print("Value= "+value)

  arr = value.split("&")

  print(arr)

  cmd = "app " + arr[0][2:] + " " + arr[1][2:] + " " + arr[2][6:]

  print("yun app command : " + cmd)
  os.system(cmd)

  urlf = open("outurl", "rt")
  qpathurl = urlf.read(1 << 15)
  urlf.close()

  print(qpathurl)
  req = s.get(qpathurl,
  cookies=qobj)

  # print(req.text)

