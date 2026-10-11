
ALTER TABLE public.ideograms
    ADD COLUMN IF NOT EXISTS meanings TEXT[] NOT NULL DEFAULT '{}',
    ADD COLUMN IF NOT EXISTS radical VARCHAR(10),
    ADD COLUMN IF NOT EXISTS stroke_count INT,
    ADD COLUMN IF NOT EXISTS example_sentence TEXT,
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

DO $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_schema = 'public'
          AND table_name = 'ideograms'
          AND column_name = 'meaning'
    ) THEN
        UPDATE public.ideograms
        SET meanings = ARRAY[meaning]
        WHERE (
            meanings IS NULL
            OR cardinality(meanings) = 0
        )
          AND meaning IS NOT NULL
          AND BTRIM(meaning) <> '';

        ALTER TABLE public.ideograms
            DROP COLUMN meaning;
    END IF;
END $$;
