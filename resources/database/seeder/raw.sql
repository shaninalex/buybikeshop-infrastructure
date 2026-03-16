insert into catalog.brands (title)
values ('test');
insert into catalog.categories (title)
values ('test');
insert into catalog.collections (name)
values ('test');

insert into catalog.products
(title, collection_id, brand_id, category_id, description, short_description)
values ('product A', 1, 1, 1, 'description', 'short description'),
       ('product B', 1, 1, 1, 'description', 'short description');
