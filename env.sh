project_name='payjp-grpc'
alias docker-compose="docker-compose -p $project_name"
alias build="docker-compose build"
alias up="docker-compose up"
alias stop="docker-compose stop"
alias go="docker-compose run --rm app go"
alias protoc="docker-compose run --rm app protoc"
