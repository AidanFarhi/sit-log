-- Seed data for 'adult' table
INSERT INTO adult (name, email, password) VALUES 
('Alice Johnson', 'alice.j@example.com', 'password123'),
('Bob Smith', 'bob.smith@example.com', 'mypassword'),
('Carol White', 'carol.w@example.com', 'securepass'),
('David Brown', 'david.b@example.com', 'safepass');

-- Seed data for 'child' table
INSERT INTO child (name, birthday) VALUES
('Emily Johnson', '2010-05-15'),
('Jake Smith', '2012-09-22'),
('Mia White', '2014-07-30'),
('Sophia Johnson', '2016-11-03');

-- Seed data for 'adult_child_relation' table
INSERT INTO adult_child_relation (adult_id, child_id, relationship) VALUES
(1, 1, 'Parent'),  -- Alice is Emily's parent
(1, 4, 'Parent'),  -- Alice is Sophia's parent
(2, 2, 'Parent'),  -- Bob is Jake's parent
(3, 3, 'Parent'),  -- Carol is Mia's parent
(4, 1, 'Sitter'),  -- David is Emily's sitter
(4, 3, 'Sitter');  -- David is Mia's sitter

-- Seed data for 'event' table
INSERT INTO event (child_id, type, description, start_time, end_time, duration) VALUES
(1, 'Playdate', 'Emily had a playdate with her friends', '2024-10-20 14:00', '2024-10-20 16:00', '02:00'),
(2, 'Homework', 'Jake finished his math homework', '2024-10-19 17:00', '2024-10-19 18:30', '01:30'),
(3, 'Swimming Lesson', 'Mia attended her weekly swimming lesson', '2024-10-21 10:00', '2024-10-21 11:00', '01:00'),
(4, 'Doctor Appointment', 'Sophia had a doctor checkup', '2024-10-22 09:00', '2024-10-22 09:45', '00:45');
