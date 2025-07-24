#!/bin/bash

echo "🧪 Testing Quiz API Endpoints"
echo "=============================="

BASE_URL="http://localhost:8080"

echo ""
echo "1. Health Check:"
curl -s -X GET $BASE_URL/health | jq .

echo ""
echo "2. Get All Questions:"
curl -s -X GET $BASE_URL/api/questions | jq '.[0:2]'  # Show first 2 questions

echo ""
echo "3. Get Question by ID:"
curl -s -X GET $BASE_URL/api/questions/1 | jq .

echo ""
echo "4. Submit Perfect Quiz (100%):"
curl -s -X POST $BASE_URL/api/quiz/submit \
  -H "Content-Type: application/json" \
  -d '{
    "answers": [
      {"questionId": 1, "answer": 2},
      {"questionId": 2, "answer": 1},
      {"questionId": 3, "answer": 1}
    ]
  }' | jq .

echo ""
echo "5. Submit Failing Quiz (33%):"
curl -s -X POST $BASE_URL/api/quiz/submit \
  -H "Content-Type: application/json" \
  -d '{
    "answers": [
      {"questionId": 1, "answer": 0},
      {"questionId": 2, "answer": 1},
      {"questionId": 3, "answer": 2}
    ]
  }' | jq .

echo ""
echo "6. Test Error Handling - Invalid Question ID:"
curl -s -X POST $BASE_URL/api/quiz/submit \
  -H "Content-Type: application/json" \
  -d '{
    "answers": [
      {"questionId": 999, "answer": 1}
    ]
  }' | jq .

echo ""
echo "✅ All tests completed!"
