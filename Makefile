dbup: 
	psql -h localhost -U postgres -d pdoro -f C:\Users\joshu\Documents\projects\pomodoro\server\db\migrations\up.sql

dbdown:
	psql -h localhost -U postgres -d pdoro -f C:\Users\joshu\Documents\projects\pomodoro\server\db\migrations\down.sql