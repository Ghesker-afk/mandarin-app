
import type { Ideogram } from '../types/ideogram'

export interface TranslationResponse {
  original_text: string
  literal_text: string
  ideograms: Ideogram[]
}

export async function translateText(
  text: string,
): Promise<TranslationResponse> {
  const params = new URLSearchParams({ text })

  const response = await fetch(`/api/translate?${params.toString()}`)

  if (!response.ok) {
    if (response.status === 404) {
      throw new Error('Não foi possível encontrar uma tradução para esse texto.')
    }

    throw new Error(`Erro ao consultar a API: ${response.status}`)
  }

  return response.json() as Promise<TranslationResponse>
}
