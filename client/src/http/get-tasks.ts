import { api } from '@/lib/ky'

export interface GetTasksResponse {
  tasks: {
    id: number
    title: string
    description: string
  }[]
}

export async function getTasks() {
  const result = await api.get('tasks').json<GetTasksResponse>()

  return result
}
