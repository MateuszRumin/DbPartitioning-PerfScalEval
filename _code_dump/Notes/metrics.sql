WITH ranked AS (
    SELECT
        q.duration_ms,
        CUME_DIST() OVER (ORDER BY q.duration_ms) AS cumulative_distribution
    FROM QueryResults AS q
    JOIN Tests AS t
      ON q.timeEnded >= t.timeStart
     AND q.timeEnded < t.timeEnd
    WHERE t.id = 61
      AND q.duration_ms IS NOT NULL
)
SELECT
    COUNT(*) AS query_count,
    MIN(CASE
        WHEN cumulative_distribution >= 0.50 THEN duration_ms
    END) AS p50_ms,
    MIN(CASE
        WHEN cumulative_distribution >= 0.95 THEN duration_ms
    END) AS p95_ms,
    MIN(CASE
        WHEN cumulative_distribution >= 0.99 THEN duration_ms
    END) AS p99_ms
FROM ranked;

SELECT COUNT(*) FROM QueryResults q JOIN Tests t ON q.timeEnded >= t.timeStart AND q.timeEnded < t.timeEnd WHERE t.id = 40;


WITH test_bounds AS (
    SELECT
        timeStart,
        timeEnd,
        TIMESTAMPDIFF(MICROSECOND, timeStart, timeEnd) / 1000000.0
            AS test_duration_s
    FROM Tests
    WHERE id = 61
),
ranked AS (
    SELECT
        q.duration_ms,
        CUME_DIST() OVER (ORDER BY q.duration_ms) AS cumulative_distribution
    FROM QueryResults AS q
    CROSS JOIN test_bounds AS tb
    WHERE q.timeEnded >= tb.timeStart
      AND q.timeEnded < tb.timeEnd
      AND q.duration_ms IS NOT NULL
)
SELECT
    COUNT(*) AS query_count,
    ROUND(
        COUNT(*) / NULLIF((SELECT test_duration_s FROM test_bounds), 0),
        3
    ) AS avg_qps,
    MIN(CASE
        WHEN cumulative_distribution >= 0.50 THEN duration_ms
    END) AS p50_ms,
    MIN(CASE
        WHEN cumulative_distribution >= 0.95 THEN duration_ms
    END) AS p95_ms,
    MIN(CASE
        WHEN cumulative_distribution >= 0.99 THEN duration_ms
    END) AS p99_ms
FROM ranked;