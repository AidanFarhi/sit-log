-- Insert data into the user table
INSERT INTO user (username, email, password) VALUES ('Alice', 'alice@example.com', 'password1');
INSERT INTO user (username, email, password) VALUES ('Bob', 'bob@example.com', 'password2');
INSERT INTO user (username, email, password) VALUES ('Carol', 'carol@example.com', 'password3');

-- Insert data into the child table
INSERT INTO child (name) VALUES ('Charlie');
INSERT INTO child (name) VALUES ('Daisy');
INSERT INTO child (name) VALUES ('Evan');

-- Insert data into the user_child_relation table
INSERT INTO user_child_relation (user_id, child_id, relationship) VALUES (1, 1, 'parent');
INSERT INTO user_child_relation (user_id, child_id, relationship) VALUES (2, 2, 'parent');
INSERT INTO user_child_relation (user_id, child_id, relationship) VALUES (1, 2, 'parent'); -- Alice has two children
INSERT INTO user_child_relation (user_id, child_id, relationship) VALUES (3, 3, 'sitter'); -- Carol is the sitter for Evan

-- Insert data into the event table
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (1, 'play', 'Playing at the park', '10:00:00', '11:00:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (2, 'study', 'Math tutoring session', '15:00:00', '16:30:00', '01:30:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (3, 'nap', 'Afternoon nap', '13:00:00', '14:00:00', '01:00:00');

-- Insert more events for Charlie (child_id = 1)
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (1, 'play', 'Playing soccer with friends', '11:30:00', '12:30:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (1, 'study', 'Reading session with parent', '14:00:00', '15:00:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (1, 'play', 'Bike riding in the park', '16:00:00', '17:00:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (1, 'play', 'Building sandcastles at the beach', '10:00:00', '11:00:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (1, 'study', 'Science homework session', '17:00:00', '18:00:00', '01:00:00');

-- Insert more events for Daisy (child_id = 2)
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (2, 'study', 'Reading comprehension practice', '09:00:00', '10:00:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (2, 'play', 'Playing with dolls', '11:00:00', '12:00:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (2, 'study', 'Writing exercise', '13:30:00', '14:30:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (2, 'play', 'Art and craft session', '15:30:00', '16:30:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (2, 'study', 'Math revision with parent', '17:00:00', '18:00:00', '01:00:00');

-- Insert more events for Evan (child_id = 3)
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (3, 'nap', 'Morning nap', '09:00:00', '10:00:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (3, 'play', 'Building blocks', '11:00:00', '12:00:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (3, 'nap', 'Afternoon nap', '13:00:00', '14:00:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (3, 'study', 'Interactive learning with parent', '14:30:00', '15:30:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (3, 'play', 'Playing with toys', '16:00:00', '17:00:00', '01:00:00');

-- Insert more random events for various children
-- For Charlie (child_id = 1)
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (1, 'study', 'Geography revision', '10:00:00', '11:00:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (1, 'play', 'Football practice', '14:00:00', '15:00:00', '01:00:00');
-- For Daisy (child_id = 2)
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (2, 'study', 'History revision', '09:00:00', '10:00:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (2, 'play', 'Gardening with parent', '11:30:00', '12:30:00', '01:00:00');
-- For Evan (child_id = 3)
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (3, 'nap', 'Post-lunch nap', '12:00:00', '13:00:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (3, 'study', 'Reading storybook', '15:30:00', '16:30:00', '01:00:00');

-- Insert events for different children
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (1, 'play', 'Fishing at the lake', '08:00:00', '09:00:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (2, 'study', 'English grammar practice', '10:30:00', '11:30:00', '01:00:00');
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES (3, 'play', 'Running at the playground', '13:00:00', '14:00:00', '01:00:00');

-- Insert valid sessions into the session table
INSERT INTO session (token, username) VALUES ('abc123sessiontoken', 'Alice');
INSERT INTO session (token, username) VALUES ('def456sessiontoken', 'Bob');
