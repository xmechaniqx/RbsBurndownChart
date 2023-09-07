CREATE TABLE t_ref_position (
    c_id SMALLSERIAL PRIMARY KEY,
    c_name VARCHAR(100) NOT NULL);
    COMMENT ON COLUMN t_ref_position.c_id IS 'ID';
    COMMENT ON COLUMN t_ref_position.c_name IS 'Должность';
--Создаем таблицу с классификацией источников задач
CREATE TABLE t_ref_tasks_source (
    c_id SMALLSERIAL PRIMARY KEY,
    c_name VARCHAR(100)
);
    COMMENT ON COLUMN t_ref_tasks_source.c_id IS 'ID';
    COMMENT ON COLUMN t_ref_tasks_source.c_name IS 'Источник задачи';
CREATE TABLE t_task (
    c_id SERIAL PRIMARY KEY,
    c_title VARCHAR(250) NOT NULL,
    c_cost SMALLINT
	);
	COMMENT ON COLUMN t_task.c_id IS 'ID';
    COMMENT ON COLUMN t_task.c_title IS 'Наименование задачи';
    COMMENT ON COLUMN t_task.c_cost IS 'Норма выполнения(в часах)';
	--Создаем таблицу со списком сотрудников
CREATE TABLE t_developer (
    c_id SMALLSERIAL PRIMARY KEY,                                                               
    c_lastname VARCHAR(70) NOT NULL,                                          
    c_firstname VARCHAR(50) NOT NULL,                                          
    c_login VARCHAR(25) NOT NULL,
    fk_position SMALLINT REFERENCES t_ref_position (c_id),
    c_mon_hour SMALLINT,
    c_tue_hour SMALLINT,                 
    c_wed_hour SMALLINT,
    c_thu_hour SMALLINT,                              
    c_fri_hour SMALLINT,
    c_sat_hour SMALLINT,
    c_sun_hour SMALLINT);
    COMMENT ON COLUMN t_developer.c_id IS 'ID сотрудника';
    COMMENT ON COLUMN t_developer.c_lastname IS 'Фамилия';
    COMMENT ON COLUMN t_developer.c_firstname IS 'Имя';
    COMMENT ON COLUMN t_developer.c_login IS 'Логин';
    COMMENT ON COLUMN t_developer.fk_position IS 'Внешний ключ - должность';
    COMMENT ON COLUMN t_developer.c_mon_hour IS 'Понедельник';
    COMMENT ON COLUMN t_developer.c_tue_hour IS 'Вторник';
    COMMENT ON COLUMN t_developer.c_wed_hour IS 'Среда';
    COMMENT ON COLUMN t_developer.c_thu_hour IS 'Четверг';
    COMMENT ON COLUMN t_developer.c_fri_hour IS 'Пятница';
    COMMENT ON COLUMN t_developer.c_sat_hour IS 'Суббота';
    COMMENT ON COLUMN t_developer.c_sun_hour IS 'Воскресенье';
--Создаем таблицу с перечнем должностей сотрудников
--Создаем таблицу которая будет хранить данных о проектах и их расположении
CREATE TABLE t_project (
    c_id SMALLSERIAL PRIMARY KEY,
    c_name VARCHAR(100) NOT NULL,
    c_sprint_start_date DATE NOT NULL,
    fk_tasks_source SMALLINT REFERENCES t_ref_tasks_source (c_id),
    c_task_list_file_path TEXT NOT NULL);
	
	
    
    COMMENT ON COLUMN t_project.c_id IS 'ID';
    COMMENT ON COLUMN t_project.c_name IS 'Название проекта';
    COMMENT ON COLUMN t_project.c_sprint_start_date IS 'Дата начала выполнения';
    COMMENT ON COLUMN t_project.fk_tasks_source IS 'Внешний ключ источника задач';
    COMMENT ON COLUMN t_project.c_task_list_file_path IS 'Расположение задач';

--Создаем связную таблицу между сотрудниками и проектами
CREATE TABLE toc_project_developer (
    c_id SERIAL PRIMARY KEY,
    fk_project SMALLINT REFERENCES t_project (c_id),
    fk_developer SMALLINT REFERENCES t_developer (c_id));
    COMMENT ON COLUMN toc_project_developer.c_id IS 'ID';
    COMMENT ON COLUMN toc_project_developer.fk_project IS 'Внешний ключ проекта';
    COMMENT ON COLUMN toc_project_developer.fk_developer IS 'Внешний ключ разработчика';


