@echo off
set commit=%1
if "%commit%"=="" (
	echo "please input git commit !"
)else (
	git add .
	git commit -m "%commit%"
	git push origin master
)
