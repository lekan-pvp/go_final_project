#/usr/bin/bash

cd frontend
npm run build

cd ..
go build -o ./go1f

mv ./web/js/style.css ./web/css

./go1f

# sudo adduser $USER docker
# docker build -t scheduler .
# docker run -it --rm -p 7541:7540 scheduler
# docker run -it --rm -e TODO_PORT=$TODO_PORT -p $TODO_PORT:$TODO_PORT scheduler
# docker run -it --rm -e TODO_PORT=$TODO_PORT -p $TODO_PORT:$TODO_PORT -v /home/ak/go/github.com/gentee/go1f-task/scheduler.db:/usr/local/app/scheduler.db scheduler
# docker run -it --rm -e TODO_PORT=$TODO_PORT -e TODO_PASSWORD=123 -p $TODO_PORT:$TODO_PORT -v /home/ak/go/github.com/gentee/go1f-task/scheduler.db:/usr/local/app/scheduler.db scheduler
