import { useEffect, useState } from 'react'
import { BookOpen, Hash, LoaderCircle, RefreshCw } from 'lucide-react'
import { listAll } from '../services/ideogramService'
import type { Ideogram } from '../types/ideogram'

export default function MyIdeogramsPage() {
  const [ideograms, setIdeograms] = useState<Ideogram[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  useEffect(() => {
    async function loadIdeograms() {
      try {
        setLoading(true)
        setError('')

        const data = await listAll()
        setIdeograms(data)
      } catch (err) {
        setError(
          err instanceof Error
            ? err.message
            : 'Não foi possível carregar seus ideogramas.',
        )
      } finally {
        setLoading(false)
      }
    }

    loadIdeograms()
  }, [])

  async function handleRetry() {
    try {
      setLoading(true)
      setError('')

      const data = await listAll()
      setIdeograms(data)
    } catch (err) {
      setError(
        err instanceof Error
          ? err.message
          : 'Não foi possível carregar seus ideogramas.',
      )
    } finally {
      setLoading(false)
    }
  }

  return (
    <section className="mx-auto max-w-5xl space-y-8">
      <header>
        <p className="text-sm font-semibold uppercase tracking-widest text-emerald-800">
          Seu vocabulário
        </p>

        <h1 className="mt-2 text-3xl font-bold tracking-tight text-slate-800 sm:text-4xl">
          Meus ideogramas
        </h1>

        <p className="mt-3 max-w-2xl leading-relaxed text-slate-500">
          Consulte os caracteres chineses cadastrados e revise seus
          significados, pronúncias e características.
        </p>
      </header>

      {!loading && !error && (
        <div className="flex items-center justify-between rounded-2xl border border-slate-200 bg-white p-5 shadow-sm">
          <div>
            <p className="text-sm text-slate-500">Total cadastrado</p>
            <p className="mt-1 text-2xl font-bold text-slate-800">
              {ideograms.length}
            </p>
          </div>

          <BookOpen size={28} className="text-emerald-800" />
        </div>
      )}

      {loading && (
        <div
          role="status"
          className="flex items-center justify-center gap-3 py-12 text-slate-500"
        >
          <LoaderCircle size={22} className="animate-spin" />
          Carregando seus ideogramas...
        </div>
      )}

      {error && (
        <div
          role="alert"
          className="rounded-xl border border-red-200 bg-red-50 p-5 text-sm text-red-700"
        >
          <p>{error}</p>

          <button
            type="button"
            onClick={handleRetry}
            className="mt-4 inline-flex items-center gap-2 rounded-lg bg-red-100 px-4 py-2 font-semibold transition hover:bg-red-200"
          >
            <RefreshCw size={16} />
            Tentar novamente
          </button>
        </div>
      )}

      {!loading && !error && ideograms.length === 0 && (
        <div className="rounded-2xl border border-dashed border-slate-300 px-6 py-12 text-center">
          <BookOpen size={32} className="mx-auto text-slate-400" />

          <h2 className="mt-4 font-semibold text-slate-700">
            Nenhum ideograma cadastrado
          </h2>

          <p className="mt-2 text-sm text-slate-500">
            Quando houver ideogramas cadastrados, eles aparecerão aqui.
          </p>
        </div>
      )}

      {!loading && !error && ideograms.length > 0 && (
        <section aria-live="polite" className="space-y-4">
          <h2 className="text-lg font-bold text-slate-800">
            Seus caracteres
          </h2>

          <div className="grid gap-4 sm:grid-cols-2">
            {ideograms.map((ideogram) => (
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
                      {ideogram.meanings.join(', ') || 'Sem significado cadastrado'}
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