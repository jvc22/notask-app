import { api } from '@/lib/ky'

export interface GetProfileResponse {
  username: string
}

export async function getProfile() {
  const result = await api.get('user/profile').json<GetProfileResponse>()

  return result
}
