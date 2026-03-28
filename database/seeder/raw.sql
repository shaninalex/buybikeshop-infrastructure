-- partner

insert into partners.roles
    (role)
values ('contractor'),
       ('supplier');

insert into partners.partner
    (title, type)
values ('Acme Corporation', 'company'),
       ('PartsParts', 'company'),
       ('John Doe', 'person'),
       ('Shimano', 'company'),
       ('QubePartner', 'company');

insert into partners.partner_roles
    (partner_id, role_id)
values (1, 2),
       (2, 2),
       (3, 1),
       (4, 2),
       (5, 2);

insert into partners.suppliers
    (partner_id)
values (1),
       (2),
       (3),
       (4);

insert into partners.partner_contacts
    (contacts, partner_id)
values ('637 Driftwood Road\nLA PINE, Oregon(OR), 97739\n 408-420-0037', 1),
       ('2419 Crowfield Road\nQUINN, South Dakota(SD), 57775\n602-914-2117', 2),
       ('131 Echo Lane\nBellevue, Michigan(MI), 49021\n269-758-7376', 3),
       ('3354 Lighthouse Drive\nSpringfield, Missouri(MO), 65804\n417-349-3810', 4),
       ('1550 Trails End Road\nFort Lauderdale, Florida(FL), 33308\n954-384-9906', 5);

-- catalog

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
values ('catalog A', 1, 1, 'description', 'short description'),
       ('catalog B', 2, 2, 'description', 'short description');

insert into catalog.product_variants
    (product_id, title, description, sku, barcode, price)
values (1, 'catalog A (vA)', 'description', 'S9SD0-SDA-94K90', '23MSD20934MA', 12.00),
       (1, 'catalog A (vB)', 'description', 'S9SD0-SDA-94K91', '23MSD20934MB', 12.00),
       (1, 'catalog A (vC)', 'description', 'S9SD0-SDA-94K92', '23MSD20934MC', 12.00),
       (2, 'catalog B (vA)', 'description', 'DIJF8-VDB-94K90', '8DS8VLPO9KMA', 12.00),
       (2, 'catalog B (vB)', 'description', 'DIJF8-VDB-94K91', '8DS8VLPO9KMB', 12.00),
       (2, 'catalog B (vC)', 'description', 'DIJF8-VDB-94K92', '8DS8VLPO9KMC', 12.00),
       (2, 'catalog B (vD)', 'description', 'DIJF8-VDB-94K93', '8DS8VLPO9KMD', 12.00);

