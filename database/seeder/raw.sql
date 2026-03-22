insert into catalog.brands
    (title)
values ('Cube'),
       ('Azimut'),
       ('Camanche'),
       ('Author');

insert into catalog.categories
    (title)
values ('MTB'),
       ('City'),
       ('Sport'),
       ('Child');

insert into catalog.products
    (title, brand_id, category_id, description, short_description)
values ('product A', 1, 1, 'description', 'short description'),
       ('product B', 2, 2, 'description', 'short description');

insert into catalog.product_variants
    (product_id, title, description, sku, barcode, price)
values (1, 'product A (vA)', 'description', 'S9SD0-SDA-94K90', '23MSD20934MA', 12.00),
       (1, 'product A (vB)', 'description', 'S9SD0-SDA-94K91', '23MSD20934MB', 12.00),
       (1, 'product A (vC)', 'description', 'S9SD0-SDA-94K92', '23MSD20934MC', 12.00),
       (2, 'product B (vA)', 'description', 'DIJF8-VDB-94K90', '8DS8VLPO9KMA', 12.00),
       (2, 'product B (vB)', 'description', 'DIJF8-VDB-94K91', '8DS8VLPO9KMB', 12.00),
       (2, 'product B (vC)', 'description', 'DIJF8-VDB-94K92', '8DS8VLPO9KMC', 12.00),
       (2, 'product B (vD)', 'description', 'DIJF8-VDB-94K93', '8DS8VLPO9KMD', 12.00);

