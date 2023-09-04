
CREATE DATABASE burndown_db;

CREATE TABLE t_task (
    c_id SERIAL,
    c_title VARCHAR(250),
    c_cost SMALLINT
)

CREATE TABLE t_project (
    c_id SMALLSERIAL,
    c_name VARCHAR(100),
    c_sprint_start_date DATE,
    fk_tasks_source SMALLINT,
    c_task_list_file_path TEXT
)

CREATE TABLE t_ref_tasks_source (
    c_id SMALLSERIAL,
    c_name VARCHAR(100)
)

CREATE TABLE toc_project_developer (
    c_id SERIAL,
    fk_project SMALLINT,
    fk_developer SMALLINT
)

CREATE TABLE t_developer (                                                      
    c_id SMALLSERIAL,                                                               
    c_lastname VARCHAR(70),                                          
    c_firstname VARCHAR(50),                                          
    c_login VARCHAR(25)
    fk_position SMALLINT,
    c_mon_hour SMALLINT,
    c_tue_hour SMALLINT,                 
    c_wed_hour SMALLINT,
    c_thu_hour SMALLINT,                              
    c_fri_hour SMALLINT,
    c_sat_hour SMALLINT,
    c_sun_hour SMALLINT,
)

CREATE TABLE t_ref_position (
    c_id SMALLSERIAL,
    c_name VARCHAR(100)
)

INSERT INTO t_task VALUES ('Создание и конфигурация репозитория', 1)
INSERT INTO t_task VALUES ('Разработка макета веб-страницы', 8)
INSERT INTO t_task VALUES ('Перенос макета в рабочий проект', 2)
INSERT INTO t_task VALUES ('Запуск макета на тестовом сервере с тестовыми данными', 12)
INSERT INTO t_task VALUES ('Подключение веб-интерфейса к БД', 10)
INSERT INTO t_task VALUES ('Запуск веб-интерфейса и отладка', 15)
INSERT INTO t_task VALUES ('Настройка визуализации параметров полученных из БД', 5)
INSERT INTO t_task VALUES ('Сборка Webpack', 14)
INSERT INTO t_task VALUES ('Тестирование', 12)
INSERT INTO t_task VALUES ('Отладка', 12)

INSERT INTO t_ref_tasks_source VALUES ('Data_base')          "нумерация для названий проектов???"
INSERT INTO t_ref_tasks_source VALUES ('File')    
INSERT INTO t_ref_tasks_source VALUES ('Web')             

INSERT INTO t_project VALUES ('website', '2023-09-04', 1, 't_task')   ???
INSERT INTO t_project VALUES ('robointellect', '2023-09-04', 2, '/home/username/go/src/RbsBurndownChart/tasks/group_2.txt')
INSERT INTO t_project VALUES ('repka', '2023-09-04', 3, 'localhost:4000')

INSERT INTO t_ref_position VALUES ('stager')
INSERT INTO t_ref_position VALUES ('developer')
INSERT INTO t_ref_position VALUES ('architect')


INSERT INTO t_developer VALUES ('Ivanov', 'Petr', 'pivanov', 1, 4, 6, 2, 4, 8, 0, 0)
INSERT INTO t_developer VALUES ('Petrov', 'Ivan', 'ipetrov', 1, 6, 2, 6, 8, 8, 0, 0)
INSERT INTO t_developer VALUES ('Sidorov','Nikolay', 'nsidorov', 1, 8, 8, 6, 4, 6, 0, 0)

INSERT INTO t_developer VALUES ('Chernetsov','Gennadiy', 'gchernetsov', 2, 8, 8, 6, 8, 6, 0, 0)
INSERT INTO t_developer VALUES ('Dubovickiy','Andrey', 'adubovickiy', 2, 8, 8, 7, 6, 8, 0, 0)

INSERT INTO t_developer VALUES ('Sidorov','Nikolay', 'nsidorov', 3, 8, 8, 8, 8, 8, 0, 0)

INSERT INTO toc_project_developer VALUES (1,1)
INSERT INTO toc_project_developer VALUES (1,2)
INSERT INTO toc_project_developer VALUES (1,3)
INSERT INTO toc_project_developer VALUES (2,4)
INSERT INTO toc_project_developer VALUES (2,5)
INSERT INTO toc_project_developer VALUES (2,6)
INSERT INTO toc_project_developer VALUES (3,4)
INSERT INTO toc_project_developer VALUES (3,6)
;