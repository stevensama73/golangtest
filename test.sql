CREATE TABLE `dbprovider`.`providers` (
  `provider` VARCHAR(32) NOT NULL, 
  `no_hp` VARCHAR(255) NOT NULL, 
  PRIMARY KEY (`provider`, `no_hp`)
) ENGINE = InnoDB;
