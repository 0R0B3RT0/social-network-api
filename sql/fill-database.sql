insert into users (name, nick, email, pass) values
('Usuário um'  , 'usuarioum'  , 'usuario-um@gmail.com'  , '$2a$10$swqWzMjISIsQdkLEeCnsFuAkIeT5h2L3vaVpmugkU9g/HhvZlPncy'),
('Usuário dois', 'usuariodois', 'usuario-dois@gmail.com', '$2a$10$swqWzMjISIsQdkLEeCnsFuAkIeT5h2L3vaVpmugkU9g/HhvZlPncy'),
('Usuário três', 'usuariotres', 'usuario-tres@gmail.com', '$2a$10$swqWzMjISIsQdkLEeCnsFuAkIeT5h2L3vaVpmugkU9g/HhvZlPncy');

insert into followers (follower_id, following_id) values
(1, 2),
(1, 3),
(3, 1);

insert into publications(user_id, title, content, likes) values
(1, 'Primeira publicação', 'Essa é a primeira de muitas publicações', 1),
(3, 'Segunda publicação', 'Essa é a primeira de muitas publicações do usuário três', 3);