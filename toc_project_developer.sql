--Заполняем таблицу задач t_task 
INSERT INTO t_task VALUES (default,'Создание и конфигурация репозитория', 1);
INSERT INTO t_task VALUES (default,'Разработка макета веб-страницы', 8);
INSERT INTO t_task VALUES (default,'Перенос макета в рабочий проект', 2);
INSERT INTO t_task VALUES (default,'Запуск макета на тестовом сервере с тестовыми данными', 12);
INSERT INTO t_task VALUES (default,'Подключение веб-интерфейса к БД', 10);
INSERT INTO t_task VALUES (default,'Запуск веб-интерфейса и отладка', 15);
INSERT INTO t_task VALUES (default,'Настройка визуализации параметров полученных из БД', 5);
INSERT INTO t_task VALUES (default,'Сборка Webpack', 14);
INSERT INTO t_task VALUES (default,'Тестирование', 12);
INSERT INTO t_task VALUES (default,'Отладка', 12);
--Заполняем таблицу t_ref_tasks_source
INSERT INTO t_ref_tasks_source (c_name) VALUES ('Data_base');        
INSERT INTO t_ref_tasks_source (c_name) VALUES ('File');
INSERT INTO t_ref_tasks_source (c_name) VALUES ('Web');          
--Заполняем таблицу t_project
INSERT INTO t_project VALUES (default,'website', '2023-09-04', 28, 't_task');
INSERT INTO t_project VALUES (default,'robointellect', '2023-09-04', 29, '/home/username/go/src/RbsBurndownChart/tasks/group_2.txt');
INSERT INTO t_project VALUES (default,'repka', '2023-09-04', 30, 'localhost:4000');
--Заполняем таблицу t_ref_position
INSERT INTO t_ref_position VALUES (default,'stager');
INSERT INTO t_ref_position VALUES (default,'developer');
INSERT INTO t_ref_position VALUES (default,'architect');
--Заполняем таблицу t_developer
INSERT INTO t_developer VALUES (default,'Ivanov', 'Petr', 'pivanov', 1, 4, 6, 2, 4, 8, 0, 0);
INSERT INTO t_developer VALUES (default,'Petrov', 'Ivan', 'ipetrov', 1, 6, 2, 6, 8, 8, 0, 0);
INSERT INTO t_developer VALUES (default,'Sidorov','Nikolay', 'nsidorov', 1, 8, 8, 6, 4, 6, 0, 0);
INSERT INTO t_developer VALUES (default,'Chernetsov','Gennadiy', 'gchernetsov', 2, 8, 8, 6, 8, 6, 0, 0);
INSERT INTO t_developer VALUES (default,'Dubovickiy','Andrey', 'adubovickiy', 2, 8, 8, 7, 6, 8, 0, 0);
INSERT INTO t_developer VALUES (default,'Sidorov','Nikolay', 'nsidorov', 3, 8, 8, 8, 8, 8, 0, 0);
--Заполняем таблицу t_ref_position toc_project_developer
INSERT INTO toc_project_developer VALUES (default,8,1);
INSERT INTO toc_project_developer VALUES (default,8,2);
INSERT INTO toc_project_developer VALUES (default,8,3);
INSERT INTO toc_project_developer VALUES (default,9,4);
INSERT INTO toc_project_developer VALUES (default,9,5);
INSERT INTO toc_project_developer VALUES (default,9,6);
INSERT INTO toc_project_developer VALUES (default,10,4);
INSERT INTO toc_project_developer VALUES (default,10,6);





-- SELECT proj.c_id, proj.c_name, CAST(proj.c_sprint_start_date AS TEXT),tref.c_name, proj.c_task_list_file_path
-- 	FROM public.t_project AS proj
-- 	LEFT JOIN t_ref_tasks_source AS tref ON proj.fk_tasks_source = tref.c_id
-- 	WHERE proj.c_id=8

-- SELECT dev.c_id, toc.fk_project, dev.c_lastname, dev.c_firstname, dev.c_login, 
-- tref.c_name
-- FROM t_developer AS dev
-- LEFT JOIN t_ref_position AS tref
-- ON dev.fk_position = tref.c_id
-- LEFT JOIN toc_project_developer AS toc
-- ON dev.c_id = toc.fk_developer
-- WHERE dev.c_login = 'gchernetsov'

-- SELECT dev.c_id, toc.fk_project, dev.c_lastname, dev.c_firstname, dev.c_login, 
-- tref.c_name,  dev.c_mon_hour, dev.c_tue_hour, dev.c_wed_hour,
-- dev.c_thu_hour, dev.c_fri_hour, dev.c_sat_hour, dev.c_sun_hour
-- FROM t_developer AS dev
-- LEFT JOIN t_ref_position AS tref
-- ON dev.fk_position = tref.c_id
-- LEFT JOIN toc_project_developer AS toc
-- ON dev.c_id = toc.fk_developer
-- WHERE dev.c_login = 'gchernetsov'

-- SELECT proj.c_id, proj.c_name, CAST(proj.c_sprint_start_date AS TEXT), 
-- tref.c_name, proj.c_task_list_file_path 
-- FROM public.t_project AS proj 
-- LEFT JOIN t_ref_tasks_source AS tref 
-- ON proj.fk_tasks_source = tref.c_id 
-- WHERE proj.c_id=8

-- SELECT * FROM t_developer

-- SELECT dev.c_id, toc.fk_project, dev.c_lastname, dev.c_firstname, dev.c_login, 
-- tref.c_name,  dev.c_mon_hour, dev.c_tue_hour, dev.c_wed_hour,
-- dev.c_thu_hour, dev.c_fri_hour, dev.c_sat_hour, dev.c_sun_hour
-- FROM t_developer AS dev
-- LEFT JOIN t_ref_position AS tref
-- ON dev.fk_position = tref.c_id
-- LEFT JOIN toc_project_developer AS toc
-- ON dev.c_id = toc.fk_developer
-- WHERE toc.fk_project =8

-- SELECT dev.c_login
-- FROM t_developer AS dev
-- LEFT JOIN t_ref_position AS tref
-- ON dev.fk_position = tref.c_id
-- LEFT JOIN toc_project_developer AS toc
-- ON dev.c_id = toc.fk_developer
-- WHERE toc.fk_project =8