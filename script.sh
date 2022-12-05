#!/bin/bash
for i in {1..9}
do
   mkdir day0$i
   touch day0$i/day0${i}_A.go day0$i/day0${i}_B.go day0$i/input.txt
done
for i in {10..25}
do
   mkdir day$i
   touch day$i/day${i}_A.go day$i/day${i}_B.go day$i/input.txt
done
