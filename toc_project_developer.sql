--Заполняем таблицу задач t_task 
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
--Заполняем таблицу t_ref_tasks_source
INSERT INTO t_ref_tasks_source VALUES ('Data_base')          "нумерация для названий проектов???"
INSERT INTO t_ref_tasks_source VALUES ('File')    
INSERT INTO t_ref_tasks_source VALUES ('Web')             
--Заполняем таблицу t_project
INSERT INTO t_project VALUES ('website', '2023-09-04', 1, 't_task')   ???
INSERT INTO t_project VALUES ('robointellect', '2023-09-04', 2, '/home/username/go/src/RbsBurndownChart/tasks/group_2.txt')
INSERT INTO t_project VALUES ('repka', '2023-09-04', 3, 'localhost:4000')
--Заполняем таблицу t_ref_position
INSERT INTO t_ref_position VALUES ('stager')
INSERT INTO t_ref_position VALUES ('developer')
INSERT INTO t_ref_position VALUES ('architect')
--Заполняем таблицу t_developer
INSERT INTO t_developer VALUES ('Ivanov', 'Petr', 'pivanov', 1, 4, 6, 2, 4, 8, 0, 0)
INSERT INTO t_developer VALUES ('Petrov', 'Ivan', 'ipetrov', 1, 6, 2, 6, 8, 8, 0, 0)
INSERT INTO t_developer VALUES ('Sidorov','Nikolay', 'nsidorov', 1, 8, 8, 6, 4, 6, 0, 0)
INSERT INTO t_developer VALUES ('Chernetsov','Gennadiy', 'gchernetsov', 2, 8, 8, 6, 8, 6, 0, 0)
INSERT INTO t_developer VALUES ('Dubovickiy','Andrey', 'adubovickiy', 2, 8, 8, 7, 6, 8, 0, 0)
INSERT INTO t_developer VALUES ('Sidorov','Nikolay', 'nsidorov', 3, 8, 8, 8, 8, 8, 0, 0)
--Заполняем таблицу t_ref_position toc_project_developer
INSERT INTO toc_project_developer VALUES (1,1)
INSERT INTO toc_project_developer VALUES (1,2)
INSERT INTO toc_project_developer VALUES (1,3)
INSERT INTO toc_project_developer VALUES (2,4)
INSERT INTO toc_project_developer VALUES (2,5)
INSERT INTO toc_project_developer VALUES (2,6)
INSERT INTO toc_project_developer VALUES (3,4)
INSERT INTO toc_project_developer VALUES (3,6)
;