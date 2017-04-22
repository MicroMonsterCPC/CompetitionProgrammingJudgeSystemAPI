#!/bin/sh
curl \
    -X POST \
    http://localhost:1323/answer-data \
    -H 'Content-Type: application/json' \
    -d '{"question_id": "1", "answer_data": "puts gets.chomp", "lang": "ruby"}'

# curl \
#     -X POST \
#     http://localhost:1323/answer-data \
#     -H 'Content-Type: application/json' \
#     -d '{"question_id": "2", "answer_data": "puts (1..gets.to_i).to_a.reduce(:+)", "lang": "ruby"}'
