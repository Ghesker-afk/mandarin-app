
import { useState } from 'react'
import {
  BookOpen,
  Home,
  Search,
  Languages,
  Layers3,
  ArrowRight,
  Sun,
  Menu,
  X,
} from 'lucide-react'
import SearchPage from './pages/SearchPage'

const navigation = [
  { label: 'Início', icon: Home },
  { label: 'Pesquisar', icon: Search },
  { label: 'Tradutor', icon: Languages },
  { label: 'Meus ideogramas', icon: Layers3 },
]

const features = [
  {
    title: 'Pesquisar ideogramas',
    description: 'Buscar por caractere ou significado.',
    icon: Search,
  },
  {
    title: 'Tradutor',
    description: 'Consultar os ideogramas conhecidos pela API.',
    icon: Languages,
  },
  {
    title: 'Meus ideogramas',
    description: 'Sua futura coleção pessoal de caracteres.',
    icon: Layers3,
  },
]

function App() {
  const [activePage, setActivePage] = useState('Início')
  const [sidebarOpen, setSidebarOpen] = useState(false)

  function navigateTo(page: string) {
    setActivePage(page)
    setSidebarOpen(false)
  }

  return (
    <div className="min-h-screen bg-[#f7f8f5] text-slate-800">
      {sidebarOpen && (
        <button
          aria-label="Fechar menu"
          onClick={() => setSidebarOpen(false)}
          className="fixed inset-0 z-40 bg-black/40 md:hidden"
        />
      )}

      <aside
        className={`fixed inset-y-0 left-0 z-50 flex w-72 flex-col
          bg-[#123b35] px-4 py-6 text-white transition-transform
          duration-200 md:translate-x-0
          ${sidebarOpen ? 'translate-x-0' : '-translate-x-full'}`}
      >
        <div className="mb-12 flex items-center justify-between px-2">
          <div className="flex items-center gap-3">
            <div className="rounded-xl bg-[#dcece2] p-3 text-[#164b40]">
              <BookOpen size={28} />
            </div>

            <div>
              <h1 className="text-xl font-bold">Mandarim</h1>
              <p className="mt-1 text-xs text-emerald-100">
                Seu espaço de aprendizagem
              </p>
            </div>
          </div>

          <button
            onClick={() => setSidebarOpen(false)}
            aria-label="Fechar menu"
            className="rounded-lg p-2 hover:bg-white/10 md:hidden"
          >
            <X size={20} />
          </button>
        </div>

        <nav className="space-y-2">
          {navigation.map(({ label, icon: Icon }) => (
            <button
              key={label}
              onClick={() => navigateTo(label)}
              className={`flex w-full items-center gap-4 rounded-xl
                px-4 py-3 text-left transition
                ${
                  activePage === label
                    ? 'bg-emerald-100/15 text-white'
                    : 'text-emerald-50/75 hover:bg-white/10 hover:text-white'
                }`}
            >
              <Icon size={21} />
              <span className="text-sm font-medium">{label}</span>
            </button>
          ))}
        </nav>

        <div className="mt-auto px-3 pb-2">
          <p className="text-sm leading-relaxed text-emerald-100/75">
            Grandes caracteres,
            <br />
            grandes histórias.
          </p>
          <div className="mt-4 h-0.5 w-8 rounded bg-emerald-200/70" />
        </div>
      </aside>

      <main className="min-h-screen md:ml-72">
        <header className="flex h-20 items-center justify-between border-b border-slate-200/80 px-5 sm:px-8 lg:px-12">
          <button
            onClick={() => setSidebarOpen(true)}
            aria-label="Abrir menu"
            className="rounded-lg p-2 hover:bg-slate-200/70 md:hidden"
          >
            <Menu size={23} />
          </button>

          <div className="hidden text-sm text-slate-500 md:block">
            Seu caminho para aprender mandarim
          </div>

          <button
            aria-label="Alternar tema (em breve)"
            title="Tema claro"
            className="rounded-full p-2 text-[#164b40] transition hover:bg-emerald-100"
          >
            <Sun size={23} />
          </button>
        </header>

        <div className="mx-auto max-w-7xl space-y-9 px-5 py-8 sm:px-8 lg:px-12 lg:py-10">
          {activePage === 'Pesquisar' ? (
            <SearchPage />
          ) : (
            <>
              <section className="relative isolate overflow-hidden rounded-3xl bg-gradient-to-br from-[#123b35] via-[#20584b] to-[#628b76] px-7 py-10 text-white sm:px-12 sm:py-14">
                <div
                  aria-hidden="true"
                  className="pointer-events-none absolute -right-12 -top-16 -z-10 h-72 w-72 rounded-full bg-emerald-100/10 blur-3xl"
                />

                <div
                  aria-hidden="true"
                  className="pointer-events-none absolute -bottom-12 right-8 -z-10 select-none text-[12rem] font-bold leading-none text-white/[0.06]"
                >
                  汉字
                </div>

                <p className="text-xs font-semibold uppercase tracking-[0.2em] text-emerald-100 sm:text-sm">
                  Aprenda um ideograma por vez
                </p>

                <h2 className="mt-5 text-4xl font-bold tracking-tight sm:text-5xl lg:text-6xl">
                  你好, mundo!
                </h2>

                <p className="mt-5 max-w-lg text-base leading-relaxed text-emerald-50/90 sm:text-lg">
                  Explore caracteres, descubra significados e construa seu
                  vocabulário, um ideograma por vez.
                </p>

                <div className="mt-8 flex items-center gap-3 text-sm text-emerald-100">
                  <span className="h-px w-8 bg-emerald-200" />
                  <span>你的中文之旅 — sua jornada no mandarim</span>
                </div>
              </section>

              <section>
                <div className="mb-5">
                  <h2 className="text-xl font-bold text-slate-800">
                    O que vamos estudar hoje?
                  </h2>
                  <p className="mt-1 text-sm text-slate-500">
                    Escolha por onde deseja começar.
                  </p>
                </div>

                <div className="space-y-4">
                  {features.map(({ title, description, icon: Icon }) => (
                    <button
                      key={title}
                      onClick={() =>
                        navigateTo(
                          title === 'Pesquisar ideogramas'
                            ? 'Pesquisar'
                            : title,
                        )
                      }
                      className="group flex w-full items-center gap-4 rounded-2xl
                        border border-slate-200 bg-white p-5 text-left
                        transition duration-200 hover:-translate-y-0.5
                        hover:border-emerald-200 hover:shadow-md
                        sm:gap-6 sm:p-6"
                    >
                      <div className="shrink-0 rounded-2xl bg-[#e2efe6] p-4 text-[#20584b] transition group-hover:bg-[#d2e8d9]">
                        <Icon size={27} />
                      </div>

                      <div className="min-w-0 flex-1">
                        <h3 className="font-semibold text-slate-800 sm:text-lg">
                          {title}
                        </h3>
                        <p className="mt-1 text-sm leading-relaxed text-slate-500">
                          {description}
                        </p>
                      </div>

                      <ArrowRight
                        size={22}
                        className="shrink-0 text-[#20584b] transition group-hover:translate-x-1"
                      />
                    </button>
                  ))}
                </div>
              </section>

              <footer className="pb-4 text-center text-xs text-slate-400">
                Mandarim App · Aprendizado contínuo, um caractere por vez.
              </footer>
            </>
          )}
        </div>
      </main>
    </div>
  )
}

export default App
