CREATE TABLE IF NOT EXISTS pokemons (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  identifier TEXT,
  species_id INTEGER,
  height INTEGER,
  weight INTEGER,
  base_experience INTEGER,
  order_c INTEGER,
  is_default INTEGER 
);
