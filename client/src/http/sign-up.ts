import { api } from '@/lib/ky'

interface SignUpRequest {
  username: string
  password: string
}

export async function signUpFn({ username, password }: SignUpRequest) {
  await api.post('auth/sign-up', {
    json: {
      username,
      password,
    },
  })
}
