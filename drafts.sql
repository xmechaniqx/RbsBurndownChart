
удалить БД

создать БД

CREATE TABLE t_group (                                                       // Таблица группы разработки
c_id integer NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,              // Идентификатор табличного значения
c_name varchar (100) NOT NULL,                                               // Имя группы
c_project varchar (100) NOT NULL,                                            // Название проекта
c_task_file_path varchar (100) NOT NULL,                                     // Путь к файлу со списком задач
c_date_sprint_start date NOT NULL                                            // Дата начала спринта
)

CREATE TABLE t_dev (                                                         // Таблица с перечнем разработчиков
c_id integer NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,              // Идентификатор табличного значения    
fk_group INTEGER REFERENCES t_group (c_id),                                  // Внешний ключ таблицы для связи с t_group
c_lastname varchar (100) NOT NULL,                                           // Фамилия 
c_firstname varchar (100) NOT NULL,                                          // Имя
c_login varchar (100) NOT NULL                                               // Логин
)

CREATE TABLE t_ref_daily (                                                   // Таблица с перечнем дней недели
c_id integer NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,              // Идентификатор табличного значения    
c_name varchar (100) NOT NULL,                                               // Название дня недели (Понедельник, вторник и т.д.)
c_alias char(3) NOT NULL,                                                    // Название дня недели сокращенно на англ. (mon, tue, wed, thu, fri, sat, sun)
c_day_order integer NOT NULL                                                 // Порядковый номер дня недели (1-понедельник, 2-вторник и т.д.)
)

CREATE TABLE toc_dev_day (
c_id integer NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,              // Идентификатор табличного значения    
fk_dev INTEGER REFERENCES t_dev (c_id),                                      // Внешний ключ таблицы для связи с t_dev
fk_day INTEGER REFERENCES t_ref_daily (c_id),                                // Внешний ключ таблицы для связи с t_ref_daily
c_hour integer NOT NULL                                                      // Часы работы
)
