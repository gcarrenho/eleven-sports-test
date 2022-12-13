## Before start

You can run this software in two ways

1- Using Docker
- Install Docker

Then run
```bash
docker build -t "elevenimage" .
docker run elevenimage
```

2- Out of docker - Installing the necessary software 
- Go 1.18.x
- Move inside src/cmd
- Then run
```bash
go run main.go
```