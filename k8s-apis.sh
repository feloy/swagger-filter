#!/bin/bash

mkdir -p output

cat k8s-apis.txt | while read i j
do
    go run . --endpoint-prefix $j swagger.json output/$i.json
done