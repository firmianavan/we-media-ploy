create database wmp default charset 'utf8';
use wmp;

CREATE TABLE IF NOT EXISTS user(
    user_id INT NOT NULL AUTO_INCREMENT ,
    account VARCHAR(30) UNIQUE COMMENT '账户',
    passwd VARCHAR(60) NOT NULL COMMENT '加密后的密码',
    name VARCHAR(30) NOT NULL COMMENT '昵称',
    email VARCHAR(60) UNIQUE,
    age int,
    address VARCHAR(60) COMMENT '现居地',
    score int DEFAULT 0 COMMENT '积分',
    introduce VARCHAR(300) COMMENT '简介',
    created TIMESTAMP DEFAULT NOW()  COMMENT '创建时间',
    PRIMARY KEY ( user_id )
    ) COMMENT='用户表';

CREATE TABLE IF NOT EXISTS category(
    category_id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL COMMENT '作者id',
    name VARCHAR(30) COMMENT '专辑名',
    introduce VARCHAR(60) COMMENT '简介',
    created TIMESTAMP DEFAULT NOW() COMMENT '创建时间',
    cover VARCHAR(200) COMMENT '封面图片url',
    visible ENUM('public','private','only-to-friend') NOT NULL COMMENT '专辑可见性',
    PRIMARY KEY (category_id)
    ) COMMENT='用户的文章专辑';
ALTER TABLE category ADD INDEX idx_special_userid (user_id);
ALTER TABLE category ADD CONSTRAINT fk_special_user FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE;
 
CREATE TABLE IF NOT EXISTS tags(
    article_id INT NOT NULL COMMENT '文章id',
    name VARCHAR(20) NOT NULL COMMENT '标签名称'
    ) COMMENT='文章标签';
ALTER TABLE tags ADD INDEX idx_tags_articleid (article_id);
ALTER TABLE tags ADD INDEX idx_tags_tag (name);


CREATE TABLE IF NOT EXISTS article(
    article_id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL COMMENT '作者id',
    title VARCHAR(150) NOT NULL COMMENT '标题',
    body TEXT NOT NULL COMMENT '文章内容(html片段)',
    category_id INT NOT NULL COMMENT '专辑id',
    created TIMESTAMP DEFAULT NOW()  COMMENT '创建时间',
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    draft ENUM('false','true') NOT NULL COMMENT '是否为草稿',
    PRIMARY KEY (article_id)
    ) COMMENT='用户文章';
ALTER TABLE article ADD INDEX idx_article_userid (user_id);
ALTER TABLE article ADD CONSTRAINT fk_article_user FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE;
ALTER TABLE article ADD INDEX idx_article_specialid (category_id);
ALTER TABLE article ADD CONSTRAINT fk_article_special FOREIGN KEY (category_id) REFERENCES category(category_id);