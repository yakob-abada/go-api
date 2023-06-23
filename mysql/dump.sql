DROP TABLE IF EXISTS `product`;

CREATE TABLE `product` (
  `sku` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `price` decimal(20,0) NOT NULL,
  `product_type` varchar(45) NOT NULL,
  `size` decimal(20,0) DEFAULT NULL,
  `weight` decimal(20,0) DEFAULT NULL,
  `height` decimal(20,0) DEFAULT NULL,
  `length` decimal(20,0) DEFAULT NULL,
  `width` decimal(20,0) DEFAULT NULL,
  PRIMARY KEY (`sku`)
);

LOCK TABLES `product` WRITE;

INSERT INTO `product` VALUES ('test','ProductName',100,'dvd',200,NULL,NULL,NULL,NULL);

UNLOCK TABLES;