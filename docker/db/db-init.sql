IF(db_id(N'BLOG_DEV') IS NULL)
BEGIN
    CREATE DATABASE BLOG_DEV;
END
GO

USE [BLOG_DEV]
GO

IF OBJECT_ID(N'dbo.blog_tag', N'U') IS NOT NULL
BEGIN
    DROP TABLE blog_tag
END

IF OBJECT_ID(N'dbo.blog_article', N'U') IS NOT NULL
BEGIN
    DROP TABLE blog_article
END

IF OBJECT_ID(N'dbo.blog_auth', N'U') IS NOT NULL
BEGIN
    DROP TABLE blog_auth
END

create table [blog_tag] (
[id] [int] identity,
[name] [varchar] (100) NULL,
[created_on] [int] NULL,
[created_by] [varchar] (100) NULL,
[modified_on] [int] NULL,
[modified_by] [varchar] (100) NULL,
[deleted_on] [int] NULL,
[state] [tinyint] NULL);
GO


set identity_insert [blog_tag] on;


insert [blog_tag] ([id],[name],[created_on],[created_by],[modified_on],[modified_by],[deleted_on],[state])
select 1,'Go',0,'admin',0,'',0,1 UNION ALL
select 2,'tag 1',0,'admin',0,'autor 2',0,0 UNION ALL
select 3,'JavaScript',0,'juan',0,'',0,1 UNION ALL
select 4,'Rust',0,'juan',0,'',0,1 UNION ALL
select 5,'C#',0,'gerardo',0,'',0,1 UNION ALL
select 6,'Java',0,'joel',0,'',0,1 UNION ALL
select 7,'Scala',0,'joel',0,'',0,1 UNION ALL
select 8,'Kotlin',0,'joel',0,'',0,1 UNION ALL
select 9,'C++',0,'alberto',0,'',0,1 UNION ALL
select 10,'C',0,'alberto',0,'',0,1 UNION ALL
select 11,'tag 1',0,'admin',0,'',0,1;

set identity_insert [blog_tag] off;
GO

create table [blog_auth] (
[id] [int] identity,
[username] [varchar] (50) NULL,
[password] [varchar] (50) NULL);
GO

set identity_insert [blog_auth] on;

insert [blog_auth] ([id],[username],[password])
select 1,'admin','admin';

set identity_insert [blog_auth] off;
GO

create table [blog_article] (
[id] [int] identity,
[tag_id] [int] NULL,
[title] [varchar] (100) NULL,
[desc] [varchar] (255) NULL,
[content] [text] NULL,
[cover_image_url] [varchar] (255) NULL,
[created_on] [int] NULL,
[created_by] [varchar] (100) NULL,
[modified_on] [int] NULL,
[modified_by] [varchar] (255) NULL,
[deleted_on] [int] NULL,
[state] [tinyint] NULL);
GO


set identity_insert [blog_article] on;

insert [blog_article] ([id],[tag_id],[title],[desc],[content],[cover_image_url],[created_on],[created_by],[modified_on],[modified_by],[deleted_on],[state])
select 2,1,'Articulo 1','Descripción corta del articulo',
'Contenido completo del articulo'
,'https://www.researchgate.net/publication/231225356/figure/fig5/AS:300440198631431@1448641933728/TALER-activity-in-mammalian-cells-A-B-Fold-reduction-of-luciferase-expression-in_Q320.jpg'
,0,'admin',0,'',0,1;

set identity_insert [blog_article] off;
GO