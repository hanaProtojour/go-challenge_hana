set -e

echo "=== Test GET ==="
curl -X GET -s localhost:3333 | rg -q "get" 

echo "=== Test POST ==="
curl -X POST -s localhost:3333 | rg -q "post"

echo "=== Test Post ==="
curl -iX POST localhost:3333 -d \
`{"seeds":["abc", "def", "xyz"]}`


echo "=== Test error in JSON Post"
curl -iX POST localhost:3333 -d \
`{"seeds":["abc", "def", "xyz"]`

echo "Success"