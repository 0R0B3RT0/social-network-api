echo "########## INSERT TWO NEW USERS ##########\n"
curl -i -X POST localhost:5000/users --header "Content-Type: application/json" --data '{"name":"Fulano um", "nick":"apelido","email":"fulano-um@gmail.com","pass":"skdjalskj"}'
echo "\n----------------------------------\n\n"

curl -i -X POST localhost:5000/users --header "Content-Type: application/json" --data '{"name":"Fulano dois", "nick":"apelido2","email":"fulano-dois@gmail.com","pass":"sdassa"}'
echo "\n----------------------------------\n\n"

echo "########## UPDATE THE USER ##########\n"
curl -i -X PUT localhost:5000/users/1 --header "Content-Type: application/json" --data '{"name":"Usuário de nome novo", "nick":"apelido novo38", "email":"usuario.novo@gmail.com"}'
echo "\n----------------------------------\n\n"

echo "########## FIND ALL USERS ##########\n"
curl -i -X GET localhost:5000/users?user=fulano --header "Content-Type: application/json"
echo "\n----------------------------------\n\n"

echo "########## FIND ONE USER ##########\n"
curl -i -X GET localhost:5000/users/4 --header "Content-Type: application/json"
echo "\n----------------------------------\n\n"

# echo "########## REMOVE THE USER ##########\n"
# curl -i -X DELETE localhost:5000/users/1 --header "Content-Type: application/json"
# echo "\n----------------------------------\n\n"