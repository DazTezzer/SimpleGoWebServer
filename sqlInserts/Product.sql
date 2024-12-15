-- Вставка данных в таблицу ProductDescription
INSERT INTO ProductDescription (color, material, gender, additional_info) VALUES
('Черный', 'Хлопок', 'Мужской', 'Удобная и стильная шляпа, придающая стиль. Она всем как раз!'),
('Черный', 'Полиэстер', 'Мужской', 'Стильный черный худи из мягкого материала идеально подходит для повседневной носки, сочетая комфорт и универсальность'),
('Коричневый', 'Кожа', 'Мужской', 'Качественные зимние ботинки для холодной погоды'),
('Белый', 'Неопрен', 'Мужской', 'Куртка с принтом скорпиона сзади придает загадочность вашему стилю. Отлично подойдет для повседневной носки и ночных поездок на машине'),
('Синий', 'Деним', 'Мужской', 'Классические синие джинсы с комфортной посадкой'),
('Белый', 'Хлопок', 'Мужской', 'Футболка с самураем для самых настоящих воинов.')

-- Вставка данных в таблицу ProductCategory
INSERT INTO ProductCategory (name) VALUES
('Футболки'),
('Куртки'),
('Штаны'),
('Головные уборы'),
('Обувь'),
('Кофты')

-- Вставка данных в таблицу ProductImage
INSERT INTO ProductImage (name, image) VALUES
('Белая кепка', '/images/whiteCap/whiteCap.jpg'),
('Коричневая кожаная куртка', '/images/brownLeatherJacket/brownLeatherJacket.jpg'),
('Кремовые брюки', '/images/creamPants/creamPants.jpg'),
('Черная футболка', '/images/blackTshirt/blackTshirt.jpg'),
('Серый свитшот', '/images/greySweatshort/greySweatshort.jpg'),
('Синие кроссовки', '/images/blueSneakers/blueSneakers.jpg');

('Футболка с самураем','/images/samuraiT-shirt/samuraiT-Shirt.jpg'),
('Шляпа', '/images/blackHat/blackHat.jpg'),
('Худи', '/images/blackHoodie/blackHoodie.jpg'),
('Зимние ботинки', '/images/brownWinterBoots/brownWinterBoots.jpg'),
('Куртка со скорпионом', '/images/scorpionJacket/scorpionJacket.jpg'),
('Джинсы', '/images/blueJeans/blueJeans.jpg');

SELECT * FROM ProductImage
DELETE FROM ProductImage;
DELETE FROM ProductImage
WHERE id BETWEEN 7 AND 12;

-- Сначала временно изменяем id на отрицательные числа или большие числа
UPDATE ProductImage
SET id = -id
WHERE id IN (3, 4, 5, 6, 7);

-- Затем обновляем id на нужные значения
UPDATE ProductImage
SET id = CASE
    WHEN id = -3 THEN 2
    WHEN id = -4 THEN 3
    WHEN id = -5 THEN 4
    WHEN id = -6 THEN 5
    WHEN id = -7 THEN 6
END
WHERE id IN (-3, -4, -5, -6, -7);

-- И наконец возвращаем отрицательные id обратно в положительные
UPDATE my_table
SET id = ABS(id)
WHERE id < 0;



-- Вставка данных в таблицу Products
INSERT INTO Products (name, image_id, price, description_id, category_id) VALUES
('Шляпа', 2, 1000, 2, 4),
('Худи', 3, 2000, 3, 6),
('Зимние ботинки', 4, 5000, 4, 5),
('Куртка со скорпионом', 5, 6000, 5, 2),
('Синие джинсы', 6, 3000, 6, 3);
('Футболка с самураем', 1, 1500, 1, 1),

INSERT INTO Size (product_id, size) VALUES
(1, 'XS'),
(1, 'S'),
(1, 'M'),
(1, 'L'),
(1, 'XL'),
(1, 'XXL'),
(27, 'OS'),
(28, 'XS'),
(28, 'S'),
(28, 'M'),
(28, 'L'),
(28, 'XL'),
(28, 'XXL'),
(29, '39EU'),
(29, '40EU'),
(29, '41EU'),
(29, '42EU'),
(29, '43EU'),
(29, '44EU'),
(30, 'XS'),
(30, 'S'),
(30, 'M'),
(30, 'L'),
(30, 'XL'),
(30, 'XXL'),
(31, 'XS'),
(31, 'S'),
(31, 'M'),
(31, 'L'),
(31, 'XL'),
(31, 'XXL');


SELECT * FROM ProductImage
SELECT * FROM  Products

SELECT * FROM Size

SELECT p.name, s.size 
FROM Products as p
JOIN Size as s ON s.product_id = p.id;