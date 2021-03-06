#!/bin/bash

cmd=${1:-up}
project_name="vcon-backend"
file_prefix=""
container_name="backend"

has-docker-sync() {
  command -v docker-sync 1>/dev/null
}

execute-docker-compose () {
  opts="-f ${file_prefix}docker-compose.yml"

  if has-docker-sync; then
    opts="$opts -f ${file_prefix}docker-compose-sync.yml"
  fi

  if [ -n "$project_name" ]; then
    opts="$opts -p $project_name"
  fi

  docker-compose $opts $@
}

execute-docker-sync () {
  if ! has-docker-sync; then
    return 0
  fi

  opts="-c ${file_prefix}docker-sync.yml"
  docker-sync $@ $opts
}

stop-docker-compose () {
  execute-docker-sync stop
  execute-docker-compose stop
}

if [ $cmd = 'up' ] && [ $# -le 1 ]; then
  execute-docker-sync start
  execute-docker-compose up
  stop-docker-compose

elif [ $cmd = 'stop' ]; then
  stop-docker-compose

elif [ $cmd = 'bash-b' ]; then
  execute-docker-compose exec $container_name /bin/bash
elif [ $cmd = 'bash-p' ]; then
  execute-docker-compose exec postgres /bin/bash

elif [ $cmd = 'repl-p' ]; then
  execute-docker-compose exec postgres /bin/bash -c \
    'psql -U vcon -d vcon'

else
  execute-docker-compose $@
fi
