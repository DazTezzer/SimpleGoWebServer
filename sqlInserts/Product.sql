-- Вставка данных в таблицу ProductDescription
INSERT INTO ProductDescription (color, material, gender, additional_info) VALUES
('Red', 'Cotton', 'Unisex', 'Soft and comfortable'),
('Blue', 'Polyester', 'Men', 'Durable for outdoor activities'),
('Green', 'Silk', 'Women', 'Elegant and stylish');

-- Вставка данных в таблицу ProductCategory
INSERT INTO ProductCategory (name) VALUES
('Clothing'),
('Accessories'),
('Footwear');

-- Вставка данных в таблицу ProductImage
INSERT INTO ProductImage (name, image) VALUES
('Red T-Shirt', 'https://example.com/images/red_tshirt.jpg'),
('Blue Jeans', 'https://example.com/images/blue_jeans.jpg'),
('Green Scarf', 'https://example.com/images/green_scarf.jpg');

-- Вставка данных в таблицу Products
INSERT INTO Products (name, image_id, price, description_id, category_id) VALUES
('Red T-Shirt', 1, 19.99, 1, 1),
('Blue Jeans', 2, 49.99, 2, 1),
('Green Scarf', 3, 15.99, 3, 2);

INSERT INTO Size (product_id, size) VALUES
(1,'M'),
(1,'L'),
(1,'S'),
(2,'M'),
(2,'L'),
(2,'XL');

SELECT * FROM ProductImage

SELECT * FROM Size

SELECT p.name, s.size 
FROM Products as p
JOIN Size as s ON s.product_id = p.id;