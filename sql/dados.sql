INSERT INTO usuarios (id, nome, nick, email, senha) VALUES
    (1, 'Adriano', 'adriano', 'adriano@devbook.com', '$2a$10$fw8EcFYm5VkI/VNpSAM9T..N8Jh17h/TKp.BZGZg2sKDw4Kk5LdYK'),
    (2, 'Pedro', 'pedro', 'pedro@devbook.com', '$2a$10$yDq/SV85Rmu6PdOlhYKehObwEe6Wd3YNkssF6O1aT9jvApatwPRMy'),
    (3, 'Lizete', 'lizete', 'lizete@devbook.com', '$2a$10$Dnsxpi14fZ9pnuNmV7gIJeT/JP2kNyeMUuWTtXN5QyxRfXGFnxR/O'),
    (4, 'Márcia', 'marcia', 'marcia@devbook.com', '$2a$10$3qRwWpCjJDLz47XMqGwq8OjVmkLW7.iOhu.hzn.fILG.t4zWcOYXq'),
    (5, 'Sandra', 'sandra', 'sandra@devbook.com', '$2a$10$3yaV1kVLYg4xUH24gCN63umzJXQwKDotQt.1YURP4qaMCVcQkOOcm');

-- usuario_id é o usuário seguido; seguidor_id é quem o segue.
INSERT INTO seguidores (usuario_id, seguidor_id) VALUES
    (1, 2), -- Pedro segue Adriano
    (1, 3), -- Lizete segue Adriano
    (2, 1), -- Adriano segue Pedro
    (2, 4), -- Márcia segue Pedro
    (3, 5), -- Sandra segue Lizete
    (4, 1), -- Adriano segue Márcia
    (5, 3); -- Lizete segue Sandra
