insert into psychologist (id, name, description)
values ('97d6a971-2ea6-499a-84f9-3b0ff360876e', 'Adam Smith', 'Practitioner with 14 years of experience, help families to achieve financial independence'),
       ('c28284ba-c9f6-4e05-8794-740f7304e4fb', 'Thomas Hobbes', 'Advanced qualification in interpersonal relations'),
       ('1485cdaf-092e-4198-8b81-40a1aec15322', 'Michel de Montaigne', 'Primary field: creation of yourself as success person'),
       ('daa2d032-0ab9-4275-a5f7-b643a4b02365', 'Albert Camus', 'Reliable psychologists that helps you to accept yourself and reality')
on conflict (id) do update
set name = excluded.name,
    description = excluded.description;

insert into course (id, name, description, price)
values ('71d53f3c-b904-4cac-b3c5-ceb994d882e9', '5 simple steps to achieve financial independence', 'This course will help you to become more efficient in financial planning', 5400),--Adam Smith
       ('a14c6a54-7adf-4eff-842e-93891b72b20f', 'Aspects of poor people', 'With this course you will understand how are different viewpoints of poor and rich people', 9800),--Adam Smith
       ('34906cbf-9e19-41c1-bb95-8108cb3501f7', 'You are not alone', 'This material teaches customers to achieve money goals involving nearby people', 1400),--Adam Smith,Thomas Hobbes
       ('c1162fb5-504a-406e-9886-75649cd97e12', 'Internal politic', 'With such knowledge you can became a master of influence', 4300),--Thomas Hobbes
       ('eb9fa6ec-e9e4-4105-96c5-f740378a3885', 'How to become a snail, or... dragon', 'This course lets you to open mind and become whatever you want', 3400),--Michel de Montaigne
       ('b5d198c5-1d71-40a3-b4e5-520dc1cc25c7', 'Understanding yourself', 'Have you already understood yourself? Nah, it''s impossible, but how far can you go?', 4600),--Michel de Montaigne,Albert Camus
       ('6b5de104-337b-44cb-acec-dd8e0ca800c6', 'Accepting reality', 'This masterpiece will immerse you into new old world', 2600)--Albert Camus
on conflict (id) do update
set name = excluded.name,
    description = excluded.description,
    price = excluded.price;

insert into courses_psychologists (course, psychologist)
values
    --Smith
    ('71d53f3c-b904-4cac-b3c5-ceb994d882e9', '97d6a971-2ea6-499a-84f9-3b0ff360876e'),
    ('a14c6a54-7adf-4eff-842e-93891b72b20f', '97d6a971-2ea6-499a-84f9-3b0ff360876e'),
    ('34906cbf-9e19-41c1-bb95-8108cb3501f7', '97d6a971-2ea6-499a-84f9-3b0ff360876e'),
    --Hobbes
    ('34906cbf-9e19-41c1-bb95-8108cb3501f7', 'c28284ba-c9f6-4e05-8794-740f7304e4fb'),
    ('c1162fb5-504a-406e-9886-75649cd97e12', 'c28284ba-c9f6-4e05-8794-740f7304e4fb'),
    --Montaigne
    ('eb9fa6ec-e9e4-4105-96c5-f740378a3885', '1485cdaf-092e-4198-8b81-40a1aec15322'),
    ('b5d198c5-1d71-40a3-b4e5-520dc1cc25c7', '1485cdaf-092e-4198-8b81-40a1aec15322'),
    --Camus
    ('b5d198c5-1d71-40a3-b4e5-520dc1cc25c7', 'daa2d032-0ab9-4275-a5f7-b643a4b02365'),
    ('6b5de104-337b-44cb-acec-dd8e0ca800c6', 'daa2d032-0ab9-4275-a5f7-b643a4b02365')
on conflict (course, psychologist) do nothing;

insert into lesson (id, name, number, course)
values
    --5 simple steps to achieve financial independence
    ('1f34f3fa-da88-463a-9478-c027d3799d79', 'Cut costs', 1, '71d53f3c-b904-4cac-b3c5-ceb994d882e9'),
    ('cd645214-f721-44dc-bc3d-526f647d5a99', 'Make investments', 2, '71d53f3c-b904-4cac-b3c5-ceb994d882e9'),
    --Aspects of poor people
    ('e4525d15-7be8-469f-a2b8-4557ca91dd13', 'Origin of poorness', 1, 'a14c6a54-7adf-4eff-842e-93891b72b20f'),
    ('58ee8ac8-8edd-4d33-8b4a-da66ce0d4bd2', 'How to get rid of bad habits', 2, 'a14c6a54-7adf-4eff-842e-93891b72b20f'),
    --You are not alone
    ('d72f565f-3911-4664-9271-f501722c0bc8', 'Look around', 1, '34906cbf-9e19-41c1-bb95-8108cb3501f7'),
    ('3727e0a9-2959-4f04-9e6c-e0497244feca', 'Make contacts', 2, '34906cbf-9e19-41c1-bb95-8108cb3501f7'),
    --Internal politic
    ('0f5d6927-297a-4b80-a481-7a3614f16778', 'Good and bad sides of machiavellianism', 1, 'c1162fb5-504a-406e-9886-75649cd97e12'),
    ('db3b412d-ba51-49f2-9d7f-66cdd36ac663', 'Learn to reassure people', 2, 'c1162fb5-504a-406e-9886-75649cd97e12'),
    --How to become a snail, or... dragon
    ('7b21f831-3b1e-4f14-8af2-cb0d4fc7da3a', 'You don''t know yourself', 1, 'eb9fa6ec-e9e4-4105-96c5-f740378a3885'),
    ('38c97134-2a28-45f0-826d-e68126a7a60a', 'How to understand yourself and not harm others', 2, 'eb9fa6ec-e9e4-4105-96c5-f740378a3885'),
    --Understanding yourself
    ('66a8f5dc-b148-49ff-9899-7b5eba36bfec', 'Have you ever had the past?', 1, 'b5d198c5-1d71-40a3-b4e5-520dc1cc25c7'),
    ('408a26c1-7402-4086-a911-288cd29e4516', 'You is you and no one another', 2, 'b5d198c5-1d71-40a3-b4e5-520dc1cc25c7'),
    --Accepting reality
    ('c70d72d9-b8f1-4062-9b57-dcd30c67780e', 'Lone stranger', 1, '6b5de104-337b-44cb-acec-dd8e0ca800c6'),
    ('d500503c-9d5a-4225-ba83-f1e3ed0694dc', 'What do we do with huge rocks?', 2, '6b5de104-337b-44cb-acec-dd8e0ca800c6')
on conflict (id) do update
set name = excluded.name,
    number = excluded.number,
    course = excluded.course;
