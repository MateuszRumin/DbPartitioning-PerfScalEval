-- Pamięć i buforowanie

--Określa rozmiar puli buforów InnoDB, która przechowuje dane i indeksy:
SHOW VARIABLES LIKE 'innodb_buffer_pool_size'; --Default:128 MB 
--ozmiar bufora logów transakcyjnych InnoDB:
SHOW VARIABLES LIKE 'innodb_log_buffer_size'; --Default:16 MB
--Rozmiar bufora indeksów dla silnika MyISAM:
SHOW VARIABLES LIKE 'key_buffer_size';--Default:8 MB
--iczba tabel, które MySQL może trzymać otwarte jednocześnie:
SHOW VARIABLES LIKE 'table_open_cache';--Default:2000
--Liczba przechowywanych w pamięci definicji tabel:
SHOW VARIABLES LIKE 'table_definition_cache';--Default:2000
--Liczba przechowywanych w pamięci wątków, które mogą być ponownie użyte:
SHOW VARIABLES LIKE 'thread_cache_size';--Default:9

-- Logi i operacje na dysku
--Określa, jak często log transakcyjny InnoDB jest zapisywany na dysk:
--  -1 – gwarantuje trwałość transakcji 
--  -0 – zapis odbywa się co sekundę, ryzyko utraty danych w przypadku awarii
--  -2 – zapis do pamięci co transakcję, ale na dysk co sekundę
SHOW VARIABLES LIKE 'innodb_flush_log_at_trx_commit';--Default:1
--Kontroluje częstotliwość zapisu binlogów na dysk:
--  -1 – pełna trwałość (wolniejsze, ale bezpieczne)
--  -0 – większa wydajność, ale możliwa utrata binlogów przy awarii
SHOW VARIABLES LIKE 'sync_binlog';--Default:1
--Określa metodę zapisu danych na dysk:
--  -O_DIRECT pozwala ominąć cache systemowy
SHOW VARIABLES LIKE 'innodb_flush_method';--Default:fsync
--iczba operacji wejścia/wyjścia na sekundę, którą InnoDB uznaje za optymalną:
SHOW VARIABLES LIKE 'innodb_io_capacity';--Default:200
SHOW VARIABLES LIKE 'innodb_io_capacity_max';

-- Połączenia i wątki
--Maksymalna liczba jednoczesnych połączeń do serwera
SHOW VARIABLES LIKE 'max_connections';--Default:151
--Czas oczekiwania na bezczynne połączenie przed jego zamknięciem
SHOW VARIABLES LIKE 'wait_timeout';--Default:28800
--Określa sposób zarządzania wątkami:
--pool-of-threads może być lepszy przy dużej liczbie połączeń
SHOW VARIABLES LIKE 'thread_handling';--Default:one-thread-per-connection

SHOW VARIABLES LIKE 'innodb_log_file_size';

-- Indeksy i sortowanie
--Rozmiar bufora używanego do sortowania wyników
SHOW VARIABLES LIKE 'innodb_sort_buffer_size';--Default:1 MB
--Rozmiar pamięci dla operacji JOIN bez indeksów
SHOW VARIABLES LIKE 'join_buffer_size';--Default:256 KB
--Rozmiar bufora sortowania dla operacji ORDER BY i GROUP BY
SHOW VARIABLES LIKE 'sort_buffer_size';---Default:256 KB
--Bufor dla operacji INSERT ... SELECT i LOAD DATA INFILE
SHOW VARIABLES LIKE 'bulk_insert_buffer_size';---Default:256 KB
--Maksymalny rozmiar tabel tymczasowych w pamięci
SHOW VARIABLES LIKE 'tmp_table_size';--Default:16 MB
--Maksymalny rozmiar tabel tymczasowych w pamięci
SHOW VARIABLES LIKE 'max_heap_table_size';--Default:16 MB
--Rozmiar bufora używanego do odczytu losowego przy sortowaniu
SHOW VARIABLES LIKE 'read_rnd_buffer_size';--Default:256 KB

