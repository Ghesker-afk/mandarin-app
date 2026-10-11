
ALTER TABLE public.ideograms
    ADD COLUMN IF NOT EXISTS meaning TEXT;

UPDATE public.ideograms
SET meaning = meanings[1]
WHERE meaning IS NULL
  AND cardinality(meanings) > 0;

ALTER TABLE public.ideograms
    ALTER COLUMN meaning SET DEFAULT '',
    ALTER COLUMN meaning SET NOT NULL;
