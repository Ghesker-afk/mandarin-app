CREATE TABLE IF NOT EXISTS public.ideograms (
    id SERIAL PRIMARY KEY,
    character VARCHAR(10) NOT NULL UNIQUE,
    pinyin VARCHAR(50) NOT NULL,
    meanings TEXT[] NOT NULL DEFAULT '{}',
    radical VARCHAR(10),
    stroke_count INT,
    example_sentence TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);