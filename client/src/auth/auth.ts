import { hasCookie } from 'cookies-next/server'
import { cookies } from 'next/headers'

export async function isAuthenticated() {
  return await hasCookie('notask-token', { cookies })
}
