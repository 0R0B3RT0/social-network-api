echo "########## INSERT TWO NEW USERS ##########\n"
curl -i -X POST localhost:5000/usuarios --header "Content-Type: application/json" --data '{"id":123,"nome":"Fulano um","email":"fulano@gmail.com"}'
echo "\n----------------------------------\n\n"

curl -i -X POST localhost:5000/usuarios --header "Content-Type: application/json" --data '{"id":123,"nome":"Fulano dois","email":"fulano-dois@gmail.com"}'
echo "\n----------------------------------\n\n"

echo "########## UPDATE THE USER ##########\n"
curl -i -X PUT localhost:5000/usuarios/1 --header "Content-Type: application/json" --data '{"nome":"Joãozinho","email":"joaozinho@gmail.com"}'
echo "\n----------------------------------\n\n"

echo "########## FIND ALL USERS ##########\n"
curl -i -X GET localhost:5000/usuarios --header "Content-Type: application/json"                                                                
echo "\n----------------------------------\n\n"

echo "########## FIND ONE USER ##########\n"
curl -i -X GET localhost:5000/usuarios/1 --header "Content-Type: application/json"
echo "\n----------------------------------\n\n"

echo "########## REMOVE THE USER ##########\n"
curl -i -X DELETE localhost:5000/usuarios/1 --header "Content-Type: application/json"
echo "\n----------------------------------\n\n"