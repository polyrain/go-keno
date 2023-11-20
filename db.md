We need two tables, one for the picks and one for the games
A pick is just some numbers from a starting game index
Games play themselves players just check against the DB


Table 1: Games
- id: primary key
- nums: board nums

```
  -- public.games definition

-- Drop table

-- DROP TABLE public.games;

CREATE TABLE public.games (
	id serial4 NOT NULL,
	board _int8 NOT NULL DEFAULT '{}'::bigint[],
	finished bool NOT NULL DEFAULT false
);
```

Table 2: Picks
- id: primary key
- startID: foreign key (Game ID)
- etc

```
  -- public.picks definition

-- Drop table

-- DROP TABLE public.picks;

CREATE TABLE public.picks (
	id serial4 NOT NULL,
	picks _int8 NOT NULL DEFAULT '{}'::bigint[],
	start_game_num int4 NOT NULL DEFAULT 0,
	number_games int4 NOT NULL DEFAULT 1,
	price_per_game int4 NOT NULL DEFAULT 1,
	picks_per_game int4 NOT NULL
);
````
