import:
	mongoimport --db='gifty'  --collection='employees' --jsonArray --file='datasets/employees.json' --uri mongodb+srv://andy:ruby12@cluster0.ih56b.mongodb.net/?retryWrites=true&w=majority \
	mongoimport --db='gifty' --collection='gifts' --jsonArray --file='datasets/gifts.json' --uri mongodb+srv://andy:ruby12@cluster0.ih56b.mongodb.net/?retryWrites=true&w=majority

server:
	go run main.go

.PHONY: import