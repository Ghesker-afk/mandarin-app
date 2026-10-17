
import { useState, type FormEvent } from 'react'
import {
  Languages,
  LoaderCircle,
  ArrowDown,
  BookOpen,
  Search,
} from 'lucide-react'
import { translateText } from '../services/translationService'
import type { TranslationResponse } from '../services/translationService'

export default function TranslatorPage() {
  const [text, setText] = useState('')
  const [result, setResult] = useState<TranslationResponse | null>(null)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  async function handleTranslate(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()

    const value = text.trim()

    if (!value) {
      setError('Digite um texto para traduzir.')
      setResult(null)
      return
    }

    setLoading(true)
    setError('')
    setResult(null)

    try {
      const translation = await translateText(value)
      setResult(translation)
    } catch (err) {
      setError(
        err instanceof Error
          ? err.message
          : 'Não foi possível realizar a tradução.',
      )
    } finally {
      setLoading(false)
    }
  }

  return (
    <section className="mx-auto max-w-5xl space-y-8">
      <header>
        <p className="text-sm font-semibold uppercase tracking-widest text-emerald-800">
          Português e mandarim
        </p>

        <h1 className="mt-2 text-3xl font-bold tracking-tight text-slate-800 sm:text-4xl">
          Tradutor
        </h1>

        <p className="mt-3 max-w-2xl leading-relaxed text-slate-500">
          Consulte o texto na API e descubra quais ideogramas cadastrados
          correspondem aos caracteres enviados.
        </p>
      </header>

      <form
        onSubmit={handleTranslate}
        className="space-y-5 rounded-2xl border border-slate-200 bg-white p-5 shadow-sm sm:p-7"
      >
        <label
          htmlFor="translation-text"
          className="block text-sm font-semibold text-slate-700"
        >
          Texto para consultar
        </label>

        <textarea
          id="translation-text"
          value={text}
          onChange={(event) => setText(event.target.value)}
          placeholder="Digite um texto, por exemplo: 山 ou 你好"
          rows={4}
          className="w-full resize-y rounded-xl border border-slate-200 bg-slate-50 p-4 outline-none transition placeholder:text-slate-400 focus:border-emerald-700 focus:ring-2 focus:ring-emerald-700/10"
        />

        <div className="flex flex-wrap items-center justify-between gap-3">
          <p className="text-xs text-slate-400">
            Consulte os caracteres conhecidos pelo seu backend.
          </p>

          <button
            type="submit"
            disabled={loading}
            className="inline-flex items-center justify-center gap-2 rounded-xl bg-[#164b40] px-6 py-3 font-semibold text-white transition hover:bg-[#20584b] disabled:cursor-not-allowed disabled:opacity-60"
          >
            {loading ? (
              <LoaderCircle size={19} className="animate-spin" />
            ) : (
              <Languages size={19} />
            )}
            {loading ? 'Consultando...' : 'Traduzir'}
          </button>
        </div>
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
          className="flex items-center justify-center gap-3 py-8 text-slate-500"
        >
          <LoaderCircle size={22} className="animate-spin" />
          Consultando a API...
        </div>
      )}

      {result && !loading && (
        <section aria-live="polite" className="space-y-5">
          <div className="flex items-center gap-2 text-slate-500">
            <ArrowDown size={18} />
            <h2 className="text-lg font-bold text-slate-800">
              Resultado
            </h2>
          </div>

          <div className="rounded-2xl border border-slate-200 bg-white p-6 shadow-sm">
            <p className="text-sm font-medium text-slate-500">
              Texto original
            </p>

            <p className="mt-2 break-words text-2xl font-semibold text-slate-800">
              {result.original_text}
            </p>

            <div className="mt-5 border-t border-slate-100 pt-5">
              <p className="text-sm font-medium text-slate-500">
                Resultado literal
              </p>

              <p className="mt-2 break-words text-lg leading-relaxed text-[#164b40]">
                {result.literal_text || 'Nenhum significado literal disponível.'}
              </p>
            </div>
          </div>

          <div>
            <div className="mb-4 flex items-center justify-between gap-3">
              <h2 className="text-lg font-bold text-slate-800">
                Ideogramas identificados
              </h2>

              <span className="text-sm text-slate-500">
                {result.ideograms.length}{' '}
                {result.ideograms.length === 1 ? 'ideograma' : 'ideogramas'}
              </span>
            </div>

            {result.ideograms.length === 0 ? (
              <div className="rounded-2xl border border-dashed border-slate-300 px-6 py-10 text-center">
                <Search size={30} className="mx-auto text-slate-400" />

                <p className="mt-3 font-medium text-slate-700">
                  Nenhum ideograma cadastrado foi identificado.
                </p>

                <p className="mt-2 text-sm text-slate-500">
                  Isso não significa necessariamente que o texto não tenha
                  tradução; pode ser que os caracteres ainda não estejam
                  cadastrados na sua base.
                </p>
              </div>
            ) : (
              <div className="grid gap-4 sm:grid-cols-2">
                {result.ideograms.map((ideogram) => (
                  <article
                    key={ideogram.id}
                    className="rounded-2xl border border-slate-200 bg-white p-5 shadow-sm"
                  >
                    <div className="flex items-center gap-4">
                      <div className="flex h-20 w-20 shrink-0 items-center justify-center rounded-2xl bg-[#e2efe6] text-4xl text-[#164b40]">
                        {ideogram.character}
                      </div>

                      <div className="min-w-0">
                        <h3 className="text-lg font-bold text-slate-800">
                          {ideogram.pinyin}
                        </h3>

                        <p className="mt-1 break-words text-sm text-slate-600">
                          {ideogram.meanings.join(', ')}
                        </p>
                      </div>
                    </div>

                    <div className="mt-4 flex items-center gap-2 border-t border-slate-100 pt-4 text-sm text-slate-500">
                      <BookOpen size={16} />
                      Radical: {ideogram.radical || '—'}
                      <span className="ml-auto">
                        Traços: {ideogram.stroke_count}
                      </span>
                    </div>
                  </article>
                ))}
              </div>
            )}
          </div>
        </section>
      )}
    </section>
  )
}
