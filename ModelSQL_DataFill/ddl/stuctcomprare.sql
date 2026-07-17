SELECT
    'users' AS tabela,
    (SELECT COUNT(*) FROM baza_zrodlowa.users) AS zrodlo,
    (SELECT COUNT(*) FROM baza_docelowa.users) AS cel
UNION ALL
SELECT
    'posts',
    (SELECT COUNT(*) FROM baza_zrodlowa.posts),
    (SELECT COUNT(*) FROM baza_docelowa.posts)
UNION ALL
SELECT
    'comments',
    (SELECT COUNT(*) FROM baza_zrodlowa.comments),
    (SELECT COUNT(*) FROM baza_docelowa.comments)
UNION ALL
SELECT
    'post_history',
    (SELECT COUNT(*) FROM baza_zrodlowa.post_history),
    (SELECT COUNT(*) FROM baza_docelowa.post_history)
UNION ALL
SELECT
    'post_links',
    (SELECT COUNT(*) FROM baza_zrodlowa.post_links),
    (SELECT COUNT(*) FROM baza_docelowa.post_links)
UNION ALL
SELECT
    'votes',
    (SELECT COUNT(*) FROM baza_zrodlowa.votes),
    (SELECT COUNT(*) FROM baza_docelowa.votes)
UNION ALL
SELECT
    'badges',
    (SELECT COUNT(*) FROM baza_zrodlowa.badges),
    (SELECT COUNT(*) FROM baza_docelowa.badges)
UNION ALL
SELECT
    'tags',
    (SELECT COUNT(*) FROM baza_zrodlowa.tags),
    (SELECT COUNT(*) FROM baza_docelowa.tags);