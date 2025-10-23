CREATE TABLE snippets (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMP NOT NULL,
    expires TIMESTAMP NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);

INSERT INTO snippets (title, content, created, expires)
VALUES (
    'She who arrival end how fertile enabled',
    'Brother she add yet see minuter natural smiling article painted. Themselves at dispatched interested insensible am be prosperous reasonably it. In either so spring wished. Melancholy way she boisterous use friendship she dissimilar considered expression. Sex quick arose mrs lived. Mr things do plenty others an vanity myself waited to. Always parish tastes at as mr father dining at.',
    NOW() AT TIME ZONE 'UTC',
    (NOW() AT TIME ZONE 'UTC') + INTERVAL '365 days'
);

insert into snippets (title, content, created, expires)
values(
'Carried nothing on am warrant towards',
'Polite in of in oh needed itself silent course. Assistance travelling so especially do prosperous appearance mr no celebrated. Wanted easily in my called formed suffer. Songs hoped sense as taken ye mirth at. Believe fat how six drawing pursuit minutes far. Same do seen head am part it dear open to. Whatever may scarcely judgment had.',
NOW() AT TIME ZONE 'UTC',
    (NOW() AT TIME ZONE 'UTC') + INTERVAL '365 days'
)

insert into snippets (title, content, created, expires)
values(
'Improved own provided blessing may peculiar domestic',
'Sight house has sex never. No visited raising gravity outward subject my cottage mr be. Hold do at tore in park feet near my case. Invitation at understood occasional sentiments insipidity inhabiting in. Off melancholy alteration principles old. Is do speedily kindness properly oh. Respect article painted cottage he is offices parlors.',
NOW() AT TIME ZONE 'UTC',
    (NOW() AT TIME ZONE 'UTC') + INTERVAL '365 days'
)