'use server'

import { revalidatePath } from 'next/cache'

import { api } from '@/lib/ky'

interface RemoveTaskRequest {
  taskId: number
}

export async function removeTask({ taskId }: RemoveTaskRequest) {
  const result = await api.delete(`tasks/${taskId}`)

  if (result.status === 200) {
    revalidatePath('/get-tasks')
  }

  return result.ok
}
