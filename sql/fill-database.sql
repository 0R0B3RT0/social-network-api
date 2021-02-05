insert into users (name, nick, email, pass) values
('Usuário um'  , 'usuarioum'  , 'usuario-um@gmail.com'  , '$2a$10$swqWzMjISIsQdkLEeCnsFuAkIeT5h2L3vaVpmugkU9g/HhvZlPncy'),
('Usuário dois', 'usuariodois', 'usuario-dois@gmail.com', '$2a$10$swqWzMjISIsQdkLEeCnsFuAkIeT5h2L3vaVpmugkU9g/HhvZlPncy'),
('Usuário três', 'usuariotres', 'usuario-tres@gmail.com', '$2a$10$swqWzMjISIsQdkLEeCnsFuAkIeT5h2L3vaVpmugkU9g/HhvZlPncy');

insert into followers (follower_id, followed_id) values
(1, 2),
(1, 3),
(3, 1);