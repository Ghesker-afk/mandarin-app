
import type { Ideogram } from '../types/ideogram'

const API_BASE_URL = '/api'

async function handleResponse<T>(response: Response): Promise<T> {
  if (!response.ok) {
    const message =
      response.status === 404
        ? 'Nenhum ideograma encontrado.'
        : `Erro na API: ${response.status}`

    throw new Error(message)
  }

  return response.json() as Promise<T>
}

export async function findByCharacter(
  character: string,
): Promise<Ideogram> {
  const response = await fetch(
    `${API_BASE_URL}/ideograms/${encodeURIComponent(character)}`,
  )

  return handleResponse<Ideogram>(response)
}

export async function searchByMeaning(
  meaning: string,
): Promise<Ideogram[]> {
  const params = new URLSearchParams({ meaning })

  const response = await fetch(
    `${API_BASE_URL}/ideograms/search?${params.toString()}`,
  )

  return handleResponse<Ideogram[]>(response)
}


export async function listAll(): Promise<Ideogram[]> {
  const response = await fetch(`${API_BASE_URL}/ideograms`);
  return handleResponse<Ideogram[]>(response);
}