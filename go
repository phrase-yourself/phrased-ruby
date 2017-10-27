#!/bin/bash

set -eu

function ensure_bundle {
  if [[ ! -d vendor ]];
  then
    bundle install --path vendor/bundle
    touch vendor
  fi

  if [[ phrased.gemspec -nt vendor ]] || [[ Gemfile.lock -nt vendor ]];
  then
    bundle update
    touch vendor
  fi
}

function task_usage {
  echo "usage: $0 test"
  exit 1
}

function task_test {
  ensure_bundle
  bundle exec rubocop -f emacs
  bundle exec ruby -Ispec spec/*_spec.rb
}

args=${1:-}
shift || true
case "$args" in
  test) task_test ;;
  *) task_usage ;;
esac
