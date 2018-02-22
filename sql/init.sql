-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema tobee
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema tobee
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `tobee` DEFAULT CHARACTER SET utf8 ;
USE `tobee` ;

-- -----------------------------------------------------
-- Table `tobee`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(20) NOT NULL,
  `phone` VARCHAR(20) NOT NULL,
  `password` VARCHAR(50) NOT NULL,
  `location` VARCHAR(50) NOT NULL,
  `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `vip` INT NOT NULL,
  `avatar_url` VARCHAR(50) NULL,
  `camera` VARCHAR(50) NULL,
  `description` VARCHAR(200) NULL,
  `occupation` VARCHAR(50) NULL,
  `college` VARCHAR(50) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC));


-- -----------------------------------------------------
-- Table `tobee`.`photos`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`photos` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `url` VARCHAR(50) NOT NULL,
  `took_time` DATETIME NOT NULL,
  `took_location` VARCHAR(50) NOT NULL,
  `release_time` DATETIME NULL,
  `category` VARCHAR(20) NOT NULL,
  `likes` INT NOT NULL,
  `reject_reason` VARCHAR(200) NULL,
  `photographer_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `photographer_id_idx` (`photographer_id` ASC),
  CONSTRAINT `photographer_id_1`
    FOREIGN KEY (`photographer_id`)
    REFERENCES `tobee`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tobee`.`photo_tags`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`photo_tags` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `tag` VARCHAR(20) NOT NULL,
  `photo_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `photo_id_idx` (`photo_id` ASC),
  CONSTRAINT `photo_id_2`
    FOREIGN KEY (`photo_id`)
    REFERENCES `tobee`.`photos` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tobee`.`organizations`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`organizations` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(20) NOT NULL,
  `phone` VARCHAR(20) NOT NULL,
  `password` VARCHAR(50) NOT NULL,
  `college` VARCHAR(20) NOT NULL,
  `logo_url` VARCHAR(50) NULL,
  `description` VARCHAR(200) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  UNIQUE INDEX `phone_UNIQUE` (`phone` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tobee`.`photo_comments`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`photo_comments` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(20) NOT NULL,
  `content` VARCHAR(200) NOT NULL,
  `user_id` INT NOT NULL,
  `photo_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `user_id_idx` (`user_id` ASC),
  INDEX `photo_id_idx` (`photo_id` ASC),
  CONSTRAINT `user_id_1`
    FOREIGN KEY (`user_id`)
    REFERENCES `tobee`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `photo_id_1`
    FOREIGN KEY (`photo_id`)
    REFERENCES `tobee`.`photos` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tobee`.`organization_departments`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`organization_departments` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `organization_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `organization_id_idx` (`organization_id` ASC),
  CONSTRAINT `organization_id_3`
    FOREIGN KEY (`organization_id`)
    REFERENCES `tobee`.`organizations` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tobee`.`activities`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`activities` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(20) NOT NULL,
  `description` VARCHAR(200) NOT NULL,
  `category` VARCHAR(10) NOT NULL,
  `poster_url` VARCHAR(50) NULL,
  `logo_url` VARCHAR(50) NULL,
  `organization_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `organization_id_idx` (`organization_id` ASC),
  CONSTRAINT `organization_id_1`
    FOREIGN KEY (`organization_id`)
    REFERENCES `tobee`.`organizations` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tobee`.`activity_stages`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`activity_stages` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `stage_num` INT NOT NULL,
  `start_time` DATETIME NOT NULL,
  `end_time` DATETIME NOT NULL,
  `location` VARCHAR(20) NOT NULL,
  `content` VARCHAR(400) NOT NULL,
  `activity_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `activity_id_idx` (`activity_id` ASC),
  CONSTRAINT `activity_id_1`
    FOREIGN KEY (`activity_id`)
    REFERENCES `tobee`.`activities` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
COMMENT = '社团活动阶段';


-- -----------------------------------------------------
-- Table `tobee`.`photo_lives`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`photo_lives` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `expect_members` INT NOT NULL,
  `ad_progress` VARCHAR(200) NOT NULL,
  `activity_stage_id` INT NULL,
  `manager_id` INT NULL,
  `photographer_manager_id` INT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `activity_stage_id_idx` (`activity_stage_id` ASC),
  INDEX `manager_id_idx` (`manager_id` ASC),
  INDEX `photographer_manager_id_idx` (`photographer_manager_id` ASC),
  CONSTRAINT `activity_stage_id_1`
    FOREIGN KEY (`activity_stage_id`)
    REFERENCES `tobee`.`activity_stages` (`id`)
    ON DELETE SET NULL
    ON UPDATE CASCADE,
  CONSTRAINT `manager_id_1`
    FOREIGN KEY (`manager_id`)
    REFERENCES `tobee`.`users` (`id`)
    ON DELETE SET NULL
    ON UPDATE CASCADE,
  CONSTRAINT `photographer_manager_id_1`
    FOREIGN KEY (`photographer_manager_id`)
    REFERENCES `tobee`.`users` (`id`)
    ON DELETE SET NULL
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tobee`.`user_login_logs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`user_login_logs` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `login_time` DATETIME NOT NULL,
  `login_location` VARCHAR(50) NOT NULL,
  `login_device` VARCHAR(20) NOT NULL,
  `user_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `user_id_idx` (`user_id` ASC),
  CONSTRAINT `user_id_3`
    FOREIGN KEY (`user_id`)
    REFERENCES `tobee`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tobee`.`organization_login_logs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`organization_login_logs` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `login_time` DATETIME NOT NULL,
  `login_location` VARCHAR(20) NOT NULL,
  `login_device` VARCHAR(20) NOT NULL,
  `organization_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `organization_id_idx` (`organization_id` ASC),
  CONSTRAINT `organization_id_4`
    FOREIGN KEY (`organization_id`)
    REFERENCES `tobee`.`organizations` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tobee`.`administrators`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`administrators` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(20) NOT NULL,
  `password` VARCHAR(50) NOT NULL,
  `level` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tobee`.`administrator_login_logs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`administrator_login_logs` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `login_time` DATETIME NOT NULL,
  `login_location` VARCHAR(20) NOT NULL,
  `login_device` VARCHAR(20) NOT NULL,
  `administrator_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `administrator_id_idx` (`administrator_id` ASC),
  CONSTRAINT `administrator_id_1`
    FOREIGN KEY (`administrator_id`)
    REFERENCES `tobee`.`administrators` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tobee`.`users_organizations`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`users_organizations` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `organization_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `organization_id_idx` (`organization_id` ASC),
  INDEX `user_id_idx` (`user_id` ASC),
  CONSTRAINT `user_id_4`
    FOREIGN KEY (`user_id`)
    REFERENCES `tobee`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `organization_id_5`
    FOREIGN KEY (`organization_id`)
    REFERENCES `tobee`.`organizations` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tobee`.`users_users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`users_users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `liked_user_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `liked_user_id_idx` (`liked_user_id` ASC),
  INDEX `user_id_idx` (`user_id` ASC),
  CONSTRAINT `user_id_6`
    FOREIGN KEY (`user_id`)
    REFERENCES `tobee`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `liked_user_id_1`
    FOREIGN KEY (`liked_user_id`)
    REFERENCES `tobee`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tobee`.`users_photos`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`users_photos` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `liked_photo_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `user_id_idx` (`user_id` ASC),
  INDEX `liked_photo_id_idx` (`liked_photo_id` ASC),
  CONSTRAINT `user_id_5`
    FOREIGN KEY (`user_id`)
    REFERENCES `tobee`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `liked_photo_id_1`
    FOREIGN KEY (`liked_photo_id`)
    REFERENCES `tobee`.`photos` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tobee`.`users_activities`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`users_activities` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `activity_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `user_id_idx` (`user_id` ASC),
  INDEX `activity_id_idx` (`activity_id` ASC),
  CONSTRAINT `user_id_2`
    FOREIGN KEY (`user_id`)
    REFERENCES `tobee`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `activity_id_2`
    FOREIGN KEY (`activity_id`)
    REFERENCES `tobee`.`activities` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tobee`.`organizations_contactors`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`organizations_contactors` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `organization_id` INT NOT NULL,
  `contact_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `organization_id_idx` (`organization_id` ASC),
  INDEX `contact_id_idx` (`contact_id` ASC),
  CONSTRAINT `organization_id_2`
    FOREIGN KEY (`organization_id`)
    REFERENCES `tobee`.`organizations` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `contact_id_1`
    FOREIGN KEY (`contact_id`)
    REFERENCES `tobee`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tobee`.`photo_lives_supervisors`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tobee`.`photo_lives_supervisors` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `photo_live_id` INT NOT NULL,
  `supervisor_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `photo_live_id_idx` (`photo_live_id` ASC),
  INDEX `supervisor_id_idx` (`supervisor_id` ASC),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  CONSTRAINT `photo_live_id_1`
    FOREIGN KEY (`photo_live_id`)
    REFERENCES `tobee`.`photo_lives` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `supervisor_id_1`
    FOREIGN KEY (`supervisor_id`)
    REFERENCES `tobee`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
activitiesactivities