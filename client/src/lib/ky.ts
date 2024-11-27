import { CookiesFn, deleteCookie, getCookie } from 'cookies-next'
import ky from 'ky'
import { redirect } from 'next/navigation'
import { toast } from 'sonner'

let hasShownSessionToast = false

export const api = ky.create({
  prefixUrl: 'http://localhost:8080',
  retry: {
    limit: 0,
  },
  hooks: {
    beforeRequest: [
      async (request) => {
        let cookieStore: CookiesFn | undefined

        if (typeof window === 'undefined') {
          const { cookies: serverCookies } = await import('next/headers')

          cookieStore = serverCookies
        }

        const token = getCookie('notask-token', { cookies: cookieStore })

        if (token) {
          request.headers.set('Authorization', `Bearer ${token}`)
        }
      },
    ],
    afterResponse: [
      async (input, options, response) => {
        let cookieStore: CookiesFn | undefined

        if (typeof window === 'undefined') {
          const { cookies: serverCookies } = await import('next/headers')

          cookieStore = serverCookies
        }

        if (response.status === 401 && !hasShownSessionToast) {
          hasShownSessionToast = true

          deleteCookie('notask-token', { cookies: cookieStore })

          toast.warning('Your session expired.', {
            action: {
              label: 'Sign in',
              onClick: () => redirect('/auth/sign-in'),
            },
          })
        }
      },
    ],
  },
})
