CREATE TABLE IF NOT EXISTS ingredients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    quantity INT NOT NULL
);
INSERT INTO ingredients (name, quantity) VALUES ('Crispy Cow', 100);
INSERT INTO ingredients (name, quantity) VALUES ('Crunchy Cow', 50);
INSERT INTO ingredients (name, quantity) VALUES ('Soggy Cow', 30);
INSERT INTO ingredients (name, quantity) VALUES ('Squishy Cow', 200);
INSERT INTO ingredients (name, quantity) VALUES ('Mad Cow', 75);

CREATE TABLE IF NOT EXISTS order_history (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    item_id INT NOT NULL,
    item_name VARCHAR(255) NOT NULL,
    cost FLOAT NOT NULL,
    quantity INT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO order_history (order_id, item_id, item_name, cost, quantity, timestamp) VALUES
    (1, 1, 'Crispy Cow', 7.12, 3, CURRENT_TIMESTAMP),
    (2, 2, 'Crunchy Cow', 12.49, 3, CURRENT_TIMESTAMP),
    (3, 3, 'Soggy Cow', 7.99, 3, CURRENT_TIMESTAMP),
    (4, 4, 'Squishy Cow', 5.49, 3, CURRENT_TIMESTAMP),
    (5, 5, 'Mad Cow', 8.99, 3, CURRENT_TIMESTAMP),
    (6, 1, 'Crispy Cow', 9.99, 3, CURRENT_TIMESTAMP),
    (7, 2, 'Crunchy Cow', 12.49, 3, CURRENT_TIMESTAMP),
    (8, 3, 'Soggy Cow', 7.99, 3, CURRENT_TIMESTAMP),
    (9, 4, 'Squishy Cow', 5.49, 3, CURRENT_TIMESTAMP),
    (10, 5, 'Mad Cow', 8.99, 3, CURRENT_TIMESTAMP),
    (11, 1, 'Crispy Cow', 9.99, 3, CURRENT_TIMESTAMP),
    (12, 2, 'Crunchy Cow', 12.49, 3, CURRENT_TIMESTAMP),
    (13, 3, 'Soggy Cow', 7.99, 3, CURRENT_TIMESTAMP),
    (14, 4, 'Squishy Cow', 5.49, 3, CURRENT_TIMESTAMP),
    (15, 5, 'Mad Cow', 8.99, 3, CURRENT_TIMESTAMP),
    (16, 1, 'Crispy Cow', 9.99, 3, CURRENT_TIMESTAMP),
    (17, 2, 'Crunchy Cow', 12.49, 3, CURRENT_TIMESTAMP),
    (18, 3, 'Soggy Cow', 7.99, 3, CURRENT_TIMESTAMP),
    (19, 4, 'Squishy Cow', 5.49, 3, CURRENT_TIMESTAMP),
    (20, 5, 'Mad Cow', 8.99, 3, CURRENT_TIMESTAMP),
    (21, 1, 'Crispy Cow', 9.99, 3, CURRENT_TIMESTAMP),
    (22, 2, 'Crunchy Cow', 12.49, 3, CURRENT_TIMESTAMP),
    (23, 3, 'Soggy Cow', 7.99, 3, CURRENT_TIMESTAMP),
    (24, 4, 'Squishy Cow', 5.49, 3, CURRENT_TIMESTAMP),
    (25, 5, 'Mad Cow', 8.99, 3, CURRENT_TIMESTAMP),
    (26, 1, 'Crispy Cow', 9.99, 3, CURRENT_TIMESTAMP),
    (27, 2, 'Crunchy Cow', 12.49, 3, CURRENT_TIMESTAMP),
    (28, 3, 'Soggy Cow', 7.99, 3, CURRENT_TIMESTAMP),
    (29, 4, 'Squishy Cow', 5.49, 3, CURRENT_TIMESTAMP),
    (30, 5, 'Mad Cow', 8.99, 3, CURRENT_TIMESTAMP),
    (31, 1, 'Crispy Cow', 9.99, 3, CURRENT_TIMESTAMP),
    (32, 2, 'Crunchy Cow', 12.49, 3, CURRENT_TIMESTAMP),
    (33, 3, 'Soggy Cow', 7.99, 3, CURRENT_TIMESTAMP),
    (34, 4, 'Squishy Cow', 5.49, 3, CURRENT_TIMESTAMP),
    (35, 5, 'Mad Cow', 8.99, 3, CURRENT_TIMESTAMP),
    (36, 1, 'Crispy Cow', 9.99, 3, CURRENT_TIMESTAMP),
    (37, 2, 'Crunchy Cow', 12.49, 3, CURRENT_TIMESTAMP),
    (38, 3, 'Soggy Cow', 7.99, 3, CURRENT_TIMESTAMP),
    (39, 4, 'Squishy Cow', 5.49, 3, CURRENT_TIMESTAMP),
    (40, 5, 'Mad Cow', 8.99, 3, CURRENT_TIMESTAMP),
    (41, 1, 'Crispy Cow', 9.99, 3, CURRENT_TIMESTAMP),
    (42, 2, 'Crunchy Cow', 12.49, 3, CURRENT_TIMESTAMP),
    (43, 3, 'Soggy Cow', 7.99, 3, CURRENT_TIMESTAMP),
    (44, 4, 'Squishy Cow', 5.49, 3, CURRENT_TIMESTAMP),
    (45, 5, 'Mad Cow', 8.99, 3, CURRENT_TIMESTAMP),
    (46, 1, 'Crispy Cow', 9.99, 3, CURRENT_TIMESTAMP),
    (47, 2, 'Crunchy Cow', 12.49, 3, CURRENT_TIMESTAMP),
    (48, 3, 'Soggy Cow', 7.99, 3, CURRENT_TIMESTAMP),
    (49, 4, 'Squishy Cow', 5.49, 3, CURRENT_TIMESTAMP),
    (50, 5, 'Mad Cow', 8.99, 3, CURRENT_TIMESTAMP);