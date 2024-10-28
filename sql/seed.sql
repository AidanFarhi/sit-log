-- Insert data into the adult table
INSERT INTO adult (name, email, password) VALUES ('Alice', 'alice@example.com', 'password1');
INSERT INTO adult (name, email, password) VALUES ('Bob', 'bob@example.com', 'password2');
INSERT INTO adult (name, email, password) VALUES ('Carol', 'carol@example.com', 'password3');

-- Insert data into the child table
INSERT INTO child (name) VALUES ('Charlie');
INSERT INTO child (name) VALUES ('Daisy');
INSERT INTO child (name) VALUES ('Evan');

-- Insert data into the adult_child_relation table
INSERT INTO adult_child_relation (adult_id, child_id, relationship) VALUES (1, 1, 'parent');
INSERT INTO adult_child_relation (adult_id, child_id, relationship) VALUES (2, 2, 'parent');
INSERT INTO adult_child_relation (adult_id, child_id, relationship) VALUES (1, 2, 'parent'); -- Alice has two children
INSERT INTO adult_child_relation (adult_id, child_id, relationship) VALUES (3, 3, 'sitter'); -- Carol is the sitter for Evan

-- Insert data into the event table
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (1, 'play', 'Playing at the park', '10:00:00', '11:00:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (2, 'study', 'Math tutoring session', '15:00:00', '16:30:00', '01:30:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (3, 'nap', 'Afternoon nap', '13:00:00', '14:00:00', '01:00:00');
