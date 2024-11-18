import { api } from '@/lib/ky'

interface RemoveTaskRequest {
  taskId: number
}

export async function removeTask({ taskId }: RemoveTaskRequest) {
  await api.delete(`tasks/${taskId}`)
}
