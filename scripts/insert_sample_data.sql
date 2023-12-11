INSERT INTO buildings (id, name) VALUES
  ('f47ee4b9-826d-4c8d-8c2c-6880cc1f68c2', 'Building A'),
  ('e11c5e21-b85d-4f0f-8693-285b9eb14847', 'Building B');

INSERT INTO toilets (id, name, score, location, image, building_id) VALUES
  ('d7380e1f-3d7e-4c3d-b42d-04a3dfc4f6fb', 'Toilet 1', 4, 'Location A', 'image1.jpg', 'f47ee4b9-826d-4c8d-8c2c-6880cc1f68c2'),
  ('a2b6f4d7-8f6a-4e9c-86a2-5e2f4d668c50', 'Toilet 2', 3, 'Location B', 'image2.jpg', 'f47ee4b9-826d-4c8d-8c2c-6880cc1f68c2'),
  ('9d9a6324-798b-4f0e-a90b-7a29ce4c7a62', 'Toilet 3', 5, 'Location C', 'image3.jpg', 'e11c5e21-b85d-4f0f-8693-285b9eb14847'),
  ('b32cf9e3-10a9-496e-a4c8-d7ddef4646d5', 'Toilet 4', 2, 'Location D', 'image4.jpg', 'e11c5e21-b85d-4f0f-8693-285b9eb14847');

INSERT INTO reviews (id, username, comment, score, toilet_id) VALUES
  ('1c6da665-1ccf-4cfb-96a8-f1a4f3f5c663', 'User1', 'Good toilet!', 5, 'd7380e1f-3d7e-4c3d-b42d-04a3dfc4f6fb'),
  ('f1f7191c-0b9a-40aa-aa5d-9dd7870ef507', 'User2', 'Could be cleaner', 3, 'd7380e1f-3d7e-4c3d-b42d-04a3dfc4f6fb'),
  ('d91a8e3d-6d53-45a1-b993-79bc7a482f0b', 'User3', 'Average experience', 4, 'a2b6f4d7-8f6a-4e9c-86a2-5e2f4d668c50'),
  ('cc0a4502-10eb-4cf3-99cf-44e0cf2fde3f', 'User4', 'Excellent facilities', 5, '9d9a6324-798b-4f0e-a90b-7a29ce4c7a62');