-- Partycjonowanie i replikacja
--Format binlogów (STATEMENT, ROW, MIXED:
-- -ROW daje większą dokładność przy replikacji, ale zużywa więcej zasobów
SHOW VARIABLES LIKE 'binlog_format';--Default:ROW
--Czy replikant zapisuje binlogi z replikowanych zmian
SHOW VARIABLES LIKE 'log_slave_updates';--Default:OFF
--Pamięć wstępnie przydzielona dla analizy zapytań
SHOW VARIABLES LIKE 'query_prealloc_size';--Default:8192 bajtów





SHOW GLOBAL STATUS LIKE 'Open_tables';
SHOW GLOBAL STATUS LIKE 'Opened_tables';
SHOW VARIABLES LIKE 'table_open_cache';
table_open_cache ≈ Open_tables * 2
SHOW VARIABLES LIKE 'open_files_limit';


SHOW ENGINE INNODB STATUS\G
SHOW GLOBAL STATUS LIKE 'Innodb_buffer_pool_pages_flushed';
SHOW GLOBAL STATUS LIKE 'Innodb_data_pending_fsyncs';
--Flushowanie (Innodb_buffer_pool_pages_flushed) wzrosło, ale pending I/O (Innodb_data_pending_fsyncs) nadal jest wysokie, zwiększ innodb_io_capacity.
--Flushowanie wzrosło, ale nie ma dużych opóźnień, wartości są optymalne.
--Serwer zużywa za dużo I/O na zapis (spowalniając inne operacje), zmniejsz innodb_io_capacity.




--kontroluje, jak często InnoDB zapisuje dane dziennika transakcji na dysku.
-- -0 Dziennik transakcji InnoDB jest zapisywany do pamięci i tylko co sekundę zapisywany na dysku.
-- -1 Dziennik transakcji jest zapisywany na dysku przy każdym zatwierdzeniu transakcji. 
-- -2 Dziennik transakcji jest zapisywany do pamięci przy każdym zatwierdzeniu transakcji, ale zapis na dysk odbywa się co sekundę.
SHOW VARIABLES LIKE 'innodb_flush_log_at_trx_commit';
--określa całkowitą pojemność bufora dziennika redo log w InnoDB, który jest używany do zapisywania zmian w bazie danych przed ich zapisaniem do danych na dysku.
SHOW VARIABLES LIKE 'innodb_redo_log_capacity';--Default:; 1GB
--kontroluje maksymalny rozmiar pakietu, który MySQL może przesłać w jednym zapytaniu. Oznacza to maksymalną wielkość pojedynczego pakietu danych (w bajtach), który może zostać wysłany lub odebrany przez serwer.
SHOW VARIABLES LIKE 'max_allowed_packet';--Default:64 MB
--kontroluje rozmiar bufora odczytu wykorzystywanego przez silnik InnoDB (i inne silniki) podczas losowego odczytu wierszy w zapytaniach, które korzystają z operacji ORDER BY, GROUP BY oraz innych operacji wymagających sortowania.
SHOW VARIABLES LIKE 'read_rnd_buffer_size';--Default:256 KB


--Dodatkowe:
--odzielenie bufora na mniejsze instancje
-- -Zaleca się ustawić  na wartość równą liczbie rdzeni CPU
SHOW VARIABLES LIKE 'innodb_buffer_pool_instances';
-- pozwala przechowywać dane każdej tabeli w osobnym pliku, co ułatwia zarządzanie tabelami i poprawia wydajność.
SHOW VARIABLES LIKE 'innodb_file_per_table';

-- -ustawienie na 0 może poprawić wydajność, ponieważ zapytania o metadane nie będą blokowały odświeżania statystyk.
SHOW VARIABLES LIKE 'innodb_stats_on_metadata';
-- -
SHOW VARIABLES LIKE 'innodb_stats_on_metadata';




SHOW VARIABLES LIKE 'innodb_read_io_threads';
SHOW VARIABLES LIKE 'innodb_write_io_threads';


SELECT post_body,post_title, COUNT(*) AS occurrences
FROM posts
GROUP BY post_body, post_title         
HAVING occurrences > 1;








