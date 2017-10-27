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

function current_version {
  ruby -Ilib -rphrased/version -e 'puts Phrased::VERSION'
}

function task_usage {
  echo "usage: $0 test | build | release"
  exit 1
}

function task_test {
  ensure_bundle
  bundle exec rubocop -f emacs
  bundle exec ruby -Ispec spec/*_spec.rb
}

function task_build {
  ensure_bundle
  local version
  version=$(current_version)

  mkdir -p build
  gem build phrased.gemspec
  mv "phrased-${version}.gem" build
}

function task_release {
  local version
  version=$(current_version)

  if gem search phrased |grep -q "phrased (${version})$";
  then
    echo "phrased (${version}) already released, doing nothing.."
  else
    echo "Releasing phrased (${version})"
    gem release "build/phrased-${version}.gem"
  fi
}

args=${1:-}
shift || true
case "$args" in
  test) task_test ;;
  build) task_build ;;
  release) task_release ;;
  *) task_usage ;;
esac
