import { api } from '@/lib/ky'

interface SignInRequest {
  username: string
  password: string
}

interface SignInResponse {
  token: string
}

export async function signInFn({ username, password }: SignInRequest) {
  const result = await api
    .post('auth/sign-in', {
      json: {
        username,
        password,
      },
    })
    .json<SignInResponse>()

  return result
}
