INSERT INTO articles (`id`, `title`, `body`,`url`) VALUES ("c1", "title1", "hello!","https://hoge.png");
INSERT INTO articles (`id`, `title`, `body`,`url`) VALUES ("c2", "title2", "hello!","https://hoge1.png");
INSERT INTO articles (`id`, `title`, `body`,`url`) VALUES ("c3", "title3", "hello!","https://hoge2.png");

-- DELIMITER //
-- CREATE PROCEDURE testInsertvol2(IN max INT)
--   BEGIN
--   DECLARE cnt INT Default 1 ;
--     simple_loop: LOOP
--       INSERT INTO articles (id, title, body, url) VALUES (CONCAT('user',cnt), CONCAT('test',cnt,'@test.com'),CONCAT('test_password',cnt),CONCAT('test_password',cnt));
--       SET cnt = cnt+1;
--       IF cnt=max THEN
--         LEAVE simple_loop;
--        END IF;
--     END LOOP simple_loop;
-- END //
-- DELIMITER ;

-- CALL testInsertvol2(1000000);

-- DROP PROCEDURE testInsertvol2;