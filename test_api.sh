#!/bin/bash

BASE_URL="http://localhost:8080"
EMAIL="testuser$(date +%s)@example.com"
PASSWORD="password123"

echo "Testing with Email: $EMAIL"

# 1. Signup
echo "---------------------------------"
echo "1. Signup..."
curl -s -X POST "$BASE_URL/signup" -H "Content-Type: application/json" -d "{\"name\":\"Test User\", \"email\":\"$EMAIL\", \"password\":\"$PASSWORD\"}"
echo ""

# 2. Login
echo "---------------------------------"
echo "2. Login..."
LOGIN_RES=$(curl -s -X POST "$BASE_URL/login" -H "Content-Type: application/json" -d "{\"email\":\"$EMAIL\", \"password\":\"$PASSWORD\"}")
TOKEN=$(echo $LOGIN_RES | grep -o '"token":"[^"]*' | cut -d'"' -f4)

if [ -z "$TOKEN" ]; then
    echo "❌ Login failed"
    exit 1
fi
echo "✅ Logged in. Token: ${TOKEN:0:10}..."

# 3. Create Task
echo "---------------------------------"
echo "3. Creating Task..."
CREATE_RES=$(curl -s -X POST "$BASE_URL/api/tasks/" -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" -d '{"title":"My First Task","description":"Something to do","completed":false}')
echo $CREATE_RES
echo ""

# Extract Task ID (simple parsing)
TASK_ID=$(echo $CREATE_RES | grep -o '"id":"[^"]*' | cut -d'"' -f4)
if [ -z "$TASK_ID" ]; then
     # Try alternative if "id" is possibly "_id" or inside "task" object
     TASK_ID=$(echo $CREATE_RES | grep -o '"id":"[^"]*' | head -1 | cut -d'"' -f4)
fi

echo "Task ID: $TASK_ID"

# 4. Get Tasks
echo "---------------------------------"
echo "4. Get Tasks..."
curl -s -X GET "$BASE_URL/api/tasks/" -H "Authorization: Bearer $TOKEN"
echo ""

# 5. Update Task
echo "---------------------------------"
echo "5. Update Task..."
if [ ! -z "$TASK_ID" ]; then
    curl -s -X PATCH "$BASE_URL/api/tasks/$TASK_ID" -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" -d '{"completed":true}'
    echo ""
else
    echo "⚠️ Skipping update (No Task ID)"
fi

# 6. Profile
echo "---------------------------------"
echo "6. Get Profile..."
curl -s -X GET "$BASE_URL/api/profile" -H "Authorization: Bearer $TOKEN"
echo ""

# 7. Delete Task
echo "---------------------------------"
echo "7. Delete Task..."
if [ ! -z "$TASK_ID" ]; then
    curl -s -X DELETE "$BASE_URL/api/tasks/$TASK_ID" -H "Authorization: Bearer $TOKEN"
    echo ""
else
    echo "⚠️ Skipping delete (No Task ID)"
fi

echo "---------------------------------"
echo "Done."
