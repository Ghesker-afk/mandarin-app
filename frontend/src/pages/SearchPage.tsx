
import { useState, type FormEvent } from 'react'
import { Search, LoaderCircle, BookOpen, Hash } from 'lucide-react'
import {
  findByCharacter,
  searchByMeaning,
} from '../services/ideogramService'
import type { Ideogram } from '../types/ideogram'

type SearchMode = 'character' | 'meaning'

export default function SearchPage() {
  const [query, setQuery] = useState('');
  const [mode, setMode] = useState<SearchMode>('character');
  const [results, setResults] = useState<Ideogram[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [searched, setSearched] = useState(false);

  async function handleSearch(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()

    const value = query.trim();

    if (!value) {
      setError('Digite um caractere ou significado para pesquisar.');
      setResults([]);
      setSearched(false);
      return;
    }

    setLoading(true);
    setError('');
    setSearched(true);

    try {
      if (mode === 'character') {
        const ideogram = await findByCharacter(value);
        setResults([ideogram]);
      } else {
        const ideograms = await searchByMeaning(value);
        setResults(ideograms);
      }
    } catch (err) {
      setResults([]);

      setError(
        err instanceof Error
          ? err.message
          : 'Não foi possível realizar a pesquisa.',
      );
    } finally {
      setLoading(false);
    }
  }

  return (
    <section className="mx-auto max-w-5xl space-y-8">
      <header>
        <p className="text-sm font-semibold uppercase tracking-widest text-emerald-800">
          Seu vocabulário
        </p>

        <h1 className="mt-2 text-3xl font-bold tracking-tight text-slate-800 sm:text-4xl">
          Pesquisar ideogramas
        </h1>

        <p className="mt-3 max-w-2xl leading-relaxed text-slate-500">
          Descubra o significado, a pronúncia e outras informações dos
          caracteres chineses cadastrados no Mandarim App.
        </p>
      </header>

      <form
        onSubmit={handleSearch}
        className="space-y-5 rounded-2xl border border-slate-200 bg-white p-5 shadow-sm sm:p-7"
      >
        <div>
          <label className="mb-3 block text-sm font-semibold text-slate-700">
            O que você deseja pesquisar?
          </label>

          <div className="flex flex-wrap gap-2">
            <button
              type="button"
              onClick={() => {
                setMode('character');
                setError('');
                setResults([]);
                setSearched(false);
              }}
              className={`rounded-xl px-4 py-2.5 text-sm font-medium transition ${
                mode === 'character'
                  ? 'bg-[#164b40] text-white'
                  : 'bg-slate-100 text-slate-600 hover:bg-slate-200'
              }`}
            >
              Por caractere
            </button>

            <button
              type="button"
              onClick={() => {
                setMode('meaning');
                setError('');
                setResults([]);
                setSearched(false);
              }}
              className={`rounded-xl px-4 py-2.5 text-sm font-medium transition ${
                mode === 'meaning'
                  ? 'bg-[#164b40] text-white'
                  : 'bg-slate-100 text-slate-600 hover:bg-slate-200'
              }`}
            >
              Por significado
            </button>
          </div>
        </div>

        <div className="flex flex-col gap-3 sm:flex-row">
          <div className="relative flex-1">
            <Search
              size={20}
              className="absolute left-4 top-1/2 -translate-y-1/2 text-slate-400"
            />

            <input
              value={query}
              onChange={(event) => setQuery(event.target.value)}
              placeholder={
                mode === 'character'
                  ? 'Ex.: 山'
                  : 'Ex.: montanha'
              }
              aria-label={
                mode === 'character'
                  ? 'Caractere para pesquisar'
                  : 'Significado para pesquisar'
              }
              className="w-full rounded-xl border border-slate-200 bg-slate-50 py-3.5 pl-11 pr-4 outline-none transition placeholder:text-slate-400 focus:border-emerald-700 focus:ring-2 focus:ring-emerald-700/10"
            />
          </div>

          <button
            type="submit"
            disabled={loading}
            className="inline-flex items-center justify-center gap-2 rounded-xl bg-[#164b40] px-6 py-3.5 font-semibold text-white transition hover:bg-[#20584b] disabled:cursor-not-allowed disabled:opacity-60"
          >
            {loading ? (
              <LoaderCircle size={19} className="animate-spin" />
            ) : (
              <Search size={19} />
            )}
            Pesquisar
          </button>
        </div>

        <p className="text-xs text-slate-400">
          {mode === 'character'
            ? 'Pesquise um caractere cadastrado, como 山.'
            : 'A busca encontra significados que correspondam ao texto digitado.'}
        </p>
      </form>

      {error && (
        <div
          role="alert"
          className="rounded-xl border border-red-200 bg-red-50 p-4 text-sm text-red-700"
        >
          {error}
        </div>
      )}

      {loading && (
        <div
          role="status"
          className="flex items-center justify-center gap-3 py-12 text-slate-500"
        >
          <LoaderCircle size={22} className="animate-spin" />
          Pesquisando ideogramas...
        </div>
      )}

      {!loading && searched && !error && results.length === 0 && (
        <div className="rounded-2xl border border-dashed border-slate-300 px-6 py-12 text-center">
          <Search size={32} className="mx-auto text-slate-400" />
          <h2 className="mt-4 font-semibold text-slate-700">
            Nenhum resultado encontrado
          </h2>
          <p className="mt-2 text-sm text-slate-500">
            Tente outro caractere ou significado.
          </p>
        </div>
      )}

      {!loading && results.length > 0 && (
        <section aria-live="polite" className="space-y-4">
          <div className="flex items-center justify-between">
            <h2 className="text-lg font-bold text-slate-800">
              Resultados
            </h2>

            <span className="text-sm text-slate-500">
              {results.length} {results.length === 1 ? 'ideograma' : 'ideogramas'}
            </span>
          </div>

          <div className="grid gap-4 sm:grid-cols-2">
            {results.map((ideogram) => (
              <article
                key={ideogram.id}
                className="rounded-2xl border border-slate-200 bg-white p-6 shadow-sm transition hover:border-emerald-200 hover:shadow-md"
              >
                <div className="flex items-start gap-5">
                  <div className="flex h-24 w-24 shrink-0 items-center justify-center rounded-2xl bg-[#e2efe6] text-5xl text-[#164b40]">
                    {ideogram.character}
                  </div>

                  <div className="min-w-0 flex-1">
                    <h3 className="text-xl font-bold text-slate-800">
                      {ideogram.pinyin}
                    </h3>

                    <p className="mt-2 break-words text-sm leading-relaxed text-slate-600">
                      {ideogram.meanings.join(', ')}
                    </p>
                  </div>
                </div>

                <div className="mt-6 flex flex-wrap gap-3 border-t border-slate-100 pt-4">
                  <span className="inline-flex items-center gap-1.5 rounded-lg bg-slate-100 px-3 py-2 text-xs text-slate-600">
                    <BookOpen size={14} />
                    Radical: {ideogram.radical || '—'}
                  </span>

                  <span className="inline-flex items-center gap-1.5 rounded-lg bg-slate-100 px-3 py-2 text-xs text-slate-600">
                    <Hash size={14} />
                    Traços: {ideogram.stroke_count}
                  </span>
                </div>
              </article>
            ))}
          </div>
        </section>
      )}
    </section>
  )
}
