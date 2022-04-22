INSERT INTO users (id, name, nick, email, passwd)
VALUES
(DEFAULT, "Usuario 1", "usuario 1", "usuario1@gmail.com", "$2a$10$.KEXtj7voYapChfqcqv/WukxqIIcJDwFnQLS4WjcCfW3eigvI5guO"),
(DEFAULT, "Usuario 2", "usuario 2", "usuario2@gmail.com", "$2a$10$.KEXtj7voYapChfqcqv/WukxqIIcJDwFnQLS4WjcCfW3eigvI5guO"),
(DEFAULT, "Usuario 3", "usuario 3", "usuario3@gmail.com", "$2a$10$.KEXtj7voYapChfqcqv/WukxqIIcJDwFnQLS4WjcCfW3eigvI5guO"),
(DEFAULT, "Usuario 4", "usuario 4", "usuario4@gmail.com", "$2a$10$.KEXtj7voYapChfqcqv/WukxqIIcJDwFnQLS4WjcCfW3eigvI5guO"),
(DEFAULT, "Usuario 5", "usuario 5", "usuario5@gmail.com", "$2a$10$.KEXtj7voYapChfqcqv/WukxqIIcJDwFnQLS4WjcCfW3eigvI5guO"),
(DEFAULT, "Usuario 6", "usuario 6", "usuario6@gmail.com", "$2a$10$.KEXtj7voYapChfqcqv/WukxqIIcJDwFnQLS4WjcCfW3eigvI5guO"),
(DEFAULT, "Usuario 7", "usuario 7", "usuario7@gmail.com", "$2a$10$.KEXtj7voYapChfqcqv/WukxqIIcJDwFnQLS4WjcCfW3eigvI5guO");

INSERT INTO followers(user_id, follower_id)
VALUES
(1, 2),
(3, 1),
(1, 3);

INSERT INTO posts(title, content, autor_id)
VALUES
("Publicação do usuário 1", "Essa é a publicação 1", 1),
("Publicação do usuário 2", "Essa é a publicação 2", 2),
("Publicação do usuário 3", "Essa é a publicação 3", 3);