USE logDB;



DROP TABLE IF EXISTS resource_usages;


CREATE TABLE resource_usages (
    timestamp           DATETIME,         -- Czas pomiaru 
    cpu_percent         DECIMAL(5,2),     -- Obciążenie CPU w % 
    ram_used            BIGINT,           -- Użyta pamięć RAM w bajtach
    ram_percent         DECIMAL(5,2),     -- Procentowe użycie RAM 
    
    -- Dane dla pierwszego dysku 
    disk1_name          VARCHAR(20),
    disk1_reads         BIGINT,           -- Liczba odczytów
    disk1_writes        BIGINT,           -- Liczba zapisów
    disk1_read_bytes    BIGINT,           -- Bajty odczytane
    disk1_write_bytes   BIGINT,           -- Bajty zapisane
    disk1_read_time     BIGINT,           -- Czas odczytu (ms)
    disk1_write_time    BIGINT,           -- Czas zapisu (ms)
    
    -- Dane dla drugiego dysku
    disk2_name          VARCHAR(20),
    disk2_reads         BIGINT,
    disk2_writes        BIGINT,
    disk2_read_bytes    BIGINT,
    disk2_write_bytes   BIGINT,
    disk2_read_time     BIGINT,
    disk2_write_time    BIGINT,
    
    network_sent        BIGINT,           -- Wysłane bajty przez sieć
    network_received    BIGINT            -- Odebrane bajty przez sieć
);


