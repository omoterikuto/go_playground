#!/bin/bash
paths=($(ls -d jobs/*))

for path in "${paths[@]}"; do
  module=$(basename "$path")
  docker build -f Dockerfile -t "${module}" --build-arg MODULE=${module} .
done
