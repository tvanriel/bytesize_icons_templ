
bytesize-icons:
	git clone https://github.com/danklammer/bytesize-icons.git
update: bytesize-icons
	go run cmd/update.go
	templ generate

