CREATE TABLE IF NOT EXISTS `basic_user_info`(
   `user_id` INT UNSIGNED AUTO_INCREMENT,
   `user_name` VARCHAR(100) NOT NULL,
   `user_password` VARCHAR(40) NOT NULL,
   `user_submission_date` DATE,
   PRIMARY KEY ( `user_id` )
)ENGINE=InnoDB;

INSERT INTO basic_user_info(
user_name, user_password,user_submission_date)VALUES(
"admin","admin_pass",NOW());

SELECT * FROM basic_user_info
