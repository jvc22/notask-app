import { api } from '@/lib/ky'

interface CreateTaskRequest {
  title: string
  description?: string
}

export async function createTask({ title, description }: CreateTaskRequest) {
  await api.post('tasks', {
    json: {
      title,
      description,
    },
  })
}
